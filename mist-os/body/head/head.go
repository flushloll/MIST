package head

import "mist-os/actuators"

type Head struct {
	Pan *actuators.Servo
}

func New(c *actuators.Controller) *Head {
	return &Head{Pan: actuators.NewServo(c, actuators.CHANNEL_HEAD_SERVO)}
}

func (h *Head) Center() error {
	return h.Pan.SetAngle(90)
}
