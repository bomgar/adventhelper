package point

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManhattanDistance(t *testing.T) {

	p1 := Point{1, 2}
	p2 := Point{10, 3}
	assert.Equal(t, 10, ManhattanDistance(p1, p2))
}
