package main

import (
	"flag"
	"fmt"
	"hash/crc64"
	"io"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		h := crc64.New(crc64.MakeTable(crc64.ECMA))
		_, err := io.Copy(h, os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
			return
		}
		fmt.Println(h.Sum64())
		return
	}
	var hasError bool
	for _, file := range flag.Args() {
		if !hashFile(file) {
			hasError = true
		}
	}
	if hasError {
		os.Exit(1)
	}
}

func hashFile(file string) bool {
	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, file, err)
		return false
	}
	defer f.Close()
	h := crc64.New(crc64.MakeTable(crc64.ECMA))
	_, err = io.Copy(h, f)
	if err != nil {
		fmt.Fprintln(os.Stderr, file, err)
		return false
	}
	fmt.Printf("%s: %d\n", file, h.Sum64())
	return true
}
