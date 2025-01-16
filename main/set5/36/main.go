package main

import (
	"crypto-pals/server"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"net"
	"os"
	"time"
)

var N, _ = new(big.Int).SetString("fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffb", 16)
var g = big.NewInt(2)
var k = big.NewInt(3)

func sha256Hash(data []byte) []byte {
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}

func sha256Int(data []byte) *big.Int {
	hash := sha256Hash(data)
	return new(big.Int).SetBytes(hash)
}

func hmacSHA256(key, data []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}

func main() {
	// Setup and connect to the server
	go server.Start()

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connected to server")

	// Sleep for 5 seconds to allow the server to start
	time.Sleep(5 * time.Second)

	fmt.Println("Sending test messasge")
	conn.Write([]byte("test test test test"))

	// Client-side private key
	a, _ := rand.Int(rand.Reader, N)
	A := new(big.Int).Exp(g, a, N)

	// Send A to the server
	conn.Write(A.Bytes())
	fmt.Println("Sent A to server")
	fmt.Printf("A: %x\n", A)

	// Receive salt and B from the server
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from server:", err)
		return
	}

	salt := buf[:16]
	B := new(big.Int).SetBytes(buf[16:n])
	fmt.Printf("Salt: %x\n", salt)
	fmt.Printf("B: %x\n", B)

	// Compute u = SHA256(A | B)
	uH := sha256Hash(append(A.Bytes(), B.Bytes()...))
	u := sha256Int(uH)

	// Compute S (client-side)
	// Normally, we would use x derived from the password
	xH := sha256Hash(append(salt, []byte("password")...))
	x := sha256Int(xH)
	fmt.Printf("Client SHA256: %x\n", x)

	S_client := new(big.Int).Sub(B, new(big.Int).Mul(k, new(big.Int).Exp(g, x, N))) // (B - k * g^x) % N
	S_client.Exp(S_client, new(big.Int).Add(a, new(big.Int).Mul(u, x)), N)          // S_client = S_client^(a + u * x) % N

	fmt.Printf("Client: a = %s, u = %s, x = %s\n", a, u, x)

	// Derive session key K from S
	K_client := sha256Hash(S_client.Bytes())
	fmt.Printf("Client session key: %x\n", K_client)

	// Receive HMAC from the server
	time.Sleep(2 * time.Second)
	hmacServer := make([]byte, 32)
	_, err = conn.Read(hmacServer)
	if err != nil {
		fmt.Println("Error reading HMAC from server:", err)
		return
	}

	// Validate HMAC
	hmacClient := hmacSHA256(K_client, salt)
	fmt.Printf("HMAC Server: %x\n", hmacServer)
	fmt.Printf("HMAC Client: %x\n", hmacClient)
	if hmac.Equal(hmacClient, hmacServer) {
		fmt.Println("***HMAC success***")
	} else {
		fmt.Println("***HMAC failed***")
	}
}
