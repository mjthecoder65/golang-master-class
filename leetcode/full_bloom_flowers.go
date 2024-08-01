package leetcode

// URL: https://leetcode.com/problems/number-of-flowers-in-full-bloom/description/
// Title: Number of Flowers in Full Bloom

// Time Complexity: O(n^2)
// Space Complexity: O(1)
func fullBloomFlowers(flowers [][]int, people []int) []int {
	answers := []int{}
	for _, time := range people {
		counter := 0
		for _, period := range flowers {
			if time >= period[0] && time <= period[1] {
				counter++
			}
		}
		answers = append(answers, counter)
	}
	return answers
}

// Time Complexity: O(n)
// Space Complexity: O(n)

// This solution exceed the time limit
func fullBloomFlowersHashTable(flowers [][]int, people []int) []int {
	complements := make(map[int]int)
	answers := []int{}
	for _, period := range flowers {
		for i := period[0]; i <= period[1]; i++ {
			complements[i]++
		}
	}
	for _, time := range people {
		answers = append(answers, complements[time])
	}
	return answers
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func fullBloomFlowersHashTableOptimized(flowers [][]int, people []int) []int {
	complements := make(map[int]int)
	answers := []int{}
	for _, period := range flowers {
		complements[period[0]]++
		complements[period[1]]--
	}
	counter := 0
	for i := 0; i <= 365; i++ {
		counter += complements[i]
		complements[i] = counter
	}
	for _, time := range people {
		answers = append(answers, complements[time])
	}
	return answers
}
