package util

import "strings"

// help functions for block ciphers
// chunkify - chunks a byte slice in 2x2 # sized blocks
// pkcs7 - pads any blocks that are not # size
// 			includes removal of pad

func Chunkify(data []byte, size int) [][]byte {
	if size < 1 {
		panic("Chunkify: Invalid chunksize")
	}

	var chunks [][]byte

	for i := 0; i < len(data); i += size {
		end := i + size

		if end > len(data) {
			end = len(data)
		}

		cchunk := make([]byte, len(data[i:end]))
		copy(cchunk, data[i:end])
		chunks = append(chunks, cchunk)
	}

	return chunks
}

func PKCS7(data []byte, size int) []byte {
	var padLen int

	if len(data) > size {
		panic("PKCS7: Data greater than block size")
	}

	if size == 16 {
		padLen = size - (len(data) & 15)
	} else {
		padLen = size - (len(data) % size)
	}

	padding := make([]byte, padLen)

	for i := range padding {
		padding[i] = byte(padLen)
	}

	result := append(data, padding...)

	return result
}

func UnPad(data []byte) []byte {
	return data[:len(data)-int(data[len(data)-1])]
}

func DetectPad(str string, size int) bool {
	if len(str)%size != 0 {
		panic("String is not a multiple of the size")
	}

	sLen := len(str)
	lastByte := str[sLen-1]
	trim := strings.TrimRight(str, string(lastByte))

	if sLen-len(trim) != int(lastByte) {
		return false
	}

	return true
}
