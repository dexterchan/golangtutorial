package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	checkFileIO()
}
func simpleCheck() {
	n, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(n)
}

func checkInput() {
	var answer1, answer2, answer3 string
	fmt.Print("Name: ")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Food: ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Sport:")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		panic(err)
	}
	fmt.Println(answer1, answer2, answer3)
}

func checkFileIO() {
	f, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))
}
