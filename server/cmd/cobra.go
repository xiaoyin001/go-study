package cmd

import (
	"fmt"
	"go-study/server/cmd/config"
	"go-study/server/cmd/server"

	"github.com/spf13/cobra"
)

// Go命令行工具 Cobra
// https://www.cnblogs.com/haiyux/p/17510318.html#command%E5%8F%82%E6%95%B0
// https://www.cnblogs.com/yrxing/p/14981190.html

// 入口Cmd【注意根Cmd不是可用命令】
var uRootCmd = &cobra.Command{
	Use:   "go-xy-study",              // 命令名称
	Short: "XiaoYin Study Cobra",      // 命令简短描述
	Long:  `这是准备学习如何弄一个命令程序，这是梦开始的地方`, // 命令详细描述
	Run:   rootCmdRun,                 // 命令执行入口
}

func init() {
	uRootCmd.AddCommand(config.ConfigCmd)
	uRootCmd.AddCommand(server.ServerCmd)
}

func rootCmdRun(cmd *cobra.Command, args []string) {
	fmt.Println("梦开始的地方", args)
}

func AppExecute() {
	mErr := uRootCmd.Execute()
	if mErr != nil {
		fmt.Println("Err:", mErr.Error())
	}
}
