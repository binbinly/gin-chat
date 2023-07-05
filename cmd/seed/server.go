package seed

import (
	"log"

	"github.com/spf13/cobra"

	"gin-chat/pkg/dbs"
)

var (
	dsn      string
	driver   string
	StartCmd = &cobra.Command{
		Use:          "seed",
		Short:        "Seed data",
		Example:      "chat-micro seed -c config/logic/default.yml",
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
	StartCmd.PersistentFlags().StringVarP(&dsn, "dsn", "d", "root:root@127.0.0.1:3306/chat", "dbs dsn data source name")
	StartCmd.PersistentFlags().StringVarP(&driver, "driver", "t", "mysql", "db driver")
}

func setup() {
	dbs.NewBasicDB(driver, dsn)
}

func run() {
	if err := SyncBQB(); err != nil {
		log.Fatalf("err:%v\n", err)
	}
	log.Println("sync emoticon success")
}
