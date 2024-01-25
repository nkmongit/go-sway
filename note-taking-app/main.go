package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

// creating an interface
type saver interface {
	Save() error
}

// type displayer interface {
//	Display()
// }

// embedded interface
type outputtable interface {
	saver
	Display()
}

// type outputtable interface {
// 	Save() error
//	Display()
// }

func main() {
	// printing any value with same function
	printSomething(1)
	printSomething(0.9)
	printSomethingAgain("Nishant")
	title, content := getNoteData()

	todoText := getUserInput("Todo text: ")

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

	err = outputData(todo)
	if err != nil {
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func add(a, b interface{}) any {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		return (aInt + bInt)
	}
	return nil
}

func printSomething(value interface{}) {
	switch value.(type) {
	case string:
		fmt.Println("Integer:", value)
	case float64:
		fmt.Println("Float:", value)
	default:
		fmt.Printf("%T:%v\n", value, value)
	}
}

func printSomethingAgain(value any) {
	fmt.Println(value)
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

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(promt string) string {
	fmt.Print(promt)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
