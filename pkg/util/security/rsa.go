package security

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"

	"github.com/sirupsen/logrus"
)

var PrivateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDZsfv1qscqYdy4vY+P4e3cAtmvppXQcRvrF1cB4drkv0haU24Y
7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0DgacdwYWd/7PeCELyEipZJL07Vro7
Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NLAUeJ6PeW+DAkmJWF6QIDAQAB
AoGBAJlNxenTQj6OfCl9FMR2jlMJjtMrtQT9InQEE7m3m7bLHeC+MCJOhmNVBjaM
ZpthDORdxIZ6oCuOf6Z2+Dl35lntGFh5J7S34UP2BWzF1IyyQfySCNexGNHKT1G1
XKQtHmtc2gWWthEg+S6ciIyw2IGrrP2Rke81vYHExPrexf0hAkEA9Izb0MiYsMCB
/jemLJB0Lb3Y/B8xjGjQFFBQT7bmwBVjvZWZVpnMnXi9sWGdgUpxsCuAIROXjZ40
IRZ2C9EouwJBAOPjPvV8Sgw4vaseOqlJvSq/C/pIFx6RVznDGlc8bRg7SgTPpjHG
4G+M3mVgpCX1a/EU1mB+fhiJ2LAZ/pTtY6sCQGaW9NwIWu3DRIVGCSMm0mYh/3X9
DAcwLSJoctiODQ1Fq9rreDE5QfpJnaJdJfsIJNtX1F+L3YceeBXtW0Ynz2MCQBI8
9KP274Is5FkWkUFNKnuKUK4WKOuEXEO+LpR+vIhs7k6WQ8nGDd4/mujoJBr5mkrw
DPwqA3N5TMNDQVGv8gMCQQCaKGJgWYgvo3/milFfImbp+m7/Y3vCptarldXrYQWO
AQjxwc71ZGBFDITYvdgJM1MTqc8xQek1FXn1vfpy2c6O
-----END RSA PRIVATE KEY-----
`)

var PublicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZsfv1qscqYdy4vY+P4e3cAtmv
ppXQcRvrF1cB4drkv0haU24Y7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0Dgacd
wYWd/7PeCELyEipZJL07Vro7Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NL
AUeJ6PeW+DAkmJWF6QIDAQAB
-----END PUBLIC KEY-----
`)

func ParsePKCS1Or8PrivKey(key []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("priv key error")
	}
	var privKey *rsa.PrivateKey
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		return privKey, nil
	}

	privKeyInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err == nil {
		if priKey, ok := privKeyInterface.(*rsa.PrivateKey); ok {
			return priKey, nil
		}
	}
	return nil, errors.New("parse private key error")
}

// ParsePublicKey 解析公钥
func ParsePublicKey(key []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("public key error")
	}
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	if publicKey, ok := publicKeyInterface.(*rsa.PublicKey); ok {
		return publicKey, nil
	}
	return nil, errors.New("public key error")
}

// RSAEncrypt RSA 加密
func RSAEncrypt(origData, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		logrus.Errorf("x509.ParsePKIXPublicKey error: %s", err)
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)

	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// RSAEncryptBase64 RSA 加密 Base64 密文
func RSAEncryptBase64(origData, publicKey []byte) (string, error) {
	cipherText, err := RSAEncrypt(origData, publicKey)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// RSAEncryptDERType RSA加密，密钥是 DER 格式
func RSAEncryptDERType(origData, publicDERKey []byte) ([]byte, error) {
	pubInterface, err := x509.ParsePKIXPublicKey(publicDERKey)
	if err != nil {
		logrus.Errorf("x509.ParsePKIXPublicKey error: %s", err)
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)

	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// RSADecrypt RSA 解密
func RSADecrypt(ciphertext, privateKey []byte) ([]byte, error) {
	priv, err := ParsePKCS1Or8PrivKey(privateKey)
	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

// RSADecryptBase64 RSA 解密 Base64 密文
func RSADecryptBase64(b64Cipher string, privateKey []byte) ([]byte, error) {
	cipherText, err := base64.StdEncoding.DecodeString(b64Cipher)
	if err != nil {
		logrus.Errorf("base64.StdEncoding.DecodeString error! b64Cipher:%s, err:%s", b64Cipher, err)
		return nil, err
	}

	return RSADecrypt(cipherText, privateKey)
}

// RSASign RSA 签名
func RSASign(text []byte, privateKey string, hashType crypto.Hash) ([]byte, error) {
	priv, err := ParsePKCS1Or8PrivKey([]byte(privateKey))
	if err != nil {
		logrus.Errorf("x509.ParsePKCS1PrivateKey error: %s", err)
		return nil, err
	}

	return rsa.SignPKCS1v15(rand.Reader, priv, hashType, text)
}

// RSASign RSA 签名 Base64
func RSASignBase64(text []byte, privateKey string, hashType crypto.Hash) (string, error) {
	sign, err := RSASign(text, privateKey, hashType)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(sign), nil
}

// RSAVerifySign RSA验证签名
func RSAVerifySign(text []byte, publicKey string, hashType crypto.Hash, sig []byte) error {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		logrus.Errorf("pem.Decode error")
		return errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		logrus.Errorf("x509.ParsePKIXPublicKey error: %s", err)
		return err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.VerifyPKCS1v15(pub, hashType, text, sig)
}

// RSAVerifySignBase64 RSA验证Base64签名
func RSAVerifySignBase64(b64Cipher, publicKey string, hashType crypto.Hash, text []byte) error {
	cipherText, err := base64.StdEncoding.DecodeString(b64Cipher)
	if err != nil {
		cipherText, err = base64.RawURLEncoding.DecodeString(b64Cipher)
		if err != nil {
			logrus.Errorf("base64.StdEncoding.DecodeString error! b64Cipher:%s, err:%s", b64Cipher, err)
			return err
		}
	}

	return RSAVerifySign(text, publicKey, hashType, cipherText)
}
