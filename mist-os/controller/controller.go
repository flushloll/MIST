package controller

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const DeadZone = 6769

type Controller struct {
	sdlID          sdl.JoystickID
	gameController *sdl.GameController

	LEFT, RIGHT, UP, DOWN           bool
	SQUARE, CIRCLE, TRIANGLE, CROSS bool
	L1, R1, L3, R3                  bool
	START, SELECT, TOUCHPAD         bool

	LeftStickX, LeftStickY   float32
	RightStickX, RightStickY float32
	L2, R2                   float32
}

func NewController() (*Controller, error) {
	if err := sdl.Init(sdl.INIT_GAMECONTROLLER); err != nil {
		return nil, fmt.Errorf("failed to init SDL controller: %v", err)
	}

	if sdl.NumJoysticks() < 1 {
		return nil, fmt.Errorf("no controllers found")
	}

	gc := sdl.GameControllerOpen(0)
	if gc == nil {
		return nil, fmt.Errorf("failed to open game controller: %v", sdl.GetError())
	}

	joystick := gc.Joystick()
	return &Controller{
		gameController: gc,
		sdlID:          joystick.InstanceID(),
	}, nil
}

func (c *Controller) Close() {
	if c.gameController != nil {
		c.gameController.Close()
	}
}

func (c *Controller) Update() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		// Buttons
		case *sdl.ControllerButtonEvent:
			if t.Which != c.sdlID {
				continue
			}
			isPressed := t.State == sdl.PRESSED
			switch t.Button {

			case sdl.CONTROLLER_BUTTON_A:
				c.CROSS = isPressed
			case sdl.CONTROLLER_BUTTON_B:
				c.CIRCLE = isPressed
			case sdl.CONTROLLER_BUTTON_X:
				c.SQUARE = isPressed
			case sdl.CONTROLLER_BUTTON_Y:
				c.TRIANGLE = isPressed

			case sdl.CONTROLLER_BUTTON_DPAD_UP:
				c.UP = isPressed
			case sdl.CONTROLLER_BUTTON_DPAD_DOWN:
				c.DOWN = isPressed
			case sdl.CONTROLLER_BUTTON_DPAD_LEFT:
				c.LEFT = isPressed
			case sdl.CONTROLLER_BUTTON_DPAD_RIGHT:
				c.RIGHT = isPressed

			case sdl.CONTROLLER_BUTTON_LEFTSHOULDER:
				c.L1 = isPressed
			case sdl.CONTROLLER_BUTTON_RIGHTSHOULDER:
				c.R1 = isPressed
			case sdl.CONTROLLER_BUTTON_LEFTSTICK:
				c.L3 = isPressed
			case sdl.CONTROLLER_BUTTON_RIGHTSTICK:
				c.R3 = isPressed

			case sdl.CONTROLLER_BUTTON_START:
				c.START = isPressed
			case sdl.CONTROLLER_BUTTON_BACK:
				c.SELECT = isPressed
			case 20:
				c.TOUCHPAD = isPressed
			}

		// Analog Sticks
		case *sdl.ControllerAxisEvent:
			if t.Which != c.sdlID {
				continue
			}
			fVal := float32(t.Value) / 32767.0
			switch t.Axis {
			case sdl.CONTROLLER_AXIS_LEFTX:
				if abs(t.Value) > DeadZone {
					c.LeftStickX = fVal
				} else {
					c.LeftStickX = 0
				}
			case sdl.CONTROLLER_AXIS_LEFTY:
				if abs(t.Value) > DeadZone {
					c.LeftStickY = fVal
				} else {
					c.LeftStickY = 0
				}
			case sdl.CONTROLLER_AXIS_RIGHTX:
				if abs(t.Value) > DeadZone {
					c.RightStickX = fVal
				} else {
					c.RightStickX = 0
				}
			case sdl.CONTROLLER_AXIS_RIGHTY:
				if abs(t.Value) > DeadZone {
					c.RightStickY = fVal
				} else {
					c.RightStickY = 0
				}

			case sdl.CONTROLLER_AXIS_TRIGGERLEFT:
				c.L2 = fVal
			case sdl.CONTROLLER_AXIS_TRIGGERRIGHT:
				c.R2 = fVal
			}
		}
	}
}

func abs(x int16) int16 {
	if x < 0 {
		return -x
	}
	return x
}

// Check Buttons (Written by AI):

// func formatBool(b bool, label string) string {
// 	if b {
// 		return fmt.Sprintf("[\033[32m%s\033[0m]", label) // Green when pressed
// 	}
// 	return fmt.Sprintf("[ %s ]", label) // Gray/Default when released
// }

// func main() {
// 	controller, err := NewController()
// 	if err != nil {
// 		fmt.Printf("\033[31mSetup Error: %v\033[0m\n", err)
// 		return
// 	}
// 	defer controller.Close()
// 	defer sdl.Quit()

// 	// Clear screen command
// 	fmt.Print("\033[H\033[2J")

// 	for {
// 		controller.Update()

// 		// ANSI sequence to reset cursor to top left (prevents rolling text)
// 		fmt.Print("\033[H")

// 		fmt.Println("==================================================")
// 		fmt.Println("        PS5 DualSense Live Testing Tool           ")
// 		fmt.Println("==================================================")
// 		fmt.Println(" Press buttons and move sticks to check inputs.    ")
// 		fmt.Println(" Press Ctrl+C in your terminal to exit.           ")
// 		fmt.Println("--------------------------------------------------")

// 		// Row 1: Shoulders & Triggers
// 		fmt.Printf(" L1: %s   R1: %s\n", formatBool(controller.L1, "L1"), formatBool(controller.R1, "R1"))
// 		fmt.Printf(" L2 Trigger: %5.2f                  R2 Trigger: %5.2f\n", controller.L2, controller.R2)
// 		fmt.Println()

// 		// Row 2: Navigation & Specials
// 		fmt.Printf(" Create/Select: %s   Options/Start: %s\n", formatBool(controller.SELECT, "SHARE"), formatBool(controller.START, "OPT "))
// 		fmt.Printf(" Touchpad Click:%s\n", formatBool(controller.TOUCHPAD, "PAD"))
// 		fmt.Println()

// 		// Row 3: D-Pad & Face Buttons
// 		fmt.Printf(" D-PAD:            FACE BUTTONS:\n")
// 		fmt.Printf("    %s                 %s\n", formatBool(controller.UP, "▲"), formatBool(controller.TRIANGLE, "▲ TRIANGLE"))
// 		fmt.Printf(" %s   %s       %s   %s\n", formatBool(controller.LEFT, "◀"), formatBool(controller.RIGHT, "▶"), formatBool(controller.SQUARE, "■ SQUARE  "), formatBool(controller.CIRCLE, "● CIRCLE  "))
// 		fmt.Printf("    %s                 %s\n", formatBool(controller.DOWN, "▼"), formatBool(controller.CROSS, "✖ CROSS   "))
// 		fmt.Println()

// 		// Row 4: Joysticks
// 		fmt.Println(" ANALOG STICKS:")
// 		fmt.Printf(" Left Stick:  X: %5.2f, Y: %5.2f  %s\n", controller.LeftStickX, controller.LeftStickY, formatBool(controller.L3, "L3 Click"))
// 		fmt.Printf(" Right Stick: X: %5.2f, Y: %5.2f  %s\n", controller.RightStickX, controller.RightStickY, formatBool(controller.R3, "R3 Click"))
// 		fmt.Println("==================================================")

// 		time.Sleep(16 * time.Millisecond) // Poll at ~60 FPS
// 	}
// }
