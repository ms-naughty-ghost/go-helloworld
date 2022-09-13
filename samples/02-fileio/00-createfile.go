package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Create("sample.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	w := io.Writer(f)
	_, err = w.Write([]byte("hello text!"))
	if err != nil {
		fmt.Println(err)
		return
	}

}
