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
All of the eyes are great! Some of the mouses are great! But rest are sort of ugly... Here is breakdown of what's ugly:
- speech -> ugly, 
- speech-nerdy -> ugly
- soft -> must be higher
- cutie -> must be higher
- speech-happy -> ugly & low
- speech-curve -> same as speech happy, why exists?
- three-dot -> too small & low
general feedback: speech modules are supposed to be like the mouth here: :D but filled. The soft modules are very similar but are more like :) but the buttom part of the shape curves like D but without a vertical line. 
Should I just design it all as svg's and we should just import them rather than draw it all in go?