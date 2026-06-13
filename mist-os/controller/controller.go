package controller

import (
	"fmt"
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const DeadZone = 6769

type Button struct {
	Pressed     bool
	LastPressed time.Time
	LastLifted  time.Time
}

func (b *Button) set(pressed bool) time.Duration {
	if pressed == b.Pressed {
		return 0
	}
	b.Pressed = pressed
	now := time.Now()
	if pressed {
		b.LastPressed = now
		return 0
	} else {
		b.LastLifted = now
		return b.LastLifted.Sub(b.LastPressed)
	}
}

func (b *Button) HeldDuration() time.Duration {
	if !b.Pressed {
		return 0
	}
	return time.Since(b.LastPressed)
}

func (b *Button) PressDuration() time.Duration {
	if b.LastLifted.Before(b.LastPressed) {
		return 0
	}
	return b.LastLifted.Sub(b.LastPressed)
}

type SensorData struct {
	GyroX, GyroY, GyroZ    float32
	AccelX, AccelY, AccelZ float32
	Roll, Pitch            float32
}

type Controller struct {
	sdlID          sdl.JoystickID
	gameController *sdl.GameController

	LEFT, RIGHT, UP, DOWN           Button
	SQUARE, CIRCLE, TRIANGLE, CROSS Button
	L1, R1, L3, R3                  Button
	START, SELECT, TOUCHPAD         Button

	LeftStickX, LeftStickY   float32
	RightStickX, RightStickY float32
	L2, R2                   float32

	Motion SensorData
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

	if gc.HasSensor(sdl.SENSOR_ACCEL) {
		if err := gc.SetSensorEnabled(sdl.SENSOR_ACCEL, true); err != nil {
			fmt.Println("accel enable error:", err)
		}
	}
	if gc.HasSensor(sdl.SENSOR_GYRO) {
		if err := gc.SetSensorEnabled(sdl.SENSOR_GYRO, true); err != nil {
			fmt.Println("gyro enable error:", err)
		}
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
		case sdl.ControllerButtonEvent:
			if t.Which != c.sdlID {
				continue
			}
			isPressed := t.State == sdl.PRESSED
			switch t.Button {

			case sdl.CONTROLLER_BUTTON_A:
				c.CROSS.set(isPressed)
			case sdl.CONTROLLER_BUTTON_B:
				c.CIRCLE.set(isPressed)
			case sdl.CONTROLLER_BUTTON_X:
				c.SQUARE.set(isPressed)
			case sdl.CONTROLLER_BUTTON_Y:
				c.TRIANGLE.set(isPressed)

			case sdl.CONTROLLER_BUTTON_DPAD_UP:
				c.UP.set(isPressed)
			case sdl.CONTROLLER_BUTTON_DPAD_DOWN:
				c.DOWN.set(isPressed)
			case sdl.CONTROLLER_BUTTON_DPAD_LEFT:
				c.LEFT.set(isPressed)
			case sdl.CONTROLLER_BUTTON_DPAD_RIGHT:
				c.RIGHT.set(isPressed)

			case sdl.CONTROLLER_BUTTON_LEFTSHOULDER:
				c.L1.set(isPressed)
			case sdl.CONTROLLER_BUTTON_RIGHTSHOULDER:
				c.R1.set(isPressed)
			case sdl.CONTROLLER_BUTTON_LEFTSTICK:
				c.L3.set(isPressed)
			case sdl.CONTROLLER_BUTTON_RIGHTSTICK:
				c.R3.set(isPressed)

			case sdl.CONTROLLER_BUTTON_START:
				c.START.set(isPressed)
			case sdl.CONTROLLER_BUTTON_BACK:
				c.SELECT.set(isPressed)
			case 20:
				c.TOUCHPAD.set(isPressed)
			}

		// Analog Sticks
		case sdl.ControllerAxisEvent:
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

		// Sensors
		case sdl.ControllerSensorEvent:
			if t.Which != c.sdlID {
				continue
			}
			switch sdl.SensorType(t.Sensor) {
			case sdl.SENSOR_ACCEL:
				c.Motion.AccelX = t.Data[0]
				c.Motion.AccelY = t.Data[1]
				c.Motion.AccelZ = t.Data[2]

				c.Motion.Roll = float32(math.Atan2(float64(c.Motion.AccelX), float64(c.Motion.AccelZ))) * 180 / math.Pi
				c.Motion.Pitch = float32(math.Atan2(float64(c.Motion.AccelY), float64(c.Motion.AccelZ))) * 180 / math.Pi
			case sdl.SENSOR_GYRO:
				c.Motion.GyroX = t.Data[0]
				c.Motion.GyroY = t.Data[1]
				c.Motion.GyroZ = t.Data[2]
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
// 		fmt.Printf(" L1: %s   R1: %s\n", formatBool(controller.L1.Pressed, "L1"), formatBool(controller.R1.Pressed, "R1"))
// 		fmt.Printf(" L2 Trigger: %5.2f                  R2 Trigger: %5.2f\n", controller.L2, controller.R2)
// 		fmt.Println()

// 		// Row 2: Navigation & Specials
// 		fmt.Printf(" Create/Select: %s   Options/Start: %s\n", formatBool(controller.SELECT.Pressed, "SHARE"), formatBool(controller.START.Pressed, "OPT "))
// 		fmt.Printf(" Touchpad Click:%s\n", formatBool(controller.TOUCHPAD.Pressed, "PAD"))
// 		fmt.Println()

// 		// Row 2.5: Motion Sensors
// 		fmt.Println(" MOTION SENSORS (GYRO / ACCELEROMETER):")
// 		fmt.Printf(" Gyro Velocity ->  X: %6.2f, Y: %6.2f, Z: %6.2f\n", controller.Motion.GyroX, controller.Motion.GyroY, controller.Motion.GyroZ)
// 		fmt.Printf(" Accel Gravity ->  X: %6.2f, Y: %6.2f, Z: %6.2f\n", controller.Motion.AccelX, controller.Motion.AccelY, controller.Motion.AccelZ)
// 		fmt.Println("--------------------------------------------------")
// 		fmt.Printf(" Calculated Steering (Roll Tilt):  %6.2f°\n", controller.Motion.Roll)
// 		fmt.Println("==================================================")

// 		// Row 3: D-Pad & Face Buttons
// 		fmt.Printf(" D-PAD:            FACE BUTTONS:\n")
// 		fmt.Printf("    %s                 %s\n", formatBool(controller.UP.Pressed, "▲"), formatBool(controller.TRIANGLE.Pressed, "▲ TRIANGLE"))
// 		fmt.Printf(" %s   %s       %s   %s\n", formatBool(controller.LEFT.Pressed, "◀"), formatBool(controller.RIGHT.Pressed, "▶"), formatBool(controller.SQUARE.Pressed, "■ SQUARE  "), formatBool(controller.CIRCLE.Pressed, "● CIRCLE  "))
// 		fmt.Printf("    %s                 %s\n", formatBool(controller.DOWN.Pressed, "▼"), formatBool(controller.CROSS.Pressed, "✖ CROSS   "))
// 		fmt.Println()

// 		// Row 4: Joysticks
// 		fmt.Println(" ANALOG STICKS:")
// 		fmt.Printf(" Left Stick:  X: %5.2f, Y: %5.2f  %s\n", controller.LeftStickX, controller.LeftStickY, formatBool(controller.L3.Pressed, "L3 Click"))
// 		fmt.Printf(" Right Stick: X: %5.2f, Y: %5.2f  %s\n", controller.RightStickX, controller.RightStickY, formatBool(controller.R3.Pressed, "R3 Click"))
// 		fmt.Println("==================================================")

// 		time.Sleep(16 * time.Millisecond) // Poll at ~60 FPS
// 	}
// }
