package main

import (
	"math/rand"
	"time"

	. "fmt"
)

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(s1)
	for i := 0; i < 100; i++ {
		randomNumber := rand.Intn(20)
		Printf("%d ", randomNumber)
	}
}
