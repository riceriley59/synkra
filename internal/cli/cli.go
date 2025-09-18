package cli

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/riceriley59/synkra/internal/version"
	"github.com/riceriley59/synkra/internal/cli/infra"
)

func Execute() {
	synkraCmd := NewSynkraCmd()
	err := synkraCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func NewSynkraCmd() *cobra.Command {
	synkraCmd := &cobra.Command{
		Use:   "synkra",
		Short: "",
		Long:  ``,
		Version: version.Version,
	}

	synkraCmd.AddCommand(infra.NewInfraCmd())

	return synkraCmd
}
