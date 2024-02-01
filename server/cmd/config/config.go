package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:     "config",
	Short:   "可以生成默认配置",
	Example: "exe文件名.exe config",
	Long:    `加上这个快代码，后面就可以不用提交配置，直接运行这个命令就可以进行生成默认的配置`,
	Run:     configCmdRun,
}

func init() {
	fmt.Println("server/cmd/config/config.go 初始化")
	// configYml := ""
	// ConfigCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
}

func configCmdRun(cmd *cobra.Command, args []string) {
	fmt.Println("这里可以弄一个生成默认配置")
}
