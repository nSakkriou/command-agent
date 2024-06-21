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
