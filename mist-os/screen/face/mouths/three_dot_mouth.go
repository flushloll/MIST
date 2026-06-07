package mouths

import (
	"image"

	"mist-os/screen/face"
)

type ThreeDotMouth struct {
	face.BaseFeature
	Radius        int
	Spacing       int
	TargetRadius  int
	TargetSpacing int
}

func (m *ThreeDotMouth) Draw(img *image.RGBA) {
	r := int(float64(m.Radius) * m.Scale)
	spacing := int(float64(m.Spacing) * m.Scale)
	center := m.Position

	leftCenter := face.RotatePoint(image.Pt(center.X-spacing, center.Y), center, m.Rotation)
	rightCenter := face.RotatePoint(image.Pt(center.X+spacing, center.Y), center, m.Rotation)

	face.DrawCircle(img, leftCenter, r, m.Color)
	face.DrawCircle(img, center, r, m.Color)
	face.DrawCircle(img, rightCenter, r, m.Color)
}

func (m *ThreeDotMouth) IsClosed() bool { return true }

func (m *ThreeDotMouth) Update(dt float64) {
	m.BaseFeature.Update(dt)
	if m.TransitionRate <= 0 {
		return
	}
	t := m.TransitionRate * dt * 60.0
	if t > 1.0 {
		t = 1.0
	}
	m.Radius = face.LerpInt(m.Radius, m.TargetRadius, t)
	m.Spacing = face.LerpInt(m.Spacing, m.TargetSpacing, t)
}
