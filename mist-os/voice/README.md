# Voice
Raspberry Pi doesn't have a build-in speaker, but USB/Bluetooth speaker should do.

Example usage:
```go
package main

import (
	"mist-os/voice"
)

func main() {
	voice := voice.NewVoiceModule()
	voice.Speak("Hello, I am MIST. My voice system is now operational.", "en+f5", 100, 150)
	for {

	}
}
```

Promising Voices for Mac:
- Karen
- Kyoko
- Meijia
- Melina
- Moira
- Ona
- Samantha
- Tina
- Yuna

Promising Voice for Linux:
speak-ng -v en+f4 -p 125 "There are some people who would like to create, dream" -w morning4.wav