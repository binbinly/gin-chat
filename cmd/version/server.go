package version

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gin-chat",
	Long:  `All software has versions. This is gin-chat`,
	Run: func(cmd *cobra.Command, args []string) {
		output, err := ExecuteCommand("git", "describe", "--tags")
		if err != nil {
			fmt.Fprintf(os.Stderr, "execute %s args:%v error:%v\n", cmd.Name(), args, err)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stdout, "gin-chat version ", output)
		fmt.Fprintln(os.Stdout, "go version ", runtime.Version())
		fmt.Fprintln(os.Stdout, "Compiler ", runtime.Compiler)
		fmt.Fprintln(os.Stdout, "Platform ", fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH))
	},
}

// ExecuteCommand 执行命令
func ExecuteCommand(name string, subName string, args ...string) (string, error) {
	args = append([]string{subName}, args...)

	cmd := exec.Command(name, args...)
	bytes, err := cmd.CombinedOutput()

	return string(bytes), err
}
