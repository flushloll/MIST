package telegraph

import (
	"fmt"
	"mist-os/controller"
	"time"
)

var morseAlphabet = map[string]string{
	".-": "A", "-...": "B", "-.-.": "C", "-..": "D", ".": "E",
	"..-.": "F", "--.": "G", "....": "H", "..": "I", ".---": "J",
	"-.-": "K", ".-..": "L", "--": "M", "-.": "N", "---": "O",
	".--.": "P", "--.-": "Q", ".-.": "R", "...": "S", "-": "T",
	"..-": "U", "...-": "V", ".--": "W", "-..-": "X", "-.--": "Y",
	"--..": "Z",
}

const (
	DotMaxDuration = 200 * time.Millisecond
	LetterTimeout  = 600 * time.Millisecond
	WordTimeout    = 1500 * time.Millisecond
)

type Telegraph struct {
	signal     *controller.Button
	lastState  bool
	sequence   string
	lastLift   time.Time
	lastLetter time.Time
	buffer     string
}

func NewTelegraph(button *controller.Button) *Telegraph {
	return &Telegraph{signal: button}
}

func (t *Telegraph) Tick() {
	if t.lastState && !t.signal.Pressed {
		duration := t.signal.PressDuration()
		if duration <= DotMaxDuration {
			t.sequence += "."
		} else {
			t.sequence += "-"
		}
		t.lastLift = time.Now()
	}
	t.lastState = t.signal.Pressed

	if t.sequence != "" && !t.signal.Pressed && time.Since(t.lastLift) > LetterTimeout {
		if letter, ok := morseAlphabet[t.sequence]; ok {
			t.buffer += letter
			fmt.Println(t.buffer)
			t.lastLetter = time.Now()
		}
		t.sequence = ""
	}
}

func (t *Telegraph) Ready() bool {
	return t.buffer != "" && !t.signal.Pressed && t.sequence == "" &&
		time.Since(t.lastLetter) > WordTimeout
}

func (t *Telegraph) Drain() string {
	cmd := t.buffer
	t.buffer = ""
	return cmd
}

// A playful function for testing
func (t *Telegraph) Execute() {
	if t.buffer == "SOS" {
		fmt.Println("send help immediately")
		t.buffer = ""
	}
}
