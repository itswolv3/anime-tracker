package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type Anime struct {
	name       string
	episodeNum string
	directory  string
	episodes   []string
}

func main() {
	fmt.Println("vim-go")

	var dir string
	var objName string
	var ep string
	fmt.Println("Please enter dir:")
	fmt.Scan(&dir)
	fmt.Println("Please enter name:")
	fmt.Scan(&objName)
	fmt.Println("Episode:")
	fmt.Scan(&ep)

	x := getUserInfo(dir, objName, ep)
	fmt.Println(x.name)
	fmt.Println(x.episodeNum)
	fmt.Println(x.directory)
	fmt.Println(x.episodes)
}
func getUserInfo(dir, name, episode string) Anime {

	slc := []string{}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		slc = append(slc, file.Name())
	}

	item := Anime{name, episode, dir, slc}

	return item
}
