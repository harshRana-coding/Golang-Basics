package note

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/investment-calculator/models"
)

type saver interface {
	Save() error
}
// type displayer interface {
// 	Display()
// }

type outputtable interface{
	saver
	Display()
}

func Init(){
	printSomething(5);
	printSomething(5.5)
	title, content:= getNoteData()
	text := getUserInput("Todo Text: ")
	todo, err := models.NewTodo(text)
	userNote, err := models.NewNote(title,content)
	if err!=nil {
		fmt.Println(err)
		return
	}
	todo.DisplayTodo()
	err = saveData(todo)
	if err!=nil {
		fmt.Println("Saving the todo failed!");
		return
	}
	fmt.Println("Saving the todo succeeded!");
	userNote.DisplayNote();
	err = saveData(userNote)
	if err!=nil {
		fmt.Println("Saving the todo failed!");
		return
	}
}

func printSomething(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Integer: ",value)
	case float64:
		fmt.Println("Float: ",value)
	}

	fmt.Println(value)
}

func outputData(data outputtable) {
	data.Display();
	saveData(data);
}

func saveData( data saver) error {
	err := data.Save()
	if err!=nil {
		fmt.Println("Saving the note failed!");
		return err
	}
	fmt.Println("Saving the note succeeded!");
	return nil
}

func getNoteData( ) (title string, content string ){
	title= getUserInput("Note title: ")
	content= getUserInput("Note content: ")
	return title, content
}

func getUserInput(prompt string) (value string) {
	fmt.Print(prompt)
	
	rd := bufio.NewReader(os.Stdin)
	value,err := rd.ReadString('\n')
	if err!= nil {
		return ""
	}

	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value,"\r")
	return value
}