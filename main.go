package main

import (
	"fmt"
	"path/filepath"
	"os"
	"archive/tar"
	"log"
	"io"
	"compress/gzip"
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
	model.SetOutput("/tmp/")

	go fetchDirectories(directories, &model)
	fetchFiles(filetypes, &model)

	fmt.Println("Filetypes retrieved: ", len(model.dataLD))
	fmt.Println("Files from specified dirs retrieved: ", len(model.dirFiles))
	fmt.Println("Tree structure retrieved: ", len(model.entireStruct))

	mergeOperation(&model)
}


// -- Operations --
func mergeOperation(m *dataDump) {
	// Here we are making the final compressed file
	var denied int
	var totalBytes int
	treeLoc := "/tmp/getmytree"

	file, err := os.Create(m.output + "output.tar.gz")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	// Setting up the gzip writer
	gw := gzip.NewWriter(file)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	// Adding files
	for i := range m.dataLD {
		if err := addFile(tw, m.dataLD[i]); err != nil {
			denied++
		}
	}
	// Adding directories
	// You can enable this if you want, but it will take a long time to compress all the files
	// for i := range m.dirFiles {
	// 	if err := addFile(tw, m.dirFiles[i]); err != nil {
	// 		denied++
	// 	}
	// }
	fmt.Printf("Didn't have permissions for %d files\n", denied)

	// Writing the tree structure
	f, err := os.Create(treeLoc)
	if err != nil {
		panic(err)
	}
	for i := range m.entireStruct {
		n3, err := f.WriteString(m.entireStruct[i] + "\n")

		switch {
		case err != nil:
			fmt.Println("Writing error in structure, ", err)
		case err == nil:
			totalBytes += n3
		}
	}
	if err := addFile(tw, treeLoc); err != nil {
		fmt.Println("No success in writing the tree to tar")
	}

	fmt.Printf("Tree file was %d bytes\n", totalBytes)
	fmt.Println("Done!")
}

func addFile(tw * tar.Writer, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if stat, err := file.Stat(); err == nil {
		// now lets create the header as needed for this file within the tarball
		header := new(tar.Header)
		header.Name = path
		header.Size = stat.Size()
		header.Mode = int64(stat.Mode())
		header.ModTime = stat.ModTime()
		// write the header to the tarball archive
		if err := tw.WriteHeader(header); err != nil {
			return err
		}
		// copy the file data to the tarball
		if _, err := io.Copy(tw, file); err != nil {
			return err
		}
	}
	return nil
}

func fetchFiles(filetypes []string, m *dataDump) {
	// Here we are fetching the important files specified above
	fmt.Println("Fetching important files")
	for index, file := range filetypes {
		files, err := WalkMatch("/", "*" + file, false)
		fmt.Printf("Index: %d, conf: %v, Files: %d, error: %v\n", index, file, len(files), err)
		m.SetDataLD(files)
	}
	fmt.Println("Fetching files done!")
}

func fetchDirectories(directories []string, m *dataDump) {
	// Here we are fetching the important directories specified above
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
	// This functions as a gateway for the above 2 functions to access the filesystem
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
