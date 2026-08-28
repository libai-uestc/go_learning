package encryption_test

import (
	"fmt"
	"libai/go/basic/phase-one/encryption"
	"log"
	"testing"
)

func TestDES(t *testing.T) {
	key := [8]byte{34, 65, 12, 125, 65, 70, 54, 27}
	plain := "这街道车水马龙，我能和谁相拥？"
	cipher, err := encryption.DesEncrypt(plain, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("密文：%s\n", cipher)

	plain, err = encryption.DesDecrypt(cipher, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("明文: %s\n", plain)
	fmt.Println("--------------------------")

	cipher, _ = encryption.DesEncryptCBC(plain, key)
	fmt.Printf("密文：%s\n", cipher)
	plain, _ = encryption.DesDecryptCBC(cipher, key)
	fmt.Printf("明文: %s\n", plain)
}
