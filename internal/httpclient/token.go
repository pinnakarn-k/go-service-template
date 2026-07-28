package httpclient

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"time"
)

type Base64Mode string

const (
	Base64Standard Base64Mode = "standard"
	Base64URL      Base64Mode = "url"
	Base64RawURL   Base64Mode = "raw_url"
)

type TokenConfig struct {
	Application string
	Requester   string
	Key         string
	Base64Mode  Base64Mode
}

type TokenHeaders struct {
	Application string
	Requester   string
	Pretoken    string
	Token       string
}

func GenerateTokenHeaders(cfg TokenConfig) (TokenHeaders, error) {
	encoder, err := getBase64Encoding(cfg.Base64Mode)
	if err != nil {
		return TokenHeaders{}, err
	}

	timestamp := time.Now().Format("20060102150405")
	iv := "KS" + timestamp

	pretoken := encoder.EncodeToString([]byte(timestamp))

	token, err := encryptAESCBC(
		timestamp,
		cfg.Key,
		iv,
		encoder,
	)
	if err != nil {
		return TokenHeaders{}, fmt.Errorf("encrypt token: %w", err)
	}

	return TokenHeaders{
		Application: cfg.Application,
		Requester:   cfg.Requester,
		Pretoken:    pretoken,
		Token:       token,
	}, nil
}

func getBase64Encoding(mode Base64Mode) (*base64.Encoding, error) {
	switch mode {
	case "", Base64Standard:
		return base64.StdEncoding, nil

	case Base64URL:
		return base64.URLEncoding, nil

	case Base64RawURL:
		return base64.RawURLEncoding, nil

	default:
		return nil, fmt.Errorf("unsupported base64 mode: %q", mode)
	}
}

func encryptAESCBC(
	plainText string,
	key string,
	iv string,
	encoder *base64.Encoding,
) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("create AES cipher: %w", err)
	}

	plainBytes := pkcs7Padding(
		[]byte(plainText),
		aes.BlockSize,
	)

	cipherText := make([]byte, len(plainBytes))

	mode := cipher.NewCBCEncrypter(block, []byte(iv))
	mode.CryptBlocks(cipherText, plainBytes)

	return encoder.EncodeToString(cipherText), nil
}

func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize

	return append(
		data,
		bytes.Repeat([]byte{byte(padding)}, padding)...,
	)
}
