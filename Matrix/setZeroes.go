package main

import (
	"fmt"
	"math"
)

func setZeroes(matrix [][]int) {
	rowf := false
	colf := false
	if matrix[0][0] == 0 {
		rowf = true
		colf = true
	}
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			matrix[i][0] = math.MaxInt32 + 1
			colf = true
		}
	}
	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			matrix[0][j] = math.MaxInt32 + 1
			rowf = true
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = math.MaxInt32 + 1
				matrix[0][j] = math.MaxInt32 + 1
			}
		}
	}
	//fmt.Println(matrix)

	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == math.MaxInt32+1 {
			for j := 0; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}

	for i := 1; i < len(matrix[0]); i++ {
		if matrix[0][i] == math.MaxInt32+1 {
			for j := 0; j < len(matrix); j++ {
				matrix[j][i] = 0
			}
		}
	}
	//fmt.Println(colf, rowf)
	if colf {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
	if rowf {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}
}

func main() {
	matrix := [][]int{{0, 1}}
	setZeroes(matrix)
	fmt.Println(matrix)
}
