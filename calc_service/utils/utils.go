package utils

import (
	"regexp"
)

func ValidateExpression(expr string) bool {
	validExpr := regexp.MustCompile(`^[0-9+\-*/().\s]+$`)
	return validExpr.MatchString(expr)
}
