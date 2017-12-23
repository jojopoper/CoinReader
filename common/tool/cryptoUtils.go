package tool

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

// CryptoUtils : crypto utils define
type CryptoUtils struct {
}

// GetHmacSHA512 : HmacHash512
func (ths *CryptoUtils) GetHmacSHA512(key string, params ...string) (string, error) {
	if len(params) < 1 {
		return "", fmt.Errorf("Enter at least one \"param\" parameter")
	}
	hmac512 := hmac.New(sha512.New, []byte(key))
	for _, s := range params {
		_, err := hmac512.Write([]byte(s))
		if err != nil {
			return "", err
		}
	}
	return hex.EncodeToString(hmac512.Sum(nil)), nil
}

// GetMd5 : md5
func (ths *CryptoUtils) GetMd5(params ...string) (string, error) {
	if len(params) < 1 {
		return "", fmt.Errorf("Enter at least one \"param\" parameter")
	}
	m := md5.New()
	for _, s := range params {
		_, err := m.Write([]byte(s))
		if err != nil {
			return "", err
		}
	}
	return hex.EncodeToString(m.Sum(nil)), nil
}
