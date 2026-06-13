package wheels

import "mist-os/actuators"

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

func (w *Wheels) Stop() error {
	return w.SetAll(0)
}
