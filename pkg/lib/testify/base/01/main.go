package main

import (
	"math/rand"
	"time"
)

func Calculate(a int) int {
	return a + 2
}

func SimpleRand(start, end int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(end) + start
}
