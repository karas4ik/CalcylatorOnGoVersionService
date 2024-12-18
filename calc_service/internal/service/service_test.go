package service

import (
	"testing"
)

func TestCalculateArithmeticExpression(t *testing.T) {
	tests := []struct {
		expression string
		expected   string
		shouldFail bool
	}{
		{"2 + 2", "4", false},
		{"10 - 5", "5", false},
		{"3 * 3", "9", false},
		{"10 / 2", "5", false},
		{"5 / 0", "", true},
		{"invalid", "", true},
	}

	for _, tt := range tests {
		result, err := CalculateArithmeticExpression(tt.expression)
		if tt.shouldFail {
			if err == nil {
				t.Errorf("Expected error for expression: %s", tt.expression)
			}
		} else {
			if result != tt.expected {
				t.Errorf("For expression %s, expected %s, got %s", tt.expression, tt.expected, result)
			}
		}
	}
}
