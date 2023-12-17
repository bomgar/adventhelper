package cardinaldirection

type CardinalDirection int

func (direction CardinalDirection) String() string {
	switch direction {
	case North:
		return "South"
	case East:
		return "West"
	case South:
		return "North"
	case West:
		return "East"
	}
	panic("switch not exhaustive")
}

const (
	North CardinalDirection = iota
	East
	South
	West
)

var AllDirections = []CardinalDirection{North, East, South, West}

// Converts the direction you're going to into the direction you're comming from
func (direction CardinalDirection) toSource() CardinalDirection {
	switch direction {
	case North:
		return South
	case East:
		return West
	case South:
		return North
	case West:
		return East
	}
	panic("switch not exhaustive")
}
