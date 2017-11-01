package main

import "os"
import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

/**
Description: main runs the entire program, including reading in a game file,
				solving the game, and printing the outputs.
Arguments: None
Returns: Nothing
*/
func main() {
	// This section of code reads in a file and produces a 9 x 9 integer matrix
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

	printMatrix(matrix, "Initial Matrix")

	// Solve the game and calculate the time it took
	startTime := time.Now()
	solved := solve(matrix)
	elapsed := time.Now().Sub(startTime).Nanoseconds()
	elapsed = elapsed / 1000000
	fmt.Printf("Time: %d ms\n", elapsed)

	// If the game does not have a solution, print that
	if solved == false {
		fmt.Printf("No solution.\n")
	}

}

/**
Description: This is the most important function: it takes a matrix and tries to fill it
				in recursively, while checking if each move is valid.
Arguments:
			- matrix: a 9 x 9 integer matrix, or game board.
Returns: This function returns the boolean value true if the Suduko game was solved
			and false if it was not.
*/
func solve(matrix [9][9]int) bool {
	// Find the next cell to try numbers at. If there are no cells left
	// then we have solved the game.
	emptyRow, emptyCol := findEmptyCell(matrix)
	if emptyRow == -1 && emptyCol == -1 {
		if checkSolution(matrix) {
			printMatrix(matrix, "Solved")
			return true
		} else {
			return false
		}
	}

	// Place numbers at the current empty cell, and recursively call
	// the solve function. If the number didn't work, put it back to
	// empty.
	var i int
	for i = 1; i <= 9; i++ {
		if canPlaceNumber(matrix, emptyRow, emptyCol, i) {
			matrix[emptyRow][emptyCol] = i
			if solve(matrix) {
				return true
			}
			matrix[emptyRow][emptyCol] = 0
		}
	}
	return false
}

/**
Description: This function checks the final soluction to make sure the game is
				actually solved.
Arguments:
			- matrix: a 9 x 9 integer matrix, or game board
Returns: True if the game is solved, false otherwise.
*/
func checkSolution(matrix [9][9]int) bool {
	var row, col int
	for row = 0; row < 9; row++ {
		for col = 0; col < 9; col++ {
			num := matrix[row][col]
			if canPlaceInRow(matrix, row, num) && canPlaceInCol(matrix, col, num) && canPlaceInArea(matrix, row, col, num) {
				continue
			} else {
				return false
			}
		}
	}
	return true
}

/**
Description: This function checks to make sure we can place a
				certain number in a given cell, with regards
				to the rules of Suduko
Arguments:
			- matrix: a 9 x 9 integer matrix, or game board
			- row/col: the current cell we want to place a number in
			- number: the number we want to place in the cell
Returns: True if we are allowed to place the number there, false
			otherwise.
*/
func canPlaceNumber(matrix [9][9]int, row, col, number int) bool {
	// The three rules of Suduko define that a number 'x' can be
	// the only 'x' in the that row, column, and 3 x 3 sub-grid.
	canPlaceRow := canPlaceInRow(matrix, row, number)
	canPlaceCol := canPlaceInCol(matrix, col, number)
	canPlaceArea := canPlaceInArea(matrix, row, col, number)
	if canPlaceRow && canPlaceCol && canPlaceArea {
		return true
	} else {
		return false
	}
}

/**
Description: This method checks if we can place a number in a
				3 x 3 sub-grid.
Arguments:
			- matrix: a 9 x 9 integer matrix, or game board
			- row/col: the cell we want to place the number in
			- the number we want to put in the cell
Returns: True if the 3 x 3 sub-grid that row/cell is in does
			not already contain 'number', false if it does.
*/
func canPlaceInArea(matrix [9][9]int, row, col, number int) bool {
	// The findAreaMinMax method returns the mins/maxes for the
	// 3 x 3 sub-grid that row/col is in.
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

/**
Description: This method finds the minimums and maximums
				of the current 3 x 3 sub-grid given the
				current row and col.
Arguments:
			- row/col: current cell
Returns: Returns the rowMin, rowMax, colMin, and colMax
*/
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

/**
Description: This method checks if we can place 'number'
				in column 'col'.
Arguments:
			- matrix: a 9 x 9 integer matrix, or game board
			- col: the column we want to put 'number' in
			- number: the number we want to put in 'col'
Returns: Returns false if 'number' is already in 'col', true
			if 'number' is NOT in the column already.
*/
func canPlaceInCol(matrix [9][9]int, col, number int) bool {
	var i int
	for i = 0; i < 9; i++ {
		if matrix[i][col] == number {
			return false
		}
	}
	return true
}

/**
Description: This method check sif we can place 'number'
				in row 'row'.
Arguments:
			- matrix: a 9 x 9 integer matrix, or game board
			- row: the row we want to put 'number' in
			- number: the number we want to put in 'row'
Returns: Return false if 'number' is already in 'row', true
			if 'number' is NOT in the row already.
*/
func canPlaceInRow(matrix [9][9]int, row, number int) bool {
	var i int
	for i = 0; i < 9; i++ {
		if matrix[row][i] == number {
			return false
		}
	}
	return true
}

/**
Description: Given a 9 x 9 matrix, this function finds the row and column
				of the nearest empty cell, starting on the first row and
				column.
Arguments:
			- matrix: a 9 x 9 integer matrix, or game board
Returns: Returns a row/col if an empty cell was found, otherwise -1,-1 is
			returned.
*/
func findEmptyCell(matrix [9][9]int) (int, int) {
	// We start at the first cell (upper-left-hand corner)
	i := 0
	j := 0
	for {
		// Only check if we are in the bounds of board
		if i < 9 && j < 9 {
			num := matrix[i][j]
			// If we found an empty cell, return the current row/col
			if num == 0 {
				return i, j
			} else {
				// Increment the col we are checking. If we get past
				// the end of the board, reset the col and increase
				//the row counter by one.
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

/**
Description: This function prints out a matrix.
Arguments:
			- matrix: a 9 x 9 integer matrix, or game board
			- message: a message to print before printing the matrix
Returns: Nothing.
*/
func printMatrix(matrix [9][9]int, message string) {
	fmt.Printf("\n%s:\n", message)

	var i, j int
	for i = 0; i < 9; i++ {
		for j = 0; j < 9; j++ {
			fmt.Printf("%d", matrix[i][j])
			if j == 2 || j == 5 {
				fmt.Printf("\t")
			} else {
				fmt.Printf(" ")
			}
			if j == 8 {
				fmt.Printf("\n")
			}
		}
		if i == 2 || i == 5 {
			fmt.Printf("\n")
		}
	}
}

/**
Description: This function reads in game file and creates
				a 9 x 9 integer matrix from the read-in values.
Arguments:
			- fileName: the name of the file to read
Returns: Returns a 9 x 9 integer matrix and an error. If the
			error is nil, then the matrix was read in successfully.
*/
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
			if number < 0 || number > 9 {
				err = errors.New("invalid number read in. Numbers must be in the range [0, 9]")
				goto errorWhileReading
			}
			if err == nil {
				matrix[rowIndex][colIndex] = number
			} else {
				goto errorWhileReading
			}
		}
		rowIndex++
	}
	goto noError
errorWhileReading:
	file.Close()
	return matrix, err

noError:
	file.Close()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return matrix, err
}
