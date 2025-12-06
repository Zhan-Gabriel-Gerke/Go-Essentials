package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Content string `json:"text"`
}

func (todo *Todo) Display() {
	fmt.Printf("Your todo has the following Content:\n\n%v\n\n", todo.Content)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("No user input")
	}
	return Todo{
		Content: content,
	}, nil
}
