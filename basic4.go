package main

import "fmt"

func main() {
	a := "hello"
	b := "world"
	a, b = b, a
	fmt.Println(a, b)
}
