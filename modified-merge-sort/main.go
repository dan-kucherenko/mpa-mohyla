package main

import (
	"bufio"
	"fmt"
	"modified-merge-sort/generator"
	"modified-merge-sort/sorter"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Enter the length of the array: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arrLength, _ := strconv.Atoi(scanner.Text())
	fmt.Println("Length of the array will be", arrLength)
	arr := generator.GenerateArr(arrLength)
	fmt.Println("Generated array:", arr, "\n")
	sorter.MergeSort(arr)
	fmt.Println("\nResulting array:", arr)
}
