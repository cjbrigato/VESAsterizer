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
