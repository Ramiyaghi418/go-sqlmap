package util

// DeleteLastChar 删除字符串最后一位
func DeleteLastChar(data string) string {
	r := []rune(data)
	return string(r[:len(r)-1])
}

// SubstringFrom 从索引处截断字符串
func SubstringFrom(data string, index int) string {
	r := []rune(data)
	return string(r[index:])
}

// GetIndexChar 获得字符串中某个索引对应的字符
func GetIndexChar(data string, index int) string {
	r := []rune(data)
	return string(r[index])
}
