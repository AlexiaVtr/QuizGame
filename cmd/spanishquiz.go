/*
Alexia Vetere
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var spanishquizCmd = &cobra.Command{
	Use:   "spanishquiz",
	Short: "Spanish quiz.",
	Long: `
	///////////////////////////////////////////////////////////////////////////////	
	
	This application is a quiz to measure your knowledge in Spanish. The rules are:
	
	* There are 5 questions.
	* You have 3 options for each question.
	* Each correct question adds 1 point.
	* At the end of the quiz you can compare your result with that of others.
	
	Enjoy and good luck!
	
	///////////////////////////////////////////////////////////////////////////////
	`,
	Run: func(cmd *cobra.Command, args []string) {

		userScore := startGame()
		compareUserPerformance(userScore)

	},
}

func init() {
	rootCmd.AddCommand(spanishquizCmd)
}
