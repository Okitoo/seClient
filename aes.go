package seClient

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"
)

func Encrypt(text_simple []byte, key string) ([]byte, error) {
	text := Pad(append([]byte(nil), text_simple...))
	block, err := aes.NewCipher(_32BytesHash([]byte(key)))
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCBCEncrypter(block, iv)
	stream.CryptBlocks(ciphertext[aes.BlockSize:], text)

	return ciphertext, nil
}

func Pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func Unpad(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])

	if unpadding > length {
		return nil, errors.New("unpad error. This could happen when incorrect encryption key is used")
	}

	return src[:(length - unpadding)], nil
}

func _32BytesHash(input []byte) (output []byte) {

	if len(input) == 0 {
		return nil
	}

	hasher := sha256.New()
	hasher.Write(input)

	// Cut the length down to 32 bytes and return.
	return hasher.Sum(nil)[:32]
}
