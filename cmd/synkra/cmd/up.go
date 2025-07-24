package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/riceriley59/synkra/engine"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Apply the infrastructure changes",
	Long:  "Apply the infrastructure changes by running the SDK and making calls to the engine.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("you must provide a Python file to run, e.g. synkra up myinfra.py")
		}
		userFile := args[0]

		// Start the engine as a subprocess
		engineProc := exec.Command(os.Args[0], "engine")
		engineProc.Stdout = os.Stdout
		engineProc.Stderr = os.Stderr
		if err := engineProc.Start(); err != nil {
			return fmt.Errorf("failed to start engine: %w", err)
		}
		defer func() {
			_ = engineProc.Process.Kill()
			_ = engineProc.Wait()
		}()

		// Pass the engine address to the SDK
		env := os.Environ()
		env = append(env, "SYNKRA_ENGINE_ADDR=localhost:50051")

		pythonPath, err := exec.LookPath("python3")
		if err != nil {
			return fmt.Errorf("python3 not found: %w", err)
		}
		proc := exec.Command(pythonPath, "sdk/python/synkra.py", userFile)
		proc.Stdout = os.Stdout
		proc.Stderr = os.Stderr
		proc.Stdin = os.Stdin
		proc.Env = env
		if err := proc.Run(); err != nil {
			return fmt.Errorf("failed to run SDK: %w", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
} 