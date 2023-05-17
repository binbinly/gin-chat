package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"gin-chat/cmd/migrate"
	"gin-chat/cmd/seed"
	"gin-chat/cmd/server"
	"gin-chat/cmd/version"
)

var rootCmd = &cobra.Command{
	Use:   "gin-chat",
	Short: "gin-chat app",
	Long:  `即时通讯学习项目`,
	Args:  args,
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatalf("unrecognized command cmd: %v args: %v", cmd.Name(), args)
	},
}

func args(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("please select option")
	}
	return nil
}

// Execute 命令执行入口
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(server.StartCmd)
	rootCmd.AddCommand(migrate.StartCmd)
	rootCmd.AddCommand(version.StartCmd)
	rootCmd.AddCommand(seed.StartCmd)
}
