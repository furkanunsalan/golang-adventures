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

// type displayer interface {
// 	Display() error
// }

// type outputable interface {
// 	Save() error
// 	Display()
// }

type outputable interface {
	saver
	Display()
	// DoSomething(string) int
	// anotherInterface
}

func main() {
	// will accept the values and display them correctly
	printSomething(1)
	printSomething(2.5)
	printSomething("test")

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	// will accept the value but won't do anything, without crashing
	printSomething(todo)

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	

	outputData(userNote)
}

// any value is allowed
func printSomething(value interface{}) {
	typedVal, ok := value.(int) 
	// typedVal -> value saved with the type, var typedVal int
	// ok -> boolean value, whether the value is the asked type or not

	if !ok {
		fmt.Println(ok)
	}

	// value + 1 -> won't work because the type of value is not known (any)
	// typedVal + 1 -> works because now we know it's int

	fmt.Println(typedVal + 1)

	switch value.(type) { // only works in switch
	case int:
		fmt.Println("int: ", value)
	case float64:
		fmt.Println("float: ", value)
	case string:
		fmt.Println("str: ", value)
	// default:
		// ...
	}


	fmt.Println(value)
}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed!")
		return err
	}

	fmt.Println("Note Saved!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title: ")
	content := getUserInput("Note Content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	// we can't use the following method for long text inputs longer than single words
	// var value string
	// fmt.Scanln(&value)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
