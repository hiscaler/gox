package filex

import "os"

func IsFile(path string) bool {
	return !IsDir(path)
}

func IsDir(path string) bool {
	fi, err := os.Stat(path)
	if os.IsExist(err) && fi.IsDir() {
		return true
	}

	return false
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}
