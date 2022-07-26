// Package flaky is an example package with flaky tests.
package flaky

import (
	"crypto/rand"
	"math/big"
)

func Flake() bool {
	x, err := rand.Int(rand.Reader, big.NewInt(1000))
	if err != nil {
		panic(err)
	}
	return big.NewInt(0).Cmp(x) == 0
}
