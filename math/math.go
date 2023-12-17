package math

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Greatest Common Divisor
func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Least common multiple
func Lcm(a, b int) int {
	return a / Gcd(a, b) * b // The order of operations avoids potential overflow
}
