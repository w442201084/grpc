package main

import "fmt"

type Demo struct {
	Name string
}

func Test(demo *Demo) {
	demo.Name = "23"
}

func main() {
	demo := &Demo{}
	Test(demo)
	fmt.Println(demo)
}
