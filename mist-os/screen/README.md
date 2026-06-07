# MIST-OS Face Styles

Eyes:
- idle
- energetic
- soft
- soft-sad
- fancy-<insert character>

Mouths:
- none
- speech
- speech-nerdy
- silent
- soft
- cutie

Usage guide:
```go
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

	sc.StartLoading(0, 1)
	sc.SetFace("soft-sad", "silent")
	sc.Run()
}
```

# TODO
- Fix many of the mouth styles (many are ugly or too low)
- Make so in SetFace, eyes and mouths change simultaneously.