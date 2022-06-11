package main

import "fmt"

func add(x int, y int) int { // int == 型指定
	return x + y
}

func main() {
	fmt.Println(add(10, 100))
}
