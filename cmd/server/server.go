package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"gin-chat/internal/server"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
	"gin-chat/pkg/config"
)

var (
	cfgDir   string
	env      string
	StartCmd = &cobra.Command{
		Use:          "server",
		Short:        "Start gin-chat server",
		Example:      "gin-chat server -c configs",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&cfgDir, "config", "c", "configs", "config path")
	StartCmd.PersistentFlags().StringVarP(&env, "env", "e", "", "Configure Runtime Environment")
}

func setup() {
	// init config
	c := config.New(cfgDir, config.WithEnv(env))
	var cfg app.Config
	if err := c.Load("app", &cfg); err != nil {
		panic(err)
	}
	app.Conf = &cfg

	gin.SetMode(cfg.Mode)
}

// run 核心业务服务启动
func run() {
	ws := server.NewWsServer(&app.Conf.Websocket)
	// init service
	service.Svc = service.New(ws,
		service.WithJwtTimeout(app.Conf.JwtTimeout),
		service.WithJwtSecret(app.Conf.JwtSecret))

	// start app
	apps := app.New(
		app.WithName(app.Conf.Name),
		app.WithServer(
			server.NewHTTPServer(&app.Conf.HTTP),
			ws,
		),
	)

	if err := apps.Run(); err != nil {
		panic(err)
	}
}
