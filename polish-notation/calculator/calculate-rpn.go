package calculator

import (
	"bufio"
	"os"
	"polish-notation/structure"
	"regexp"
	"strconv"
	"strings"
)

func Calculate(filePath string) int64 {
	stack := structure.Stack{}
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	query := scanner.Text()
	tokens := strings.Fields(query)
	for _, tok := range tokens {
		charIsNum, _ := isNum(tok)
		if charIsNum {
			stack.Push(tok)
		} else if isOperator(tok) {
			operand2, _ := strconv.ParseInt(stack.Pop(), 10, 16)
			operand1, _ := strconv.ParseInt(stack.Pop(), 10, 16)
			stack.Push(strconv.FormatInt(int64(makeOperation(int16(operand1), int16(operand2), tok)), 10))
		}
	}
	res, _ := strconv.ParseInt(stack.Peek(), 10, 16)
	return res
}

func makeOperation(operand1, operand2 int16, operator string) int16 {
	switch operator {
	case "+":
		return operand1 + operand2
	case "-":
		return operand1 - operand2
	case "*":
		return operand1 * operand2
	case "/":
		return operand1 / operand2
	}
	return 0
}

func isOperator(curChar string) bool {
	return curChar == "+" || curChar == "-" || curChar == "*" || curChar == "/"
}

func isNum(s string) (bool, error) {
	return regexp.MatchString("\\d", s)
}
