package cmd

import (
	"fmt"
	"testing"
)

func TestControlScanOptions(t *testing.T) {
	result := ControlScanOptions(1)

	if result != 1 {
		t.Errorf("ControlScanOptions() FAILED. Expected %d, got %d", 1, result)
	} else {
		t.Logf("ControlScanOptions() PASSED. Expected %d, got %d", 1, result)
	}
}

func TestCalculateUserAverage(t *testing.T) {
	score := make([]int, 5, 5)
	score = append(score, 1, 1, 1, 2, 5)

	result := CalculateUserAverage(score, 5)
	if result != 50 {
		t.Errorf("CalculateUserAverage() FAILED. Expected %d, got %d", 50, result)
	} else {
		t.Logf("CalculateUserAverage() PASSED. Expected %d, got %d", 50, result)
	}

	score = nil
	score = append(score, 2, 2, 0)
	result = CalculateUserAverage(score, 4)
	if result != 100 {
		t.Errorf("CalculateUserAverage() FAILED. Expected %d, got %d", 100, result)
	} else {
		t.Logf("CalculateUserAverage() PASSED. Expected %d, got %d", 100, result)
	}

}

func TestCompareUserPerformance(t *testing.T) {
	scores := make([]int, 3, 3)
	scores = append(scores, 5, 5, 4)
	response := "Your score is below average, try again!"

	result := CompareUserPerformance(1, scores)
	if result != response {

		t.Errorf("CompareUserPerformance() FAILED. Expected %s, got %s", response, result)
	} else {
		t.Logf("CompareUserPerformance() PASSED. Expected %s, got %s", response, result)
	}

	result = CompareUserPerformance(0, scores)
	if result != response {

		t.Errorf("CompareUserPerformance() FAILED. Expected %s, got %s", response, result)
	} else {
		t.Logf("CompareUserPerformance() PASSED. Expected %s, got %s", response, result)
	}

	response = fmt.Sprintf("You scored higher than %d%s of all quizzers!", 100, "%")
	result = CompareUserPerformance(5, []int{1, 1, 1})
	if result != response {

		t.Errorf("CompareUserPerformance() FAILED. Expected %s, got %s", response, result)
	} else {
		t.Logf("CompareUserPerformance() PASSED. Expected %s, got %s", response, result)
	}

	scores = nil
	response = "You are the first quizzer! Share the game or play again to compare your performance."
	result = CompareUserPerformance(1, scores)
	if result != response {

		t.Errorf("CompareUserPerformance() FAILED. Expected %s, got %s", response, result)
	} else {
		t.Logf("CompareUserPerformance() PASSED. Expected %s, got %s", response, result)
	}
}

func TestReturnResult(t *testing.T) {
	result := FirstQuestion.ReturnResult(2)
	response := "The answerd was correct."
	if result != response {

		t.Errorf("ReturnResult() FAILED. Expected %s, got %s", response, result)
	} else {
		t.Logf("ReturnResult() PASSED. Expected %s, got %s", response, result)
	}

	result = FirstQuestion.ReturnResult(1)
	response = "The answerd was wrong."
	if result != response {

		t.Errorf("ReturnResult() FAILED. Expected %s, got %s", response, result)
	} else {
		t.Logf("ReturnResult() PASSED. Expected %s, got %s", response, result)
	}
}

func TestCalculateUserScore(t *testing.T) {
	result := FirstQuestion.CalculateUserScore(2)
	if result != 1 {

		t.Errorf("CalculateUserScore() FAILED. Expected %d, got %d", 1, result)
	} else {
		t.Logf("CalculateUserScore() PASSED. Expected %d, got %d", 1, result)
	}

	result = FirstQuestion.CalculateUserScore(1)
	if result != 0 {

		t.Errorf("CalculateUserScore() FAILED. Expected %d, got %d", 0, result)
	} else {
		t.Logf("CalculateUserScore() PASSED. Expected %d, got %d", 0, result)
	}
}

func TestToShow(t *testing.T) {
	FirstQuestion.ToShow(1)
}

func TestAppendScore(t *testing.T) {
	GameScores.AppendScore(1)
}

func TestShowUserPerformance(t *testing.T) {
	answer := UserAnswer{
		question: "What does 'arma' mean in Spanish?",
		answer:   2,
		result:   "The answerd was correct.",
	}
	us := UserPerformance{
		answers: []UserAnswer{
			answer,
		},
	}
	ShowUserPerformance(&us)
}
