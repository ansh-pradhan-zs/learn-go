package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main(){
	file, err := os.Open("test.txt");

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	fmt.Println("The file has", len(data), "bytes")
	file.Close()

}