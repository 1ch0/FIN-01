package security

import (
	"encoding/base64"
	"testing"
)

var b64cipher = "Mnwr4OoDjAJhO45Xl6H4SPH7zaGHIuSFwlXcPxb0ztTOskmfpMZenh6RFYCqDCBVz2W5CS7R3kAbebvg42/vVWs4smY4k3d+5i839ujJEvYF6NqyITXY4yD3U6ClZerKF5mcqa0znccdECyM7hUyqJUHH5Qs6RODsWGwqDDLmqTFWeCaNm1g1m1PCsUk/92lTxkKNEPjCG+vDCf9EnsEmr8D4b/t1XVCPq83rYHE8ezdINasOsRKaeLid0Tdo4vO3POLJnv56KUdgAgBcvnCZsu4/JdHDYEiNC6aSXXj3l2yQms8gpEQB/+IcZtrZrTV7zOnZstvSS2cSQabI73+Yw=="
var origData = "Aa123456"

func TestRSAEncrypt(t *testing.T) {
	ciphertext, err := RSAEncryptBase64([]byte(origData), PublicKey)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(ciphertext)
	// 相同的源文，每次加密后的密文都不相同，不能根据加密后的数据是否相等来判断加密算法是否正确
	// 要把加密后得数据再解码，如果和源文一致，说明加密算法正确
	actual, err := RSADecryptBase64(ciphertext, PrivateKey)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(string(actual))
	if string(actual) != origData {
		t.Errorf("decrypt failed:  expected=%s, actual=%s", origData, actual)
	}
}

func TestRSADecrypt(t *testing.T) {
	ciphertext, err := base64.StdEncoding.DecodeString(b64cipher)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	actual, err := RSADecrypt(ciphertext, PrivateKey)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if string(actual) != origData {
		t.Errorf("decrypt failed: ciphertext=%s, expected=%s, actual=%s", ciphertext, origData, actual)
	}
}
