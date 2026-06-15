package brain

import (
	"fmt"
	"mist-os/body"
	"mist-os/controller"
	"mist-os/screen"
	"mist-os/telegraph"
	"time"
)

const tickRate = 33 * time.Millisecond

type faceCommand struct {
	eye, mouth string
}

var faceCommands = map[string]faceCommand{
	"H": {"energetic", "speech-happy"},
	"S": {"soft-sad", "speech-sad"},
	"O": {"soft", "soft"},
	"I": {"idle", "silent"},
	"D": {"fancy-X", "speech-sad"},
}

type Brain struct {
	Body       *body.Body
	Controller *controller.Controller
	Screen     *screen.Screen
	Telegraph  *telegraph.Telegraph
}

func New() (*Brain, error) {
	b, err := body.New()
	if err != nil {
		return nil, fmt.Errorf("brain error: body init error: %w", err)
	}

	c, err := controller.NewController()
	if err != nil {
		b.Close()
		return nil, fmt.Errorf("brain error: controller init error: %w", err)
	}

	sc := screen.NewScreen(800, 400)
	if sc == nil {
		c.Close()
		b.Close()
		return nil, fmt.Errorf("brain error: screen init error: %w", err)
	}

	return &Brain{
		Body:       b,
		Controller: c,
		Screen:     sc,
		Telegraph:  telegraph.NewTelegraph(&c.CIRCLE),
	}, nil
}

func (b *Brain) Close() {
	b.Screen.Close()
	b.Controller.Close()
	b.Body.Close()
}

func (b *Brain) applyFaceCommand(cmd string) {
	face, ok := faceCommands[cmd]
	if !ok {
		fmt.Printf("brain: unknown telegraph command %q\n", cmd)
		return
	}
	b.Screen.SetFace(face.eye, face.mouth)
}

func (b *Brain) Run(stop <-chan struct{}) {
	b.Screen.StartLoading(0, 5)
	b.Screen.SetFace("idle", "silent")

	ticker := time.NewTicker(tickRate)
	defer ticker.Stop()

	for {
		select {
		case <-stop:
			_ = b.Body.Stop()
			return
		case <-ticker.C:
		}

		b.Controller.Update()
		if b.Controller.Quit {
			_ = b.Body.Stop()
			return
		}

		b.Telegraph.Tick()
		if b.Telegraph.Ready() {
			b.applyFaceCommand(b.Telegraph.Drain())
		}

		rx := float64(b.Controller.RightStickX)
		ry := float64(b.Controller.RightStickY)
		lx := float64(b.Controller.LeftStickX)
		ly := float64(b.Controller.LeftStickY)

		if err := b.Body.Ride(rx, ry, lx, ly); err != nil {
			fmt.Printf("brain error: ride error: %v\n", err)
		}

		b.Screen.Update(tickRate.Seconds())
		b.Screen.DrawAndPresent()
	}
}
