package main

import "fmt"

var a = 10
const b = 100

func main() {
	if true {
		a := 11
		const b = 101
		fmt.Println(a, b)
		if a > 10 {
			var a = 12
			fmt.Println(a, b)
		}
	}
	fmt.Println(a, b)
}