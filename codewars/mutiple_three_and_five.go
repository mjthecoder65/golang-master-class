package codewars

// Question: https://www.codewars.com/kata/514b92a657cdc65150000006/train/go

func Multiple3And5(number int) int {

	if number < 0 {
		return 0
	}

	var sum int

	for num := 0; num < number; num++ {

		if (num%3 == 0) || (num%5 == 0) {
			sum += num
		}
	}

	return sum
}
