package main

import (
	"io/ioutil"
	"math/rand"
	"time"
)

func readFromFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func getRandIntInRange(randRange int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(randRange)
}
