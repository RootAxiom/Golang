package main

import "fmt"

func main(){
   // if-else statement
//    a := 21
//    b := 20

//    if a > b {
//        println("a is greater than b")
//    } else {
//        println("b is greater than a")
//    }
// }

var role ="admin"
var haspermissions = true

if role == "admin" && haspermissions {
    fmt.Println("User is an admin and he/she can access the system.")
} else {
    fmt.Println("User does not have admin permissions.")
}

}