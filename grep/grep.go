package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var DEBUG = false

// A tester function to experiment with regular expression handling
func tester() {
	m, e := regexp.Match(`c*`, []byte("This is a test"))
	fmt.Println(m, e)

	re := regexp.MustCompile(`This..`)
	var output = re.Find([]byte(`This is a test`))
	if output == nil {
		fmt.Println("Match not found")
	}
	fmt.Printf("%q\n", output)
}

// Determines whether this program is being used as part of a pipe
func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

// Primary program loop
func main() {
	if DEBUG {
		tester()
	}

	args := os.Args[1:]

	var pattern = ""
	var file = ""

	for _, arg := range args {
		if arg[0] == '-' {
			fmt.Println("Options not implemented yet")
		} else if pattern == "" {
			pattern = arg
		} else if file == "" {
			file = arg
		} else {
			fmt.Println("Extra argument:", arg)
		}
	}

	var matched bool
	var err error

	if isInputFromPipe() {
		scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
		for scanner.Scan() {
			// Matching with one method that just says true or false:
			/*
				matched, err = regexp.Match(pattern, []byte(scanner.Text()))
				fmt.Println("Match?", matched)
			*/

			// Matching with a method that says what part of the input matches:
			re := regexp.MustCompile(pattern)
			var output = re.Find([]byte(scanner.Text()))
			if output == nil || string(output) == "" {
				fmt.Println("Match not found")
			} else {
				fmt.Printf("Match: %q\n", output)
			}
		}
	} else {
		if pattern == "" || file == "" {
			fmt.Println("Insufficient arguments")
		}
		matched, err = regexp.Match(pattern, []byte(file))
		fmt.Println(matched)
	}

	fmt.Println("Errors:", err)
}
