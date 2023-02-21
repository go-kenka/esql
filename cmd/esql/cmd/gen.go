/*
Copyright © 2023 go-kenka <1107015496@qq.com>
*/
package cmd

import (
	"fmt"
	"github.com/go-kenka/esql/dsl/ast"
	"github.com/go-kenka/esql/gen"
	"github.com/go-kenka/esql/uitls"

	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "通过DSL生成CRUD代码模型",
	Long:  `通过DSL生成CRUD代码模型`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("正在生成代码 %+v\n", args)

		// ast 路径
		schemaPath := args[0]
		// 代码生成路径
		targetPath, _ := cmd.Flags().GetString("target")
		// 读取schema定义
		tbs := ast.ReadDir(schemaPath)
		// 获取当前项目path路径
		pkg := uitls.PkgPath(targetPath)
		// 开始生成代码
		err := gen.GenClient(targetPath, pkg, tbs)
		if err != nil {
			panic(err)
		}

		fmt.Println("代码生成完成")
		fmt.Println("正在使用gofmt格式化代码")
		err = uitls.GoFmt(targetPath)
		if err != nil {
			fmt.Println("格式化出错了，请安装gofmt,并将bin目录设置到环境变量中")
		}
		fmt.Println("格式化完成")
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	genCmd.Flags().StringP("target", "t", ".", "生成代码文件路径")
}
