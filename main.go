package main

import (
	"fmt"
	sudoku "github.com/camen6ert/sudokuGenerator/sudokuGen"
)

func main() {
	for {
		s := sudoku.SudokuGen()
		b := sudoku.SudokuCorrect(s)
		fmt.Println(b)
		if b == true {
			break
		}
	}
}
