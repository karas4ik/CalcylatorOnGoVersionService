package service

import (
	"calc_service/internal/errors"
	"calc_service/internal/model"
	"regexp"
	"strconv"
	"strings"
)

func CalculateArithmeticExpression(expr string) (string, error) {
	if !isValidExpression(expr) {
		return "", &errors.InvalidExpressionError{Expression: expr}
	}

	postfix, err := infixToPostfix(expr)
	if err != nil {
		return "", err
	}
	result, err := evaluatePostfixExpression(postfix)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(result), nil
}

func isValidExpression(expr string) bool {
	validExpr := regexp.MustCompile(`^[0-9+\-*/().\s]+$`)
	return validExpr.MatchString(expr)
}

func infixToPostfix(infix string) ([]model.Token, error) {
	var output []model.Token
	stack := []string{}
	tokens := tokenize(infix)

	for _, token := range tokens {
		switch token.Type {
		case "NUMBER":
			output = append(output, token)
		case "OPERATOR":
			for len(stack) > 0 && precedence(token.Value) <= precedence(stack[len(stack)-1]) {
				output = append(output, model.Token{"OPERATOR", stack[len(stack)-1]})
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token.Value)
		case "LPAREN":
			stack = append(stack, token.Value)
		case "RPAREN":
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				output = append(output, model.Token{"OPERATOR", stack[len(stack)-1]})
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				return nil, &errors.MismatchedParenthesesError{}
			}
			stack = stack[:len(stack)-1]
		}
	}

	for len(stack) > 0 {
		if stack[len(stack)-1] == "(" {
			return nil, &errors.MismatchedParenthesesError{}
		}
		output = append(output, model.Token{"OPERATOR", stack[len(stack)-1]})
		stack = stack[:len(stack)-1]
	}

	return output, nil
}

func evaluatePostfixExpression(postfix []model.Token) (int, error) {
	stack := []int{}

	for _, token := range postfix {
		if token.Type == "NUMBER" {
			num, err := strconv.Atoi(token.Value)
			if err != nil {
				return 0, &errors.InvalidExpressionError{Expression: token.Value}
			}
			stack = append(stack, num)
		} else if token.Type == "OPERATOR" {
			if len(stack) < 2 {
				return 0, &errors.InsufficientOperandsError{}
			}
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result int
			switch token.Value {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					return 0, &errors.DivisionByZeroError{}
				}
				result = a / b
			default:
				return 0, &errors.UnsupportedOperatorError{Operator: token.Value}
			}
			stack = append(stack, result)
		}
	}

	if len(stack) != 1 {
		return 0, &errors.InvalidExpressionError{Expression: "invalid expression"}
	}

	return stack[0], nil
}

func tokenize(expression string) []model.Token {
	var tokens []model.Token
	var currentToken strings.Builder

	for _, char := range expression {
		if char == ' ' {
			continue
		}
		if isOperator(char) || char == '(' || char == ')' {
			if currentToken.Len() > 0 {
				tokens = append(tokens, model.Token{"NUMBER", currentToken.String()})
				currentToken.Reset()
			}
			if char == '(' {
				tokens = append(tokens, model.Token{"LPAREN", string(char)})
			} else if char == ')' {
				tokens = append(tokens, model.Token{"RPAREN", string(char)})
			} else {
				tokens = append(tokens, model.Token{"OPERATOR", string(char)})
			}
		} else {
			currentToken.WriteRune(char)
		}
	}
	if currentToken.Len() > 0 {
		tokens = append(tokens, model.Token{"NUMBER", currentToken.String()})
	}

	return tokens
}

func isOperator(char rune) bool {
	return char == '+' || char == '-' || char == '*' || char == '/'
}

func precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	}
	return 0
}
