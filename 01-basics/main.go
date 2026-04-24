package main

import "fmt"

func main() {
	DoDefer() // 1 2 3
}

func DoDefer() {
	defer fmt.Println("3")
	defer fmt.Println("2")
	fmt.Println("1")
}
