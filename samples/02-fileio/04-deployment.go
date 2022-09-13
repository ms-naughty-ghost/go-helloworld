package main

import (
	"compress/gzip"
	"io"
	"os"
)

func main() {
	dist, err := os.Create("hello.txt")
	if err != nil {
		panic(err)
	}
	defer dist.Close()
	// dist2, err := os.Create("hello2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer dist2.Close()

	src, err := os.Open("hello.txt.gz")
	if err != nil {
		panic(err)
	}
	defer src.Close()
	// src2, err := os.Open("hello.txt.zip")
	// if err != nil {
	// 	panic(err)
	// }
	// defer src2.Close()

	gr, err := gzip.NewReader(src)
	if err != nil {
		panic(err)
	}
	defer gr.Close()

	if _, err := io.CopyN(dist, gr, 1024); err != nil {
		panic(err)
	}
	// zr, err := zlib.NewReader(src2)
	// if err != nil {
	// 	panic(err)
	// }
	// defer zr.Close()

	// if _, err := io.Copy(dist2, zr); err != nil {
	// 	panic(err)
	// }
}
