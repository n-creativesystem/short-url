package apps

import (
	"sync"

	"github.com/n-creativesystem/short-url/pkg/utils"
)

var (
	serviceName     = ""
	initServiceName sync.Once

	appVersion = "1.0.0"

	serverRoot     = ""
	initServerRoot sync.Once
)

func ServiceName() string {
	initServiceName.Do(func() {
		serviceName = utils.Getenv("SERVICE_NAME", "QUICK_LINK")
	})
	return serviceName
}

func Version() string {
	return appVersion
}

func TrackingEnvironment() string {
	return utils.Getenv("TRACKING_ENV", utils.AppEnv())
}

func ServerRoot() string {
	initServerRoot.Do(func() {
		serverRoot = utils.Getenv("SERVER_ROOT", "github.com/n-creativesystem/short-url")
	})
	return serverRoot
}
