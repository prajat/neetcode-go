package main

import "strconv"

/*
You are given an array of strings tokens that represents a valid arithmetic expression in Reverse Polish Notation.
Return the integer that represents the evaluation of the expression.
The operands may be integers or the results of other operations.
The operators include '+', '-', '*', and '/'.
Assume that division between integers always truncates toward zero.
*/
func evalRPN(tokens []string) int {
	stack := make([]string, 0)

	for i := range tokens {
		if !isOperator(tokens[i]) {
			stack = append(stack, tokens[i])
		} else {
			operator := tokens[i]

			operand2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			operand1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			stack = append(stack, operate(operand1, operand2, operator))
		}
	}

	result, _ := strconv.Atoi(stack[0])
	return result
}

func isOperator(token string) bool {
	if token == "+" || token == "-" || token == "*" || token == "/" {
		return true
	}
	return false
}

func operate(a, b, o string) string {
	aa, _ := strconv.Atoi(a)
	bb, _ := strconv.Atoi(b)

	if o == "+" {
		return strconv.FormatInt(int64(aa+bb), 10)
	} else if o == "-" {
		return strconv.FormatInt(int64(aa-bb), 10)
	} else if o == "*" {
		return strconv.FormatInt(int64(aa*bb), 10)
	} else {
		return strconv.FormatInt(int64(aa/bb), 10)
	}
}
