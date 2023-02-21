package calculator

import (
	"bufio"
	"io"
	"log"
	"os"
	"polish-notation/structure"
	"regexp"
	"strconv"
	"strings"
)

func Calculate(filePath string) float64 {
	stack := structure.Stack{}
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	variablesMap := scanVariables(file)
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}
	scanner.Scan()
	query := scanner.Text()
	tokens := strings.Fields(query)
	for _, tok := range tokens {
		if val, has := variablesMap[tok]; has {
			stack.Push(val)
		} else if charIsNum, _ := isNum(tok); charIsNum {
			stack.Push(tok)
		} else if isOperator(tok) {
			operand2, _ := strconv.ParseFloat(stack.Pop(), 64)
			operand1, _ := strconv.ParseFloat(stack.Pop(), 64)
			stack.Push(strconv.FormatFloat(makeOperation(operand1, operand2, tok), 'f', -2, 64))
		}
	}
	res, _ := strconv.ParseFloat(stack.Peek(), 64)
	return res
}

func scanVariables(f *os.File) map[string]string {
	varmap := make(map[string]string)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.ContainsRune(line, '=') {
			parts := strings.Split(line, "=")
			if len(parts) == 2 {
				varName := strings.TrimSpace(parts[0])
				val := strings.TrimSpace(parts[1])
				varmap[varName] = val
			}
		}
	}
	return varmap
}

func makeOperation(operand1, operand2 float64, operator string) float64 {
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
