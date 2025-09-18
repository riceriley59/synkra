package utils

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func LogPlaceholder(actionName string, cmd *cobra.Command, args []string) {
	log.Printf("%s command called", actionName)

	if len(args) > 0 {
		log.Printf("Arguments: %v", args)
	}

	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if flag.Changed {
			log.Printf("%s: %s", flag.Name, flag.Value.String())
		}
	})

	log.Printf("%s not yet implemented", actionName)
}

func ValidateNonEmptyArgs(args []string, argType string) error {
	for i, arg := range args {
		if strings.TrimSpace(arg) == "" {
			if len(args) == 1 {
				return fmt.Errorf("%s cannot be empty", argType)
			}

			return fmt.Errorf("argument %d (%s) cannot be empty", i + 1, argType)
		}
	}

	return nil
}
