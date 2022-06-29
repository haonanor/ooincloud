package cmd

import (
	"errors"
	"fmt"
	"os"
	"study/cmd/api"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "RobotCoder",
	Short:        "RobotCoder",
	SilenceUsage: true,
	Long:         `RobotCoder`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
}

func tip() {
	usageStr := `欢迎使用   查看命令`
	usageStr1 := `也可以参考 https://doc.go-admin.dev/guide/ksks.html 里边的【启动】章节`
	fmt.Printf("%s\n", usageStr)
	fmt.Printf("%s\n", usageStr1)
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
