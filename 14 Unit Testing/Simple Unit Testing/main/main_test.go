package main

import (
	"fmt"
	"testing"
)

func TestMath(t *testing.T) {
	if Addition(6, 2) != 8 {
		t.Errorf(fmt.Sprintf("Addition 6 and 2 got %d, want 8", Addition(6, 2)))
	}

	if Subtraction(6, 2) != 4 {
		t.Errorf(fmt.Sprintf("Subtraction 6 and 2 got %d, want 4", Subtraction(6, 2)))
	}

	if Division(6, 2) != 3 {
		t.Errorf(fmt.Sprintf("Division 6 and 2 got %d, want 3", Division(6, 2)))
	}

	if Multiplication(6, 2) != 12 {
		t.Errorf(fmt.Sprintf("Division 6 and 2 got %d, want 12", Multiplication(6, 2)))
	}
}
