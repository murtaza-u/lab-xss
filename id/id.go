package id

import (
	"crypto/rand"
	"io"
)

var chars = []byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C',
	'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
	'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c',
	'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p',
	'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
}

func NewID() (string, error) {
	b := make([]byte, 10)
	_, err := io.ReadAtLeast(rand.Reader, b, 10)
	if err != nil {
		return "", err
	}

	for i := 0; i < len(b); i++ {
		b[i] = chars[int(b[i])%len(chars)]
	}

	return string(b), nil
}
