package direction

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func (direction Direction) String() string {
	switch direction {
	case Right:
		return "R"
	case Left:
		return "L"
	case Up:
		return "U"
	case Down:
		return "D"
	}
	panic("switch not exhaustive")
}

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
