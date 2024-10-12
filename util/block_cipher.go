package util

func Chunkify(data []byte, size int) [][]byte {
	if size < 1 {
		panic("Chunkify: Invalid chunksize")
	}

	var chunks [][]byte

	/*
		for i := 0; i < len(data); i += size {
			end := i + size

			if end > len(data) {
				end = len(data)
			}

			cchunk := make([]byte, len(data[i:end]))
			copy(cchunk, data[i:end])
			chunk = append(chunk, cchunk)
		}
	*/

	for i := 0; i < len(data); i += size {
		end := i + size
		if end > len(data) {
			end = len(data)
		}
		chunks = append(chunks, data[i:end])
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
