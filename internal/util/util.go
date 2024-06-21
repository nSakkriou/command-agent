package util

import (
	"os"
)

func IsEmpty(str string) bool {
	return str == ""
}

func ArrayContains(array []string, elem string) bool {
	for _, arrayElem := range array {
		if elem == arrayElem {
			return true
		}
	}

	return false
}

func FileExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}

		return false
	}
	return true
}

func Prefix(s, char string) string {
	if len(char) != 1 {
		return s
	}
	if len(s) > 0 && s[0] == char[0] {
		return s
	}
	return char + s
}
