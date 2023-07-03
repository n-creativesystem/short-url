package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/n-creativesystem/short-url/pkg/cmd/flags"
	"github.com/n-creativesystem/short-url/pkg/domain/config"
	config_infra "github.com/n-creativesystem/short-url/pkg/infrastructure/config"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/session"
	"github.com/n-creativesystem/short-url/pkg/interfaces/router"
	"github.com/n-creativesystem/short-url/pkg/utils"
	"github.com/n-creativesystem/short-url/pkg/utils/credentials/crypto"
	"github.com/n-creativesystem/short-url/pkg/utils/logging"
	"github.com/sethvargo/go-envconfig"
	"github.com/spf13/cobra"
)

func serverCommand() *cobra.Command {
	var (
		driver string
	)
	f := flags.New(gin.ReleaseMode, gin.DebugMode, gin.TestMode)
	cmd := cobra.Command{
		Use:           "server",
		SilenceUsage:  true,
		SilenceErrors: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := crypto.Init(); err != nil {
				return err
			}
			cmd.Root().PersistentPreRun(cmd, args)
			gin.SetMode(f.String())
			config.SetDriver(config.ConvertDriverFromString(driver))
			return nil
		},
	}
	pflags := cmd.PersistentFlags()
	pflags.Var(f, "mode", "handler mode")
	pflags.Int("port", 8080, "port number")
	pflags.StringVar(&driver, "driver", "mysql", "driver name")

	cmd.AddCommand(apiModeCommand())
	cmd.AddCommand(serviceModeCommand())
	cmd.AddCommand(webUIModeCommand())
	return &cmd
}

type serverMode int

const (
	api serverMode = iota
	urlService
	webUI
)

func apiModeCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:           "api",
		SilenceUsage:  true,
		SilenceErrors: true,
		Run: func(cmd *cobra.Command, args []string) {
			pFlags := cmd.Flags()
			port, err := pFlags.GetInt("port")
			if err != nil {
				logging.Default().Error(err)
				return
			}
			utils.RunAPI()
			executeServer(cmd.Context(), port, api)
		},
	}
	return &cmd
}

func webUIModeCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:           "web-ui",
		SilenceUsage:  true,
		SilenceErrors: true,
		Run: func(cmd *cobra.Command, args []string) {
			pFlags := cmd.Flags()
			port, err := pFlags.GetInt("port")
			if err != nil {
				logging.Default().Error(err)
				return
			}
			utils.RunUI()
			executeServer(cmd.Context(), port, webUI)
		},
	}
	return &cmd
}

func serviceModeCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:           "service",
		SilenceUsage:  true,
		SilenceErrors: true,
		Run: func(cmd *cobra.Command, args []string) {
			pFlags := cmd.Flags()
			port, err := pFlags.GetInt("port")
			if err != nil {
				logging.Default().Error(err)
				return
			}
			utils.RunService()
			executeServer(cmd.Context(), port, urlService)
		},
	}
	return &cmd
}

func executeServer(ctx context.Context, port int, mode serverMode) {
	var stop context.CancelFunc
	ctx, stop = signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGKILL)
	defer stop()
	appConfigRepo := config_infra.NewApplication()
	appConfig, err := appConfigRepo.Get(ctx, config.WithEnvConfigLookuper(envconfig.OsLookuper()))
	if err != nil {
		logging.Default().Error(err)
		return
	}
	input, closer, err := getInput(ctx, appConfig)
	if err != nil {
		logging.Default().Error(err)
		return
	}
	defer closer()
	var handler http.Handler
	switch mode {
	case api:
		handler = router.NewAPI(input)
	case urlService:
		handler = router.NewMainService(input)
	case webUI:
		sessionCfg := config_infra.NewSession()
		repo := config_infra.NewSocialConfig()
		mpConfig, err := repo.GetProviders(ctx)
		if err != nil {
			logging.Default().Error(err)
			return
		}
		sessionStore, err := session.New(ctx, sessionCfg)
		if err != nil {
			logging.Default().Error(err)
			return
		}
		cfg := config_infra.NewWebUI(ctx)
		cfg.Store = sessionStore
		cfg.Providers = mpConfig
		handler = router.NewWebUI(input, cfg)
	default:
		logging.Default().Error("An unexpected server mode is specified.")
		return
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logging.Default().Errorf("listen: %s", err)
		}
	}()
	logging.Default().Infof("Start server :%d", port)
	<-ctx.Done()
	logging.Default().Info("Shutdown server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logging.Default().Errorf("Server shutdown error: %s", err)
	}
	logging.Default().Info("Server exiting")
}

func getInput(ctx context.Context, appConfig *config.Application) (*router.RouterInput, func(), error) {
	var (
		input  *router.RouterInput
		closer func()
		err    error
	)
	switch config.GetDriver() {
	case config.MySQL, config.PostgreSQL, config.SQLite:
		input, closer, err = getRDBInput(ctx)
	default:
		panic(config.ErrNoSupportDriver())
	}
	if err != nil {
		return nil, nil, err
	}
	input.AppConfig = appConfig
	return input, closer, err
}
