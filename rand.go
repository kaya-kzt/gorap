package gorap

import (
	"math/rand"
	"time"
)

// GetRandInt returns random int value
func GetRandInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()
}

// GetRandFloat64 returns random int value
func GetRandFloat64() float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()
}

// GetRandIntWithRange returns random int value between specific range
func GetRandIntWithRange(r int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(r)
}
