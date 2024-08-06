package leetcode_50

// Rolling up some shit with Golang

// URL: https://leetcode.com/problems/set-matrix-zeroes/

// Rolling up some shit with Golang
type Pair struct {
	Row    int
	Column int
}

func SetZeroes(matrix [][]int) {
	pairs := []Pair{}

	for row := 0; row < len(matrix); row++ {
		for column := 0; column < len(matrix[0]); column++ {
			if matrix[row][column] == 0 {
				pairs = append(pairs, Pair{Row: row, Column: column})
			}
		}
	}

	for _, pair := range pairs {
		// Set column to zeroes
		setColumnToZeroes(matrix, pair.Column)
		// Set row to zeroes.
		setRowToZeroes(matrix, pair.Row)

	}
}

func setColumnToZeroes(matrix [][]int, column int) {
	for row := 0; row < len(matrix); row++ {
		matrix[row][column] = 0
	}
}

func setRowToZeroes(matrix [][]int, row int) {
	for column := 0; column < len(matrix[0]); column++ {
		matrix[row][column] = 0
	}
}
