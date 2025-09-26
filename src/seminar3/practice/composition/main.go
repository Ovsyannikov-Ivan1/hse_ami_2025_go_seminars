package main

import "fmt"

type Reader interface {
	Read([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type ReadCloser interface {
	Reader
	Closer
}

type File struct{}

func (f File) Read(data []byte) (int, error) {
	return len(data), nil
}

func (f File) Close() error {
	fmt.Println("File closed")
	return nil
}

func ProcessFile(rc ReadCloser) {
	data := make([]byte, 100)
	_, err := rc.Read(data)
	if err != nil {
		return
	}
	err = rc.Close()
	if err != nil {
		return
	}
}
