package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type Anime struct {
	name      string
	episode   int
	directory string
}

func main() {
	fmt.Println("vim-go")

	var dir string
	var objName string
	fmt.Println("Please enter dir:")
	fmt.Scan(&dir)
	fmt.Println("Please enter name:")
	fmt.Scan(&objName)

	newFile := getUserInfo(dir, objName)
	fmt.Println(newFile)
}

func getUserInfo(dir, name string) []string {

	slc := []string{}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		slc = append(slc, file.Name())
	}

	return slc
}
