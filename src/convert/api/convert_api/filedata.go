package convertapi

import (
	"fmt"
	"os"
)


func FileData (file string) (os.FileInfo, error) {
	fd, err := os.Stat(file)
	if err != nil {
		return fd, fmt.Errorf("%v", err.Error())
	}
	return fd, nil
}