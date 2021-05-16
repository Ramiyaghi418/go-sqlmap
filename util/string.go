package util

func DeleteLastChar(data string) string {
	r := []rune(data)
	return string(r[:len(r)-1])
}

func SubstringFrom(data string, index int) string {
	r := []rune(data)
	return string(r[index:])
}

func GetIndexChar(data string, index int) string {
	r := []rune(data)
	return string(r[index])
}
