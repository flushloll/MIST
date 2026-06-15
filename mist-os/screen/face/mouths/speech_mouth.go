package mouths

import (
	"image"

	"mist-os/screen/face"
)

type SpeechMouth struct {
	face.BaseFeature
	Width        int
	Height       int
	TargetWidth  int
	TargetHeight int
}

func (m *SpeechMouth) Draw(img *image.RGBA) {
	w := int(float64(m.Width) * m.Scale)
	h := int(float64(m.Height) * m.Scale)

	r1 := 0.1
	r2 := 0.1
	r3 := 1.0
	r4 := 1.0

	face.DrawRoundedRotatedRect(img, m.Position, w, h, m.Rotation, r1, r2, r3, r4, m.Color)
}

func (m *SpeechMouth) IsClosed() bool { return false }

func (m *SpeechMouth) Update(dt float64) {
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
