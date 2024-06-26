package server

import (
	"fmt"
	"xyserver/route"

	"github.com/gin-gonic/gin"
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
	mGin := gin.Default()

	route.InitRoute(mGin)

	mGin.Run(":8080")
}
