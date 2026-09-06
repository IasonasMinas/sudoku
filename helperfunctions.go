package main // Declares the main package — the entry point of the program

import "github.com/01-edu/z01" // Imports the z01 library for printing runes (characters) to the terminal

// Prints an error message character by character
func printError(msg string) { // This function takes one parameter called msg of type string
	for _, char := range msg { // Loops through each character in the string
		z01.PrintRune(char) // Prints each character individually
	}
	z01.PrintRune('\n') // Adds a newline after the message
}

// Prints the Sudoku board in a formatted way
func printBoard(board [9][9]rune) { // // This function takes one parameter, a 2D array of runes board that consists of a 9x9 grid
	for row := 0; row < 9; row++ { // Loops through each row
		for col := 0; col < 9; col++ { // Loops through each column
			z01.PrintRune(board[row][col]) // Prints the cell value
			z01.PrintRune(' ')             // Adds a space between numbers
		}
		z01.PrintRune('\n') // Prints a newline after each row
	}
}

// Checks for duplicates in rows, columns, and 3x3 boxes
func hasDuplicates(board [9][9]rune) bool { // This function takes a 9x9 grid of runes as input and returns a boolean
	for i := 0; i < 9; i++ { // Loops through each row and column index
		row := [9]bool{} // Tracks seen digits in the current row
		col := [9]bool{} // Tracks seen digits in the current column
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' { // Skips empty cells
				n := board[i][j] - '1' // Converts rune to index (0-8)
				if row[n] {            // If already seen in row
					return true // Duplicate found
				}
				row[n] = true // Marks as seen
			}
			if board[j][i] != '.' { // Same for column
				n := board[j][i] - '1'
				if col[n] {
					return true
				}
				col[n] = true
			}
		}
	}

	// Checks each 3x3 box
	for row := 0; row < 9; row += 3 { // Loops through box starting rows
		for col := 0; col < 9; col += 3 { // Loops through box starting columns
			box := [9]bool{}               // Tracks seen digits in the current box
			for i := row; i < row+3; i++ { // Loops through the 3 rows
				for j := col; j < col+3; j++ { // Loops through the 3 columns
					if board[i][j] != '.' { // If the current cell is not empty
						n := board[i][j] - '1' // Converts the character to an index
						if box[n] {            // If already seen in box
							return true // Duplicate found
						}
						box[n] = true // Marks as seen
					}
				}
			}
		}
	}
	return false // No duplicates found
}

// Checks if a number can be placed in a specific cell
func isValid(board [9][9]rune, row, col int, num rune) bool { // This function takes four parameters, the 9x9 grid in a 2D board, integers row and col, num character and returns a boolean
	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false // Number already exists in row or column
		}
	}
	startRow := (row / 3) * 3 // Starting row of the 3x3 box
	startCol := (col / 3) * 3 // Starting column of the 3x3 box
	for row := startRow; row < startRow+3; row++ {
		for col := startCol; col < startCol+3; col++ {
			if board[row][col] == num {
				return false // Number already exists in box
			}
		}
	}
	return true // Number is valid in this cell
}

// Solves the Sudoku puzzle using backtracking
func solve(board *[9][9]rune) bool { // This function takes a pointer to the 9x9 grid and returns a boolean
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' { // If the current cell is empty
				for num := '1'; num <= '9'; num++ { // Tries all possible numbers from 1 to 9
					if isValid(*board, row, col, num) { // Checks if the number is valid in this cell
						board[row][col] = num // Places the number in the cell
						if solve(board) {     // Recursively tries to solve the rest of the puzzle
							return true // If the board is solved, returns true
						}
						board[row][col] = '.' // If the number doesn't work, resets the cell and tries next
					}
				}
				return false // If no number can be placed in this cell, returns false
			}
		}
	}
	return true // If all cells are filled, returns true (the board is solved)
}
