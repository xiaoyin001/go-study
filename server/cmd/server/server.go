package server

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:     "server",
	Short:   "启动服务",
	Example: "exe文件名.exe server",
	Long:    `通过这里启动服务`,
	Run:     severCmdRun,
}

func init() {
	fmt.Println("server/cmd/server/server.go 初始化")
	// configYml := ""
	// ConfigCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
}

func severCmdRun(cmd *cobra.Command, args []string) {
	fmt.Println("这里就是服务的入口，假装起了一个服务")
}
