package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

func main() {
	modelType := flag.String("type", "torus", "Model type: torus, icosphere, helix, sculpture")
	output := flag.String("output", "", "Output file path")
	detail := flag.Int("detail", 32, "Level of detail/subdivisions")
	flag.Parse()

	if *output == "" {
		fmt.Fprintf(os.Stderr, "Error: -output flag is required\n")
		os.Exit(1)
	}

	var vertices []Vec3
	var faces []Face

	switch *modelType {
	case "torus":
		vertices, faces = generateTorus(*detail, *detail/2)
	case "icosphere":
		vertices, faces = generateIcosphere(*detail)
	case "helix":
		vertices, faces = generateHelix(*detail)
	case "sculpture":
		vertices, faces = generateSculpture(*detail)
	default:
		fmt.Fprintf(os.Stderr, "Unknown model type: %s\n", *modelType)
		os.Exit(1)
	}

	if err := writeOBJ(*output, vertices, faces); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing OBJ: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated %s with %d vertices and %d faces\n", *modelType, len(vertices), len(faces))
}

type Vec3 struct {
	X, Y, Z float64
}

type Face struct {
	V0, V1, V2 int
}

// generateTorus creates a torus (donut) - classic demoscene object
func generateTorus(majorSegments, minorSegments int) ([]Vec3, []Face) {
	majorRadius := 1.0
	minorRadius := 0.4

	vertices := make([]Vec3, 0)
	faces := make([]Face, 0)

	// Generate vertices
	for i := 0; i < majorSegments; i++ {
		theta := float64(i) * 2.0 * math.Pi / float64(majorSegments)
		cosTheta := math.Cos(theta)
		sinTheta := math.Sin(theta)

		for j := 0; j < minorSegments; j++ {
			phi := float64(j) * 2.0 * math.Pi / float64(minorSegments)
			cosPhi := math.Cos(phi)
			sinPhi := math.Sin(phi)

			x := (majorRadius + minorRadius*cosPhi) * cosTheta
			y := (majorRadius + minorRadius*cosPhi) * sinTheta
			z := minorRadius * sinPhi

			vertices = append(vertices, Vec3{x, y, z})
		}
	}

	// Generate faces
	for i := 0; i < majorSegments; i++ {
		for j := 0; j < minorSegments; j++ {
			v0 := i*minorSegments + j
			v1 := i*minorSegments + (j+1)%minorSegments
			v2 := ((i+1)%majorSegments)*minorSegments + (j+1)%minorSegments
			v3 := ((i+1)%majorSegments)*minorSegments + j

			faces = append(faces, Face{v0, v1, v2})
			faces = append(faces, Face{v0, v2, v3})
		}
	}

	return vertices, faces
}

// generateIcosphere creates a subdivided icosahedron for smooth shading
func generateIcosphere(subdivisions int) ([]Vec3, []Face) {
	// Start with icosahedron
	t := (1.0 + math.Sqrt(5.0)) / 2.0

	vertices := []Vec3{
		{-1, t, 0}, {1, t, 0}, {-1, -t, 0}, {1, -t, 0},
		{0, -1, t}, {0, 1, t}, {0, -1, -t}, {0, 1, -t},
		{t, 0, -1}, {t, 0, 1}, {-t, 0, -1}, {-t, 0, 1},
	}

	faces := []Face{
		{0, 11, 5}, {0, 5, 1}, {0, 1, 7}, {0, 7, 10}, {0, 10, 11},
		{1, 5, 9}, {5, 11, 4}, {11, 10, 2}, {10, 7, 6}, {7, 1, 8},
		{3, 9, 4}, {3, 4, 2}, {3, 2, 6}, {3, 6, 8}, {3, 8, 9},
		{4, 9, 5}, {2, 4, 11}, {6, 2, 10}, {8, 6, 7}, {9, 8, 1},
	}

	// Normalize vertices to unit sphere
	for i := range vertices {
		length := math.Sqrt(vertices[i].X*vertices[i].X +
			vertices[i].Y*vertices[i].Y +
			vertices[i].Z*vertices[i].Z)
		vertices[i].X /= length
		vertices[i].Y /= length
		vertices[i].Z /= length
	}

	// Subdivide
	for sub := 0; sub < subdivisions; sub++ {
		newFaces := make([]Face, 0)
		midpointCache := make(map[[2]int]int)

		getMidpoint := func(v1, v2 int) int {
			key := [2]int{min(v1, v2), max(v1, v2)}
			if idx, exists := midpointCache[key]; exists {
				return idx
			}

			mid := Vec3{
				(vertices[v1].X + vertices[v2].X) / 2,
				(vertices[v1].Y + vertices[v2].Y) / 2,
				(vertices[v1].Z + vertices[v2].Z) / 2,
			}
			// Normalize to sphere
			length := math.Sqrt(mid.X*mid.X + mid.Y*mid.Y + mid.Z*mid.Z)
			mid.X /= length
			mid.Y /= length
			mid.Z /= length

			vertices = append(vertices, mid)
			idx := len(vertices) - 1
			midpointCache[key] = idx
			return idx
		}

		for _, face := range faces {
			m0 := getMidpoint(face.V0, face.V1)
			m1 := getMidpoint(face.V1, face.V2)
			m2 := getMidpoint(face.V2, face.V0)

			newFaces = append(newFaces,
				Face{face.V0, m0, m2},
				Face{face.V1, m1, m0},
				Face{face.V2, m2, m1},
				Face{m0, m1, m2},
			)
		}

		faces = newFaces
	}

	return vertices, faces
}

