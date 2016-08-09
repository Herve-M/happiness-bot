package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "happiness-bot",
	Short: "Slack happiness-bot for fun.",
	Long:  "",
  Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error trying to execute default command: %v\n", err)
		os.Exit(1)
	}
}
