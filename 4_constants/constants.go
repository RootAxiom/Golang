package main

import "fmt"

func main() {
	const name = "Golang"
	const version = 1.18
	const isOpenSource = true

	// name = "Python" // This will cause a compile-time error
	fmt.Println("Name:", name)
	fmt.Println("Version:", version)
	fmt.Println("Is Open Source:", isOpenSource)

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println("Host:", host)
	fmt.Println("Port:", port)
}
