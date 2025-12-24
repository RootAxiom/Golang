package main

import "fmt"

func main() {
    var a int = 4
	fmt.Println("Integer variable a:", a)

	var b float64 = 9.99
	fmt.Println("Float variable b:", b)

	// type inference
	var c = "Hello, Go!"
	fmt.Println("String variable c:", c)	



	// ShortHand declaration
	d := "golang"
	fmt.Println("string variable d:", d)
} 