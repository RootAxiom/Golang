package main


import "fmt"

// for --> loop in Go is used to execute a block of code repeatedly until a specified condition is met.
// In Go, the only loop construct available is the 'for' loop, which can be used in several different ways.
func main(){
     // while style loop
	 fmt.Println("While style loop:")
	 i := 1

	 for i <= 5 {
		 fmt.Println(i)
		 i++
	 }

	 // infinite loop
	//  for {
	// 	 fmt.Println("Infinite loop iteration")
	//  }

     // traditional for loop
	 fmt.Println("Traditional for loop:")
	 for j := 0; j <= 5; j++ {
		 
		//  if j == 1 {
		// 	 continue // skip the rest of the loop when j is 3
		//  }
		//  if j == 4{
		// 	 break // exit the loop when j is 4
		// }
		 fmt.Println(j)
	 }

	 // for range loop
	 fmt.Println("For range loop:")
	 arr := []string{"apple", "banana", "cherry"}
	 for index, value := range arr {
		 fmt.Printf("Index: %d, Value: %s\n", index, value)
	 }
}