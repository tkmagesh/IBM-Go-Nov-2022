package main

import "fmt"

type MyStr string

func (s MyStr) Length() int {
	return len(s)
}

func main() {
	str := MyStr("This is an example string")
	fmt.Println(str.Length())
}
