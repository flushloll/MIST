package actuators

import (
	"fmt"

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

func (c *Controller) Close() {
	c.driver.Halt()
	c.adaptor.Finalize()
}
