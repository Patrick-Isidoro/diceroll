package main

import (
	"fmt"
	"math/rand"
	"time"
)

const sides = 6

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		fmt.Println(rand.Intn(sides) + 1)
	}
}
