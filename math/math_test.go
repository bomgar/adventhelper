package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGcd(t *testing.T) {

	assert.Equal(t, 5, Gcd(10, 5))
}

func TestLcm(t *testing.T) {
	assert.Equal(t, 21, Lcm(7, 3))
}
