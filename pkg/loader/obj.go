package loader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cjbrigato/VESAsterizer/pkg/math3d"
)

// LoadOBJ loads a Wavefront OBJ file
func LoadOBJ(filename string) (*Mesh, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	mesh := NewMesh()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "v": // Vertex
			if len(parts) >= 4 {
				x, _ := strconv.ParseFloat(parts[1], 64)
				y, _ := strconv.ParseFloat(parts[2], 64)
				z, _ := strconv.ParseFloat(parts[3], 64)
				mesh.AddVertex(math3d.NewVec3(x, y, z))
			}

		case "f": // Face
			// Parse face indices (supports v, v/vt, v/vt/vn, v//vn formats)
			indices := make([]int, 0)
			for i := 1; i < len(parts); i++ {
				// Split by '/' and take first value (vertex index)
				indexParts := strings.Split(parts[i], "/")
				idx, err := strconv.Atoi(indexParts[0])
				if err != nil {
					continue
				}
				// OBJ indices are 1-based, convert to 0-based
				indices = append(indices, idx-1)
			}

			// Triangulate polygon (simple fan triangulation)
			if len(indices) >= 3 {
				for i := 1; i < len(indices)-1; i++ {
					mesh.AddTriangle(indices[0], indices[i], indices[i+1])
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return mesh, nil
}
