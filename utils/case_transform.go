package utils

import (
	"strings"
	"unicode"
)

// ToCamelCase converts a string to camel case (e.g., "myUser" -> "MyUser")
func ToCamelCase(s string) string {
	var result strings.Builder
	upperNext := true

	for _, r := range s {
		if r == '_' || r == ' ' || r == '-' {
			upperNext = true
		} else {
			if upperNext {
				result.WriteRune(unicode.ToUpper(r))
				upperNext = false
			} else {
				result.WriteRune(r) // Preserve the original case
			}
		}
	}
	return result.String()
}

// ToSnakeCase converts a string to snake case (e.g., "MyUser" -> "my_user")
func ToSnakeCase(s string) string {
	if len(s) == 0 {
		return s
	}

	// Convert the first character to lowercase
	result := strings.Builder{}
	result.WriteRune(unicode.ToLower(rune(s[0])))

	// Append the rest of the string as-is
	result.WriteString(s[1:])

	return result.String()
}

// TransformString transforms the input string to camel case or snake case based on the input format
func TransformString(s string) (camelCase, snakeCase string) {
	// Detect if the input is already in camel case or snake case
	isCamelCase := false
	for _, r := range s {
		if unicode.IsUpper(r) {
			isCamelCase = true
			break
		}
	}

	if isCamelCase {
		// Input is in camel case, convert to snake case
		snakeCase = ToSnakeCase(s)
		camelCase = ToCamelCase(snakeCase) // Ensure camel case format
	} else {
		// Input is in snake case or lowercase, convert to camel case
		camelCase = ToCamelCase(s)
		snakeCase = ToSnakeCase(camelCase) // Ensure snake case format
	}

	return camelCase, snakeCase
}
