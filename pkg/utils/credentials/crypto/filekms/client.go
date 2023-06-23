package filekms

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/google/tink/go/core/registry"
	"github.com/google/tink/go/tink"
)

const filePrefix = "file:///"

type client struct {
	keyPrefix string
}

func NewClient(uriPrefix string) (registry.KMSClient, error) {
	return newClient(uriPrefix)
}

func newClient(uriPrefix string) (registry.KMSClient, error) {
	if !strings.HasPrefix(strings.ToLower(uriPrefix), filePrefix) {
		return nil, fmt.Errorf("uriPrefix must start with %s, but got %s", filePrefix, uriPrefix)
	}

	return &client{
		keyPrefix: uriPrefix,
	}, nil
}

func (c *client) Supported(keyURI string) bool {
	return strings.HasPrefix(keyURI, c.keyPrefix)
}

func (c *client) GetAEAD(keyURI string) (tink.AEAD, error) {
	if !c.Supported(keyURI) {
		return nil, fmt.Errorf("keyURI must start with prefix %s, but got %s", c.keyPrefix, keyURI)
	}

	uri := strings.TrimPrefix(keyURI, filePrefix)

	return newAEAD(filepath.FromSlash(uri))
}
