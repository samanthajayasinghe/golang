package plus

import "testing"

func TestPlusReturnsZeroForZeroInput(t *testing.T) {
	got := Plus(0, 0)
	if got != 0 {
		t.Errorf("Plus(0,0) = %d; want 0", got)
	}
}

func TestReturnsSumOfTwoPositiveNumbers(t *testing.T) {
	got := Plus(1, 2)
	if got != 3 {
		t.Errorf("Plus(1,2) = %d; want 3", got)
	}
}

func TestReturnsSumOfTwoNegativeNumbers(t *testing.T) {
	got := Plus(-1, -2)
	if got != -3 {
		t.Errorf("Plus(-1,-2) = %d; want -3", got)
	}
}
