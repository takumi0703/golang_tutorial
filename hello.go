package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(math.pi)=> error
	fmt.Println(math.Pi) // => 大文字から始まるものが、外部パッケージから参照できる
}
