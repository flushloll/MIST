package mouths

import (
	"image"
	"math"

	"mist-os/screen/face"
)

type SoftMouth struct {
	face.BaseFeature
	Width       int
	TargetWidth int
}

func (m *SoftMouth) Draw(img *image.RGBA) {
	w := int(float64(m.Width) * m.Scale)
	// Rotation directed to the top (-Pi/2) with a gap of Pi
	// draws a full semi-circle (U)
	face.DrawArc(img, m.Position, w/2, m.LineWidth, m.Rotation - math.Pi/2, math.Pi, m.Color)
}

func (m *SoftMouth) IsClosed() bool { return false }

func (m *SoftMouth) Update(dt float64) {
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
