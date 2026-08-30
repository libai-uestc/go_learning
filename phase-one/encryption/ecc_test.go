package encryption_test

import (
	"fmt"
	"libai/go/basic/phase-one/encryption"
	"testing"
)

func TestECC(t *testing.T) {
	prvKey, err := encryption.GenPrivateKey()
	if err != nil {
		t.Fatalf("genPrivateKey fail: %s\n", err)
	}
	pubKey := prvKey.PublicKey
	plain := "会当凌绝顶，一览众山小"
	cipher, err := encryption.ECCEncrypt(plain, pubKey)
	if err != nil {
		t.Fatalf("ECCEncrypt fail: %s\n", err)
	}
	plain, err = encryption.ECCDecrypt(cipher, prvKey)
	if err != nil {
		t.Fatalf("ECCDecrypt fail: %s\n", err)
	}
	fmt.Printf("明文: %s\n", plain)
}
