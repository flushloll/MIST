package wheels

import (
	"errors"
	"math"
	"mist-os/actuators"
)

type Wheels struct {
	FrontRight *actuators.Motor
	FrontLeft  *actuators.Motor
	BackRight  *actuators.Motor
	BackLeft   *actuators.Motor
}

func New(c *actuators.Controller) *Wheels {
	return &Wheels{
		FrontRight: actuators.NewMotor(c, actuators.CHANNEL_WHEEL_FR),
		FrontLeft:  actuators.NewMotor(c, actuators.CHANNEL_WHEEL_FL),
		BackRight:  actuators.NewMotor(c, actuators.CHANNEL_WHEEL_BR),
		BackLeft:   actuators.NewMotor(c, actuators.CHANNEL_WHEEL_BL),
	}
}

func (w *Wheels) all() []*actuators.Motor {
	return []*actuators.Motor{w.FrontRight, w.FrontLeft, w.BackRight, w.BackLeft}
}

func (w *Wheels) SetAll(throttle float64) error {
	for _, m := range w.all() {
		if err := m.SetThrottle(throttle); err != nil {
			return err
		}
	}
	return nil
}

// Implements R3 controls:
// Horizontal x: rotation in place
// Vertical y: driving forward/backward
func (w *Wheels) DriveR3(x, y float64) error {
	left, right := y+x, y-x

	if max := math.Max(math.Abs(left), math.Abs(right)); max > 1.0 {
		left /= max
		right /= max
	}

	var errs []error

	run := func(m *actuators.Motor, speed float64) {
		if err := m.SetThrottle(speed); err != nil {
			errs = append(errs, err)
		}
	}

	run(w.FrontRight, right)
	run(w.BackRight, right)
	run(w.FrontLeft, left)
	run(w.BackLeft, left)

	if len(errs) > 0 {
		_ = w.Stop()
		return errors.Join(errs...)
	}
	return nil
}

func (w *Wheels) Stop() error {
	return w.SetAll(0)
}
