package util

import (
	"fmt"
	"log"
)

func CheckError(err error, msg string) {
	if err != nil {
		if msg == "" {
			log.Printf("%s", err.Error())
			return
		}
		log.Printf("%s: %s", msg, err.Error())
	}
}

func StringToBool(input string) (bool, error) {
	switch input {
	case "1":
		return true, nil
	case "0":
		return false, nil
	default:
		return false, fmt.Errorf("input '%s' is not a valid boolean", input)
	}
}
