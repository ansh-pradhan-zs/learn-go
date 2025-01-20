package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("test.txt")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if _,err := io.Copy(os.Stdout, file); err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	file.Close()
}