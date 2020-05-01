package main

import "fmt"

var a = "G"

// result: G O O
func main()  {
	m()
	n()
	m()
}

func m()  {
	fmt.Println(a)
}

func n()  {
	a = "O"
	fmt.Println(a)
}