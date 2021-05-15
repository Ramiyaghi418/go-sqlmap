package util

func DeleteLastChar(data string) string {
	r := []rune(data)
	return string(r[:len(r)-1])
}
