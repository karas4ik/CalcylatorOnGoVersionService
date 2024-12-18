package errors

import "fmt"

type InvalidExpressionError struct {
	Expression string
}

func (e *InvalidExpressionError) Error() string {
	return fmt.Sprintf("invalid expression: %s", e.Expression)
}

type DivisionByZeroError struct{}

func (e *DivisionByZeroError) Error() string {
	return "division by zero"
}

type MalformedRequestError struct {
	Message string
}

func (e *MalformedRequestError) Error() string {
	return fmt.Sprintf("malformed request: %s", e.Message)
}

type InsufficientOperandsError struct{}

func (e *InsufficientOperandsError) Error() string {
	return "insufficient operands for operation"
}

type UnsupportedOperatorError struct {
	Operator string
}

func (e *UnsupportedOperatorError) Error() string {
	return fmt.Sprintf("unsupported operator: %s", e.Operator)
}

type MismatchedParenthesesError struct{}

func (e *MismatchedParenthesesError) Error() string {
	return "mismatched parentheses"
}
