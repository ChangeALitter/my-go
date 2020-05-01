package main

import "fmt"

var a  string = "G"

//  result: G O G
func main() {
	m()
	n()
	m()
}

func m() {
	fmt.Println(a)
}

func n()  {
	a := "O"
	fmt.Println(a)
}