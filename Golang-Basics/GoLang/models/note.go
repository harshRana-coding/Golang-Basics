package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewNote(title, content string) (Note, error) {
	if(title == "" || content=="") {
		return Note{}, errors.New("Invalid Input !")
	}
	return Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) DisplayNote() {
	fmt.Printf("Your note titled %s has following contents : \n\n%v", note.Title,note.Content)
}

func (note Note) Save() error{
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)
	if err !=nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}