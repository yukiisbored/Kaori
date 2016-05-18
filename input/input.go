package input

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	joysticks            map[sdl.JoystickID]*sdl.Joystick = make(map[sdl.JoystickID]*sdl.Joystick)
	joystickAxises       map[sdl.JoystickID][]int16       = make(map[sdl.JoystickID][]int16)
	joystickButtons      map[sdl.JoystickID][]bool        = make(map[sdl.JoystickID][]bool)
	joystickHats         map[sdl.JoystickID][]uint8       = make(map[sdl.JoystickID][]uint8)
	joysticksInitialised bool
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
