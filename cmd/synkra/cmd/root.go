package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
	Use: "synkra",
	Short: "synkra manages infra provisioning & configuration management",
	Long:  "synkra manages infra provisioning & configuration management", // TODO: Write long description
	Run: func (cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World!")
	},
}

func init() {
	// TODO: define flags and configuration
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
