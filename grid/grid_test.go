package grid

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseGrid(t *testing.T) {

	grid, err := ParseRuneGrid(strings.NewReader("AB\nCD"))
	assert.Nil(t, err)
	assert.Equal(t, RuneGrid{{'A', 'B'}, {'C', 'D'}}, grid)
}

func TestGridToString(t *testing.T) {

	grid := RuneGrid{{'A', 'B'}, {'C', 'D'}}

	assert.Equal(t, "AB\nCD", grid.String())
}
