package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
	Use: "synkra",
	Short: "Synkra manages infra provisioning & configuration management",
	Long:  "Synkra manages infra provisioning & configuration management", // TODO: Write long description
}

func init() {
	// TODO: Define flags and configuration
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
