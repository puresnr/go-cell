package crand

import (
	"github.com/puresnr/go-cell/celldef"
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

func RandIdxByWeight(weight []float64) celldef.INT {
	value := rand.Float64()

	for idx, v := range weight {
		if value < v {
			return idx
		}
	}

	return -1
}
