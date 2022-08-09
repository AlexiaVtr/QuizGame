package cmd

import (
	"errors"
	"fmt"
)

var (
	GameScores GameScore
)

type QuestionsList struct {
	questions []Question
}

type Question struct {
	question string
	options  []QuestionOptions
}

type QuestionOptions struct {
	option string
	value  bool
}

type UserAnswer struct {
	question string
	answer   int
	result   string
}

type UserPerformance struct {
	answers []UserAnswer
	score   int
}

type GameScore struct {
	scores []int
}

func (gs *GameScore) AppendScore(userScore int) {
	gs.scores = append(gs.scores, userScore)
}

func (up *UserPerformance) AppendUserAnswers(ua UserAnswer) {
	up.answers = append(up.answers, ua)
}

func (q *Question) CalculateUserScore(userResponse int) int {
	if q.options[userResponse-1].value == true {
		return 1
	} else {
		return 0
	}
}

func (q *Question) ReturnResult(userResponse int) string {
	if q.options[userResponse-1].value == true {
		return "The answerd was correct."
	} else {
		return "The answerd was wrong."
	}
}

func (q *Question) ToShow(questionNumber int) {

	fmt.Printf("\nQuestion %d: \n%s \n", (questionNumber + 1), q.question)

	for optionNumber, options := range q.options {
		fmt.Printf("%d - %s\n", (optionNumber + 1), options.option)
	}
}

func ControlScanOptions(userResponse int) int {
	for userResponse > 3 || userResponse < 1 {
		fmt.Scan(&userResponse)
		fmt.Println("\nPlease, select a correct option")
	}
	return userResponse
}

func QuestionIteration(q *QuestionsList) (UserPerformance, error) {
	userPoints := 0
	userOption := 0
	var userResponse UserAnswer
	var userPerformance UserPerformance
	var err error

	for questionNumber, question := range q.questions {

		question.ToShow(questionNumber)

		fmt.Scan(&userResponse)

		chosenOption := ControlScanOptions(userOption)
		userPoints += question.CalculateUserScore(chosenOption)

		userResponse = UserAnswer{
			question: question.question,
			answer:   chosenOption,
			result:   question.ReturnResult(chosenOption),
		}

		userPerformance.answers = append(userPerformance.answers, userResponse)
	}

	userPerformance.score = userPoints

	if err != nil {
		err = errors.New("Error iterating the questions")
	}
	return userPerformance, err
}

func CalculateUserAverage(quizzersScore []int, userScore int) int {
	totalScore := userScore
	quizzers := 1
	quizzersHighterAverage := 0

	for _, score := range quizzersScore {
		totalScore += score
		quizzers++
	}

	average := totalScore / quizzers
	if userScore > average {

		for _, score := range quizzersScore {

			if score > average {
				quizzersHighterAverage++
			}
		}
		userAverage := 100 / (quizzersHighterAverage + 1/quizzers)
		return userAverage
	} else {
		// The result 0 will be used to return a text response indicating that it does not exceed the average.
		return 0
	}
}

func ShowUserPerformance(up *UserPerformance) {
	fmt.Printf(`
	.
	.
	.

Your results are:
`)

	for questionNumber, questions := range up.answers {
		fmt.Printf("\nQuestion %d: \n%s \n", (questionNumber + 1), questions.question)
		fmt.Printf("\nYour answer: %d\n", questions.answer)
		fmt.Printf("\nResult: %s\n", questions.result)
	}

	fmt.Printf(`
	.
	.
	.

Your total score is: %d
	`, up.score)
}

// Start the game and return the score of the quizer:
func StartGame() int {
	fmt.Println("...")

	quizerPerformance, err := QuestionIteration(&Questions)
	quizerscore := quizerPerformance.score

	ShowUserPerformance(&quizerPerformance)

	if err != nil {
		fmt.Println(err)
	}
	return quizerscore
}

// Compare the result of the user with to the players and print the percentile of that:
func CompareUserPerformance(userScore int, scores []int) string {

	if userScore == 0 {
		return "Your score is below average, try again!"

	} else {

		// This block print the result, if CalculareUserAverage return "0" the score is below to the average:
		userAverage := CalculateUserAverage(scores, userScore)
		if userAverage != 0 {
			return "You scored higher than" + string(rune(userAverage)) + "%" + "of all quizzers!"

		} else {
			return "Your score is below average, try again!"
		}
	}
}
