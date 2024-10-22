package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// user input
	fmt.Println("Enter name")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("hello", name)

	// fmt.Println("welcome to go")
	// var name string = "lovish"
	// fmt.Println(name)
	// var version int = 1
	// var currency = 42
	// fmt.Println(version)
	// fmt.Println(currency)
	// var dim float64 = 43.424
	// fmt.Println(dim)
	// var sscc bool = true
	// fmt.Println(sscc)
	// const pi = 43.32
	// fmt.Println(pi)
	// name := "Lovish Goyal"
	// // popular method to write variable in Golang
	// fmt.Println(name)
	// age := 21
	// weight := 82.32
	// fmt.Println(age, weight)
	// var Publicvar = "global var"
	// fmt.Println(Publicvar)

	// age := 21
	// name := "lovish"
	// weight := 82.22786686
	// fmt.Printf("age is %T\n", age)
	// fmt.Printf("name is %T\n", name) // printf specifies the specific var
	// fmt.Printf("weigth is %T", weight)

}
