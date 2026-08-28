package encryption

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
)

func DesEncrypt(text string, key [8]byte) (string, error) {
	src := []byte(text)
	block, err := des.NewCipher(key[:])
	if err != nil {
		return "", nil
	}
	blockSize := block.BlockSize()
	src = PKCS7.Padding(src, blockSize)
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:blockSize])
		src = src[blockSize:]
		dst = dst[blockSize:]
	}
	return hex.EncodeToString(out), nil
}

func DesDecrypt(text string, key [8]byte) (string, error) {
	src, err := hex.DecodeString(text)
	if err != nil {
		return "", err
	}
	block, err := des.NewCipher(key[:])
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Decrypt(dst, src[:blockSize])
		src = src[blockSize:]
		dst = dst[blockSize:]
	}
	out, _ = PKCS7.Unpadding(out, blockSize)
	return string(out), nil
}

func DesEncryptCBC(text string, key [8]byte) (string, error) {
	src := []byte(text)
	block, err := des.NewCipher(key[:])
	if err != nil {
		return "", nil
	}
	blockSize := block.BlockSize()
	src = PKCS7.Padding(src, blockSize)

	out := make([]byte, len(src))
	encrypter := cipher.NewCBCEncrypter(block, key[:])
	encrypter.CryptBlocks(out, src)
	return hex.EncodeToString(out), nil
}

func DesDecryptCBC(text string, key [8]byte) (string, error) {
	src, err := hex.DecodeString(text)
	if err != nil {
		return "", nil
	}
	block, err := des.NewCipher(key[:])
	if err != nil {
		return "", nil
	}

	out := make([]byte, len(src))
	encrypter := cipher.NewCBCDecrypter(block, key[:])
	encrypter.CryptBlocks(out, src)
	out, _ = PKCS7.Unpadding(out, block.BlockSize())
	return string(out), nil
}
