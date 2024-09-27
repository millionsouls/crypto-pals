package util

func Chunkify(data []byte, size int) [][]byte {
	if size < 1 {
		panic("Invalid chunksize")
	}

	var chunk [][]byte

	for i := 0; i < len(data); i += size {
		end := i + size

		if end > len(data) {
			end = len(data)
		}
		nChunk := make([]byte, len(data[i:end]))

		copy(nChunk, data[i:end])

		chunk = append(chunk, nChunk)
	}

	return chunk
}
