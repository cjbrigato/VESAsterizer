package input

import (
	"fmt"
	"os"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/cjbrigato/VESAsterizer/pkg/math3d"
	"github.com/cjbrigato/VESAsterizer/pkg/renderer"
)

type CameraControlType int

const (
	CameraControlTypePosition CameraControlType = iota
	CameraControlTypeTarget
	CameraControlTypeUp
	CameraControlTypeMax
)

func (c CameraControlType) String() string {
	switch c {
	case CameraControlTypePosition:
		return "Position"
	case CameraControlTypeTarget:
		return "Target"
	case CameraControlTypeUp:
		return "Up"
	}
	return "Unknown"
}

//                T  --> PTCH+
//F---> YAW-      G-->PTCH-       H  --> YAW+

const (
	YAW_PLUS    = "h"
	YAW_MINUS   = "f"
	PITCH_PLUS  = "t"
	PITCH_MINUS = "g"
)

const (
	QuitKey                   = keys.Escape
	VECX_PLUS                 = "a"
	VECX_MINUS                = "q"
	VECY_PLUS                 = "z"
	VECY_MINUS                = "s"
	VECZ_PLUS                 = "e"
	VECZ_MINUS                = "d"
	STEP_PLUS                 = "+"
	STEP_MINUS                = "-"
	CHANGE_CONTROL_TYPE_PLUS  = "*"
	CHANGE_CONTROL_TYPE_MINUS = "/"
)

type CameraControls struct {
	Camera      *renderer.Camera
	Step        float64
	ControlType CameraControlType
}

func NewCameraControls(camera *renderer.Camera) *CameraControls {
	return &CameraControls{
		Camera:      camera,
		Step:        1.0,
		ControlType: CameraControlTypePosition,
	}
}

func (cc *CameraControls) ChangeControlType(key string) {
	switch key {
	case CHANGE_CONTROL_TYPE_PLUS:
		cc.ControlType++
	case CHANGE_CONTROL_TYPE_MINUS:
		cc.ControlType--
	}
	if cc.ControlType >= CameraControlTypeMax {
		cc.ControlType = CameraControlTypePosition
	}
	if cc.ControlType < 0 {
		cc.ControlType = CameraControlTypeMax - 1
	}
}

func (cc *CameraControls) UpdateCamera(key string) {

	var vec3ToUpdate *math3d.Vec3
	switch cc.ControlType {
	case CameraControlTypePosition:
		vec3ToUpdate = &cc.Camera.Position
	case CameraControlTypeTarget:
		vec3ToUpdate = &cc.Camera.Target
	case CameraControlTypeUp:
		vec3ToUpdate = &cc.Camera.Up
	default:
		vec3ToUpdate = &cc.Camera.Position
	}

	switch key {
	case VECX_PLUS:
		vec3ToUpdate.X += cc.Step
	case VECX_MINUS:
		vec3ToUpdate.X -= cc.Step
	case VECY_PLUS:
		vec3ToUpdate.Y += cc.Step
	case VECY_MINUS:
		vec3ToUpdate.Y -= cc.Step
	case VECZ_PLUS:
		vec3ToUpdate.Z += cc.Step
	case VECZ_MINUS:
		vec3ToUpdate.Z -= cc.Step
	case STEP_PLUS:
		cc.Step += 5.0
	case STEP_MINUS:
		cc.Step -= 5.0
	}

	switch key {
	case YAW_PLUS:
		cc.Camera.RotateView(0.05, 0)
	case YAW_MINUS:
		cc.Camera.RotateView(-0.05, 0)
	case PITCH_PLUS:
		cc.Camera.RotateView(0, 0.05)
	case PITCH_MINUS:
		cc.Camera.RotateView(0, -0.05)
	}
}

func QuitFromKey() {
	fmt.Println("\rQuitting application")
	os.Exit(0) // Exit application
}

func ListenForKeyPress(cc *CameraControls) {
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		switch key.Code {
		case QuitKey:
			return true, nil // Return true to stop listener
		default:
			cc.UpdateCamera(key.String())
			cc.ChangeControlType(key.String())
		}
		return false, nil // Return false to continue listening
	})
	QuitFromKey()
}
