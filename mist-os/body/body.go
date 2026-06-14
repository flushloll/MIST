package body

import (
	"context"
	"mist-os/actuators"
	"mist-os/body/head"
	"mist-os/body/hips"
	"mist-os/body/knees"
	"mist-os/body/wheels"
	"mist-os/imu"
	"time"
)

const (
	maxPitchDeg = 15.0
	imuPeriod   = 20 * time.Millisecond
)

type Body struct {
	controller *actuators.Controller
	imu        *imu.IMU
	cancel     context.CancelFunc

	Head   *head.Head
	Wheels *wheels.Wheels
	Knees  *knees.Knees
	Hips   *hips.Hips
}

func New() (*Body, error) {
	c, err := actuators.NewController()
	if err != nil {
		return nil, err
	}

	i, err := imu.New()
	if err != nil {
		c.Close()
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	go i.Stream(ctx, imuPeriod)

	return &Body{
		controller: c,
		imu:        i,
		cancel:     cancel,
		Head:       head.New(c),
		Wheels:     wheels.New(c),
		Knees:      knees.New(c),
		Hips:       hips.New(c),
	}, nil
}

func (b *Body) Arm() error {
	for _, m := range []*actuators.Motor{
		b.Wheels.FrontRight, b.Wheels.FrontLeft,
		b.Wheels.BackRight, b.Wheels.BackLeft,
	} {
		if err := m.Arm(); err != nil {
			return err
		}
	}
	return nil
}

func (b *Body) tiltCorrection() float64 {
	pitch := b.imu.Latest().Pitch()
	tilt := -pitch / maxPitchDeg
	if tilt > 1 {
		return 1
	}
	if tilt < -1 {
		return -1
	}
	return tilt
}

// Ride-drone controls
// R3: rx - rotation in place, ry - forward/backward
// L3: lx - steering angle, ry - height
func (b *Body) Ride(rx, ry, lx, ly float64) error {
	if err := b.Wheels.DriveR3(rx, ry); err != nil {
		return err
	}
	if err := b.Hips.Steer(lx); err != nil {
		return err
	}
	return b.Knees.SetStance(ly, b.tiltCorrection())
}

func (b *Body) Stop() error {
	if err := b.Wheels.Stop(); err != nil {
		return err
	}
	if err := b.Hips.Steer(0); err != nil {
		return err
	}
	return b.Knees.SetStance(0, 0)
}

func (b *Body) Close() {
	b.cancel()
	b.imu.Close()
	b.controller.Close()
}
