package terminal

import (
	"fmt"
	"math"
	"strings"
)

// CharSet represents different character set modes for rendering
type CharSet int

const (
	ASCII CharSet = iota
	UnicodeBlocks
	UnicodeShade
	UnicodeDense
)

// Brightness gradients for different character sets
var (
	// ASCII - Classic ASCII art characters
	asciiRamp = []rune{' ', '.', ':', '-', '=', '+', '*', '#', '%', '@'}

	// Unicode Block Elements - Smooth gradient using block characters
	unicodeBlockRamp = []rune{' ', '░', '▒', '▓', '█'}

	// Unicode Shade - Extended shading with more granularity
	unicodeShadeRamp = []rune{' ', '·', '⋅', '∘', '○', '●', '◉', '⬤', '⬛'}

	// Unicode Dense - Maximum density gradient using various Unicode blocks
	unicodeDenseRamp = []rune{
		' ', '·', '⋅', '░', '▒', '▓', '█', '▓', '▒', '░',
		'▀', '▄', '▌', '▐', '█', '▓', '▒', '░', '▪', '▫',
		'■', '□', '▢', '▣', '▤', '▥', '▦', '▧', '▨', '▩',
		'▬', '▭', '▮', '▯', '▰', '▱', '◆', '◇', '◈', '◉',
		'◊', '○', '◌', '◍', '◎', '●', '◐', '◑', '◒', '◓',
		'◔', '◕', '◖', '◗', '◘', '◙', '◚', '◛', '◜', '◝',
		'◞', '◟', '◠', '◡', '◢', '◣', '◤', '◥', '◦', '◧',
		'◨', '◩', '◪', '◫', '◬', '◭', '◮', '◯', '⬛', '⬜',
	}
)

// Framebuffer represents a character-based framebuffer with depth buffer
type Framebuffer struct {
	Width       int
	Height      int
	colorBuffer []float64 // Brightness values [0, 1]
	depthBuffer []float64 // Depth values for z-buffering
	charset     CharSet   // Character set to use for rendering
}

// NewFramebuffer creates a new framebuffer
func NewFramebuffer(width, height int) *Framebuffer {
	size := width * height
	return &Framebuffer{
		Width:       width,
		Height:      height,
		colorBuffer: make([]float64, size),
		depthBuffer: make([]float64, size),
		charset:     ASCII, // Default to ASCII
	}
}

// SetCharSet sets the character set for rendering
func (fb *Framebuffer) SetCharSet(cs CharSet) {
	fb.charset = cs
}

// Clear resets the framebuffer
func (fb *Framebuffer) Clear() {
	for i := range fb.colorBuffer {
		fb.colorBuffer[i] = 0.0
		fb.depthBuffer[i] = math.Inf(1) // Infinite depth
	}
}

// SetPixel sets a pixel with depth testing
// brightness: [0, 1] where 0 is black and 1 is white
// depth: z-coordinate (smaller = closer to camera)
func (fb *Framebuffer) SetPixel(x, y int, brightness, depth float64) {
	if x < 0 || x >= fb.Width || y < 0 || y >= fb.Height {
		return
	}

	idx := y*fb.Width + x

	// Depth test
	if depth < fb.depthBuffer[idx] {
		fb.depthBuffer[idx] = depth
		fb.colorBuffer[idx] = clamp(brightness, 0, 1)
	}
}

// GetPixel returns the brightness value at a pixel
func (fb *Framebuffer) GetPixel(x, y int) float64 {
	if x < 0 || x >= fb.Width || y < 0 || y >= fb.Height {
		return 0
	}
	return fb.colorBuffer[y*fb.Width+x]
}

// Render converts the framebuffer to a string for terminal output
func (fb *Framebuffer) Render() string {
	var sb strings.Builder

	// Terminal clear sequence
	sb.WriteString("\033[2J\033[H")

	for y := 0; y < fb.Height; y++ {
		for x := 0; x < fb.Width; x++ {
			brightness := fb.GetPixel(x, y)
			char := fb.brightnessToChar(brightness)
			sb.WriteRune(char)
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

// DrawLine draws a line using Bresenham's algorithm
func (fb *Framebuffer) DrawLine(x0f, y0f, z0, x1f, y1f, z1, brightness float64) {
	// Convert to int for Bresenham
	x0i, y0i := int(x0f), int(y0f)
	x1i, y1i := int(x1f), int(y1f)

	dx := abs(x1i - x0i)
	dy := abs(y1i - y0i)

	sx := 1
	if x0i > x1i {
		sx = -1
	}
	sy := 1
	if y0i > y1i {
		sy = -1
	}

	err := dx - dy
	x, y := x0i, y0i

	// Linear interpolation for depth
	length := math.Sqrt(float64(dx*dx + dy*dy))
	if length == 0 {
		fb.SetPixel(x0i, y0i, brightness, z0)
		return
	}

	for {
		t := math.Sqrt(float64((x-x0i)*(x-x0i)+(y-y0i)*(y-y0i))) / length
		z := z0 + t*(z1-z0)

		fb.SetPixel(x, y, brightness, z)

		if x == x1i && y == y1i {
			break
		}

		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x += sx
		}
		if e2 < dx {
			err += dx
			y += sy
		}
	}
}

// brightnessToChar converts a brightness value to a character based on the selected charset
func (fb *Framebuffer) brightnessToChar(brightness float64) rune {
	var ramp []rune

	switch fb.charset {
	case ASCII:
		ramp = asciiRamp
	case UnicodeBlocks:
		ramp = unicodeBlockRamp
	case UnicodeShade:
		ramp = unicodeShadeRamp
	case UnicodeDense:
		ramp = unicodeDenseRamp
	default:
		ramp = asciiRamp
	}

	idx := int(brightness * float64(len(ramp)-1))
	if idx < 0 {
		idx = 0
	}
	if idx >= len(ramp) {
		idx = len(ramp) - 1
	}
	return ramp[idx]
}

// Helper functions
func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// RenderWithInfo renders the framebuffer with additional info
func (fb *Framebuffer) RenderWithInfo(info string) string {
	render := fb.Render()
	return render + "\n" + info
}

// Print outputs the framebuffer to stdout
func (fb *Framebuffer) Print() {
	fmt.Print(fb.Render())
}
