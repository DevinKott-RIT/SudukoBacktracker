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
	solved := solve(matrix, 0, 0)
	if solved == false {
		fmt.Printf("No solution.\n")
	}
}

func solve(matrix [9][9]int, curRow int, curCol int) bool {
	emptyRow, emptyCol := findEmptyCell(matrix, curRow, curCol)
	if emptyRow == -1 && emptyCol == -1 {
		fmt.Printf("Solved:\n")
		printMatrix(matrix)
		return true
	}

	var i int
	for i = 1; i <= 9; i++ {
		if canPlaceNumber(matrix, emptyRow, emptyCol, i) {
			matrix[emptyRow][emptyCol] = i
			if solve(matrix, emptyRow, emptyCol) {
				return true
			}
			matrix[emptyRow][emptyCol] = 0
		}
	}
	return false
}

func canPlaceNumber(matrix [9][9]int, row int, col int, number int) bool {
	canPlaceRow := canPlaceInRow(matrix, row, number)
	canPlaceCol := canPlaceInCol(matrix, col, number)
	canPlaceArea := canPlaceInArea(matrix, row, col, number)
	if canPlaceRow && canPlaceCol && canPlaceArea {
		return true
	} else {
		return false
	}
}

func canPlaceInArea(matrix [9][9]int, row, col, number int) bool {
	rowMin, rowMax, colMin, colMax := findAreaMinMax(row, col)
	var i, j int
	for i = rowMin; i < rowMax; i++ {
		for j = colMin; j < colMax; j++ {
			if matrix[i][j] == number {
				return false
			}
		}
	}
	return true
}

func findAreaMinMax(row, col int) (int, int, int, int) {
	if row < 3 {
		if col < 3 {
			return 0, 3, 0, 3
		} else if col < 6 {
			return 0, 3, 3, 6
		} else {
			return 0, 3, 6, 9
		}
	} else if row < 6 {
		if col < 3 {
			return 3, 6, 0, 3
		} else if col < 6 {
			return 3, 6, 3, 6
		} else {
			return 3, 6, 6, 9
		}
	} else {
		if col < 3 {
			return 6, 9, 0, 3
		} else if col < 6 {
			return 6, 9, 3, 6
		} else {
			return 6, 9, 6, 9
		}
	}
}

func canPlaceInCol(matrix [9][9]int, col, number int) bool {
	var i int
	for i = 0; i < 9; i++ {
		if matrix[i][col] == number {
			return false
		}
	}
	return true
}

func canPlaceInRow(matrix [9][9]int, row, number int) bool {
	var i int
	for i = 0; i < 9; i++ {
		if matrix[row][i] == number {
			return false
		}
	}
	return true
}

func findEmptyCell(matrix [9][9]int, row, col int) (int, int) {
	i := row
	j := col
	for {
		if i < 9 && j < 9 {
			num := matrix[i][j]
			if num == 0 {
				return i, j
			} else {
				j++
				if j >= 9 {
					i++
					j = 0
				}
			}
		} else {
			break
		}
	}
	return -1, -1
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
		file.Close()
		return matrix, err
	}

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
		for colIndex = 0; colIndex < partsLength; colIndex++ {
			number, err = strconv.Atoi(parts[colIndex])
			if err == nil {
				matrix[rowIndex][colIndex] = number
			} else {
				file.Close()
				return matrix, err
			}
		}
		rowIndex++
	}

	file.Close()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return matrix, err
}
