package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

func main() {
	fmt.Println("Using writers and readers is fun!")
	fmt.Fprintln(os.Stdout, "Using writers and readers is fun!")
	io.WriteString(os.Stdout, "Using writers and readers is fun!")

	// Should try to figure out how to get this path without hardcoding:
	err := ioutil.WriteFile("./section-18-application/48-writer-interface/so-much-fun.txt", []byte("Using writers and readers is fun!"), 0100644)
	if err != nil {
		fmt.Println("There was an error writing to the file: ", err)
	}

	_, filename, _, ok := runtime.Caller(1)
	if ok {
		fmt.Println(path.Dir(filename))
	}
	fmt.Println(os.Getwd())
}
