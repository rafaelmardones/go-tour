package helpers

import (
	"math/rand"
	"time"
)

func RandomNumber(numPool int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(numPool)
}
