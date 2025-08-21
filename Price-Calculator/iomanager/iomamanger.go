package iomanager

type IOManager interface {
	ReadLines() ([]string, error)
	WriteLines(data interface{}) error
}