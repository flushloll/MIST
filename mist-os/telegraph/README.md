# Morse
This module allows to use a Button from a Controller as a telegraph, converting morse code into text, that can be used as commands.

Usage guide:
```go
package main

import (
	"mist-os/controller"
	"mist-os/telegraph"
	"time"
)

func main() {
	controller, _ := controller.NewController()
	defer controller.Close()

	telegraph := telegraph.NewTelegraph(&controller.CIRCLE)

	for {
		controller.Update()
		telegraph.Tick()
		telegraph.Execute()
		time.Sleep(10 * time.Millisecond)
	}
}
```