package simple

import "fmt"

type Connection struct {
	*File
}

func (connection Connection) Close() {
	fmt.Println("Close connection" + connection.File.Name)
}

func NewConnection(file *File) (*Connection, func()) {
	connection := &Connection{File: file}
	return connection, func() {
		connection.Close()
	}
}
