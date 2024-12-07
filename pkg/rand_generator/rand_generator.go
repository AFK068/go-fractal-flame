package randgenerator

import (
	"crypto/rand"
	"math/big"
)

func Float64() float64 {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(int64(1<<53)))

	return float64(nBig.Int64()) / float64(1<<53)
}

func Int64(maximum int64) int64 {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(maximum))

	return nBig.Int64()
}

func Uint8(maximum uint8) uint8 {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(int64(maximum)))

	return uint8(nBig.Uint64()) //nolint
}
