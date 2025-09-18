package infra

import (
	"github.com/riceriley59/synkra/internal/cli/utils"
	"github.com/spf13/cobra"
)

func NewInfraCmd() *cobra.Command {
	newInfraCmd := &cobra.Command{
		Use:   "infra",
		Short: "",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := utils.ValidateNonEmptyArgs(args, "infrastructure command"); err != nil {
				return err
			}

			utils.LogPlaceholder("Infrastructure Command", cmd, args)
			return nil
		},
	}

	return newInfraCmd
}
