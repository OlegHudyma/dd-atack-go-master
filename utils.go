package main

import (
	crypto_rand "crypto/rand"
	"io/ioutil"
	"math/rand"
)

func readFromFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func getRandIntInRange(randRange int) (int, error) {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		return 0, err
	}
	return rand.Intn(randRange), err
}
