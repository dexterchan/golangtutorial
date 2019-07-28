package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")

	if err != nil {
		log.Println("err happened", err)
		log.Panic("panic", err)
		//panic("panic")
		log.Fatal("fatal error", err)

	}
	defer f2.Close()

	fmt.Println("Check the log.txt in the directory")
}
