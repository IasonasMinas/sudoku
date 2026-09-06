package main // Declares the main package — the entry point of the program

import "os" // Imports the os package to access command-line arguments

func main() { // Main function where execution begins
	if len(os.Args) != 10 { // Expecting 9 rows (+ program name) = 10 arguments
		printError("Error: Invalid number of arguments") // If invalid, prints an error message
		return                                           // Stops execution
	}

	var board [9][9]rune // Initializes a 9x9 Sudoku board using rune (character) type
	filledcells := 0     // Counter to track how many cells are filled

	for i := 0; i < 9; i++ { // Loops through each row
		row := os.Args[i+1] // Gets the i-th row from command-line arguments
		if len(row) != 9 {  // Checks if the row has exactly 9 characters
			printError("Error: Invalid number of characters") // If invalid, prints an error message
			return                                            // Stops execution
		}
		for j := 0; j < 9; j++ { // Loops through each character in the row
			if row[j] != '.' && (row[j] < '1' || row[j] > '9') { // Checks if the character is valid
				printError("Error: Invalid character(s)") // If invalid, prints an error message
				return                                    // Stops execution
			}
			board[i][j] = rune(row[j]) // Converts the character to rune and stores it in the board
			if row[j] != '.' {         // Counts the filled cells
				filledcells++ // Increment the filled cell counter
			}
		}
	}

	if filledcells < 17 { // Checks if there are at least 17 filled cells
		printError("Error: Less than 17 filled cells") // If less found, prints an error message
		return                                         // Stops execution
	}

	if hasDuplicates(board) { // Checks for duplicates in rows, columns, or boxes
		printError("Error: Duplicates found in row(s), column(s) or 3x3 box(es)") // If found, prints an error message
		return                                                                    // Stops execution
	}

	if solve(&board) { //If the Sudoku is solvable,
		printBoard(board) // Prints the solved board
	} else { //If unsolvable,
		printError("Error: No solution") // Prints error
	}
}
