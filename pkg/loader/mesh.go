package loader

import (
	"github.com/cjbrigato/VESAsterizer/pkg/math3d"
)

// Triangle represents a single triangle face
type Triangle struct {
	V0, V1, V2 math3d.Vec3 // Vertex positions
}

// Mesh represents a 3D model
type Mesh struct {
	Vertices  []math3d.Vec3
	Triangles []Triangle
}

// NewMesh creates a new empty mesh
func NewMesh() *Mesh {
	return &Mesh{
		Vertices:  make([]math3d.Vec3, 0),
		Triangles: make([]Triangle, 0),
	}
}

// NewGridPlane creates an XZ-plane grid centered at the origin on Y=0.
// halfCells defines how many cells extend from the origin along +X/-X and +Z/-Z (total cells per axis = 2*halfCells).
// cellSize defines the world-space size of each cell.
func NewGridPlane(halfCells int, cellSize float64) *Mesh {
	m := NewMesh()
	if halfCells < 1 {
		halfCells = 1
	}
	side := 2*halfCells + 1 // number of vertices per axis

	// Generate vertices
	for z := -halfCells; z <= halfCells; z++ {
		for x := -halfCells; x <= halfCells; x++ {
			m.AddVertex(math3d.NewVec3(float64(x)*cellSize, 0, float64(z)*cellSize))
		}
	}

	// Helper to compute vertex index in the grid
	index := func(x, z int) int {
		return (z+halfCells)*side + (x + halfCells)
	}

	// Generate triangles (two per cell)
	for z := -halfCells; z < halfCells; z++ {
		for x := -halfCells; x < halfCells; x++ {
			i0 := index(x, z)
			i1 := index(x+1, z)
			i2 := index(x, z+1)
			i3 := index(x+1, z+1)

			// Triangle 1
			m.AddTriangle(i0, i1, i2)
			// Triangle 2
			m.AddTriangle(i2, i1, i3)
		}
	}

	return m
}

// AddVertex adds a vertex to the mesh
func (m *Mesh) AddVertex(v math3d.Vec3) {
	m.Vertices = append(m.Vertices, v)
}

// AddTriangle adds a triangle to the mesh using vertex indices
func (m *Mesh) AddTriangle(i0, i1, i2 int) {
	if i0 < len(m.Vertices) && i1 < len(m.Vertices) && i2 < len(m.Vertices) {
		m.Triangles = append(m.Triangles, Triangle{
			V0: m.Vertices[i0],
			V1: m.Vertices[i1],
			V2: m.Vertices[i2],
		})
	}
}

// ComputeNormal computes the normal of a triangle
func (t Triangle) ComputeNormal() math3d.Vec3 {
	edge1 := t.V1.Sub(t.V0)
	edge2 := t.V2.Sub(t.V0)
	return edge1.Cross(edge2).Normalize()
}

// Centroid returns the center point of the triangle
func (t Triangle) Centroid() math3d.Vec3 {
	return math3d.Vec3{
		X: (t.V0.X + t.V1.X + t.V2.X) / 3.0,
		Y: (t.V0.Y + t.V1.Y + t.V2.Y) / 3.0,
		Z: (t.V0.Z + t.V1.Z + t.V2.Z) / 3.0,
	}
}
