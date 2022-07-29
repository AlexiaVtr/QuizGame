package cmd

import "testing"

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

}
