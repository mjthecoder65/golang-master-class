package leetcode

// URL: https://leetcode.com/problems/pascals-triangle/
// Title : Pascal's Triangle

// func generate(numRows int) [][]int {
// 	pascalTriangle := make([][]int, 0)
// 	previousRow := []int{}

// 	for i := 0; i < numRows; i++ {
// 		currentRow := []int{1}
// 		if i == 0 {
// 			previousRow = currentRow
// 			pascalTriangle = append(pascalTriangle, currentRow)
// 			continue
// 		}

// 		var j int

// 		for j := 1; j < i; j++ {
// 			currentRow = append(currentRow, previousRow[j-1]+previousRow[j])
// 		}

// 		currentRow = append(currentRow, previousRow[j])
// 		pascalTriangle = append(pascalTriangle, currentRow)
// 		previousRow = currentRow
// 	}

// 	return pascalTriangle
// }

func generate(numRows int) [][]int {
	pascalTriangle := make([][]int, 0)

	for i := 0; i < numRows; i++ {
		currentRow := make([]int, i+1)
		currentRow[0] = 1
		currentRow[i] = 1

		for j := 1; j < i; j++ {
			currentRow[j] = pascalTriangle[i-1][j-1] + pascalTriangle[i-1][j]
		}

		pascalTriangle = append(pascalTriangle, currentRow)
	}

	return pascalTriangle
}
