package simple

import "fmt"

type File struct {
	Name string
}

func (file *File) Close() {
	fmt.Println("Close file", file.Name)
}

func NewFile(name string) (*File, func()) {
	file := &File{Name: name}
	return file, func() {
		file.Close()
	}
}
