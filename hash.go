package main

import (
	"crypto/md5"
	"encoding/base64"
	"io"
	"os"
)

type HashFile struct {
	path string
	info os.FileInfo
	hash string
	err  error
}

func Hash(input chan *File, output chan *HashFile) {
	done := make(chan int, 0)
	count := 0
	for file := range input {
		go handleFile(file, output, done)
		count++
	}
	for ; count > 0; count-- {
		<-done
	}
	close(output)
}

func handleFile(file *File, output chan *HashFile, done chan int) {
	defer func() {
		done <- 1
	}()
	if file.err != nil {
		output <- &HashFile{file.path, file.info, "", file.err}
		return
	}
	if file.info.Mode().IsRegular() == false {
		output <- &HashFile{file.path, file.info, "", nil}
		return
	}

	f, e := os.Open(file.path)
	defer f.Close()
	if e != nil {
		output <- &HashFile{file.path, file.info, "", e}
		return
	}

	hasher := md5.New()
	io.Copy(hasher, f)
	hash := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	output <- &HashFile{file.path, file.info, hash, nil}
	return
}
