package main

import "fmt"

// assigned an alias to the string built in type
type customString string

// adding methods to custom type, because we can't use it on local type
func (text customString) log() {
	fmt.Println(text)
}

func (text customString) hello() {
	fmt.Println("Hello,", text)
}

func main() {
	var name customString
	name = "Nishant"

	name.log()
	name.hello()
}
