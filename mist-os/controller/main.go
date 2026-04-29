package main

import (
	"fmt"
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	runtime.LockOSThread()
	if err := sdl.Init(sdl.INIT_GAMECONTROLLER); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	var controller *sdl.GameController

	fmt.Println("waiting for ps5 controller")

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.ControllerDeviceEvent:
				if t.Type == sdl.CONTROLLERDEVICEADDED {
					controller = sdl.GameControllerOpen(int(t.Which))
					fmt.Printf("controller connected: %s\n", controller.Name())
				}
			case *sdl.ControllerButtonEvent:
				state := "Released"
				if t.State == sdl.PRESSED {
					state = "Pressed"
				}
				fmt.Printf("Button %d %s\n", t.Button, state)
			case *sdl.ControllerAxisEvent:
				if t.Value > 5000 || t.Value < -5000 { // range: -32768 to 32767
					fmt.Printf("Axis %d moved to %d\n", t.Axis, t.Value)
				}
			}
		}
		sdl.Delay(20)
	}
}
