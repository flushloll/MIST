package hips

import "mist-os/actuators"

const (
	SteerRest = 90.0
	SteerMin  = 60.0
	SteerMax  = 120.0

	maxSteerDeg = 30.0
)

type Hips struct {
	Right *actuators.Servo
	Left  *actuators.Servo
}

func New(c *actuators.Controller) *Hips {
	return &Hips{
		Right: actuators.NewServo(c, actuators.CHANNEL_HIP_RIGHT),
		Left:  actuators.NewServo(c, actuators.CHANNEL_HIP_LEFT),
	}
}

func (h *Hips) Steer(x float64) error {
	x = clamp(x, -1, 1)
	angle := SteerRest + x*maxSteerDeg

	if err := h.Right.SetAngle(angle); err != nil {
		return err
	}
	return h.Left.SetAngle(angle)
}

func clamp(v, lo, hi float64) float64 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
