package cli

import (
	"fmt"
	"os"

	hcmd "github.com/Herve-M/happiness-bot/cmd"
	"github.com/Herve-M/happiness-bot/core"

	"github.com/spf13/cobra"
)

var rootFlag = &core.Flags{}
var rootCmd = &cobra.Command{
	Use:   "happiness-bot",
	Short: "Slack happiness-bot for fun.",
	Long:  "",
  Run: func(cmd *cobra.Command, args []string) {
	},
}

var initCmd = &cobra.Command{
	Use: "init",
	Short: "initialize happiness-bot",
	Long:  "",
  Run: func(cmd *cobra.Command, args []string) {
		hcmd.InitCmd()
	},
}

var runCmd = &cobra.Command{
	Use: "run",
	Short: "run happiness-bot",
	Long:  "",
  Run: func(cmd *cobra.Command, args []string) {
		hcmd.RunCmd()
	},
}

var testFlag = &core.SlackTestFlags{}
var testCmd = &cobra.Command{
	Use: "test",
	Short: "test happiness-bot",
	Long:  "",
  Run: func(cmd *cobra.Command, args []string) {
		hcmd.TestCmd(testFlag)
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&rootFlag.Verbose, "verbose", "v", false, "verbose output")
	testCmd.PersistentFlags().StringVar(&testFlag.Message, "msg", "Hi\n, I'am happiness-bot, a slave of Matthieu and Hervé with only one goal: to bring you happiness every morning!", "Message")
	testCmd.PersistentFlags().StringVar(&testFlag.BotName, "name", "happiness-bot", "Bot Name")
	testCmd.PersistentFlags().StringVar(&testFlag.BotIcon, "icon", ":gift_heart:", "Bot Icon")
	rootCmd.AddCommand(initCmd, runCmd, testCmd)
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error trying to execute default command: %v\n", err)
		os.Exit(1)
	}
}
