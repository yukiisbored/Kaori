package input

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/yukiisbored/Kaori/common"
)

var (
	// Joystick
	joysticks            map[sdl.JoystickID]*sdl.Joystick = make(map[sdl.JoystickID]*sdl.Joystick)
	joystickAxises       map[sdl.JoystickID][]int16       = make(map[sdl.JoystickID][]int16)
	joystickButtons      map[sdl.JoystickID][]bool        = make(map[sdl.JoystickID][]bool)
	joystickHats         map[sdl.JoystickID][]uint8       = make(map[sdl.JoystickID][]uint8)
	joysticksInitialised bool

	// Mouse
	mouseLocation common.Vector2D = common.Vector2D{0, 0}
	mouseState    []bool          = make([]bool, 3)
)

const (
	MOUSE_LEFT   = 0
	MOUSE_MIDDLE = 1
	MOUSE_RIGHT  = 2

	JOYSTICK_HAT_N  = 1
	JOYSTICK_HAT_NE = 3
	JOYSTICK_HAT_E  = 2
	JOYSTICK_HAT_SE = 6
	JOYSTICK_HAT_S  = 4
	JOYSTICK_HAT_SW = 9
	JOYSTICK_HAT_W  = 8
	JOYSTICK_HAT_NW = 12
)

func Init() {
	if sdl.WasInit(sdl.INIT_JOYSTICK) == 0 {
		sdl.InitSubSystem(sdl.INIT_JOYSTICK)
	}

	if sdl.NumJoysticks() > 0 {
		for i := 0; i < sdl.NumJoysticks(); i++ {
			id := sdl.JoystickID(i)

			addJoystick(id)
		}

		sdl.JoystickEventState(sdl.ENABLE)

		joysticksInitialised = true
	}
}

func HandleEvents(e sdl.Event) {
	switch t := e.(type) {
	case *sdl.JoyDeviceEvent:
		if t.Type == sdl.JOYDEVICEADDED {
			addJoystick(t.Which)
		} else if t.Type == sdl.JOYDEVICEREMOVED {
			remJoystick(t.Which)
		}
		break
	case *sdl.JoyAxisEvent:
		joystickAxises[t.Which][t.Axis] = t.Value
		break
	case *sdl.JoyButtonEvent:
		if t.State == 1 {
			joystickButtons[t.Which][t.Button] = true
		} else {
			joystickButtons[t.Which][t.Button] = false
		}
		break
	case *sdl.JoyHatEvent:
		joystickHats[t.Which][t.Hat] = t.Value
		break
	case *sdl.MouseMotionEvent:
		mouseLocation.X = float64(t.X)
		mouseLocation.Y = float64(t.Y)
		break
	case *sdl.MouseButtonEvent:
		if t.Type == sdl.MOUSEBUTTONDOWN {
			if t.Button == sdl.BUTTON_LEFT {
				mouseState[MOUSE_LEFT] = true
			}

			if t.Button == sdl.BUTTON_MIDDLE {
				mouseState[MOUSE_MIDDLE] = true
			}

			if t.Button == sdl.BUTTON_RIGHT {
				mouseState[MOUSE_RIGHT] = true
			}
		} else {
			if t.Button == sdl.BUTTON_LEFT {
				mouseState[MOUSE_LEFT] = false
			}

			if t.Button == sdl.BUTTON_MIDDLE {
				mouseState[MOUSE_MIDDLE] = false
			}

			if t.Button == sdl.BUTTON_RIGHT {
				mouseState[MOUSE_RIGHT] = false
			}
		}
		break
	}
}

func Axis(id sdl.JoystickID, axis uint) int16 {
	return joystickAxises[id][axis]
}

func Axisf(id sdl.JoystickID, axis uint) float32 {
	return float32(Axis(id, axis)) / 65536
}

func Button(id sdl.JoystickID, button uint) bool {
	return joystickButtons[id][button]
}

func Hat(id sdl.JoystickID, hat uint) uint8 {
	return joystickHats[id][hat]
}

func MouseLocation() common.Vector2D {
	return mouseLocation
}

func Mouse(button uint8) bool {
	return mouseState[button]
}

func Key(key sdl.Scancode) bool {
	keyState := sdl.GetKeyboardState()

	if keyState[key] == 1 {
		return true
	} else {
		return false
	}
}

func Clean() {
	for k := range joysticks {
		remJoystick(k)
	}
}

func addJoystick(id sdl.JoystickID) {
	if joy := sdl.JoystickOpen(id); joy != nil {
		id = joy.InstanceID()

		joysticks[id] = joy
		joystickAxises[id] = make([]int16, joy.NumAxes())
		joystickButtons[id] = make([]bool, joy.NumButtons())
		joystickHats[id] = make([]uint8, joy.NumHats())

		log.Printf("Input // Added %s as Joystick %d\n", joy.Name(), id)
	}
}

func remJoystick(id sdl.JoystickID) {
	if joy := joysticks[id]; joy != nil {
		joy.Close()

		delete(joysticks, id)
		delete(joystickAxises, id)
		delete(joystickButtons, id)

		log.Printf("Input // Removed Joystick %d\n", id)
	}
}
