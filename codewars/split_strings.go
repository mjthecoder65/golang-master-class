package codewars

// Question: https://www.codewars.com/kata/515de9ae9dcfc28eb6000001/train/go

func SplitString(str string) []string {
	result := make([]string, 0)

	if len(str)%2 != 0 {
		str += "_"
	}

	for index := 0; index < len(str); index += 2 {
		result = append(result, str[index:index+2])
	}

	return result
}
