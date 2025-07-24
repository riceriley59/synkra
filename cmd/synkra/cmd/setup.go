package cmd

import(
	"fmt"

	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command {
	Use: "setup",
	Short: "",
	Long: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
