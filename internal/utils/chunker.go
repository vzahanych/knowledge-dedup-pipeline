package utils

import (
	"bufio"
	"io"
)

// ShingleStream chunks a stream into overlapping windows.
func ShingleStream(r io.Reader, shingleSize int, overlap int) [][]byte {
	var shingles [][]byte
	reader := bufio.NewReader(r)
	window := make([]byte, 0, shingleSize)
	buf := make([]byte, shingleSize)

	for {
		n, err := reader.Read(buf)
		if n > 0 {
			window = append(window, buf[:n]...)
			for len(window) >= shingleSize {
				shingles = append(shingles, append([]byte{}, window[:shingleSize]...))
				window = window[shingleSize-overlap:]
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
	}
	return shingles
}
