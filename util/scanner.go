package util

import (
	"bufio"
	"fmt"
	"os"
)

type FileScanner struct {
	file    *os.File
	scanner *bufio.Scanner
}

func NewFileScanner(path string) *FileScanner {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return &FileScanner{
		file:    file,
		scanner: bufio.NewScanner(file),
	}
}

func NewScannerForInput(day int, readSample bool) *FileScanner {
	file := "input"
	if readSample {
		file = "sample"
	}

	return NewFileScanner(fmt.Sprintf("input/day%02d/%s.txt", day, file))
}

func (f *FileScanner) Close() {
	err := f.file.Close()
	if err != nil {
		panic(err)
	}
}

func (f *FileScanner) Scan() bool {
	return f.scanner.Scan()
}

func (f *FileScanner) Text() string {
	return f.scanner.Text()
}
