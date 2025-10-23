package math3d

import "math"

// Mat4 represents a 4x4 matrix in row-major order
type Mat4 [16]float64

// Identity returns an identity matrix
func Identity() Mat4 {
	return Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

// Mul multiplies two matrices
func (m Mat4) Mul(other Mat4) Mat4 {
	var result Mat4
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			sum := 0.0
			for k := 0; k < 4; k++ {
				sum += m[row*4+k] * other[k*4+col]
			}
			result[row*4+col] = sum
		}
	}
	return result
}

// MulVec4 multiplies a matrix by a Vec4
func (m Mat4) MulVec4(v Vec4) Vec4 {
	return Vec4{
		m[0]*v.X + m[1]*v.Y + m[2]*v.Z + m[3]*v.W,
		m[4]*v.X + m[5]*v.Y + m[6]*v.Z + m[7]*v.W,
		m[8]*v.X + m[9]*v.Y + m[10]*v.Z + m[11]*v.W,
		m[12]*v.X + m[13]*v.Y + m[14]*v.Z + m[15]*v.W,
	}
}

// Translation creates a translation matrix
func Translation(x, y, z float64) Mat4 {
	return Mat4{
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	}
}

// Scale creates a scaling matrix
func Scale(x, y, z float64) Mat4 {
	return Mat4{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	}
}

// RotationX creates a rotation matrix around the X axis
func RotationX(angle float64) Mat4 {
	c := math.Cos(angle)
	s := math.Sin(angle)
	return Mat4{
		1, 0, 0, 0,
		0, c, -s, 0,
		0, s, c, 0,
		0, 0, 0, 1,
	}
}

// RotationY creates a rotation matrix around the Y axis
func RotationY(angle float64) Mat4 {
	c := math.Cos(angle)
	s := math.Sin(angle)
	return Mat4{
		c, 0, s, 0,
		0, 1, 0, 0,
		-s, 0, c, 0,
		0, 0, 0, 1,
	}
}

// RotationZ creates a rotation matrix around the Z axis
func RotationZ(angle float64) Mat4 {
	c := math.Cos(angle)
	s := math.Sin(angle)
	return Mat4{
		c, -s, 0, 0,
		s, c, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

// Perspective creates a perspective projection matrix
// fov: field of view in radians
// aspect: aspect ratio (width/height)
// near, far: near and far clipping planes
func Perspective(fov, aspect, near, far float64) Mat4 {
	f := 1.0 / math.Tan(fov/2.0)
	nf := 1.0 / (near - far)

	return Mat4{
		f / aspect, 0, 0, 0,
		0, f, 0, 0,
		0, 0, (far + near) * nf, (2 * far * near) * nf,
		0, 0, -1, 0,
	}
}

// LookAt creates a view matrix
// eye: camera position
// center: point to look at
// up: up vector
func LookAt(eye, center, up Vec3) Mat4 {
	f := center.Sub(eye).Normalize()
	s := f.Cross(up).Normalize()
	u := s.Cross(f)

	return Mat4{
		s.X, s.Y, s.Z, -s.Dot(eye),
		u.X, u.Y, u.Z, -u.Dot(eye),
		-f.X, -f.Y, -f.Z, f.Dot(eye),
		0, 0, 0, 1,
	}
}

// Viewport creates a viewport transformation matrix
// x, y: viewport origin
// width, height: viewport dimensions
func Viewport(x, y, width, height float64) Mat4 {
	halfW := width / 2.0
	halfH := height / 2.0

	return Mat4{
		halfW, 0, 0, x + halfW,
		0, -halfH, 0, y + halfH,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}
