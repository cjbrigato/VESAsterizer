package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/cjbrigato/VESAsterizer/pkg/audio"
	"github.com/cjbrigato/VESAsterizer/pkg/input"
	"github.com/cjbrigato/VESAsterizer/pkg/loader"
	"github.com/cjbrigato/VESAsterizer/pkg/math3d"
	"github.com/cjbrigato/VESAsterizer/pkg/renderer"
	"github.com/cjbrigato/VESAsterizer/pkg/terminal"
	"golang.org/x/term"
)

func main() {
	// Command-line flags
	modelFile := flag.String("model", "examples/models/demo/torus.obj", "Path to OBJ model file")
	musicFile := flag.String("music", "examples/music/demo2.vtm", "Path to VTM music file (optional)")
	width := flag.Int("width", 120, "Terminal width in characters")
	height := flag.Int("height", 40, "Terminal height in characters")
	mode := flag.String("mode", "solid", "Render mode: wireframe, solid, both")
	charset := flag.String("charset", "unicode", "Character set: ascii, unicode, shade, dense")
	animate := flag.Bool("animate", true, "Enable rotation animation")
	fps := flag.Int("fps", 60, "Frames per second for animation")
	flag.Parse()

	// Load the model
	mesh, err := loader.LoadOBJ(*modelFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading model: %v\n", err)
		os.Exit(1)
	}

	termW, termH, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting terminal size: %v\n", err)
		os.Exit(1)
	}

	nbInfosLines := 4

	fmt.Printf("Terminal size: %d x %d\n", termW, termH)
	*width = termW - 1
	*height = termH - nbInfosLines
	fmt.Printf("Loaded model: %d vertices, %d triangles\n", len(mesh.Vertices), len(mesh.Triangles))

	// Load and start music if specified
	var audioPlayback *audio.AudioPlayback
	if *musicFile != "" {
		module, err := audio.LoadVTM(*musicFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Could not load music: %v\n", err)
		} else {
			fmt.Printf("Loaded music: %s (%d BPM)\n", module.Title, module.Tempo)
			audioPlayback, err = audio.NewAudioPlayback(module, 44100)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Warning: Could not initialize audio: %v\n", err)
				audioPlayback = nil
			} else {
				audioPlayback.Play()
				defer audioPlayback.Stop()
			}
		}
	}

	//time.Sleep(1 * time.Second)

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
	camera.Position = math3d.NewVec3(0, 0, 0)
	camera.Rotation = math3d.NewVec3(0, 0, 0)

	cameraControls := input.NewCameraControls(camera)
	go input.ListenForKeyPress(cameraControls)
	// Create renderer
	r := renderer.NewRenderer(fb)
	// Create world and add instances
	world := renderer.NewWorld(camera)

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
		// Single frame render using World
		r.Clear()
		inst := renderer.NewMeshInstance(mesh)
		inst.Position = math3d.NewVec3(0, 100, 0)
		world.AddInstance(inst)
		r.RenderWorld(world)
		fb.Print()
		// Ensure audioPlayback is considered live until here
		runtime.KeepAlive(audioPlayback)
		return
	}

	// Animation loop
	frameDuration := time.Second / time.Duration(*fps)
	angle := 0.0
	startTime := time.Now()
	frameCount := 0

	// Create instances for animation
	inst := renderer.NewMeshInstance(mesh)
	world.AddInstance(inst)
	for {
		frameStart := time.Now()

		// Update rotation
		angle += 0.02

		// Update instance rotation (spin around Y)
		inst.Rotation.Y = angle

		// Render
		r.Clear()
		r.RenderWorld(world)

		// Display info
		frameCount++
		elapsed := time.Since(startTime).Seconds()
		currentFPS := float64(frameCount) / elapsed

		info := fmt.Sprintf("VESAsterizer | FPS: %.1f | Vertices: %d | Triangles: %d | Mode: %s | CharSet: %s",
			currentFPS, len(mesh.Vertices), len(mesh.Triangles), *mode, *charset)

		if audioPlayback != nil && !audioPlayback.IsDone() {
			info += " | \u266B MUSIC PLAYING"
		}
		var vec3ToShow *math3d.Vec3
		switch cameraControls.ControlType {
		case input.CameraControlTypePosition:
			vec3ToShow = &camera.Position
		case input.CameraControlTypeRotation:
			vec3ToShow = &camera.Rotation
		}
		info += fmt.Sprintf(" | %s: %.2f, %.2f, %.2f", cameraControls.ControlType.String(), vec3ToShow.X, vec3ToShow.Y, vec3ToShow.Z)
		info += fmt.Sprintf(" | Step: %.2f", cameraControls.Step)

		fmt.Print(fb.RenderWithInfo(info))

		// Ensure audioPlayback is considered live while the loop runs
		runtime.KeepAlive(audioPlayback)

		// Frame timing
		frameTime := time.Since(frameStart)
		if frameTime < frameDuration {
			time.Sleep(frameDuration - frameTime)
		}
	}
}
