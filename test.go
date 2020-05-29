package main

import (
	"fmt"
	"path/filepath"
	"os"
)

type yeet struct {
	tekashi []string
	name string
}
func main() {
	tekashi := yeet{}
	tekashi.name = "69"

	directories, err := WalkMatch("/", "*")
	fmt.Println("Done!")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(len(directories))
	}

	fmt.Println("Done")
}

func WalkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// if err != nil {
		// 	return err
		// }
		// if info.IsDir() {
		// 	matches = append(matches, path)
		// }
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched{
			fi, err := os.Stat(path)
			if err != nil {
				fmt.Println("Skip")
			} else if fi.IsDir() {
				matches = append(matches, path)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}
