package twitch

import (
	"fmt"
	"testing"
)

func TestGoalAmount(t *testing.T) {
	testCases := []struct {
		Value         int
		DecimalPlaces int
		Expected      float64
	}{
		{550, 2, 5.5},
		{100, 2, 1},
		{10000, 2, 100},
		{12345, 4, 1.2345},
		{9999999999, 1, 999999999.9},
		{9999999999, 0, 9999999999},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d-%d", tc.Value, tc.DecimalPlaces), func(t *testing.T) {
			amount := GoalAmount{
				Value:         tc.Value,
				DecimalPlaces: tc.DecimalPlaces,
			}

			actual := amount.Amount()

			if actual != tc.Expected {
				t.Errorf("expected %f got %f", tc.Expected, actual)
			}
		})
	}
}
