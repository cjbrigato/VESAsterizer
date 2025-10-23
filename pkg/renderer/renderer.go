package renderer

import (
	"math"

	"github.com/cjbrigato/VESAsterizer/pkg/loader"
	"github.com/cjbrigato/VESAsterizer/pkg/math3d"
	"github.com/cjbrigato/VESAsterizer/pkg/terminal"
)

// RenderMode defines how to render the mesh
type RenderMode int

const (
	Wireframe RenderMode = iota
	Solid
	SolidWireframe
)

// Renderer handles the rasterization of 3D meshes to a framebuffer
type Renderer struct {
	fb   *terminal.Framebuffer
	mode RenderMode
}

// MeshInstance represents a mesh placed in world space with TRS
type MeshInstance struct {
	Mesh     *loader.Mesh
	Position math3d.Vec3
	Rotation math3d.Vec3 // Euler angles in radians: X(pitch), Y(yaw), Z(roll)
	Scale    math3d.Vec3
}

// NewMeshInstance creates a new instance with default transforms
func NewMeshInstance(mesh *loader.Mesh) *MeshInstance {
	return &MeshInstance{
		Mesh:     mesh,
		Position: math3d.NewVec3(0, 0, 0),
		Rotation: math3d.NewVec3(0, 0, 0),
		Scale:    math3d.NewVec3(1, 1, 1),
	}
}

// ModelMatrix builds the model matrix from TRS
func (mi *MeshInstance) ModelMatrix() math3d.Mat4 {
	t := math3d.Translation(mi.Position.X, mi.Position.Y, mi.Position.Z)
	rx := math3d.RotationX(mi.Rotation.X)
	ry := math3d.RotationY(mi.Rotation.Y)
	rz := math3d.RotationZ(mi.Rotation.Z)
	s := math3d.Scale(mi.Scale.X, mi.Scale.Y, mi.Scale.Z)

	r := rz.Mul(ry).Mul(rx)
	return t.Mul(r).Mul(s)
}

// NewRenderer creates a new renderer
func NewRenderer(fb *terminal.Framebuffer) *Renderer {
	return &Renderer{
		fb:   fb,
		mode: Wireframe,
	}
}

// SetRenderMode sets the rendering mode
func (r *Renderer) SetRenderMode(mode RenderMode) {
	r.mode = mode
}

// Clear clears the framebuffer
func (r *Renderer) Clear() {
	r.fb.Clear()
}

// World represents a collection of instances and a camera
type World struct {
	Camera    *Camera
	Instances []*MeshInstance
}

// NewWorld creates a new world with the given camera
func NewWorld(camera *Camera) *World {
	return &World{Camera: camera, Instances: make([]*MeshInstance, 0)}
}

// AddInstance adds a mesh instance to the world
func (w *World) AddInstance(mi *MeshInstance) {
	w.Instances = append(w.Instances, mi)
}

// RenderWorld renders all instances in the world from the world's camera POV
func (r *Renderer) RenderWorld(w *World) {
	view := w.Camera.ViewMatrix()
	projection := w.Camera.ProjectionMatrix()
	viewport := math3d.Viewport(0, 0, float64(r.fb.Width), float64(r.fb.Height))

	// Light direction (simple directional light)
	lightDir := math3d.NewVec3(0, 0, -1).Normalize()

	for _, mi := range w.Instances {
		modelMatrix := mi.ModelMatrix()
		mvp := projection.Mul(view).Mul(modelMatrix)
		mvpViewport := viewport.Mul(mvp)

		for _, tri := range mi.Mesh.Triangles {
			v0 := mvpViewport.MulVec4(math3d.Vec3ToVec4(tri.V0, 1)).ToVec3()
			v1 := mvpViewport.MulVec4(math3d.Vec3ToVec4(tri.V1, 1)).ToVec3()
			v2 := mvpViewport.MulVec4(math3d.Vec3ToVec4(tri.V2, 1)).ToVec3()

			normal := tri.ComputeNormal()
			brightness := math.Max(0, normal.Dot(lightDir))

			screenNormal := v1.Sub(v0).Cross(v2.Sub(v0))
			if screenNormal.Z <= 0 {
				continue
			}

			switch r.mode {
			case Wireframe:
				r.drawWireframe(v0, v1, v2)
			case Solid:
				r.fillTriangle(v0, v1, v2, brightness)
			case SolidWireframe:
				r.fillTriangle(v0, v1, v2, brightness)
				r.drawWireframe(v0, v1, v2)
			}
		}
	}
}

// drawWireframe draws the edges of a triangle
func (r *Renderer) drawWireframe(v0, v1, v2 math3d.Vec3) {
	r.fb.DrawLine(v0.X, v0.Y, v0.Z, v1.X, v1.Y, v1.Z, 1.0)
	r.fb.DrawLine(v1.X, v1.Y, v1.Z, v2.X, v2.Y, v2.Z, 1.0)
	r.fb.DrawLine(v2.X, v2.Y, v2.Z, v0.X, v0.Y, v0.Z, 1.0)
}

// fillTriangle fills a triangle using scanline rasterization
func (r *Renderer) fillTriangle(v0, v1, v2 math3d.Vec3, brightness float64) {
	// Sort vertices by Y coordinate
	if v0.Y > v1.Y {
		v0, v1 = v1, v0
	}
	if v0.Y > v2.Y {
		v0, v2 = v2, v0
	}
	if v1.Y > v2.Y {
		v1, v2 = v2, v1
	}

	// Convert to integer coordinates
	x0, y0, z0 := int(v0.X), int(v0.Y), v0.Z
	x1, y1, z1 := int(v1.X), int(v1.Y), v1.Z
	x2, y2, z2 := int(v2.X), int(v2.Y), v2.Z

	// Degenerate triangle
	if y0 == y2 {
		return
	}

	// Scanline fill
	for y := y0; y <= y2; y++ {
		// Determine if we're in the top or bottom half
		secondHalf := y > y1 || y1 == y0

		segmentHeight := y2 - y0
		if segmentHeight == 0 {
			continue
		}

		alpha := float64(y-y0) / float64(segmentHeight)
		xA := float64(x0) + alpha*float64(x2-x0)
		zA := z0 + alpha*(z2-z0)

		var xB float64
		var zB float64

		if secondHalf {
			if y1 == y2 {
				continue
			}
			beta := float64(y-y1) / float64(y2-y1)
			xB = float64(x1) + beta*float64(x2-x1)
			zB = z1 + beta*(z2-z1)
		} else {
			if y1 == y0 {
				continue
			}
			beta := float64(y-y0) / float64(y1-y0)
			xB = float64(x0) + beta*float64(x1-x0)
			zB = z0 + beta*(z1-z0)
		}

		if xA > xB {
			xA, xB = xB, xA
			zA, zB = zB, zA
		}

		// Fill the scanline
		for x := int(xA); x <= int(xB); x++ {
			t := 0.0
			if xB-xA != 0 {
				t = (float64(x) - xA) / (xB - xA)
			}
			z := zA + t*(zB-zA)
			r.fb.SetPixel(x, y, brightness, z)
		}
	}
}
