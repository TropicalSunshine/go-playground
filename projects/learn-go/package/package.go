package packagee

import (
	"errors"
	"fmt"
)

func Packagee(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}

	fmt.Println("packageee")
	return "test", errors.New("empty name")
}
