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

## Demo Models

VESAsterizer includes procedurally generated demo models specifically crafted to showcase the rasterizer's capabilities:

### Running the Demo

```bash
./demo.sh
```

This runs through all demo models with various rendering modes and character sets.

### Included Demo Models

**Torus** (`examples/models/demo/torus.obj`)
- Classic demoscene donut shape
- 1,152 vertices, 2,304 triangles
- Perfect for showing smooth curves and shading
- Best with: `-charset unicode -mode solid`

**Icosphere** (`examples/models/demo/icosphere.obj`)
- Subdivided icosahedron (low-poly sphere)
- 162 vertices, 320 triangles
- Demonstrates smooth shading on curved surfaces
- Best with: `-charset shade -mode solid`

**Helix** (`examples/models/demo/helix.obj`)
- Twisted 3D helix/spring
- 520 vertices, 1,024 triangles
- Showcases depth buffer with overlapping geometry
- Best with: `-charset unicode -mode wireframe`

**Sculpture** (`examples/models/demo/sculpture.obj`)
- Abstract geometric art piece with spikes
- 384 vertices, 1,104 triangles
- Complex geometry for testing dense Unicode rendering
- Best with: `-charset dense -mode solid`

## Procedural Model Generator

VESAsterizer includes `modelgen`, a tool to create custom procedural models:

```bash
# Build the generator
go build -o modelgen ./cmd/modelgen

# Generate models
./modelgen -type torus -output my_torus.obj -detail 64
./modelgen -type icosphere -output my_sphere.obj -detail 3
./modelgen -type helix -output my_helix.obj -detail 48
./modelgen -type sculpture -output my_art.obj -detail 32
```

### Model Types

- **torus**: Classic donut shape (detail controls major/minor segments)
- **icosphere**: Subdivided icosahedron (detail = subdivision levels)
- **helix**: Twisted 3D spring (detail controls segments)
- **sculpture**: Abstract spiky structure (detail controls layers)

The `-detail` parameter controls geometric complexity. Higher values = more triangles = smoother but slower rendering.

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

## Chiptune Music System

Because every demoscene production needs a soundtrack. VESAsterizer includes a complete chiptune tracker and synthesizer engine.

### VTM Format (VESAsterizer Tracker Module)

Human-readable tracker format inspired by classic .MOD/.XM formats:

```
TITLE My Amazing Track
TEMPO 140

# Define instruments with synthesis parameters
INSTRUMENT Lead SQUARE 0.010 0.100 0.600 0.200
INSTRUMENT Bass SAW 0.005 0.150 0.700 0.150

# Pattern-based sequencing
PATTERN 16 4
CH 0: C-4 ... E-4 ... G-4 ... E-4 ... C-4 ... D-4 ... E-4 ... ... ---
CH 1: C-2 --- E-2 --- G-2 --- E-2 --- C-2 --- D-2 --- E-2 --- G-2 ---
ENDPATTERN

# Define playback order
SEQUENCE 0 1 0 2
```

### Music Player Tool

Test and preview VTM files:

```bash
# Build the music player
go build -o musicplayer ./cmd/musicplayer

# Play a track (shows track info and pattern data)
./musicplayer -music examples/music/demo1.vtm
./musicplayer -music examples/music/demo2.vtm
```

### VTM Format Specification

**Header Commands:**
- `TITLE <name>` - Track title
- `TEMPO <bpm>` - Tempo in beats per minute

**Instrument Definition:**
```
INSTRUMENT <name> <wavetype> <attack> <decay> <sustain> <release>
```

Wave types: `SQUARE`, `SAW`, `TRIANGLE`, `SINE`, `NOISE`

ADSR values in seconds (attack, decay, release) and 0.0-1.0 (sustain level)

**Pattern Definition:**
```
PATTERN <rows> <channels>
CH <num>: <note> <note> <note> ...
ENDPATTERN
```

Note format:
- `C-4` - Note C in octave 4
- `C#5` - C sharp in octave 5
- `---` - Rest (no note)
- `...` - Continue previous note

**Sequence:**
```
SEQUENCE <pattern> <pattern> <pattern> ...
```

Defines the order patterns are played.

### Synthesis Engine Features

- **5 Waveform Types**: Square, Saw, Triangle, Sine, Noise
- **ADSR Envelopes**: Full Attack/Decay/Sustain/Release control
- **Multi-channel**: Up to 8 simultaneous voices
- **Pattern Sequencer**: Classic tracker-style pattern system
- **Human-Readable Format**: Edit tracks in any text editor

### Included Demo Tracks

**demo1.vtm** - "Terminal Dreams"
- 140 BPM, melodic chiptune
- Demonstrates arpeggio patterns and multi-channel composition
- 6 patterns, 8-pattern sequence

**demo2.vtm** - "Raster Madness"
- 160 BPM, aggressive chiptune
- Fast arpeggios and heavy bass
- 3 patterns, 8-pattern sequence

### Creating Your Own Music

1. Start with an existing VTM file as template
2. Define instruments with different waveforms and ADSR
3. Create patterns (16-32 rows typical)
4. Compose melodies using note notation (C-4, D#5, etc.)
5. Arrange patterns in sequence
6. Test with `musicplayer` tool

**Pro Tips:**
- Use SQUARE waves for leads and basses
- SAW waves for rich, harmonic sounds
- TRIANGLE for softer melodies
- Short attack times (0.001-0.010) for percussive sounds
- Longer release times for sustained pads
- Pattern length of 16 rows = 1 bar at typical tempo

## Why?

Because sometimes you want to appreciate the fundamentals of computer graphics AND computer music without the abstraction of modern APIs. This is for the purists who remember when every pixel counted, every voice mattered, and understanding the math was essential.

## License

MIT License - See LICENSE file for details