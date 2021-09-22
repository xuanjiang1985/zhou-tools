package util

import (
	"io"
	"os"
)

func closeFile() (err error) {
	file, err := os.Create("./a")
	if err != nil {
		return
	}

	//01 defer file.Close()

	//02 do not ignore err
	defer func() {
		closeErr := file.Close()
		if err == nil {
			err = closeErr
		}
	}()

	_, err = io.WriteString(file, "hello gopher")
	return
}
