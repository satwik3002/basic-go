package main

import "fmt"

const myConst = 3.14

func main() {
	const myString = "世界"
	fmt.Println("Hello", myString)
	fmt.Println("Happy", myConst, "Day")

	const myBool = true
	fmt.Println("Go rules?", myBool)
}
