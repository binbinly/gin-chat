package server

import (
	"gin-chat/internal/router"
	"gin-chat/internal/server"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
	"gin-chat/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
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
	// websocket server
	ws := server.NewWsServer(&app.Conf.Websocket)
	// http server
	http := server.NewHTTPServer(&app.Conf.HTTP)

	// init router
	r := router.NewRouter(app.Conf.Debug)

	// set proxy to http://[host]/ws -> ws://[host]
	if app.Conf.Proxy {
		r.Any("/ws", app.ProxyGinHandler("http://127.0.0.1"+app.Conf.Websocket.Addr))
	}
	http.Handler = r

	// init service
	service.Svc = service.New(ws,
		service.WithJwtTimeout(app.Conf.JwtTimeout),
		service.WithJwtSecret(app.Conf.JwtSecret))

	// init app
	apps := app.New(
		app.WithName(app.Conf.Name),
		app.WithServer(http, ws),
	)

	// run app
	if err := apps.Run(); err != nil {
		panic(err)
	}
}
