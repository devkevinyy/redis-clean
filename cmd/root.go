/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/chujieyang/redis-clean/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ConnectionString string
var Auth string
var Pattern string
var Db int
var Count int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "redis-clean",
	Short: "redis数据清理工具",
	Long:  `Redis实例数据清理工具，可支持无阻塞对 BigKey 进行清理，具体设置见执行参数。`,
	Args:  cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {
		utils.InitRedisPool(ConnectionString, Db, Auth)
		if err := utils.RemoveRedisKeys(Pattern, Count); err != nil {
			fmt.Println(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&ConnectionString, "host", "", "", "redis实例连接地址")
	rootCmd.MarkFlagRequired("host")
	viper.BindPFlag("host", rootCmd.Flags().Lookup("host"))

	rootCmd.Flags().StringVarP(&Auth, "auth", "", "", "用户名和密码")
	viper.BindPFlag("auth", rootCmd.Flags().Lookup("auth"))

	rootCmd.Flags().StringVarP(&Pattern, "pattern", "", "", "Key匹配模式")
	rootCmd.MarkFlagRequired("pattern")
	viper.BindPFlag("pattern", rootCmd.Flags().Lookup("pattern"))

	rootCmd.Flags().IntVarP(&Db, "db", "", 0, "指定DB")
	rootCmd.MarkFlagRequired("db")
	viper.BindPFlag("db", rootCmd.Flags().Lookup("db"))

	rootCmd.Flags().IntVarP(&Count, "count", "", 100, "scan数量")
	rootCmd.MarkFlagRequired("count")
	viper.BindPFlag("count", rootCmd.Flags().Lookup("count"))
}
