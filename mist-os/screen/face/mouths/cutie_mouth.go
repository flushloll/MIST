package mouths

import (
	"image"

	"mist-os/screen/face"
)

type CutieMouth struct {
	face.BaseFeature
	Size       int
	TargetSize int
}

func (m *CutieMouth) Draw(img *image.RGBA) {
	s := int(float64(m.Size) * m.Scale)
	center := m.Position
	rot := m.Rotation - 1.57 // directs gap to the top

	leftCenter := face.RotatePoint(image.Pt(center.X-s/2, center.Y), center, m.Rotation)
	rightCenter := face.RotatePoint(image.Pt(center.X+s/2, center.Y), center, m.Rotation)

	face.DrawArc(img, leftCenter, s/2, m.LineWidth, rot, 3.14, m.Color)
	face.DrawArc(img, rightCenter, s/2, m.LineWidth, rot, 3.14, m.Color)
}

func (m *CutieMouth) IsClosed() bool { return false }

func (m *CutieMouth) Update(dt float64) {
	m.BaseFeature.Update(dt)
	if m.TransitionRate <= 0 {
		return
	}
	t := m.TransitionRate * dt * 60.0
	if t > 1.0 {
		t = 1.0
	}
	m.Size = face.LerpInt(m.Size, m.TargetSize, t)
}
