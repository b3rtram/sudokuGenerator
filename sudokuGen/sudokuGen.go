package sudokugen

import (
	"fmt"
	"math/rand"
	"time"
)

func SudokuGen() [9][9]uint {

	var s [9][9]uint

	rand.Seed(time.Now().UTC().UnixNano())
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			s[y][x] = randInt(1, 9)
			fmt.Printf("%d ", s[y][x])
		}
		fmt.Printf("\n")
	}
	///fmt.Printf("%v %v %v\n", s[0:2][0:2], s[0:2][3:5], s[0:2][2:3])
	return s
}

func SudokuFitness(s [9][9]uint) int {

	fit := 0
	var nr [9]bool
	//check lines
	for y := 0; y <= 9; y++ {
		for x := 0; x <= 9; x++ {
			v := nr[s[y][x]]
			if v == true {
				fit++
			}

			nr[s[y][x]] = true
		}
	}

	var nr2 [9]bool

	for x := 0; x <= 9; x++ {
		for y := 0; y <= 9; y++ {
			v := nr2[s[y][x]]
			if v == true {
				fit++
			}

			nr2[s[y][x]] = true
		}
	}

	var nr3 [9]bool

	for i := 0; i < 3; i++ {
		for y := i * 3; y < (i + 1*3); y++ {
			for x := y; x < y+3; x++ {
				v := nr3[s[y][x]]
				if v == true {
					fit++
				}

				nr3[s[y][x]] = true
			}
		}
	}

	return fit

}

func randInt(min int, max int) uint {
	return uint(min + rand.Intn(max-min))
}
