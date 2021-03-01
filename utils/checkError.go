package utils

import "errors"

func CheckEmpty(data ...interface{}) error {
	for _, value := range data {
		switch value {
		case "":
			return errors.New("make sure input not empty")
		case nil:
			return errors.New("make sure input not a nil")
		}
	}
	return nil
}