// generateHelix creates a twisted helix structure
func generateHelix(segments int) ([]Vec3, []Face) {
	vertices := make([]Vec3, 0)
	faces := make([]Face, 0)

	height := 3.0
	radius := 0.8
	tubeRadius := 0.2
	turns := 3.0
	tubeSegments := 8

	// Generate helix path
	for i := 0; i <= segments; i++ {
		t := float64(i) / float64(segments)
		angle := t * turns * 2.0 * math.Pi
		y := t*height - height/2

		centerX := radius * math.Cos(angle)
		centerZ := radius * math.Sin(angle)

		// Generate tube around helix path
		for j := 0; j < tubeSegments; j++ {
			tubeAngle := float64(j) * 2.0 * math.Pi / float64(tubeSegments)

			// Calculate tube offset
			normal := Vec3{math.Cos(angle), 0, math.Sin(angle)}
			binormal := Vec3{-math.Sin(angle), 0, math.Cos(angle)}

			offset := Vec3{
				tubeRadius * (math.Cos(tubeAngle)*normal.X + math.Sin(tubeAngle)*binormal.X),
				tubeRadius * (math.Cos(tubeAngle)*normal.Y + math.Sin(tubeAngle)*binormal.Y),
				tubeRadius * (math.Cos(tubeAngle)*normal.Z + math.Sin(tubeAngle)*binormal.Z),
			}

			vertices = append(vertices, Vec3{
				centerX + offset.X,
				y + offset.Y,
				centerZ + offset.Z,
			})
		}
	}

	// Generate faces
	for i := 0; i < segments; i++ {
		for j := 0; j < tubeSegments; j++ {
			v0 := i*tubeSegments + j
			v1 := i*tubeSegments + (j+1)%tubeSegments
			v2 := (i+1)*tubeSegments + (j+1)%tubeSegments
			v3 := (i+1)*tubeSegments + j

			faces = append(faces, Face{v0, v1, v2})
			faces = append(faces, Face{v0, v2, v3})
		}
	}

	return vertices, faces
}

// generateSculpture creates an abstract geometric sculpture
func generateSculpture(detail int) ([]Vec3, []Face) {
	vertices := make([]Vec3, 0)
	faces := make([]Face, 0)

	// Create a twisted spiky structure
	layers := detail
	spikesPerLayer := 8

	for layer := 0; layer < layers; layer++ {
		t := float64(layer) / float64(layers-1)
		y := t*3.0 - 1.5
		baseRadius := 0.3 + 0.5*math.Sin(t*math.Pi)
		twist := t * math.Pi * 2

		// Inner ring
		for spike := 0; spike < spikesPerLayer; spike++ {
			angle := float64(spike)*2.0*math.Pi/float64(spikesPerLayer) + twist
			x := baseRadius * math.Cos(angle)
			z := baseRadius * math.Sin(angle)
			vertices = append(vertices, Vec3{x, y, z})

			// Outer spike point
			spikeLength := 0.3 + 0.2*math.Sin(float64(layer)*0.5)
			xOuter := (baseRadius + spikeLength) * math.Cos(angle)
			zOuter := (baseRadius + spikeLength) * math.Sin(angle)
			vertices = append(vertices, Vec3{xOuter, y, zOuter})
		}
	}

	// Generate faces
	verticesPerLayer := spikesPerLayer * 2
	for layer := 0; layer < layers-1; layer++ {
		for spike := 0; spike < spikesPerLayer; spike++ {
			baseIdx := layer * verticesPerLayer
			nextLayerIdx := (layer + 1) * verticesPerLayer

			v0Inner := baseIdx + spike*2
			v0Outer := baseIdx + spike*2 + 1
			v1Inner := baseIdx + ((spike+1)%spikesPerLayer)*2
			v1Outer := baseIdx + ((spike+1)%spikesPerLayer)*2 + 1

			v2Inner := nextLayerIdx + spike*2
			v2Outer := nextLayerIdx + spike*2 + 1
			v3Inner := nextLayerIdx + ((spike+1)%spikesPerLayer)*2
			v3Outer := nextLayerIdx + ((spike+1)%spikesPerLayer)*2 + 1

			// Inner surface
			faces = append(faces, Face{v0Inner, v2Inner, v1Inner})
			faces = append(faces, Face{v1Inner, v2Inner, v3Inner})

			// Outer surface
			faces = append(faces, Face{v0Outer, v1Outer, v2Outer})
			faces = append(faces, Face{v1Outer, v3Outer, v2Outer})

			// Spike faces
			faces = append(faces, Face{v0Inner, v0Outer, v2Outer})
			faces = append(faces, Face{v0Inner, v2Outer, v2Inner})
		}
	}

	return vertices, faces
}

func writeOBJ(filename string, vertices []Vec3, faces []Face) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintf(file, "# Generated by VESAsterizer modelgen\n")
	fmt.Fprintf(file, "# Vertices: %d, Faces: %d\n\n", len(vertices), len(faces))

	for _, v := range vertices {
		fmt.Fprintf(file, "v %.6f %.6f %.6f\n", v.X, v.Y, v.Z)
	}

	fmt.Fprintf(file, "\n")

	for _, f := range faces {
		fmt.Fprintf(file, "f %d %d %d\n", f.V0+1, f.V1+1, f.V2+1)
	}

	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
