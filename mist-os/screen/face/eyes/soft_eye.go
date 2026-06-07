package eyes

import (
	"image"
	"math"

	"mist-os/screen/face"
)

type SoftEye struct {
	face.BaseFeature
	Width       int
	TargetWidth int
}

func (e *SoftEye) Draw(img *image.RGBA) {
	w := int(float64(e.Width) * e.Scale)
	// Rotation of Pi/2 directs the gap to the bottom (forming an "n" shape facing up)
	face.DrawArc(img, e.Position, w/2, e.LineWidth, math.Pi/2, math.Pi, e.Color)
}

func (e *SoftEye) IsClosed() bool { return true }

func (e *SoftEye) Update(dt float64) {
	e.BaseFeature.Update(dt)
	if e.TransitionRate <= 0 {
		return
	}
	t := e.TransitionRate * dt * 60.0
	if t > 1.0 {
		t = 1.0
	}
	e.Width = face.LerpInt(e.Width, e.TargetWidth, t)
}
