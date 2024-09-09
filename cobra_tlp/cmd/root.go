package cmd

import (
	"github.com/spf13/cobra"
)

// 新命令通过这里注入
func init() {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
	rootCmd.AddCommand(jsonCmd)
	rootCmd.AddCommand(sqlCmd)
}

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

func Execute() error {
	return rootCmd.Execute()
}
