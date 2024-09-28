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

		cchunk := make([]byte, len(data[i:end]))
		copy(cchunk, data[i:end])
		chunk = append(chunk, cchunk)
	}

	return chunk
}
