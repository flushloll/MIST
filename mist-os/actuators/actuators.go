package actuators

import (
	"fmt"
	"time"

	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

const (
	pwmFreq       = 50   // Hz
	pwmResolution = 4096 // PCA9685 is 12bit

	servoMin = 1000 // 1ms, 0 degrees
	servoMax = 2000 // 2ms, 180 degree

	escMin = 1000 // 1ms, stop
	escMax = 2000 // 2ms, full throttle
)

type Controller struct {
	driver  *i2c.PCA9685Driver
	adaptor *raspi.Adaptor
}

func NewController() (*Controller, error) {
	adaptor := raspi.NewAdaptor()
	if err := adaptor.Connect(); err != nil {
		return nil, fmt.Errorf("actuator connection error: %w", err)
	}

	driver := i2c.NewPCA9685Driver(adaptor)
	if err := driver.Start(); err != nil {
		return nil, fmt.Errorf("driver start error: %w", err)
	}

	if err := driver.SetPWMFreq(pwmFreq); err != nil {
		return nil, fmt.Errorf("actuator set frequency error: %w", err)
	}

	return &Controller{driver: driver, adaptor: adaptor}, nil
}

// converts pulse in microseconds to PCA9685 tick count.
func usToTick(us int) int {
	periodUs := 1_000_000 / pwmFreq
	return us * pwmResolution / periodUs
}

// moves servo on a given channel to given angle (0 to 180)
func (c *Controller) SetServo(channel int, angleDeg float64) error {
	if angleDeg < 0 {
		angleDeg = 0
	}
	if angleDeg > 180 {
		angleDeg = 180
	}

	pulseUs := servoMin + int((angleDeg/180)*float64(servoMax-servoMin))
	tick := usToTick(pulseUs)

	return c.driver.SetPWM(channel, 0, uint16(tick))
}

// sets motor throttle on the given channel (0.0 to 1.0)
func (c *Controller) SetThrottle(channel int, throttle float64) error {
	if throttle < 0 {
		throttle = 0
	}
	if throttle > 1 {
		throttle = 1
	}

	pulseUs := escMin + int(throttle*float64(escMax-escMin))
	tick := usToTick(pulseUs)

	return c.driver.SetPWM(channel, 0, uint16(tick))
}

// heats up the esc
func (c *Controller) ArmESC(channel int) error {
	if err := c.SetThrottle(channel, 0); err != nil {
		return fmt.Errorf("arm channel %d error: %w", channel, err)
	}
	time.Sleep(2 * time.Second)
	return nil
}

func (c *Controller) Close() {
	c.driver.Halt()
	c.adaptor.Finalize()
}
