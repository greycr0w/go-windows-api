package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/greycr0w/go-windows-api/background"
)

func processPathArg(args []string) (string, error) {
	if len(args) <= 1 {
		return "", errors.New("[-] You need to provide a path argument")
	}

	if args[1] == "" {
		return "", errors.New("[-] You need to provide a path argument that is not empty")
	}

	path := args[1]

	//Check if the path provided is valid
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		errorMsg := fmt.Sprintf("[-] %s does not exist", path)
		return "", errors.New(errorMsg)
	}
	return path, nil
}

func main() {
	path, err := processPathArg(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	background.ChangeBackground(path)
}
