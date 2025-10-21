package math3d

import "math"

// Vec3 represents a 3D vector
type Vec3 struct {
	X, Y, Z float64
}

// NewVec3 creates a new 3D vector
func NewVec3(x, y, z float64) Vec3 {
	return Vec3{X: x, Y: y, Z: z}
}

// Add performs vector addition
func (v Vec3) Add(other Vec3) Vec3 {
	return Vec3{v.X + other.X, v.Y + other.Y, v.Z + other.Z}
}

// Sub performs vector subtraction
func (v Vec3) Sub(other Vec3) Vec3 {
	return Vec3{v.X - other.X, v.Y - other.Y, v.Z - other.Z}
}

// Mul performs scalar multiplication
func (v Vec3) Mul(scalar float64) Vec3 {
	return Vec3{v.X * scalar, v.Y * scalar, v.Z * scalar}
}

// Dot computes the dot product
func (v Vec3) Dot(other Vec3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

// Cross computes the cross product
func (v Vec3) Cross(other Vec3) Vec3 {
	return Vec3{
		v.Y*other.Z - v.Z*other.Y,
		v.Z*other.X - v.X*other.Z,
		v.X*other.Y - v.Y*other.X,
	}
}

// Length returns the magnitude of the vector
func (v Vec3) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Normalize returns a unit vector in the same direction
func (v Vec3) Normalize() Vec3 {
	length := v.Length()
	if length == 0 {
		return Vec3{0, 0, 0}
	}
	return Vec3{v.X / length, v.Y / length, v.Z / length}
}

// Negate returns the negated vector
func (v Vec3) Negate() Vec3 {
	return Vec3{-v.X, -v.Y, -v.Z}
}
