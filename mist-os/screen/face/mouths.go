package face

import (
	"image"
	"math"
)

// None; in a case of only-eyes.
type NoneMouth struct {
	BaseFeature
}

func (m *NoneMouth) Draw(img *image.RGBA) {}
func (m *NoneMouth) Update(dt float64)    { m.BaseFeature.Update(dt) }

// Silent(count); line(s) or dot(s).
type SilentMouth struct {
	BaseFeature
	Width   int
	Height  int
	Spacing int
	Count   int

	TargetWidth   int
	TargetHeight  int
	TargetSpacing int
	TargetCount   int
}

func (m *SilentMouth) Draw(img *image.RGBA) {
	s := m.Scale
	w := int(float64(m.Width) * s)
	h := int(float64(m.Height) * s)
	spacing := int(float64(m.Spacing) * s)
	center := m.Position

	for i := 0; i < m.Count; i++ {
		offsetX := (i - m.Count/2) * spacing
		if m.Count%2 == 0 {
			offsetX = int(float64(i-m.Count/2+1) * float64(spacing))
		}

		segCenter := RotatePoint(image.Pt(center.X+offsetX, center.Y), center, m.Rotation)

		if h <= m.LineWidth {
			p1 := RotatePoint(image.Pt(center.X+offsetX-w/2, center.Y), center, m.Rotation)
			p2 := RotatePoint(image.Pt(center.X+offsetX+w/2, center.Y), center, m.Rotation)
			DrawLine(img, p1, p2, m.LineWidth, m.Color)
		} else {
			DrawRotatedRect(img, segCenter, w, h, m.Rotation, m.Color)
		}
	}
}

func (m *SilentMouth) Update(dt float64) {
	m.BaseFeature.Update(dt)
	if m.TransitionRate <= 0 {
		return
	}
	t := m.TransitionRate * dt * 60.0
	if t > 1.0 {
		t = 1.0
	}
	m.Width = LerpInt(m.Width, m.TargetWidth, t)
	m.Height = LerpInt(m.Height, m.TargetHeight, t)
	m.Spacing = LerpInt(m.Spacing, m.TargetSpacing, t)
	m.Count = m.TargetCount // Count snaps
}

// Speech(height, corner_radius...); filled in rectangle with rounded corners.
type SpeechMouth struct {
	BaseFeature
	Width    int
	Height   int
	TLRadius float64
	TRRadius float64
	BRRadius float64
	BLRadius float64

	TargetWidth    int
	TargetHeight   int
	TargetTLRadius float64
	TargetTRRadius float64
	TargetBRRadius float64
	TargetBLRadius float64
}

func (m *SpeechMouth) Draw(img *image.RGBA) {
	w := int(float64(m.Width) * m.Scale)
	h := int(float64(m.Height) * m.Scale)

	maxR := math.Min(float64(w), float64(h)) / 2.0
	r1 := m.TLRadius * maxR
	r2 := m.TRRadius * maxR
	r3 := m.BRRadius * maxR
	r4 := m.BLRadius * maxR

	DrawRoundedRotatedRect(img, m.Position, w, h, m.Rotation, r1, r2, r3, r4, m.Color)
}

func (m *SpeechMouth) Update(dt float64) {
	m.BaseFeature.Update(dt)
	if m.TransitionRate <= 0 {
		return
	}
	t := m.TransitionRate * dt * 60.0
	if t > 1.0 {
		t = 1.0
	}
	m.Width = LerpInt(m.Width, m.TargetWidth, t)
	m.Height = LerpInt(m.Height, m.TargetHeight, t)
	m.TLRadius = Lerp(m.TLRadius, m.TargetTLRadius, t)
	m.TRRadius = Lerp(m.TRRadius, m.TargetTRRadius, t)
	m.BRRadius = Lerp(m.BRRadius, m.TargetBRRadius, t)
	m.BLRadius = Lerp(m.BLRadius, m.TargetBLRadius, t)
}

// Soft(corner_radius); an outline similar to AND or OR in set theory.
type SoftMouth struct {
	BaseFeature
	Width  int
	Height int

	TargetWidth  int
	TargetHeight int
}

func (m *SoftMouth) Draw(img *image.RGBA) {
	w := int(float64(m.Width) * m.Scale)
	DrawArc(img, m.Position, w/2, m.LineWidth, m.Rotation+3.14, 0.5, m.Color)
}

func (m *SoftMouth) Update(dt float64) {
	m.BaseFeature.Update(dt)
	if m.TransitionRate <= 0 {
		return
	}
	t := m.TransitionRate * dt * 60.0
	if t > 1.0 {
		t = 1.0
	}
	m.Width = LerpInt(m.Width, m.TargetWidth, t)
	m.Height = LerpInt(m.Height, m.TargetHeight, t)
}

// Cutie(); uwu-styled w.
type CutieMouth struct {
	BaseFeature
	Size int

	TargetSize int
}

func (m *CutieMouth) Draw(img *image.RGBA) {
	s := int(float64(m.Size) * m.Scale)
	center := m.Position
	rot := m.Rotation - 1.57

	leftCenter := RotatePoint(image.Pt(center.X-s/2, center.Y), center, m.Rotation)
	rightCenter := RotatePoint(image.Pt(center.X+s/2, center.Y), center, m.Rotation)

	DrawArc(img, leftCenter, s/2, m.LineWidth, rot, 3.14, m.Color)
	DrawArc(img, rightCenter, s/2, m.LineWidth, rot, 3.14, m.Color)
}

func (m *CutieMouth) Update(dt float64) {
	m.BaseFeature.Update(dt)
	if m.TransitionRate <= 0 {
		return
	}
	t := m.TransitionRate * dt * 60.0
	if t > 1.0 {
		t = 1.0
	}
	m.Size = LerpInt(m.Size, m.TargetSize, t)
}
