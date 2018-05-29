package util

import (
	"bytes"
)

const (
	str = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// ConvertTo62 covert i to 62 system string
func ConvertTo62(i int64) string {
	var buf = bytes.NewBuffer(nil)
	f(i, 62, buf)
	b := buf.Bytes()

	return string(transpose(b))
}

func f(i, j int64, b *bytes.Buffer) {
	if i == 0 {
		return
	}
	b.WriteByte(str[i%j])
	f(i/j, j, b)
}

func transpose(b []byte) []byte {
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}
