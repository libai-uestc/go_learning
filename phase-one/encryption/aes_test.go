package encryption_test

import (
	"fmt"
	"libai/go/basic/phase-one/encryption"
	"log"
	"testing"
)

func TestAES(t *testing.T) {
	key := [16]byte{'l', 'i', 'b', 'a', 'i', 'u', 'e', 's', 't', 'c', '0', 'x', '1', '7', '1', '6'}
	plain := "这城市那么空，这回忆那么凶"
	cipher, err := encryption.AesEncrypt(plain, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("密文：%s\n", cipher)

	plain, err = encryption.AesDecrypt(cipher, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("明文: %s\n", plain)
}

// go test -v .\phase-one\encryption\ -run=^TestAES$ -count=1
