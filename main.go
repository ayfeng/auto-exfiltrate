package main

import (
	"fmt"
	"path/filepath"
)


type dataDump struct {
	dataLD []string
	entireStruct string
	output string
	root bool
}

// -- Setters --
func (d *dataDump) SetDataLD(data []string) {
	d.dataLD = data
}

func (d *dataDump) SetEntireStruct(entirity string) {
	d.entireStruct = entirity
}

func (d *dataDump) SetOutput(output string) {
	d.output = output
}

func (d *dataDump) SetRoot(root bool) {
	d.root = root
}
// -- Setters --


func main() {
	model := dataDump{}
	directories := []string{
		"/home",
		"....",
	}
	filetypes := []string{}

	fetchFiles(directories, filetypes, &model)
}

func fetchFiles(directories []string, filetypes []string, m *dataDump) {
	fmt.Println("INSERT HERE")
}

// The full program will stay private and not pushed here, due to the potential of scammed that way
// I will keep you updated on the program by sending you videos
// After everything has been completed, and our transaction has been fullfilled. I will push everything here.
