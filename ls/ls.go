package main

import (
	"fmt"
	"log"
	"os"
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

/* Prints the names of every file in the current directory
 * Long version, so additional information is printed
 * TODO: Color code directories (IsDir() == true), just like the real `ls` does
 */
func printFileInfo(dirToRead string) {
	files, err := os.ReadDir(dirToRead)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name(), f.Type().IsDir(), f.Type().Perm())
	}
}

func main() {
	args := os.Args[1:]

	var defaultDir = true
	var longOutput = false

	for _, arg := range args {
		//		fmt.Println(arg)
		if arg[0] == '-' {
			fmt.Println("Parse the options!")
			if arg[1] == 'l' {
				longOutput = true
			}
		} else {
			if longOutput {
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
	}
}
