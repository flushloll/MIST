package actuators

import (
	"math"
	"time"

	"gobot.io/x/gobot/drivers/i2c"
)

type Servo struct {
	driver  *i2c.PCA9685Driver
	channel int
	angle   float64
}

func NewServo(c *Controller, channel int) *Servo {
	return &Servo{driver: c.driver, channel: channel}
}

func (s *Servo) SetAngle(angle float64) error {
	angle = normaliseAngle(angle)

	pulseUs := servoMin + int((angle/180)*float64(servoMax-servoMin))
	tick := usToTick(pulseUs)

	return s.driver.SetPWM(s.channel, 0, uint16(tick))
}

func (s *Servo) MoveTo(angle float64, degPerSec float64, stepDelay time.Duration) error {
	angle = normaliseAngle(angle)

	stepDeg := degPerSec * stepDelay.Seconds()
	if stepDeg <= 0 {
		stepDeg = 1
	}

	for {
		diff := angle - s.angle
		if diff == 0 {
			break
		}

		step := stepDeg
		if step > math.Abs(diff) {
			step = math.Abs(diff)
		}
		if diff < 0 {
			step = -step
		}

		if err := s.SetAngle(s.angle + step); err != nil {
			return err
		}
		time.Sleep(stepDelay)

		if s.angle == angle {
			break
		}
	}

	return nil
}

func normaliseAngle(angle float64) float64 {
	if angle < 0 {
		return 0
	}
	if angle > 180 {
		return 180
	}
	return angle
}
