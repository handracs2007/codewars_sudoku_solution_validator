package main

import "fmt"

func CheckCluster(i, j int) int {
	const size = 3

	return size*(i/size) + (j / size)
}

func ValidateSolution(m [][]int) bool {
	var rowBit [9][9]int
	var colBit [9][9]int
	var clusterBit [9][9]int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			var number = m[i][j]
			var cluster = CheckCluster(i, j)

			if number == 0 {
				return false
			} else if rowBit[i][number-1] == 1 || colBit[j][number-1] == 1 || clusterBit[cluster][number-1] == 1 {
				return false
			} else {
				rowBit[i][number-1] = 1
				colBit[j][number-1] = 1
				clusterBit[cluster][number-1] = 1
			}
		}
	}

	return true
}

func main() {
	var testTrue = [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	var testFalse = [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 0, 3, 4, 8},
		{1, 0, 0, 3, 4, 2, 5, 6, 0},
		{8, 5, 9, 7, 6, 1, 0, 2, 0},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 0, 1, 5, 3, 7, 2, 1, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 0, 0, 4, 8, 1, 1, 7, 9},
	}

	fmt.Println(ValidateSolution(testTrue))
	fmt.Println(ValidateSolution(testFalse))
}
