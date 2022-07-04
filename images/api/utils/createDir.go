package utils

import "os"

func CreateDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0775)
		if err != nil {
			panic(err)
		}
	}
}
