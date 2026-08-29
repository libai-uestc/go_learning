package encryption_test

import (
	"fmt"
	"libai/go/basic/phase-one/encryption"
	"testing"
)

func TestFileEncryption(t *testing.T) {
	keyAES := []byte("0123456789ABCDEF")
	plainFile := "../data/verse.txt"

	encryptFileAES := "../data/verse.aes"
	plainFileAES := "../data/verse(aes已解密).txt"
	if err := encryption.FileEncryption(plainFile, encryptFileAES, encryption.AES, keyAES); err != nil {
		fmt.Println(err)
	} else {
		if err = encryption.FileDecryption(encryptFileAES, plainFileAES, encryption.AES, keyAES); err != nil {
			fmt.Println(err)
		}
	}

	keyDES := []byte("01234567")
	encryptFileDES := "../data/verse.des"
	plainFileDES := "../data/verse(des已解密).txt"
	if err := encryption.FileEncryption(plainFile, encryptFileDES, encryption.DES, keyDES); err != nil {
		fmt.Println(err)
	} else {
		if err = encryption.FileDecryption(encryptFileDES, plainFileDES, encryption.DES, keyDES); err != nil {
			fmt.Println(err)
		}
	}

}

//  go test -v .\phase-one\encryption\ -run=^TestFileEncryption$ -count=1
