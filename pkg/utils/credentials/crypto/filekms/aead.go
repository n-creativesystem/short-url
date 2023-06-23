package filekms

import (
	"os"

	"github.com/google/tink/go/aead"
	"github.com/google/tink/go/keyset"
	"github.com/google/tink/go/tink"
)

func readFile(file string) (*keyset.Handle, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	bin := keyset.NewBinaryReader(f)
	return keyset.Read(bin, NewMasterKey())
}

func newAEAD(file string) (tink.AEAD, error) {
	kh, err := readFile(file)
	if err != nil {
		return nil, err
	}
	return aead.New(kh)
}
