package grid

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type RuneGrid [][]rune

func (c RuneGrid) String() string {
	var sb strings.Builder

	for y, row := range c {
		for _, elem := range row {
			sb.WriteRune(elem)
		}
		if y < len(c)-1 {
			sb.WriteRune('\n')
		}
	}

	return sb.String()
}

func ParseRuneGrid(r io.Reader) (RuneGrid, error) {
	grid, err := parseGrid(r, func(r rune) (rune, error) {
		return r, nil
	})
	if err != nil {
		return RuneGrid{}, fmt.Errorf("Parser encountered an error: %w", err)
	}

	return grid, nil
}

func parseGrid[T any](r io.Reader, parse func(rune) (T, error)) ([][]T, error) {

	scanner := bufio.NewScanner(r)
	result := [][]T{}
	for scanner.Scan() {
		line := scanner.Text()
		row := []T{}

		for _, c := range line {
			parsed, err := parse(c)
			if err != nil {
				return [][]T{}, err
			}

			row = append(row, parsed)
		}

		result = append(result, row)
	}

	return result, nil
}
