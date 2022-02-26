package main

import (
	"math/rand"
	"os"
	"time"
)

func readFromFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func getRandIntInRange(randRange int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(randRange)
}
