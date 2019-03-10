package cmd

import "github.com/spf13/cobra"

var startCmd = &cobra.Command{
	Use:  "start",
	RunE: startCmdEntrypointE,
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func startCmdEntrypointE(cmd *cobra.Command, args []string) error {

	return nil
}
