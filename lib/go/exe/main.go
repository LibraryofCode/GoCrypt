package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// HandleErrorFMT Handles an error.
func HandleErrorFMT(err error, message string) {
	fmt.Println("hi!")
	type output struct {
		Err     string
		Message string
	}
	var i output
	i.Err = err.Error()
	i.Message = message
	encoded, errJ := json.Marshal(i)
	if errJ != nil {
		panic(errJ)
	}
	fmt.Println(string(encoded))
	os.Exit(1)
}

func main() {
	if (len(os.Args) < 1) {
		HandleErrorFMT(errors.New("invalid command or command not found"), "[GO] Invalid command or command not found.")
	}
	if (os.Args[1] == "pkinfo") {
		PrivateKeyInfo()
	} else {
		HandleErrorFMT(errors.New("invalid command or command not found"), "[GO] Invalid command or command not found.")
	}
}
