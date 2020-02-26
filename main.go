package main

import (
	//"fmt"
	sudoku "github.com/camen6ert/sudokuGenerator/sudokuGen"
)

func main() {

	//init population

	population := make([][9][9]uint, 0)

	for i := 0; i < 1000; i++ {
		s := sudoku.SudokuGen()
		population = append(population, s)
	}

	for i:=0; i<len(population)

}
