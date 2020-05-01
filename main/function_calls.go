package main

import "fmt"

var a string


// result: G O G
func main()  {
	a = "G"
	fmt.Println(a)
	f1()
}

func f1()  {
	a := "O"
	fmt.Println(a)
	f2()
}

func f2()  {
	fmt.Println(a)
}