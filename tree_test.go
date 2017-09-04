package main

import (
	"fmt"
	"os"
)

func ExampleList() {
	os.Mkdir("./foobar", os.ModePerm)
	defer os.Remove("./foobar")
	os.Mkdir("./foobar/HelloWorld", os.ModePerm)
	defer os.Remove("./foobar/HelloWorld")
	os.Create("./foobar/HelloWorld.txt")
	defer os.Remove("./foobar/HelloWorld.txt")
	os.Create("./foobar/HelloWorld/fizzbuzz.txt")
	defer os.Remove("./foobar/HelloWorld/fizzbuzz.txt")
	output := make(chan *File, 5)
	go List("./foobar", output)
	for file := range output {
		fmt.Println(file.path)
	}
	// Output:
	// ./foobar
	// foobar/HelloWorld
	// foobar/HelloWorld/fizzbuzz.txt
	// foobar/HelloWorld.txt
}
