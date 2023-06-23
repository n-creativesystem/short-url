package filekms

import (
	"bytes"
	"context"
	"errors"
	"os"
	"testing"

	"github.com/google/tink/go/aead"
	"github.com/google/tink/go/core/registry"
	"github.com/google/tink/go/keyset"
	"github.com/google/tink/go/subtle/random"
	"github.com/google/tink/go/tink"
	"github.com/stretchr/testify/require"
)

func setupKMS(t *testing.T, ctx context.Context) string {
	t.Helper()
	require := require.New(t)
	if v := os.Getenv("CRYPTO_URI"); v != "" {
		c, err := NewClient(v)
		require.NoError(err)
		registry.ClearKMSClients()
		registry.RegisterKMSClient(c)
		return v
	} else {
		t.Fatal("KEYSET_FILE is empty")
	}
	return ""
}

func TestBasicAead(t *testing.T) {
	require := require.New(t)
	keyURI := setupKMS(t, context.Background())
	dek := aead.AES256GCMKeyTemplate()
	kh, err := keyset.NewHandle(aead.KMSEnvelopeAEADKeyTemplate(keyURI, dek))
	require.NoError(err)
	a, err := aead.New(kh)
	require.NoError(err)
	err = basicAEADTest(t, a)
	require.NoError(err)
}

func TestBasicAeadWithoutAdditionalData(t *testing.T) {
	require := require.New(t)
	keyURI := setupKMS(t, context.Background())
	dek := aead.AES128CTRHMACSHA256KeyTemplate()
	kh, err := keyset.NewHandle(aead.KMSEnvelopeAEADKeyTemplate(keyURI, dek))
	require.NoError(err)
	a, err := aead.New(kh)
	require.NoError(err)
	// Only test 10 times (instead of 100) because each test makes HTTP requests to AWS.
	err = basicAEADTestWithOptions(t, a, 10 /*loopCount*/, false /*withAdditionalData*/)
	require.NoError(err)
}

func basicAEADTest(t *testing.T, a tink.AEAD) error {
	t.Helper()
	return basicAEADTestWithOptions(t, a, 100 /*loopCount*/, true /*withAdditionalData*/)
}

func basicAEADTestWithOptions(t *testing.T, a tink.AEAD, loopCount int, withAdditionalData bool) error {
	t.Helper()
	for i := 0; i < loopCount; i++ {
		pt := random.GetRandomBytes(20)
		var ad []byte = nil
		if withAdditionalData {
			ad = random.GetRandomBytes(20)
		}
		ct, err := a.Encrypt(pt, ad)
		if err != nil {
			return err
		}
		dt, err := a.Decrypt(ct, ad)
		if err != nil {
			return err
		}
		if !bytes.Equal(dt, pt) {
			return errors.New("decrypt not inverse of encrypt")
		}
	}
	return nil
}
