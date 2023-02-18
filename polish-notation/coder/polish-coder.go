package coder

import (
	"bufio"
	"bytes"
	"os"
	"polish-notation/structure"
	"regexp"
	"strconv"
	"strings"
)

func Code(filePath string) string {
	var sb strings.Builder
	stack := structure.Stack{}
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	query := scanner.Text()
	tokens := strings.Fields(prepareTheQuery(query))

	for _, tok := range tokens {
		if res, _ := isLetterOrNum(tok); res {
			sb.WriteString(tok)
			sb.WriteString(" ")
		} else if tok == "(" {
			stack.Push(tok)
		} else if tok == ")" {
			for stack.Peek() != "(" && stack.Len() != 0 {
				poppedToken := stack.Pop()
				sb.WriteString(poppedToken)
				sb.WriteString(" ")
			}
			stack.Pop()
		} else if isOperator(tok) {
			for lessPrioritized(tok, stack.Peek()) && stack.Len() != 0 {
				poppedToken := stack.Pop()
				sb.WriteString(poppedToken)
				sb.WriteString(" ")
			}
			stack.Push(tok)
		}
	}
	for stack.Len() > 0 {
		poppedEl := stack.Pop()
		sb.WriteString(poppedEl)
		sb.WriteString(" ")
	}
	_, err = writeToFile(sb)
	if err != nil {
		return ""
	}
	return sb.String()
}

func isLetterOrNum(s string) (bool, error) {
	return regexp.MatchString("\\w", s)
}

func prepareTheQuery(input string) string {
	buf := &bytes.Buffer{}
	for _, char := range input {
		if isOperator(strconv.Itoa(int(char))) || char == 40 || char == 41 {
			buf.WriteRune(' ')
			buf.WriteRune(char)
			buf.WriteRune(' ')
		} else {
			buf.WriteRune(char)
		}
	}
	return buf.String()
}

func writeToFile(sb strings.Builder) (bool, error) {
	file, err := os.Create("files/Expr_pol.txt")
	if err != nil {
		return false, err
	}
	defer file.Close()
	_, err = file.WriteString(sb.String())
	if err != nil {
		return false, err
	}
	return true, nil
}

func isOperator(curChar string) bool {
	return curChar == "+" || curChar == "43" || curChar == "-" || curChar == "45" || curChar == "*" || curChar == "42" || curChar == "/" || curChar == "47"
}

func lessPrioritized(oper1, oper2 string) bool {
	return priority(oper1) <= priority(oper2)
}

func priority(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}
