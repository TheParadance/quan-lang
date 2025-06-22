package utils

import "os"

func ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path) // replace with your file path
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ArrayItemContain[T comparable](array []T, item T) bool {
	for _, v := range array {
		if v == item {
			return true
		}
	}
	return false
}
