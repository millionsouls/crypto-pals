package main

//PKCS7 verification
import (
	"crypto-pals/lib/util"
	"fmt"
)

func main() {
	first := "ICE ICE BABY\x04\x04\x04\x04"
	sec := "ICE ICE BAB\x05\x05\x05\x05\x05"
	third := "ICE ICE BABY\x01\x02\x03\x04"

	fmt.Println(util.DetectPad(first, 16))
	fmt.Println(util.DetectPad(sec, 16))
	fmt.Println(util.DetectPad(third, 16))
}
