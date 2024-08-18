package crand

import (
	"math/rand"
	"runtime"
	"time"
)

func init() {
	if runtime.Version() >= "go1.20" {
		return
	}
	rand.Seed(time.Now().UnixNano())
}

func RandIdxByWeight(weight []float64) int {
	value := rand.Float64()

	for idx, v := range weight {
		if value < v {
			return idx
		}
	}

	return -1
}
