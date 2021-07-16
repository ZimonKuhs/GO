package utility

func Reverse(original string) string {
	var result string
	length := len(original)

	for i := 0; i < length; i++ {
		result += string(original[length-i-1])
	}

	return result
}
