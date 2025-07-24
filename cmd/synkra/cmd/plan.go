package cmd

import(
	"fmt"

	"github.com/spf13/cobra"
)

var planCmd = &cobra.Command {
	Use: "plan",
	Short: "",
	Long: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("plan")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(planCmd)
}
