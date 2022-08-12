package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
)

// AESCFBMode CFB 模式的 AES 加密
type AESCFBMode struct {
	Key []byte
	Err error
}

// NewAESCFBEncrypt 创建AES加密处理对象
func NewAESCFBEncrypt(sKey string) *AESCFBMode {
	return &AESCFBMode{
		Key: []byte(sKey),
		Err: nil,
	}
}

// Encrypt aesCFBEncrypt aes 加密  对商户敏感信息加密
func (a *AESCFBMode) Encrypt(pt string) string {
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		a.Err = err
		logrus.Error("create the block error")
		return ""
	}
	plaintext := []byte(pt)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		a.Err = err
		logrus.Error("generate the rand num error")
		return ""
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return hex.EncodeToString(ciphertext)
}

// Decrypt aes 解密
func (a *AESCFBMode) Decrypt(ct string) string {
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		a.Err = err
		logrus.Error("create the block error")
		return ""
	}

	ciphertext, err := hex.DecodeString(ct)
	if err != nil {
		a.Err = err
		logrus.Errorf("Decode string error, the string is %s, error is %s", ct, err.Error())
		return ""
	}

	if len(ciphertext) < aes.BlockSize {
		err = fmt.Errorf("the length of the ciphertext is too short")
		a.Err = err
		logrus.Errorf("%s", err)
		return ""
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext)
}
