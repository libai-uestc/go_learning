package encryption_test

import (
	"fmt"
	"libai/go/basic/phase-one/encryption"
	"testing"
	"time"
)

var (
	BigFile = "C:\\Users\\18101\\Downloads\\go1.26.7.windows-amd64.msi"
)

func TestHash(t *testing.T) {
	data := "123456"
	hs := encryption.Sha1(data)
	fmt.Println("SHA-1", hs, len(hs))
	hm := encryption.Md5(data)
	fmt.Println("MD5", hm, len(hm))
}

func TestCreateSha256OfSmallFile(t *testing.T) {
	hash, err := encryption.CreateSha256OfSmallFile(BigFile)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("CreateSha256OfSmallFile", hash)
}

func TestCreateSha256OfBigFile(t *testing.T) {
	begin := time.Now()
	hash, err := encryption.CreateSha256OfBigFile(BigFile, 10<<20)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("CreateSha256OfBigFile", hash, "use time", time.Since(begin).Milliseconds())

}

// go test -v .\phase-one\encryption\ -run=^TestHash$ -count=1
