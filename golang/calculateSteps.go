package main

import "unicode"

func calculateSteps(password string) int {
	length := len(password)
	hasLower, hasUpper, hasDigit := false, false, false
	repeatCount := 0
	steps := 0

	var missingTypes int
	for _, char := range password {
		if unicode.IsLower(char) {
			hasLower = true
		} else if unicode.IsUpper(char) {
			hasUpper = true
		} else if unicode.IsDigit(char) {
			hasDigit = true
		}
	}

	if !hasLower {
		missingTypes++
	}
	if !hasUpper {
		missingTypes++
	}
	if !hasDigit {
		missingTypes++
	}

	for i := 2; i < length; i++ {
		if password[i] == password[i-1] && password[i-1] == password[i-2] {
			repeatCount++
			i++
		}
	}

	if length < 6 {
		steps = max(6-length, missingTypes)
	} else if length <= 20 {
		steps = max(repeatCount, missingTypes)
	} else {
		overLength := length - 20
		steps = overLength + max(repeatCount, missingTypes)
	}

	return steps
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
