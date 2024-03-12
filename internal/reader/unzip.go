package reader

import (
	"bytes"
	"compress/flate"
	"compress/zlib"
	"fmt"
	"io"
)

func DecompressZlib(compressedData []byte) ([]byte, error) {
	reader, err := zlib.NewReader(bytes.NewReader(compressedData))
	if err != nil {
		return nil, fmt.Errorf("failed to create zlib reader: %v", err)
	}
	defer func(reader io.ReadCloser) {
		err := reader.Close()
		if err != nil {
			fmt.Println("Error closing reader:", err)
		}
	}(reader)

	decompressedData, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read decompressed data: %v", err)
	}

	return decompressedData, nil
}

func DecompressDeflate(compressedData []byte) ([]byte, error) {
	reader := flate.NewReader(bytes.NewReader(compressedData))
	defer func(reader io.ReadCloser) {
		err := reader.Close()
		if err != nil {
			fmt.Println("Error closing reader:", err)
		}
	}(reader)

	decompressedData, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read decompressed data: %v", err)
	}

	return decompressedData, nil
}
