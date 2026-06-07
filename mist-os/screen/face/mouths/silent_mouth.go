package mouths

import (
	"image"

	"mist-os/screen/face"
)

type SilentMouth struct {
	face.BaseFeature
	Width       int
	TargetWidth int
}

func (m *SilentMouth) Draw(img *image.RGBA) {
	w := int(float64(m.Width) * m.Scale)
	center := m.Position

	p1 := face.RotatePoint(image.Pt(center.X-w/2, center.Y), center, m.Rotation)
	p2 := face.RotatePoint(image.Pt(center.X+w/2, center.Y), center, m.Rotation)
	face.DrawLine(img, p1, p2, m.LineWidth, m.Color)
}

func (m *SilentMouth) IsClosed() bool { return true }

func (m *SilentMouth) Update(dt float64) {
	m.BaseFeature.Update(dt)
	if m.TransitionRate <= 0 {
		return
	}
	t := m.TransitionRate * dt * 60.0
	if t > 1.0 {
		t = 1.0
	}
	m.Width = face.LerpInt(m.Width, m.TargetWidth, t)
}
