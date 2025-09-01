package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

func baseOperations(a, b float64, op rune) float64 {
	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		if b == 0 {
			panic("division by zero!")
		} else {
			return a / b
		}
	default:
		return 0
	}
}

func getTokens(expression string) []string {
	var tokens []string
	number := ""
	for _, c := range expression {
		if c != '+' && c != '-' && c != '*' && c != '/' && c != '(' && c != ')' && c != ' ' {
			number += string(c)
		} else {
			if number != "" {
				tokens = append(tokens, number)
				number = ""
			}
			if c != ' ' {
				tokens = append(tokens, string(c))
			}
		}
	}
	if number != "" {
		tokens = append(tokens, number)
	}
	return tokens
}

func makeOperation(numeric []float64, operator []rune) ([]float64, []rune, error) {
	if len(numeric) < 2 {
		return numeric, operator, errors.New("insufficient values in numeric stack for operation")
	}

	b := numeric[len(numeric)-1]
	a := numeric[len(numeric)-2]
	numeric = numeric[:len(numeric)-2]

	op := operator[len(operator)-1]
	operator = operator[:len(operator)-1]

	res := baseOperations(a, b, op)
	numeric = append(numeric, res)

	return numeric, operator, nil
}

func Calc(expression string) (float64, error) {
	priority := map[rune]int{'+': 1, '-': 1, '*': 2, '/': 2}
	var numericStack []float64
	var operatorStack []rune
	var containsNum bool

	if len(expression) == 0 {
		return 0, errors.New("empty expression")
	}

	for _, c := range expression {
		if unicode.IsDigit(c) {
			containsNum = true
		}
	}
	if !containsNum {
		return 0, errors.New("not contain numbers")
	}

	i := len(expression) - 1
	for i >= 0 && expression[i] == ' ' {
		i--
	}
	if i >= 0 {
		check := rune(expression[i])
		if !unicode.IsDigit(check) && check != ')' {
			return 0, errors.New("ends incorrectly")
		}
	}

	arr := getTokens(expression)
	for _, c := range arr {
		if num, err := strconv.ParseFloat(c, 64); err == nil {
			numericStack = append(numericStack, num)
		} else if c == "(" {
			operatorStack = append(operatorStack, '(')
		} else if c == ")" {
			for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1] != '(' {
				var err error
				numericStack, operatorStack, err = makeOperation(numericStack, operatorStack)
				if err != nil {
					return 0, err
				}
			}
			if len(operatorStack) == 0 {
				return 0, errors.New("mismatched parentheses")
			}
			operatorStack = operatorStack[:len(operatorStack)-1]
		} else {
			op := rune(c[0])
			for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1] != '(' && priority[op] <= priority[operatorStack[len(operatorStack)-1]] {
				var err error
				numericStack, operatorStack, err = makeOperation(numericStack, operatorStack)
				if err != nil {
					return 0, err
				}
			}
			operatorStack = append(operatorStack, op)
		}
	}
	for len(operatorStack) > 0 {
		var err error
		numericStack, operatorStack, err = makeOperation(numericStack, operatorStack)
		if err != nil {
			return 0, err
		}
	}

	if len(numericStack) != 1 {
		return 0, errors.New("error in expression evaluation")
	}

	return numericStack[0], nil
}

func main() {
	var expr string
	fmt.Scan(&expr)
	res, err := Calc(expr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
