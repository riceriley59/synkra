package cli_tests

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"

	"github.com/riceriley59/synkra/internal/cli"
)

func TestCLISmoke(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CLI Smoke Test Suite")
}

var _ = Describe("Synkra CLI Smoke Tests", func() {
	var synkraCmd *cobra.Command

	BeforeEach(func() {
		synkraCmd = cli.NewSynkraCmd()
	})

	Describe("Command Registration", func() {
		It("Should have all commands registered", func() {
			commands := synkraCmd.Commands()
			commandNames := make([]string, len(commands))
			for i, cmd := range(commands) {
				commandNames[i] = cmd.Use
			}

			Expect(commandNames).To(ContainElement("infra"))
		})
	})
})
