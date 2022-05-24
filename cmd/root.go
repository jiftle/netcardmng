/*
Copyright © 2022 jiftle

*/
package cmd

import (
	"fmt"
	"os"

	"netcardmng/netcard"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "netcardmng",
	Short: "netcard info read and write",
	Long:  `网卡信息读写`,
	Run: func(cmd *cobra.Command, args []string) {
		byt, err := netcard.GetInfo()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(byt))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
