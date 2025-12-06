package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type outputable interface {
	saver
	Display()
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed")
		return err
	}

	fmt.Println("Saving the note succeeded")
	return nil
}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
}

func printSomething(value interface{}) {
	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("Integer:", floatVal)
		return
	}

	stringVal, ok := value.(string)

	if ok {
		fmt.Println("String:", stringVal)
		return
	}

	//switch value.(type) {
	//case int:
	//	fmt.Println("Integer", value)
	//
	//case float64:
	//	fmt.Println("Float", value)
	//
	//case string:
	//	fmt.Println(value)
	//}
	//
	//fmt.Println(value)
}

func main() {
	printSomething(1)
	printSomething(1.4)
	printSomething("Hello")
	title, content := getNoteData()
	todoText := getTodoData()

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(&todo)

	if err != nil {
		return
	}

	err = outputData(&userNote)

	if err != nil {
		return
	}
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}

func getTodoData() string {
	return getUserInput("Todo text:")
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n") //to delete enter
	text = strings.TrimSuffix(text, "\r") //to delete enter

	return text
}
