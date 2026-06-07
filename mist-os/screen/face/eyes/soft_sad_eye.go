package eyes

import (
	"image"
	"math"

	"mist-os/screen/face"
)

type SoftSadEye struct {
	face.BaseFeature
	Width       int
	TargetWidth int
}

func (e *SoftSadEye) Draw(img *image.RGBA) {
	w := int(float64(e.Width) * e.Scale)
	// Rotation of 3*Pi/2 directs the gap to the top (forming a "U" shape facing down)
	face.DrawArc(img, e.Position, w/2, e.LineWidth, 3*math.Pi/2, math.Pi, e.Color)
}

func (e *SoftSadEye) IsClosed() bool { return true }

func (e *SoftSadEye) Update(dt float64) {
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
