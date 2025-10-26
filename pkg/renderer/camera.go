package renderer

import (
	"math"

	"github.com/cjbrigato/VESAsterizer/pkg/math3d"
)

// Camera represents a 3D camera
type Camera struct {
	Position math3d.Vec3
	Rotation math3d.Vec3 // Euler angles in radians: X(pitch), Y(yaw), Z(roll)
	FOV      float64     // Field of view in radians
	Aspect   float64     // Aspect ratio (width/height)
	Near     float64     // Near clipping plane
	Far      float64     // Far clipping plane
}

// NewCamera creates a new camera with default settings
func NewCamera() *Camera {
	return &Camera{
		Position: math3d.NewVec3(0, 0, 5),
		Rotation: math3d.NewVec3(0, 0, 0),
		FOV:      math.Pi / 4.0, // 60 degrees
		Aspect:   16.0 / 9.0,
		Near:     0.1,
		Far:      100.0,
	}
}

// ViewMatrix returns the view matrix for this camera
func (c *Camera) ViewMatrix() math3d.Mat4 {
	// Build camera basis from Euler rotation
	rx := math3d.RotationX(c.Rotation.X)
	ry := math3d.RotationY(c.Rotation.Y)
	rz := math3d.RotationZ(c.Rotation.Z)
	rot := rz.Mul(ry).Mul(rx)

	// Forward is -Z in camera local space, Up is +Y
	f := rot.MulVec4(math3d.Vec3ToVec4(math3d.NewVec3(0, 0, -1), 0)).ToVec3().Normalize()
	u := rot.MulVec4(math3d.Vec3ToVec4(math3d.NewVec3(0, 1, 0), 0)).ToVec3().Normalize()
	center := c.Position.Add(f)
	return math3d.LookAt(c.Position, center, u)
}

// ProjectionMatrix returns the projection matrix for this camera
func (c *Camera) ProjectionMatrix() math3d.Mat4 {
	return math3d.Perspective(c.FOV, c.Aspect, c.Near, c.Far)
}

// OrbitAround rotates the camera around the target point
// Legacy orbit/rotate methods removed in favor of Rotation-based camera
