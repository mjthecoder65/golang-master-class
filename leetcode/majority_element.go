package leetcode

// URL: https://leetcode.com/problems/majority-element/description/
// Title: Majority Element I

func MajorityElement(nums []int) int {
	majority := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == majority {
			count++
		} else {
			count--
		}

		if count == 0 {
			majority = nums[i]
			count = 1
		}
	}

	return majority
}

// func majorityElement(nums []int) int {
//     counter := make(map[int]int)

//     for _, num := range nums {
//         counter[num]++
//     }

//     for num, count := range counter {
//         if count > len(nums) / 2 {
//             return num
//         }
//     }

//     return -1
// }
