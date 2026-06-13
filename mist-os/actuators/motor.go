package actuators

import (
	"fmt"
	"time"

	"gobot.io/x/gobot/drivers/i2c"
)

type Motor struct {
	driver  *i2c.PCA9685Driver
	channel int
}

func NewMotor(c *Controller, channel int) *Motor {
	return &Motor{driver: c.driver, channel: channel}
}

// sets motor throttle on the given channel (0.0 to 1.0)
func (m *Motor) SetThrottle(throttle float64) error {
	if throttle < 0 {
		throttle = 0
	}
	if throttle > 1 {
		throttle = 1
	}

	pulseUs := escMin + int(throttle*float64(escMax-escMin))
	tick := usToTick(pulseUs)

	return m.driver.SetPWM(m.channel, 0, uint16(tick))
}

// heats up the esc
func (m *Motor) Arm() error {
	if err := m.SetThrottle(0); err != nil {
		return fmt.Errorf("arm channel %d error: %w", m.channel, err)
	}
	time.Sleep(2 * time.Second)
	return nil
}

func (m *Motor) Stop() error {
	return m.SetThrottle(0)
}
