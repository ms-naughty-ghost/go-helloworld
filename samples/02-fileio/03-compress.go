package main

import (
	"compress/gzip"
	"compress/zlib"
	"io"
	"os"
)

func main() {
	dist, err := os.Create("hello.txt.gz")
	if err != nil {
		panic(err)
	}
	defer dist.Close()
	dist2, err := os.Create("hello.txt.zip")
	if err != nil {
		panic(err)
	}
	defer dist2.Close()

	// 第二引数で圧縮率を指定できる
	// 指定できる値については以下を参照
	// https://golang.org/pkg/compress/gzip/#pkg-constants
	gw, err := gzip.NewWriterLevel(dist, gzip.BestCompression)
	if err != nil {
		panic(err)
	}
	defer gw.Close()
	zw, err := zlib.NewWriterLevel(dist2, zlib.BestCompression)
	if err != nil {
		panic(err)
	}
	defer zw.Close()
	src, err := os.Open("hello.txt")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	if _, err := io.Copy(gw, src); err != nil {
		panic(err)
	}
	if _, err := io.Copy(zw, src); err != nil {
		panic(err)
	}
}
