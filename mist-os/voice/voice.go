package voice

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
)

type VoiceModule struct {
	queue chan []string
}

func NewVoiceModule() *VoiceModule {
	vm := &VoiceModule{
		queue: make(chan []string, 10),
	}
	go vm.worker()
	return vm
}

func (vm *VoiceModule) worker() {
	for task := range vm.queue {
		text, voice, speed, pitch := task[0], task[1], task[2], task[3]
		var cmd *exec.Cmd
		switch runtime.GOOS {
		case "darwin":
			pitchVal, _ := strconv.Atoi(pitch)
			formattedText := fmt.Sprintf("[[pbas %d]] %s", pitchVal*10, text)
			cmd = exec.Command("say", "-v", voice, "-r", speed, formattedText)
		case "linux":
			// espeak-ng -v en-us+f2+whisper "MIST core initialization complete." -w mist_ghostly.wav
			cmd = exec.Command("espeak", "-v", voice, "-s", speed, "-p", pitch, text)
		default:
			fmt.Println("unsupported platform :(")
			continue
		}
		cmd.Run()
	}
}

func (vm *VoiceModule) Speak(text string, voice string, speed int, pitch int) {
	select {
	case vm.queue <- []string{text, voice, strconv.Itoa(speed), strconv.Itoa(pitch)}:
	default:
		fmt.Println("speech queue full... \ndropping message.")
	}
}
