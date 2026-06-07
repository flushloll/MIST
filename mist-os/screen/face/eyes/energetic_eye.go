package eyes

import (
	"image"

	"mist-os/screen/face"
)

type EnergeticEye struct {
	face.BaseFeature
	Size       int
	TargetSize int
}

func (e *EnergeticEye) Draw(img *image.RGBA) {
	s := int(float64(e.Size) * e.Scale)
	center := e.Position

	p1 := face.RotatePoint(image.Pt(center.X+s/2, center.Y-s/2), center, e.Rotation)
	p2 := face.RotatePoint(image.Pt(center.X-s/2, center.Y), center, e.Rotation)
	p3 := face.RotatePoint(image.Pt(center.X+s/2, center.Y+s/2), center, e.Rotation)

	face.DrawLine(img, p1, p2, e.LineWidth, e.Color)
	face.DrawLine(img, p2, p3, e.LineWidth, e.Color)
}

func (e *EnergeticEye) IsClosed() bool { return true }

func (e *EnergeticEye) Update(dt float64) {
	e.BaseFeature.Update(dt)
	if e.TransitionRate <= 0 {
		return
	}
	t := e.TransitionRate * dt * 60.0
	if t > 1.0 {
		t = 1.0
	}
	e.Size = face.LerpInt(e.Size, e.TargetSize, t)
}
