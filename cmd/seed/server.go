package seed

import (
	"log"

	"github.com/spf13/cobra"

	"gin-chat/pkg/mysql"
)

var (
	host     string
	user     string
	password string
	name     string
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
	StartCmd.PersistentFlags().StringVarP(&host, "host", "a", "127.0.0.1", "mysql host")
	StartCmd.PersistentFlags().StringVarP(&user, "user", "u", "root", "mysql user")
	StartCmd.PersistentFlags().StringVarP(&password, "password", "p", "root", "mysql password")
	StartCmd.PersistentFlags().StringVarP(&name, "name", "d", "chat", "mysql db name")
}

func setup() {
	mysql.NewBasicDB(host, user, password, name)
}

func run() {
	if err := SyncBQB(); err != nil {
		log.Fatalf("err:%v\n", err)
	}
	log.Println("sync emoticon success")
}
