package router

import (
	"testing"

	"github.com/n-creativesystem/short-url/pkg/utils/logging"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHandlers(t *testing.T) {
	logging.SetFormat("json")
	RegisterFailHandler(Fail)
	RunSpecs(t, "Route Suite")
}
