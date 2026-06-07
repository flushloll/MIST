package mouths

import (
	"image"
	"math"

	"mist-os/screen/face"
)

type SpeechCurveMouth struct {
	face.BaseFeature
	Width       int
	TargetWidth int
}

func (m *SpeechCurveMouth) Draw(img *image.RGBA) {
	w := int(float64(m.Width) * m.Scale)
	// A wide open speech curve facing up
	face.DrawArc(img, m.Position, w/2, m.LineWidth, m.Rotation - math.Pi/2, 3.14, m.Color)
}

func (m *SpeechCurveMouth) IsClosed() bool { return false }

func (m *SpeechCurveMouth) Update(dt float64) {
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
