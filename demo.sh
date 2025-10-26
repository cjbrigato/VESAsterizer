#!/bin/bash

# VESAsterizer Demo Script
# Full audiovisual demoscene experience!

echo "==========================================="
echo "  VESAsterizer - AUDIOVISUAL DEMO MODE"
echo "  Terminal 3D Rasterizer + Chiptune Tracker"
echo "==========================================="
echo ""
echo "Press Ctrl+C to exit at any time"
echo ""
sleep 2

# Demo 1: Torus + Terminal Dreams (melodic chiptune)
echo "Demo 1: TORUS + 'Terminal Dreams'"
echo "Visual: Classic demoscene donut with Unicode blocks"
echo "Audio: Melodic chiptune at 140 BPM"
echo "Detail: 1152 vertices, 2304 triangles"
sleep 2
./vesasterizer -model examples/models/demo/torus.obj -music examples/music/demo1.vtm -charset unicode -mode solid -fps 30 &
PID=$!
sleep 12
kill $PID 2>/dev/null
wait $PID 2>/dev/null

sleep 1

# Demo 2: Icosphere + Raster Madness (aggressive chiptune)
echo ""
echo "Demo 2: ICOSPHERE + 'Raster Madness'"
echo "Visual: Smooth shaded sphere with shade characters"
echo "Audio: Aggressive chiptune at 160 BPM"
echo "Detail: 162 vertices, 320 triangles"
sleep 2
./vesasterizer -model examples/models/demo/icosphere.obj -music examples/music/demo2.vtm -charset shade -mode solid -fps 30 &
PID=$!
sleep 12
kill $PID 2>/dev/null
wait $PID 2>/dev/null

sleep 1

# Demo 3: Helix with Terminal Dreams (wireframe showcase)
echo ""
echo "Demo 3: HELIX + 'Terminal Dreams'"
echo "Visual: Twisted 3D helix wireframe (depth buffer showcase)"
echo "Audio: Melodic patterns"
echo "Detail: 520 vertices, 1024 triangles"
sleep 2
./vesasterizer -model examples/models/demo/helix.obj -music examples/music/demo1.vtm -charset unicode -mode wireframe -fps 30 &
PID=$!
sleep 12
kill $PID 2>/dev/null
wait $PID 2>/dev/null

sleep 1

# Demo 4: Sculpture with Raster Madness (dense Unicode)
echo ""
echo "Demo 4: SCULPTURE + 'Raster Madness'"
echo "Visual: Abstract geometric art with dense Unicode"
echo "Audio: Fast arpeggios and heavy bass"
echo "Detail: 384 vertices, 1104 triangles"
sleep 2
./vesasterizer -model examples/models/demo/sculpture.obj -music examples/music/demo2.vtm -charset dense -mode solid -fps 30 &
PID=$!
sleep 12
kill $PID 2>/dev/null
wait $PID 2>/dev/null

sleep 1

# Demo 5: Torus high detail with both render modes
echo ""
echo "Demo 5: TORUS + 'Terminal Dreams' - Combined rendering"
echo "Visual: Wireframe + solid rendering with Unicode"
echo "Audio: Full melodic composition"
sleep 2
./vesasterizer -model examples/models/demo/torus.obj -music examples/music/demo1.vtm -charset unicode -mode both -fps 30 &
PID=$!
sleep 12
kill $PID 2>/dev/null
wait $PID 2>/dev/null

sleep 1

# Demo 6: Icosphere high FPS audiovisual sync
echo ""
echo "Demo 6: ICOSPHERE - High FPS audiovisual experience"
echo "Visual: 60 FPS smooth rendering"
echo "Audio: Synchronized chiptune playback"
sleep 2
./vesasterizer -model examples/models/demo/icosphere.obj -music examples/music/demo2.vtm -charset unicode -mode solid -fps 60 &
PID=$!
sleep 12
kill $PID 2>/dev/null
wait $PID 2>/dev/null

echo ""
echo "==========================================="
echo "  AUDIOVISUAL DEMO COMPLETE!"
echo "==========================================="
echo ""
echo "Run your own audiovisual demos:"
echo "  ./vesasterizer -model examples/models/demo/torus.obj -music examples/music/demo1.vtm -charset unicode"
echo "  ./vesasterizer -model examples/models/demo/icosphere.obj -music examples/music/demo2.vtm -charset shade"
echo ""
echo "Visual-only (no music):"
echo "  ./vesasterizer -model examples/models/demo/helix.obj -mode wireframe"
echo "  ./vesasterizer -model examples/models/demo/sculpture.obj -charset dense"
echo ""
echo "Generate your own models:"
echo "  ./modelgen -type torus -output my_torus.obj -detail 64"
echo "  ./modelgen -type icosphere -output my_sphere.obj -detail 3"
echo ""
echo "Create your own music:"
echo "  Edit examples/music/demo1.vtm or demo2.vtm as templates"
echo "  Test with: ./musicplayer -music your_track.vtm"
echo ""
echo "Full demoscene experience unlocked! ░▒▓█ ♫"
echo ""
