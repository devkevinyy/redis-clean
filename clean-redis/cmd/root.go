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
  "github.com/chujieyang/redis-clean/clean-redis/utils"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
  "os"
)

var ConnectionString string
var Auth string
var Pattern string
var Db int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "clean-redis",
  Short: "redis数据清理工具",
  Long: `福禄网络Redis实例数据清理工具，可支持无阻塞对 BigKey 进行清理，具体设置见执行参数。`,
  Args: cobra.OnlyValidArgs,
  // Uncomment the following line if your bare application
  // has an action associated with it:
  Run: func(cmd *cobra.Command, args []string) {
  	utils.InitRedisPool(ConnectionString, Db, Auth)
    if err := utils.RemoveRedisKeys(Pattern); err != nil {
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
  rootCmd.Flags().StringVarP(&ConnectionString, "ConnectionString", "c", "", "redis实例连接地址")
  rootCmd.MarkFlagRequired("ConnectionString")
  viper.BindPFlag("ConnectionString", rootCmd.Flags().Lookup("ConnectionString"))

  rootCmd.Flags().StringVarP(&Auth, "Auth", "a", "", "用户名和密码")
  viper.BindPFlag("Auth", rootCmd.Flags().Lookup("Auth"))

  rootCmd.Flags().StringVarP(&Pattern, "Pattern", "p", "", "匹配字符串")
  rootCmd.MarkFlagRequired("Pattern")
  viper.BindPFlag("Pattern", rootCmd.Flags().Lookup("Pattern"))

  rootCmd.Flags().IntVarP(&Db, "Db", "d", 0, "指定DB")
  rootCmd.MarkFlagRequired("Db")
  viper.BindPFlag("Db", rootCmd.Flags().Lookup("Db"))
}

