package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath string
	OutputFilePath string
}

func (fm FileManager)ReadLines() ([]string,error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		fmt.Println("Error opening prices file:", err)
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines,scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		fmt.Println("Error reading prices file:", err)
		// file.Close()
		return nil, err
	}
	// file.Close()
	return lines, nil
}

func (fm FileManager)WriteLines(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("error creating file: " + err.Error())
	}

	defer file.Close()
	time.Sleep(3 * time.Second)
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		// file.Close()
		return errors.New("Failed to convert data to JSON: " + err.Error())
	}
	// file.Close()
	return nil
}

func New(inputPath , outputPath string) *FileManager {
	return &FileManager{
		InputFilePath: inputPath,
		OutputFilePath: outputPath,
	}
}