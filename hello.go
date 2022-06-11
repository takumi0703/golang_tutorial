package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Seedを設定しないと、乱数が生成されない
	fmt.Println("my favorite number is", rand.Intn(10))
}
