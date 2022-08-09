package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	user                    string
	userScore               int
	quizzersLimitPerSession = 5

	spanishquizCmd = &cobra.Command{
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

			for i := 0; i < quizzersLimitPerSession; i++ {

				userScore = StartGame()
				fmt.Println(CompareUserPerformance(userScore, GameScores.scores))
				GameScores.scores = append(GameScores.scores, userScore)
				fmt.Println("Desde Game:", GameScores.scores)

			}

		},
	}
)

func init() {
	rootCmd.AddCommand(spanishquizCmd)
}
