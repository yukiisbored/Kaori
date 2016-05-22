package common

import "math"

type Vector2D struct {
	X float64
	Y float64
}

func (v1 *Vector2D) Add(v2 Vector2D) {
	v1.X = v1.X + v2.X
	v1.Y = v1.Y + v2.Y
}

func (v1 *Vector2D) Sub(v2 Vector2D) {
	v1.X = v1.X - v2.X
	v1.Y = v1.Y - v2.Y
}

func (v1 *Vector2D) Mul(scalar float64) {
	v1.X = v1.X * scalar
	v1.Y = v1.Y * scalar
}

func (v1 *Vector2D) Div(scalar float64) {
	v1.X = v1.X / scalar
	v1.Y = v1.Y / scalar
}

func (v *Vector2D) Normalize() {
	l := v.Length()

	if l > 0 {
		v.Mul(1 / l)
	}
}

func (v *Vector2D) Length() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}
