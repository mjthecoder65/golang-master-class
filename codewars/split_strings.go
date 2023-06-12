package codewars

// Question: https://www.codewars.com/kata/515de9ae9dcfc28eb6000001/train/go

func Solution(str string) []string {
	// Solution I created myself.
	var temp string

	result := make([]string, 0)

	for _, char := range str {

		if len(temp) == 2 {
			result = append(result, temp)
			temp = ""
		}

		temp += string(char)
	}

	if len(temp) == 1 {
		temp = temp + "_"
	}

	result = append(result, temp)

	return result
}

func Solution2(str string) []string {
	result := make([]string, 0)

	if len(str)%2 != 0 {
		str += "_"
	}

	for index := 0; index < len(str); index += 2 {
		result = append(result, str[index:index+2])
	}

	return result
}
