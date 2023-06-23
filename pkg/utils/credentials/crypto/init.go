package crypto

import (
	"errors"
	"os"

	"github.com/google/tink/go/aead"
	"github.com/google/tink/go/core/registry"
	"github.com/google/tink/go/keyset"
	"github.com/n-creativesystem/short-url/pkg/utils/credentials/crypto/awskms"
	"github.com/n-creativesystem/short-url/pkg/utils/credentials/crypto/filekms"
	"github.com/n-creativesystem/short-url/pkg/utils/credentials/crypto/gcpkms"
)

func Init() error {
	v := os.Getenv("CRYPTO_URI")
	if v == "" {
		return errors.New("CRYPTO_URI is empty")
	}
	if v != "" {
		tpl := aead.AES256GCMKeyTemplate()
		if v, err := filekms.NewClient(v); err == nil {
			registry.RegisterKMSClient(v)
		}
		if v, err := awskms.NewClient(v); err == nil {
			registry.RegisterKMSClient(v)
		}
		if v, err := gcpkms.NewClient(v); err == nil {
			registry.RegisterKMSClient(v)
		}
		handle, err := keyset.NewHandle(aead.KMSEnvelopeAEADKeyTemplate(v, tpl))
		if err != nil {
			return err
		}
		if aeadCipher, err = aead.New(handle); err != nil {
			return err
		}
	}
	return nil
}
