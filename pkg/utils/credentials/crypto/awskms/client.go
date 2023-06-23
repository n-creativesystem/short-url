// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
////////////////////////////////////////////////////////////////////////////////

package awskms

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/google/tink/go/core/registry"
	"github.com/google/tink/go/tink"
	infra_aws "github.com/n-creativesystem/short-url/pkg/infrastructure/aws"
)

const (
	awsPrefix = "aws-kms://"
)

// awsClient represents a client that connects to the AWS KMS backend.
type awsClient struct {
	keyURIPrefix string
	kms          *kms.Client
}

// NewClient returns a new AWS KMS client which will use default
// credentials to handle keys with uriPrefix prefix.
// uriPrefix must have the following format: 'aws-kms://arn:<partition>:kms:<region>:[:path]'.
// See http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html.
func NewClient(uriPrefix string) (registry.KMSClient, error) {
	endpoint := infra_aws.NewEndpoint()
	cfg, err := infra_aws.NewConfig(context.Background(), endpoint.EndpointResolver())
	if err != nil {
		return nil, err
	}

	client := kms.NewFromConfig(cfg)

	return NewClientWithKMS(uriPrefix, client)
}

func NewClientWithKMS(uriPrefix string, kms *kms.Client) (registry.KMSClient, error) {
	if !strings.HasPrefix(strings.ToLower(uriPrefix), awsPrefix) {
		return nil, fmt.Errorf("uriPrefix must start with %s, but got %s", awsPrefix, uriPrefix)
	}

	return &awsClient{
		keyURIPrefix: uriPrefix,
		kms:          kms,
	}, nil
}

// Supported true if this client does support keyURI
func (c *awsClient) Supported(keyURI string) bool {
	return strings.HasPrefix(keyURI, c.keyURIPrefix)
}

// GetAEAD gets an AEAD backend by keyURI.
// keyURI must have the following format: 'aws-kms://arn:<partition>:kms:<region>:[:path]'.
// See http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html.
func (c *awsClient) GetAEAD(keyURI string) (tink.AEAD, error) {
	if !c.Supported(keyURI) {
		return nil, fmt.Errorf("keyURI must start with prefix %s, but got %s", c.keyURIPrefix, keyURI)
	}

	uri := strings.TrimPrefix(keyURI, awsPrefix)
	return newAWSAEAD(uri, c.kms), nil
}
