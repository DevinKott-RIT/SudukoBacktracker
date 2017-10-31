package main

import "os"
import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	arguments := os.Args[1:]
	numArguments := len(arguments)
	if numArguments != 1 {
		fmt.Printf("usage: ./SudukoSolver <FILE_NAME>\n")
		os.Exit(1)
	}

	fileName := arguments[0]
	var matrix [9][9]int
	var err error
	matrix, err = readFile(fileName)
	if err != nil {
		fmt.Printf("Encountered an error, exiting the program: %s\n", err)
		os.Exit(1)
	}
	printMatrix(matrix)
}

func printMatrix(matrix [9][9]int) {
	fmt.Printf("Matrix:\n")

	var i, j int
	for i = 0; i < 9; i++ {
		for j = 0; j < 9; j++ {
			fmt.Printf("%d ", matrix[i][j])
			if j == 8 {
				fmt.Printf("\n")
			}
		}
	}
}

func readFile(fileName string) ([9][9]int, error) {
	var matrix [9][9]int
	var err error

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rowIndex := 0
	colIndex := 0
	number := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		parts := strings.Fields(line)
		partsLength := len(parts)
		var i int
		for i = 0; i < partsLength; i++ {
			number, err = strconv.Atoi(parts[colIndex])
			if err == nil {
				matrix[rowIndex][colIndex] = number
				colIndex++
			} else {
				return matrix, err
			}
		}
		colIndex = 0
		rowIndex++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return matrix, err
}
