package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

/**


生成自签名证书：
生成证书  openssl req -x509 -new -nodes -key data/rsa_private_key.pem -subj "/CN=localhost" -addext "subjectAltName=DNS:localhost" -days 3650 -out data/server.crt
*/

var (
	publicKey  []byte
	privateKey []byte
)

func ReadFile(keyFile string) ([]byte, error) {
	if f, err := os.Open(keyFile); err != nil {
		return nil, err
	} else {
		content := make([]byte, 4096)
		if n, err := f.Read(content); err != nil {
			return nil, err
		} else {
			return content[:n], nil
		}
	}
}

func ReadRSAKey(publicKeyFile, privateKeyFile string) (err error) {
	if publicKey, err = ReadFile(publicKeyFile); err != nil {
		return err
	}
	if privateKey, err = ReadFile(privateKeyFile); err != nil {
		return err
	}
	return
}

// RSA加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	// 解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes) // 目前的数字证书一般都是基于ITU（国际电信联盟）制定的X.509标准
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	// 加密
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, pub, origData, nil)
}

func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	privInf, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	priv := privInf.(*rsa.PrivateKey)
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, priv, ciphertext, nil)
}
