package main

import (
	"mist-os/screen"
)

func main() {
	sc := screen.NewScreen(800, 480)
	if sc == nil {
		return
	}
	defer sc.Close()

	sc.StartLoading(0, 3.0)
	sc.Run()
}
