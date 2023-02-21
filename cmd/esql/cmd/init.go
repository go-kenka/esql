/*
Copyright © 2023 go-kenka <1107015496@qq.com>
*/
package cmd

import (
	"fmt"
	"github.com/go-kenka/esql/gen"
	"github.com/go-kenka/esql/uitls"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化DSL定义模板",
	Long:  `初始化DSL定义模板`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("正在生成代码 %+v\n", args)

		// 代码生成路径
		targetPath, _ := cmd.Flags().GetString("target")

		for _, arg := range args {
			err := gen.Init(targetPath, arg)
			if err != nil {
				panic(err)
			}
		}

		fmt.Println("代码生成完成")
		fmt.Println("正在使用gofmt格式化代码")
		err := uitls.GoFmt(targetPath)
		if err != nil {
			fmt.Println("格式化出错了，请安装gofmt,并将bin目录设置到环境变量中")
		}
		fmt.Println("格式化完成")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	initCmd.Flags().StringP("target", "t", ".", "生成代码文件路径")
}
