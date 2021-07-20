package main

import (
	"fmt"
	"strings"
)

func reverseInParentheses(inputString string) string {
	hasParenthesis := strings.Contains(inputString, "(")
	newString := inputString

	fmt.Println("hasParenthesis: ", hasParenthesis)

	for hasParenthesis {
		newString = replaceParenthesis(newString)

		fmt.Println("newString: ", newString)

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

	fmt.Println("beginning: ", beginning)
	fmt.Println("sub: ", sub)
	fmt.Println("end: ", end)

	newSub := ""
	for i := 1; i < len(sub)-1; i++ {
		char := sub[len(sub)-i-1 : len(sub)-i]
		// fmt.Println("char: ", char)

		if char == "(" {
			char = ")"
		} else if char == ")" {
			char = "("
		}

		newSub += char
	}

	fmt.Println("newSub: ", newSub)

	return beginning + newSub + end
}

func findClosingParenthesis(s string) int {
	// fmt.Println("findClosingParenthesis - s: ", s)

	countOpen := 0
	for index, iterChar := range s {
		if iterChar == '(' {
			countOpen++
			// fmt.Println("countOpen UP! ", countOpen)
		} else if iterChar == ')' {
			countOpen--
			// fmt.Println("countOpen DOWN! ", countOpen)
			if countOpen == 0 {
				// fmt.Println("THIS IS IT: ", index)
				return index
			}
		}
	}

	return 0
}

func main() {
	test1 := "(bar)"
	test2 := "foo(bar)baz"
	test3 := "foo(bar)baz(blim)"
	test4 := "foo(bar(baz))blim"
	test5 := "((foo(bar(baz))blim))"
	test6 := "(foo(bar(baz))blim)"
	test7 := "((((roo))))"
	test8 := "(((((roo)))))"
	test9 := "((r))(roo)"
	test10 := "(())(((roo)))"

	fmt.Println("RESULT: ", reverseInParentheses(test1))
	fmt.Println("RESULT: ", reverseInParentheses(test2))
	fmt.Println("RESULT: ", reverseInParentheses(test3))  //foorabbazmilb
	fmt.Println("RESULT: ", reverseInParentheses(test4))  //foobazrabblim
	fmt.Println("RESULT: ", reverseInParentheses(test5))  //foobazrabblim
	fmt.Println("RESULT: ", reverseInParentheses(test6))  //foobazrabblim INVERSE
	fmt.Println("RESULT: ", reverseInParentheses(test7))  //roo
	fmt.Println("RESULT: ", reverseInParentheses(test8))  //oor
	fmt.Println("RESULT: ", reverseInParentheses(test9))  //roor
	fmt.Println("RESULT: ", reverseInParentheses(test10)) //oor
}
