package main

import "strings"

func reverseInParentheses(inputString string) string {
	hasParenthesis := strings.Contains(inputString, "(")
	newString := inputString

	for hasParenthesis {
		newString = replaceParenthesis(newString)

		hasParenthesis = strings.Contains(newString, "(")
	}

	return newString
}

func replaceParenthesis(inputString string) string {
	open := strings.Index(inputString, "(")
	close := strings.Index(inputString, ")")

	close = findClosingParenthesis(inputString[open:]) + open

	beginning := inputString[:open]
	sub := inputString[open : close+1]
	end := inputString[close+1:]

	newSub := ""
	for i := 1; i < len(sub)-1; i++ {
		char := sub[len(sub)-i-1 : len(sub)-i]

		if char == "(" {
			char = ")"
		} else if char == ")" {
			char = "("
		}

		newSub += char
	}

	return beginning + newSub + end
}

func findClosingParenthesis(s string) int {
	countOpen := 0
	for index, iterChar := range s {
		if iterChar == '(' {
			countOpen++
		} else if iterChar == ')' {
			countOpen--
			if countOpen == 0 {
				return index
			}
		}
	}

	return 0
}
