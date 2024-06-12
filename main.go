package main

import "fmt"

func Hello(name string) string {
	return "Hello, " + name + "!"
}

func main() {
	name := "world"
	fmt.Println(Hello(name))
}
