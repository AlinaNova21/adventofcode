package aoc

import (
	"log"
	"os"
)

type Input struct {
	file *os.File
}

func NewInput(filename string) *Input {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return &Input{file}
}

func (i *Input) Close() {
	i.file.Close()
}

func (i *Input) Read(p []byte) (int, error) {
	return i.file.Read(p)
}

type Output struct {
	Part1 interface{}
	Part2 interface{}
}

type Func func(input *Input) Output
