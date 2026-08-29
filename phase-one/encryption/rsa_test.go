package encryption_test

import (
	"fmt"
	"libai/go/basic/phase-one/encryption"
	"testing"
)

func TestRSA(t *testing.T) {
	encryption.ReadRSAKey("../data/rsa_public_key.pem", "../data/rsa_private_key.pem")

	plain := "长风破浪会有时，直挂云帆济沧海"
	cipher, err := encryption.RsaEncrypt([]byte(plain))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("密文: %v\n", cipher)
		bPlain, err := encryption.RsaDecrypt(cipher)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("明文: %s\n", string(bPlain))
		}
	}
}

// go test -v .\encryption\ -run=^TestRSA$ -count=1
