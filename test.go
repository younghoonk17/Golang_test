package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Printf("hello world")
	data, err := ioutil.ReadFile("./input.txt")
	check(err)
	string_data := string(data)
	split_data := strings.Split(string_data, "\n")

	fmt.Print(split_data)

}

func logic(input []string) string {

	//if it start with integer
	//if it only contains integer
	//number
	//else
	//verse
	//else if its blank
	//nothing
	//else
	//chorus

	return "hi"
}
