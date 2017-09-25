package main

import (
	"os"
	"path/filepath"
)

type File struct {
	path string
	info os.FileInfo
	err  error
}

func List(root string, output chan *File) {
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() == false {
			return nil
		}
		output <- &File{path, info, err}
		return nil
	})
	close(output)
}
