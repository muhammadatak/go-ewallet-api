package util

import "math/rand/v2"

func GenerateOTPCode() string {

	numbers := []rune("0123456789")
	code := make([]rune, 4)
	for i := range code {
		code[i] = numbers[rand.IntN(len(numbers))]
	}
	return string(code)
}

func GenerateId() string {
	numbers := []rune("1234567890qwertyuiopasdfghjklzxcvbnm")
	code := make([]rune, 16)
	for i := range code {
		code[i] = numbers[rand.IntN(len(numbers))]
	}
	return string(code)
}
