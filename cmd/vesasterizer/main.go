package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/cjbrigato/VESAsterizer/pkg/loader"
	"github.com/cjbrigato/VESAsterizer/pkg/math3d"
	"github.com/cjbrigato/VESAsterizer/pkg/renderer"
	"github.com/cjbrigato/VESAsterizer/pkg/terminal"
)

func main() {
	// Command-line flags
	modelFile := flag.String("model", "examples/models/cube.obj", "Path to OBJ model file")
	width := flag.Int("width", 120, "Terminal width in characters")
	height := flag.Int("height", 40, "Terminal height in characters")
	mode := flag.String("mode", "wireframe", "Render mode: wireframe, solid, both")
	charset := flag.String("charset", "unicode", "Character set: ascii, unicode, shade, dense")
	animate := flag.Bool("animate", true, "Enable rotation animation")
	fps := flag.Int("fps", 30, "Frames per second for animation")
	flag.Parse()

	// Load the model
	mesh, err := loader.LoadOBJ(*modelFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading model: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Loaded model: %d vertices, %d triangles\n", len(mesh.Vertices), len(mesh.Triangles))
	time.Sleep(1 * time.Second)

	// Create framebuffer
	fb := terminal.NewFramebuffer(*width, *height)

	// Set character set
	switch *charset {
	case "ascii":
		fb.SetCharSet(terminal.ASCII)
	case "unicode":
		fb.SetCharSet(terminal.UnicodeBlocks)
	case "shade":
		fb.SetCharSet(terminal.UnicodeShade)
	case "dense":
		fb.SetCharSet(terminal.UnicodeDense)
	default:
		fb.SetCharSet(terminal.UnicodeBlocks)
	}

	// Create camera
	camera := renderer.NewCamera()
	camera.Aspect = float64(*width) / float64(*height)
	camera.Position = math3d.NewVec3(3, 2, 5)

	// Create renderer
	r := renderer.NewRenderer(fb, camera)

	// Set render mode
	switch *mode {
	case "solid":
		r.SetRenderMode(renderer.Solid)
	case "both":
		r.SetRenderMode(renderer.SolidWireframe)
	default:
		r.SetRenderMode(renderer.Wireframe)
	}

	if !*animate {
		// Single frame render
		r.Clear()
		modelMatrix := math3d.Identity()
		r.RenderMesh(mesh, modelMatrix)
		fb.Print()
		return
	}

	// Animation loop
	frameDuration := time.Second / time.Duration(*fps)
	angle := 0.0
	startTime := time.Now()
	frameCount := 0

	for {
		frameStart := time.Now()

		// Update rotation
		angle += 0.02

		// Create model transformation matrix
		rotY := math3d.RotationY(angle)
		rotX := math3d.RotationX(angle * 0.5)
		modelMatrix := rotY.Mul(rotX)

		// Render
		r.Clear()
		r.RenderMesh(mesh, modelMatrix)

		// Display info
		frameCount++
		elapsed := time.Since(startTime).Seconds()
		currentFPS := float64(frameCount) / elapsed

		info := fmt.Sprintf("VESAsterizer | FPS: %.1f | Vertices: %d | Triangles: %d | Mode: %s | CharSet: %s",
			currentFPS, len(mesh.Vertices), len(mesh.Triangles), *mode, *charset)
		fmt.Print(fb.RenderWithInfo(info))

		// Frame timing
		frameTime := time.Since(frameStart)
		if frameTime < frameDuration {
			time.Sleep(frameDuration - frameTime)
		}
	}
}
