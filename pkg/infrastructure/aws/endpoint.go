package aws

import (
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

type Endpoint struct {
	endpoint map[string]aws.Endpoint
	mu       sync.RWMutex
}

var (
	_ aws.EndpointResolverWithOptions = (*Endpoint)(nil)
)

func NewEndpoint() *Endpoint {
	e := &Endpoint{
		endpoint: make(map[string]aws.Endpoint),
		mu:       sync.RWMutex{},
	}
	e.KMSEndpoint("")
	return e
}

func (e *Endpoint) AddEndpoint(service string, endpoint aws.Endpoint) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.endpoint[service] = endpoint
}

func (e *Endpoint) ResolveEndpoint(service, region string, options ...interface{}) (aws.Endpoint, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	v, ok := e.endpoint[service]
	if ok {
		return v, nil
	}
	return aws.Endpoint{}, &aws.EndpointNotFoundError{}
}

func (e *Endpoint) EndpointResolver() config.LoadOptionsFunc {
	return config.WithEndpointResolverWithOptions(e)
}

func (e *Endpoint) KMSEndpoint(endpoint string) {
	if endpoint == "" {
		endpoint = os.Getenv("KMS_ENDPOINT")
	}
	if endpoint == "" {
		// awsのデフォルトを使うために何もしない
		return
	}
	e.AddEndpoint(kms.ServiceID, aws.Endpoint{URL: endpoint})
}
