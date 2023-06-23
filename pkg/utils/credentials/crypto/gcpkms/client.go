package gcpkms

import (
	"context"

	"github.com/google/tink/go/core/registry"
	"github.com/google/tink/go/integration/gcpkms"
	"google.golang.org/api/option"
)

func NewClient(uri string, opts ...option.ClientOption) (registry.KMSClient, error) {
	gcpClient, err := gcpkms.NewClientWithOptions(context.Background(), uri, opts...)
	if err != nil {
		return nil, err
	}
	return gcpClient, nil
}
