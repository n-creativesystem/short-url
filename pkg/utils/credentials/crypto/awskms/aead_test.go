package awskms

import (
	"bytes"
	"context"
	"errors"
	"net/url"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/google/tink/go/aead"
	"github.com/google/tink/go/core/registry"
	"github.com/google/tink/go/keyset"
	"github.com/google/tink/go/subtle/random"
	"github.com/google/tink/go/tink"
	infra_aws "github.com/n-creativesystem/short-url/pkg/infrastructure/aws"
	"github.com/stretchr/testify/require"
)

func getTestID(ctx context.Context, cfg aws.Config) (string, error) {
	client := kms.NewFromConfig(cfg)
	output, err := client.DescribeKey(ctx, &kms.DescribeKeyInput{
		KeyId: aws.String("alias/local-kms-key"),
	})
	if err != nil {
		return "", err
	}
	u, err := url.Parse("aws-kms://")
	if err != nil {
		return "", err
	}
	u.Path = *output.KeyMetadata.Arn
	return u.String(), nil
}

func setupKMS(t *testing.T, ctx context.Context) string {
	t.Helper()
	require := require.New(t)
	endpoint := infra_aws.NewEndpoint()
	cfg, err := infra_aws.NewConfig(ctx, endpoint.EndpointResolver())
	require.NoError(err)
	arn, err := getTestID(ctx, cfg)
	require.NoError(err)
	c, err := NewClient("aws-kms://")
	require.NoError(err)
	registry.ClearKMSClients()
	registry.RegisterKMSClient(c)
	return arn
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

func TestBasicAead(t *testing.T) {
	require := require.New(t)
	keyURI := setupKMS(t, context.Background())
	dek := aead.AES128CTRHMACSHA256KeyTemplate()
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
	err = basicAEADTestWithOptions(t, a, 10 /*loopCount*/, false /*withAdditionalData*/)
	require.NoError(err)
}
