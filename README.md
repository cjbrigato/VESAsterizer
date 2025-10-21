# VESAsterizer

A terminal-based 3D software rasterizer written in Go. Experience the beauty of 3D graphics rendered entirely with ASCII/Unicode characters in your terminal.

## Features

- **Pure Software Rasterization**: No GPU required, all rendering done in CPU
- **Unicode Character Sets**: Choose from ASCII, Unicode blocks, shading, or dense mode for ultra-sharp rendering
- **Terminal Output**: Renders to character art with depth-aware brightness mapping
- **OBJ Model Loading**: Load and display standard Wavefront OBJ files
- **Real-time Animation**: Smooth rotation and transformation of 3D models
- **Multiple Render Modes**: Wireframe, solid, or combined rendering
- **Depth Buffering**: Proper Z-buffering for correct occlusion
- **Perspective Projection**: Full 3D perspective with configurable camera

## Installation

```bash
go build -o vesasterizer ./cmd/vesasterizer
```

## Usage

Basic usage with default cube:
```bash
./vesasterizer
```

Load a custom model:
```bash
./vesasterizer -model path/to/model.obj
```

### Command-line Options

- `-model <path>` - Path to OBJ model file (default: "examples/models/cube.obj")
- `-width <int>` - Terminal width in characters (default: 120)
- `-height <int>` - Terminal height in characters (default: 40)
- `-mode <mode>` - Render mode: wireframe, solid, both (default: "wireframe")
- `-charset <set>` - Character set: ascii, unicode, shade, dense (default: "unicode")
- `-animate` - Enable rotation animation (default: true)
- `-fps <int>` - Frames per second for animation (default: 30)

### Examples

Render a pyramid in solid mode with Unicode blocks:
```bash
./vesasterizer -model examples/models/pyramid.obj -mode solid -charset unicode
```

Static wireframe render of a tetrahedron with dense Unicode:
```bash
./vesasterizer -model examples/models/tetrahedron.obj -animate=false -charset dense
```

High FPS solid rendering with wireframe overlay and shaded characters:
```bash
./vesasterizer -mode both -fps 60 -charset shade
```

Classic ASCII mode (for the purists):
```bash
./vesasterizer -charset ascii
```

Ultra-sharp Unicode rendering:
```bash
./vesasterizer -mode solid -charset unicode -width 160 -height 50
```

## Technical Details

### Architecture

- `pkg/math3d/` - 3D math library (vectors, matrices, transformations)
- `pkg/terminal/` - Terminal framebuffer with depth buffer
- `pkg/loader/` - OBJ file parser and mesh structures
- `pkg/renderer/` - Rasterization pipeline and camera system
- `cmd/vesasterizer/` - Main application

### Rendering Pipeline

1. **Model Transform**: Apply rotation/scale/translation to vertices
2. **View Transform**: Transform to camera space
3. **Projection Transform**: Apply perspective projection
4. **Viewport Transform**: Map to screen coordinates
5. **Rasterization**: Convert triangles to pixels with depth testing
6. **Character Conversion**: Map brightness values to characters

### Character Sets

VESAsterizer supports multiple character sets for rendering:

**ASCII** (`-charset ascii`) - Classic ASCII art:
```
' ' . : - = + * # % @
```

**Unicode Blocks** (`-charset unicode`) - Smooth block gradient (default):
```
' ' ░ ▒ ▓ █
```

**Unicode Shade** (`-charset shade`) - Extended shading characters:
```
' ' · ⋅ ∘ ○ ● ◉ ⬤ ⬛
```

**Unicode Dense** (`-charset dense`) - Maximum variety with 80+ Unicode characters:
```
Various blocks, geometric shapes, circles, and shading characters
for ultra-fine brightness gradation
```

Unicode modes provide significantly sharper and more detailed rendering
compared to classic ASCII.

## Creating Custom Models

VESAsterizer supports Wavefront OBJ files. Create your own models in Blender, Maya, or any 3D modeling tool and export as OBJ.

Basic OBJ format example:
```
# Vertices
v 0.0 0.0 0.0
v 1.0 0.0 0.0
v 0.0 1.0 0.0

# Faces (triangles)
f 1 2 3
```

## Why?

Because sometimes you want to appreciate the fundamentals of computer graphics without the abstraction of modern APIs. This is for the purists who remember when every pixel counted and understanding the math mattered.

## License

MIT License - See LICENSE file for details