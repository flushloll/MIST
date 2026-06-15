package mouths

import (
	"image"

	"mist-os/screen/face"
)

type SpeechNerdyMouth struct {
	face.BaseFeature
	Width        int
	Height       int
	TargetWidth  int
	TargetHeight int
}

func (m *SpeechNerdyMouth) Draw(img *image.RGBA) {
	w := int(float64(m.Width) * m.Scale)
	h := int(float64(m.Height) * m.Scale)

	r1 := 0.0
	r2 := 0.0
	r3 := 0.5
	r4 := 0.5

	face.DrawRoundedRotatedRect(img, m.Position, w, h, m.Rotation, r1, r2, r3, r4, m.Color)
}

func (m *SpeechNerdyMouth) IsClosed() bool { return false }

func (m *SpeechNerdyMouth) Update(dt float64) {
	m.BaseFeature.Update(dt)
	if m.TransitionRate <= 0 {
		return
	}
	t := m.TransitionRate * dt * 60.0
	if t > 1.0 {
		t = 1.0
	}
	m.Width = face.LerpInt(m.Width, m.TargetWidth, t)
	m.Height = face.LerpInt(m.Height, m.TargetHeight, t)
}
