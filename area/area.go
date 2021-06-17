package area

import "math"

const Pi = math.Pi

func Circle(radius float64) float64 {
	return math.Pow(radius, 2) * Pi
}

func Rectangle(width, height float64) float64 {
	return width * height
}

// Not visible
func _EquilateralTriangle(base, height float64) float64 {
	return (base * height) / 2
}
