package camera

import (
	"fmt"
	"time"

	"gocv.io/x/gocv"
)

type Camera struct {
	Device *gocv.VideoCapture
}

func NewCamera(id int) (*Camera, error) {
	webcam, err := gocv.VideoCaptureDevice(id)
	if err != nil {
		return nil, err
	}
	time.Sleep(time.Second) // gives time for a sensor to adjust to the brightness
	return &Camera{Device: webcam}, nil
}

func (c *Camera) Shoot() ([]byte, error) {
	img := gocv.NewMat()
	defer img.Close()

	if ok := c.Device.Read(&img); !ok {
		return nil, fmt.Errorf("camera read failed")
	}
	if img.Empty() {
		return nil, fmt.Errorf("empty frame")
	}

	buf, err := gocv.IMEncode(".png", img)
	if err != nil {
		return nil, err
	}
	return buf.GetBytes(), nil
}

func (c *Camera) Close() {
	c.Device.Close()
}
