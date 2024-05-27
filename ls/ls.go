package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var DEBUG = false

/* Prints the names of every file in the current directory
 */
func printFileNames(dirToRead string) {
	files, err := os.ReadDir(dirToRead)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
}

func printFileInfo(dirToRead string) {
	files, err := os.ReadDir(dirToRead)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name(), f.Type().IsDir(), f.Type().Perm())
	}
}

func visitAllFiles() {
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path, info.Size())
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}

/* Walks recursively through the current file path
 * For each file in the path, print out the size of the file
 */
func betterVisitAllFiles() {
	err := filepath.WalkDir(".",
		func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}

			i, err := d.Info()
			if err != nil {
				return err
			}

			fmt.Println(path, i.Size())
			return nil
		})

	if err != nil {
		log.Println(err)
	}
}

func main() {
	args := os.Args[1:]

	var defaultDir = true
	var recursive = false

	for _, arg := range args {
		//		fmt.Println(arg)
		if arg[0] == '-' {
			fmt.Println("Parse the options!")
			if arg[1] == 'l' {
				recursive = true
			}
		} else {
			if recursive {
				printFileInfo(arg)
			} else {
				printFileNames(arg)
			}

			defaultDir = false
		}
	}

	if defaultDir {
		printFileNames(".")
	}

	if DEBUG {
		fmt.Println("testing functionality!")
		printFileNames(".")
		visitAllFiles()
		betterVisitAllFiles()
	}
}
