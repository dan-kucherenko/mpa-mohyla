package main

import (
	"fmt"
	"polish-notation/calculator"
	"polish-notation/coder"
	"strconv"
)

func main() {
	fmt.Println("Encoded example to RPN: " + coder.Code("files/Expr_ar.txt"))
	fmt.Println("Result of RPN calculation: " +
		strconv.FormatInt(calculator.Calculate("files/Expr_pol.txt"), 10))
}
