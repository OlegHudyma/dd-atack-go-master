package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"io/ioutil"
	"log"
	"math/rand"
)

func readFromFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func getRandIntInRange(randGenerator *rand.Rand, randRange int) int {
	return randGenerator.Intn(randRange)
}

type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) {}

func (s cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
