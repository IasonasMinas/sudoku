This is a command-line Sudoku solver written in Go programming language that uses recursive backtracking to solve standard 9×9 puzzles. It validates input, checks for rule violations, and prints the solved board in a clean format.

It runs only if the input rules apply:

-Exactly 9 arguments, each with 9 characters
-Valid characters: 1 – 9 and . (for empty cells)
-Minimum of 17 filled cells required for a valid puzzle with only one possible solution
-No duplicates allowed in rows, columns, or 3x3 boxes

or else it returns an "Error:" message.

To run it, copy the "go run" commands from the "allpossibleresults.txt" file in the Terminal.
