package zxcvbn_math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNchoseK100and2(t *testing.T) {
	nCk := NChoseK(100, 2)

	assert.Equal(t, float64(4950), nCk, "100 chose 2 should equal 4950")
}

func TestNChoseKwereNlessK(t *testing.T) {
	nCk := NChoseK(1, 2)

	assert.Equal(t, float64(0), nCk, "When n is less than k always 0")
}

func TestNChoseKWereKis0(t *testing.T) {
	nCk := NChoseK(50, 0)

	assert.Equal(t, float64(1), nCk, "When K is 0 then 1")
}
