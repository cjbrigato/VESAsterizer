package input

import (
	"fmt"
	"os"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/cjbrigato/VESAsterizer/pkg/renderer"
)

const (
	QuitKey       = "q"
	CAMPOSX_PLUS  = "a"
	CAMPOSX_MINUS = "q"
	CAMPOSY_PLUS  = "z"
	CAMPOSY_MINUS = "s"
	CAMPOSZ_PLUS  = "x"
	CAMPOSZ_MINUS = "d"
	STEP_PLUS     = "+"
	STEP_MINUS    = "-"
)

type CameraControls struct {
	Camera *renderer.Camera
	Step   float64
}

func NewCameraControls(camera *renderer.Camera) *CameraControls {
	return &CameraControls{
		Camera: camera,
		Step:   1.0,
	}
}

func (cc *CameraControls) UpdateCamera(key string) {
	switch key {
	case CAMPOSX_PLUS:
		cc.Camera.Position.X += cc.Step
	case CAMPOSX_MINUS:
		cc.Camera.Position.X -= cc.Step
	case CAMPOSY_PLUS:
		cc.Camera.Position.Y += cc.Step
	case CAMPOSY_MINUS:
		cc.Camera.Position.Y -= cc.Step
	case CAMPOSZ_PLUS:
		cc.Camera.Position.Z += cc.Step
	case CAMPOSZ_MINUS:
		cc.Camera.Position.Z -= cc.Step
	case STEP_PLUS:
		cc.Step += 10.0
	case STEP_MINUS:
		cc.Step -= 10.0
	}
}

func QuitFromKey() {
	fmt.Println("\rQuitting application")
	os.Exit(0) // Exit application
}

func ListenForKeyPress(cc *CameraControls) {
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		switch key.Code {
		case keys.CtrlC, keys.Escape:
			QuitFromKey()
			return true, nil // Return true to stop listener
		default:
			if key.String() == QuitKey {
				QuitFromKey()
			}
			cc.UpdateCamera(key.String())
			fmt.Printf("\rYou pressed the key: %s\n", key)
		}

		return false, nil // Return false to continue listening
	})
}
