# Controller
This is a component that abstracts PS5 controller movements into simple package.

Usage guide:
```go
package main

import (
	"fmt"
	"mist-os/controller"
	"time"
)

func main() {
	controller, _ := controller.NewController()
	for {
		controller.Update()
		if controller.CROSS { // all of the buttons and axis are in Controller struct.
			fmt.Printf("CROSS is hled down %v\n", time.Now())
		}
	}
}
```