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
	// Rotation directed to the top (-Pi/2) with a larger gap (4.7 radians)
	// so only a gentle bottom arc is drawn, forming a natural smile (U)
	face.DrawArc(img, m.Position, w/2, m.LineWidth, m.Rotation - math.Pi/2, 4.7, m.Color)
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
