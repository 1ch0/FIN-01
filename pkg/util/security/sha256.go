package security

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
)

// SHA256 哈希
func SHA256(data []byte) string {
	sig := sha256.Sum256(data)
	return fmt.Sprintf("%x", sig[:])
}

// SHA256WithRSABase64 SHA256WithRSA签名算法签名，返回base64编码后的签名
func SHA256WithRSABase64(data, privateKey []byte) (string, error) {
	sign, err := SHA256WithRSA(data, privateKey, false)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(sign), nil
}

// SHA256WithRSA SHA256WithRSA签名算法签名
func SHA256WithRSA(data, privateKey []byte, noneWithRsa bool) ([]byte, error) {
	priKey, err := ParsePKCS1Or8PrivKey(privateKey)
	if err != nil {
		return nil, err
	}

	hashed := sha256.Sum256(data)
	if noneWithRsa {
		return rsa.SignPKCS1v15(rand.Reader, priKey, crypto.Hash(0), hashed[:])
	}
	return rsa.SignPKCS1v15(rand.Reader, priKey, crypto.SHA256, hashed[:])
}

// VerifySHA256WithRSABase64 SHA256WithRSA签名算法验签，如果验签通过，则err 值为 nil
func VerifySHA256WithRSABase64(origin []byte, b64Sign, publicKey string) error {
	pubKey, err := ParsePublicKey([]byte(publicKey))
	if err != nil {
		return err
	}

	sign, err := base64.StdEncoding.DecodeString(b64Sign)
	if err != nil {
		return err
	}

	hashed := sha256.Sum256(origin)
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed[:], sign)

	return err
}

// SHA256Sign 共用的签名算法
func SHA256Sign(httpMethod, url, dateTime, signKey string, body []byte) (string, error) {
	if signKey == "" {
		return "", errors.New("empty sign key")
	}

	// 1. 获得 sha256签名的原文
	signContext := fmt.Sprintf("%s\n%s\n%s\n%s\n%s", httpMethod, url, dateTime, body, signKey)
	// 2. 进行 SHA256签名
	signature := SHA256([]byte(signContext))
	return signature, nil
}
