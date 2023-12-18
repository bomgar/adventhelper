package direction

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func FromCapitalLetter(r rune) Direction {
	switch r {
	case 'R':
		return Right
	case 'L':
		return Left
	case 'U':
		return Up
	case 'D':
		return Down
	}
	panic("Unknown direction")
}
