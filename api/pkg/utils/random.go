package utils

import (
	"math/rand"
	"time"
)

const (
	roomIDSize = 10
)

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandSeq is well named
func RandSeq() string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, roomIDSize)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
