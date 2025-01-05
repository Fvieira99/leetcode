package main

import "fmt"

func main() {
	fmt.Println(isValid("MCMXCIV"))
}

func isValid(s string) bool {
	stack := []string{}
	open := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}
	close := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}
	for _, v := range s {
		char := string(v)
		if open[char] != "" {
			stack = append(stack, char)
			continue
		}

		if len(stack) == 0 || stack[len(stack)-1] != close[char] {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
