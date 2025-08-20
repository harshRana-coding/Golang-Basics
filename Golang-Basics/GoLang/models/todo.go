package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text   string `json:"text"`
}

func NewTodo(text string) (Todo, error) {
	if(text == "") {
		return Todo{}, errors.New("Invalid Input !")
	}
	return Todo{
		Text: text,
	}, nil
}

func (todo Todo) DisplayTodo() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error{
	json, err := json.Marshal(todo)
	if err !=nil {
		return err
	}
	return os.WriteFile("Todo.json", json, 0644)
}