// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// file line count
	fileCounts := make(map[string]map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		fileCounts["stdin"] = make(map[string]int)
		countLines(os.Stdin, fileCounts["stdin"])
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			fileCounts[arg] = make(map[string]int)
			countLines(f, fileCounts[arg])
			f.Close()
		}
	}
	for fileName, line := range fileCounts {
		for ehhLine, n := range line {
			if n > 1 {
				fmt.Printf("%s\t%d\t%s\n", fileName, n, ehhLine)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

// jdeli@TRS-80 MSYS /c/RemoteRepos/gopl.io/ch1_exercises/1.4/dup2 (master)
// $ go run main.go ./test1 ./test2
// ./test2 2       this
// ./test2 2       is a stupid
// ./test2 2       test file
// ./test2 4       hello
// ./test1 2       hello

// jdeli@TRS-80 MSYS /c/RemoteRepos/gopl.io/ch1_exercises/1.4/dup2 (master)
// $ go run main.go < test1
// stdin   2       hello

// jdeli@TRS-80 MSYS /c/RemoteRepos/gopl.io/ch1_exercises/1.4/dup2 (master)
// $ cat test1
// hello
// this
// is a stupid
// test file
// hello

// jdeli@TRS-80 MSYS /c/RemoteRepos/gopl.io/ch1_exercises/1.4/dup2 (master)
// $ cat test2
// hello
// this
// is a stupid
// test file
// hello
// hello
// this
// is a stupid
// test file
// hello

// jdeli@TRS-80 MSYS /c/RemoteRepos/gopl.io/ch1_exercises/1.4/dup2 (master)
// $ go run main.go ./test1 ./test2
// ./test1 2       hello
// ./test2 2       test file
// ./test2 4       hello
// ./test2 2       this
// ./test2 2       is a stupid

// jdeli@TRS-80 MSYS /c/RemoteRepos/gopl.io/ch1_exercises/1.4/dup2 (master)
// $ go run main.go < test1
// stdin   2       hello

// jdeli@TRS-80 MSYS /c/RemoteRepos/gopl.io/ch1_exercises/1.4/dup2 (master)
// $
