package mouths

import (
	"image"

	"mist-os/screen/face"
)

type DotMouth struct {
	face.BaseFeature
	Radius       int
	TargetRadius int
}

func (m *DotMouth) Draw(img *image.RGBA) {
	r := int(float64(m.Radius) * m.Scale)
	face.DrawCircle(img, m.Position, r, m.Color)
}

func (m *DotMouth) IsClosed() bool { return true }

func (m *DotMouth) Update(dt float64) {
	m.BaseFeature.Update(dt)
	if m.TransitionRate <= 0 {
		return
	}
	t := m.TransitionRate * dt * 60.0
	if t > 1.0 {
		t = 1.0
	}
	m.Radius = face.LerpInt(m.Radius, m.TargetRadius, t)
}
