package utils

import (
	"bytes"
	"compress/gzip"
)

// Gzip compresses data using gzip compression
func Gzip(data []byte) ([]byte, error) {
	buf := new(bytes.Buffer)
	gz := gzip.NewWriter(buf)

	if _, err := gz.Write(data); err != nil {
		return nil, err
	}

	if err := gz.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
