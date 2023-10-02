package leetcode

// October 2, 2023
// URL: https://leetcode.com/problems/remove-colored-pieces-if-both-neighbors-are-the-same-color/description/
// Title: Remove Colored Pieces if Both Neighbors are the Same Color

// func WinnerOfGame(colors string) bool {
// 	if len(colors) <= 2 {
// 		return false
// 	}

// 	visited := make(map[int]bool)
// 	aliceLastMove, bobLastMove := 1, 1
// 	var player string = "Alice"

// 	for {
// 		// Variable to check if a player was able to make a move.
// 		madeMove := false

// 		// Alice should start to make a move.
// 		if player == "Alice" {
// 			for i := aliceLastMove; i < len(colors)-1; i++ {
// 				if colors[i] == 'A' && colors[i] == colors[i-1] && colors[i] == colors[i+1] {
// 					if _, ok := visited[i]; !ok {
// 						aliceLastMove = i
// 						visited[i] = true
// 						madeMove = true
// 						player = "Bob"
// 						break
// 					}
// 				}
// 			}
// 		} else {
// 			// Bob make a move after Alice.
// 			for i := bobLastMove; i < len(colors)-1; i++ {
// 				if colors[i] == 'B' && colors[i] == colors[i-1] && colors[i] == colors[i+1] {
// 					if _, ok := visited[i]; !ok {
// 						bobLastMove = i
// 						visited[i] = true
// 						madeMove = true
// 						player = "Alice"
// 						break
// 					}
// 				}
// 			}
// 		}

// 		if !madeMove {
// 			break
// 		}
// 	}

// 	if player == "Bob" {
// 		return true
// 	} else {
// 		return false
// 	}
// }

func WinnerOfGame(colors string) bool {
	aliceCount, bobCount := 0, 0

	for i := 1; i < len(colors)-1; i++ {
		if colors[i] == colors[i-1] && colors[i] == colors[i+1] {
			if colors[i] == 'A' {
				aliceCount++
			} else {
				bobCount++
			}
		}

	}

	return aliceCount > bobCount
}
