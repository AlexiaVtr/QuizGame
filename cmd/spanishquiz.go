package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
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

				fmt.Println(".\n.\n.")

				userPerformance, err := QuestionIteration(&Questions)
				userScore = userPerformance.score
				ShowUserPerformance(&userPerformance)

				fmt.Println("\n" + CompareUserPerformance(userScore, GameScores.scores))
				GameScores.AppendScore(userScore)

				if err != nil {
					fmt.Println(err)
				}

			}

		},
	}
)

func init() {
	rootCmd.AddCommand(spanishquizCmd)
}
