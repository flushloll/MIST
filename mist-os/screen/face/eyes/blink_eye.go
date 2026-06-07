package eyes

import (
	"image"

	"mist-os/screen/face"
)

type BlinkEye struct {
	face.BaseFeature
	Width       int
	TargetWidth int
}

func (e *BlinkEye) Draw(img *image.RGBA) {
	w := int(float64(e.Width) * e.Scale)
	center := e.Position

	p1 := face.RotatePoint(image.Pt(center.X-w/2, center.Y), center, e.Rotation)
	p2 := face.RotatePoint(image.Pt(center.X+w/2, center.Y), center, e.Rotation)
	face.DrawLine(img, p1, p2, e.LineWidth, e.Color)
}

func (e *BlinkEye) IsClosed() bool { return true }

func (e *BlinkEye) Update(dt float64) {
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
