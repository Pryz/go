package main

import (
	"fmt"
	"path/filepath"
	"io/fs"
)

func main() {
	err := filepath.Walk("/", func(path string, info fs.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
	if err != nil {
		panic(err)
	}
}
