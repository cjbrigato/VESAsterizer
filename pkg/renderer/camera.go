package renderer

import (
	"math"

	"github.com/cjbrigato/VESAsterizer/pkg/math3d"
)

// Camera represents a 3D camera
type Camera struct {
	Position math3d.Vec3
	Target   math3d.Vec3
	Up       math3d.Vec3
	FOV      float64 // Field of view in radians
	Aspect   float64 // Aspect ratio (width/height)
	Near     float64 // Near clipping plane
	Far      float64 // Far clipping plane
}

// NewCamera creates a new camera with default settings
func NewCamera() *Camera {
	return &Camera{
		Position: math3d.NewVec3(0, 0, 5),
		Target:   math3d.NewVec3(0, 0, 0),
		Up:       math3d.NewVec3(0, 1, 0),
		FOV:      math.Pi / 3.0, // 60 degrees
		Aspect:   16.0 / 9.0,
		Near:     0.1,
		Far:      100.0,
	}
}

// ViewMatrix returns the view matrix for this camera
func (c *Camera) ViewMatrix() math3d.Mat4 {
	return math3d.LookAt(c.Position, c.Target, c.Up)
}

// ProjectionMatrix returns the projection matrix for this camera
func (c *Camera) ProjectionMatrix() math3d.Mat4 {
	return math3d.Perspective(c.FOV, c.Aspect, c.Near, c.Far)
}

// OrbitAround rotates the camera around the target point
func (c *Camera) OrbitAround(yaw, pitch float64) {
	// Calculate distance from target
	offset := c.Position.Sub(c.Target)
	distance := offset.Length()

	// Convert to spherical coordinates
	x := distance * math.Cos(pitch) * math.Sin(yaw)
	y := distance * math.Sin(pitch)
	z := distance * math.Cos(pitch) * math.Cos(yaw)

	c.Position = c.Target.Add(math3d.NewVec3(x, y, z))
}
//                T  --> PTCH+
//F---> YAW-      G-->PTCH-       H  --> YAW+
// RotateView rotates the camera's point of view (turn head) around its position.
// Positive yaw rotates to the right around the camera up vector.
// Positive pitch rotates up around the camera right vector.
func (c *Camera) RotateView(yaw, pitch float64) {
	// Direction from camera to target
	dir := c.Target.Sub(c.Position).Normalize()

	// Rotate direction by yaw around current up vector
	dir = rotateAroundAxis(dir, c.Up.Normalize(), yaw)

	// Recompute right vector after yaw
	right := dir.Cross(c.Up).Normalize()

	// Rotate direction by pitch around right vector
	dir = rotateAroundAxis(dir, right, pitch)

	// Recompute an orthonormal up vector and update target
	c.Up = right.Cross(dir).Normalize()
	c.Target = c.Position.Add(dir)
}

// rotateAroundAxis rotates vector v around an arbitrary axis by angle (radians)
// using Rodrigues' rotation formula. Axis is assumed to be non-zero.
func rotateAroundAxis(v, axis math3d.Vec3, angle float64) math3d.Vec3 {
	k := axis.Normalize()
	c := math.Cos(angle)
	s := math.Sin(angle)
	// v*cosθ + (k×v)*sinθ + k*(k·v)*(1-cosθ)
	term1 := v.Mul(c)
	term2 := k.Cross(v).Mul(s)
	term3 := k.Mul(k.Dot(v) * (1 - c))
	return term1.Add(term2).Add(term3)
}
