package coder

import (
	"bufio"
	"os"
	"polish-notation/structure"
	"regexp"
	"strings"
)

func Code(filePath string) {
	var sb strings.Builder
	stack := structure.Stack{}
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)
	var prevChar string
	var prevPrevChar string
	for scanner.Scan() || stack.Len() != 0 {
		char := scanner.Text()
		if res, _ := isLetterOrNum(char); res {
			prevIsNum, _ := isNum(prevChar)
			if prevIsNum || prevChar == "" {
				sb.WriteString(char)
			} else if prevChar == "(" && prevPrevChar == "" {
				sb.WriteString(char)
			} else {
				sb.WriteString(" ")
				sb.WriteString(char)
			}
		} else if char == "(" {
			stack.Push(char)
		} else if char == ")" {
			for stack.Peek() != "(" && stack.Len() != 0 {
				sb.WriteString(" ")
				sb.WriteString(stack.Pop())
			}
			stack.Pop()
		} else if isOperator(char) {
			for lessPrioritized(char, stack.Peek()) && stack.Len() != 0 {
				sb.WriteString(stack.Pop())
			}
			stack.Push(char)
		} else {
			sb.WriteString(" ")
			sb.WriteString(stack.Pop())
		}
		prevPrevChar = prevChar
		prevChar = char
	}
	_, err = writeToFile(sb)
	if err != nil {
		return
	}
}

func isNum(s string) (bool, error) {
	return regexp.MatchString("\\d", s)
}

func isLetterOrNum(s string) (bool, error) {
	return regexp.MatchString("\\w", s)
}

func writeToFile(sb strings.Builder) (bool, error) {
	file, err := os.Create("files/coded_polish_notation.txt")
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
	return curChar == "+" || curChar == "-" || curChar == "*" || curChar == "/"
}

func lessPrioritized(oper1 string, oper2 string) bool {
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
