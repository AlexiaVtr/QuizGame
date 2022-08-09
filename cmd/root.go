package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "go_quiz",
		Short: "Question and answer game.",
		Long: `
	/////////////////////////////////////////////////////////////////

	This is a question and answer game.
	Use the comand "spanishquiz" to start to play the game.
	
	/////////////////////////////////////////////////////////////////
	`,
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
