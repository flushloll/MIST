package knees

import (
	"errors"
	"mist-os/actuators"
)

const (
	maxHeightDelta = 0.02
	maxTiltDelta   = 0.015
)

type Knees struct {
	FrontRight, FrontLeft, BackRight, BackLeft Leg
}

func New(c *actuators.Controller) *Knees {
	return &Knees{
		FrontRight: Leg{Close: actuators.NewServo(c, actuators.CHANNEL_KNEE_SERVO_FRC), Far: actuators.NewServo(c, actuators.CHANNEL_KNEE_SERVO_FRF)},
		FrontLeft:  Leg{Close: actuators.NewServo(c, actuators.CHANNEL_KNEE_SERVO_FLC), Far: actuators.NewServo(c, actuators.CHANNEL_KNEE_SERVO_FLF)},
		BackRight:  Leg{Close: actuators.NewServo(c, actuators.CHANNEL_KNEE_SERVO_BRC), Far: actuators.NewServo(c, actuators.CHANNEL_KNEE_SERVO_BRF)},
		BackLeft:   Leg{Close: actuators.NewServo(c, actuators.CHANNEL_KNEE_SERVO_BLC), Far: actuators.NewServo(c, actuators.CHANNEL_KNEE_SERVO_BLF)},
	}
}

// Implements L3 controls:
// Horizontal x: tilt, shifts front/back oppositely.
// Vertical: height, shifts all feet equally.
func (k *Knees) SetStance(height, tilt float64) error {
	height = clamp(height, -1, 1)
	tilt = clamp(tilt, -1, 1)

	dy := height * maxHeightDelta
	dTilt := tilt * maxTiltDelta

	legs := []struct {
		leg      *Leg
		tiltSign float64
	}{
		{&k.FrontRight, +1}, {&k.FrontLeft, +1},
		{&k.BackRight, -1}, {&k.BackLeft, -1},
	}

	var errs []error
	for _, l := range legs {
		y := restY + dy + dTilt*l.tiltSign
		if err := l.leg.SetFoot(restX, y); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
