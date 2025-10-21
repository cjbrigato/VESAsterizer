#!/bin/bash

# VESAsterizer Demo Script
# Showcases all procedurally generated demo models with various effects

echo "==================================="
echo "  VESAsterizer - DEMO MODE"
echo "  Terminal 3D Rasterizer"
echo "==================================="
echo ""
echo "Press Ctrl+C to exit at any time"
echo ""
sleep 2

# Demo 1: Torus with Unicode blocks (classic demoscene donut!)
echo "Demo 1: TORUS - Classic demoscene donut with Unicode blocks"
echo "Detail: 1152 vertices, 2304 triangles"
sleep 2
./vesasterizer -model examples/models/demo/torus.obj -charset unicode -mode solid -fps 30 &
PID=$!
sleep 8
kill $PID 2>/dev/null
wait $PID 2>/dev/null

sleep 1

# Demo 2: Icosphere with smooth shading
echo ""
echo "Demo 2: ICOSPHERE - Smooth shaded sphere"
echo "Detail: 162 vertices, 320 triangles"
sleep 2
./vesasterizer -model examples/models/demo/icosphere.obj -charset shade -mode solid -fps 30 &
PID=$!
sleep 8
kill $PID 2>/dev/null
wait $PID 2>/dev/null

sleep 1

# Demo 3: Helix with wireframe
echo ""
echo "Demo 3: HELIX - Twisted 3D helix (depth buffer showcase)"
echo "Detail: 520 vertices, 1024 triangles"
sleep 2
./vesasterizer -model examples/models/demo/helix.obj -charset unicode -mode wireframe -fps 30 &
PID=$!
sleep 8
kill $PID 2>/dev/null
wait $PID 2>/dev/null

sleep 1

# Demo 4: Sculpture with dense Unicode
echo ""
echo "Demo 4: SCULPTURE - Abstract geometric art with dense Unicode"
echo "Detail: 384 vertices, 1104 triangles"
sleep 2
./vesasterizer -model examples/models/demo/sculpture.obj -charset dense -mode solid -fps 30 &
PID=$!
sleep 8
kill $PID 2>/dev/null
wait $PID 2>/dev/null

sleep 1

# Demo 5: Torus with combined wireframe+solid
echo ""
echo "Demo 5: TORUS - Combined wireframe + solid rendering"
sleep 2
./vesasterizer -model examples/models/demo/torus.obj -charset unicode -mode both -fps 30 &
PID=$!
sleep 8
kill $PID 2>/dev/null
wait $PID 2>/dev/null

sleep 1

# Demo 6: Icosphere high FPS
echo ""
echo "Demo 6: ICOSPHERE - High FPS rendering (60 FPS)"
sleep 2
./vesasterizer -model examples/models/demo/icosphere.obj -charset unicode -mode solid -fps 60 &
PID=$!
sleep 8
kill $PID 2>/dev/null
wait $PID 2>/dev/null

echo ""
echo "==================================="
echo "  Demo Complete!"
echo "==================================="
echo ""
echo "Try running models yourself:"
echo "  ./vesasterizer -model examples/models/demo/torus.obj -charset unicode"
echo "  ./vesasterizer -model examples/models/demo/icosphere.obj -charset shade"
echo "  ./vesasterizer -model examples/models/demo/helix.obj -mode wireframe"
echo "  ./vesasterizer -model examples/models/demo/sculpture.obj -charset dense"
echo ""
echo "Generate your own models:"
echo "  ./modelgen -type torus -output my_torus.obj -detail 64"
echo "  ./modelgen -type icosphere -output my_sphere.obj -detail 3"
echo ""
