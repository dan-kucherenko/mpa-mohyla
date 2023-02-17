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
	for scanner.Scan() || stack.Len() != 0 {
		char := scanner.Text()
		isLetterOrNum, _ := regexp.MatchString("\\w", char)
		if isLetterOrNum {
			prevIsNum, _ := regexp.MatchString("\\d", prevChar)
			if prevIsNum {
				sb.WriteString(char)
			} else {
				switch prevChar {
				case "":
					sb.WriteString(char)
					break
				default:
					sb.WriteString(" ")
					sb.WriteString(char)
				}
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
		prevChar = char
	}
	_, err = writeToFile(sb)
	if err != nil {
		return
	}
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
	operators := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}
	return operators[curChar]
}

func lessPrioritized(oper1 string, oper2 string) bool {
	if (oper1 == "-" || oper1 == "+") && (oper2 == "/" || oper2 == "*") || (oper1 == oper2) ||
		(oper1 == "+" && oper2 == "-") || (oper1 == "-" && oper2 == "+") {
		return true
	} else {
		return false
	}
}
