package leetcode

// URL: https://leetcode.com/problems/pascals-triangle-ii/description/
// Title : Pascal's Triangle II

// func getRow(rowIndex int) []int {
//     lastRow := []int{}
//     for i := 0; i <= rowIndex; i++ {
//         currentRow := make([]int, i + 1)
//         currentRow[0] = 1

//         if i == 0 {
//             lastRow = currentRow
//             continue
//         }
//         var j int
//         for j = 1 ; j < i; j++ {
//             currentRow[j] = lastRow[j-1] + lastRow[j]
//         }

//         currentRow[j] = 1

//         lastRow = currentRow
//     }

//     return lastRow
// }

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1

	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			row[j] += row[j-1]
		}
	}

	return row
}
