package server

import (
	"crypto-pals/lib/crysuite"
	"crypto/rand"
	"fmt"
	"math/big"
	"net"
	"os"
)

var N, _ = new(big.Int).SetString("fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffb", 16)
var g = big.NewInt(2)
var k = big.NewInt(3)

func passVerify(salt []byte, password string) *big.Int {
	xH := crysuite.SHA256(append(salt, []byte(password)...))
	x := new(big.Int).SetBytes(crysuite.SHA256(xH))
	v := new(big.Int).Exp(g, x, N)
	return v
}

func generateServerB(b *big.Int, v *big.Int) *big.Int {
	B := new(big.Int).Set(k)
	B.Mul(B, v)
	temp := new(big.Int).Exp(g, b, N)
	B.Add(B, temp)
	B.Mod(B, N)
	return B
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	tbuf := make([]byte, 1024)
	msg, _ := conn.Read(tbuf)
	fmt.Println("Received message:", string(tbuf[:msg]))

	// Begin SRP
	// Simulate password and salt
	password := "password"
	salt, err := crysuite.GenerateSalt()
	if err != nil {
		fmt.Println("Error generating salt:", err)
		return
	}

	// Server-side private key b
	b, _ := rand.Int(rand.Reader, N)
	v := passVerify(salt, password)
	B := generateServerB(b, v)

	// Send salt and B to client
	fmt.Printf("Server Salt: %x\n", salt)
	conn.Write(append(salt, B.Bytes()...))
	fmt.Println("Sent salt and B to client")
	fmt.Printf("Server B: %x\n", B)

	// Receive the client's A (g^a % N)
	A := new(big.Int)
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from client:", err)
		return
	}
	A.SetBytes(buf[:n])
	fmt.Printf("Received A: %x\n", A)

	// Step 3: Server computes u = SHA256(A | B)
	uH := crysuite.SHA256(append(A.Bytes(), B.Bytes()...))
	u := new(big.Int).SetBytes(crysuite.SHA256(uH))

	// Calculate x for server-side computation of S
	xH := crysuite.SHA256(append(salt, []byte(password)...))
	x := new(big.Int).SetBytes(crysuite.SHA256(xH))
	fmt.Printf("Server SHA256: %x\n", x)

	// Compute the session key S (server-side)
	S_server := new(big.Int).Mul(A, new(big.Int).Exp(v, u, N)) // A * v^u % N
	S_server.Exp(S_server, b, N)                               // S_server = S_server^b % N
	fmt.Printf("Server: b = %s, u = %s, x = %s\n", b, u, x)

	// Derive the session key K from S
	K_server := crysuite.SHA256(S_server.Bytes())
	fmt.Printf("Server session key: %x\n", K_server)

	// Send HMAC to client for validation
	hmacServer := crysuite.HMAC256(K_server, salt)
	conn.Write(hmacServer)
	fmt.Println("Sent HMAC to client")
}

func Start() {
	// Start TCP server
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer ln.Close()
	fmt.Println("Server is listening on port 8080...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		fmt.Println("Received connection from:", conn.RemoteAddr())
		go handleConnection(conn)
	}
}
