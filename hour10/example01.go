package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	var file []byte
	var err error
	file, err = ioutil.ReadFile("README.md")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", file)
}
