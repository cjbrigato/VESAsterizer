package math3d

// Vec4 represents a 4D vector (homogeneous coordinates)
type Vec4 struct {
	X, Y, Z, W float64
}

// NewVec4 creates a new 4D vector
func NewVec4(x, y, z, w float64) Vec4 {
	return Vec4{X: x, Y: y, Z: z, W: w}
}

// Vec3ToVec4 converts a Vec3 to Vec4 (w=1 for positions, w=0 for directions)
func Vec3ToVec4(v Vec3, w float64) Vec4 {
	return Vec4{X: v.X, Y: v.Y, Z: v.Z, W: w}
}

// ToVec3 converts Vec4 to Vec3 (perspective divide if W != 1)
func (v Vec4) ToVec3() Vec3 {
	if v.W == 0 {
		return Vec3{v.X, v.Y, v.Z}
	}
	return Vec3{v.X / v.W, v.Y / v.W, v.Z / v.W}
}

// Add performs vector addition
func (v Vec4) Add(other Vec4) Vec4 {
	return Vec4{v.X + other.X, v.Y + other.Y, v.Z + other.Z, v.W + other.W}
}

// Mul performs scalar multiplication
func (v Vec4) Mul(scalar float64) Vec4 {
	return Vec4{v.X * scalar, v.Y * scalar, v.Z * scalar, v.W * scalar}
}
