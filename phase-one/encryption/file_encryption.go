package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
	"io"
	"os"
)

const (
	_ = iota
	DES
	AES
)

// 文件加密
func FileEncryption(infile string, outfile string, algo int, key []byte) error {
	fin, err := os.Open(infile)
	if err != nil {
		return err
	}
	defer fin.Close()
	fout, err := os.OpenFile(outfile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fout.Close()

	content, err := io.ReadAll(fin)
	if err != nil {
		return err
	}
	var block cipher.Block
	switch algo {
	case AES:
		block, err = aes.NewCipher(key)
	case DES:
		block, err = des.NewCipher(key)
	default:
		return fmt.Errorf("unsurported encrypt algo %d", algo)
	}
	if err != nil {
		return err
	}
	encrypter := cipher.NewCBCEncrypter(block, key)
	src := PKCS7.Padding(content, block.BlockSize())
	dest := make([]byte, len(src))
	encrypter.CryptBlocks(dest, src)
	fout.Write(dest)
	return nil
}

// 文件解密
func FileDecryption(infile string, outfile string, algo int, key []byte) error {
	fin, err := os.Open(infile)
	if err != nil {
		return err
	}
	defer fin.Close()
	fout, err := os.OpenFile(outfile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fout.Close()
	content, err := io.ReadAll(fin)
	if err != nil {
		return err
	}

	var block cipher.Block
	switch algo {
	case AES:
		block, err = aes.NewCipher(key)
	case DES:
		block, err = des.NewCipher(key)
	default:
		return fmt.Errorf("unsurported encrypt algo %d", algo)
	}
	if err != nil {
		return err
	}
	decrypter := cipher.NewCBCDecrypter(block, key)

	decrypted := make([]byte, len(content))
	decrypter.CryptBlocks(decrypted, content)
	out, err := PKCS7.Unpadding(decrypted, block.BlockSize())
	if err != nil {
		return err
	}
	fout.Write(out)
	return nil
}
