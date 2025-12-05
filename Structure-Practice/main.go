package main

import (
	"errors"
	"fmt"
)

func main() {
	title, content, err := getNoteData()

	if err != nil {
		fmt.Println(err)
		return
	}
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Note title:")

	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	content, err := getUserInput("Note content:")

	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	return title, content, nil
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var value string
	fmt.Scan(&value)

	if value == "" {
		return "", errors.New("No user input")
	}

	return value, nil
}
