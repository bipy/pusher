package utils

import (
	"bytes"
	"compress/gzip"
)

func Gzip(data []byte) ([]byte, error) {
	buf := new(bytes.Buffer)
	gz := gzip.NewWriter(buf)
	_, err := gz.Write(data)
	if err != nil {
		return nil, err
	}
	err = gz.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
