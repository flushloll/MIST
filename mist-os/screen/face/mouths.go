package face

import (
	"image"
)

// None; in a case of only-eyes.
type NoneMouth struct {
	BaseFeature
}

func (m *NoneMouth) Draw(img *image.RGBA) {}

// Silent(count); line(s) or dot(s).
type SilentMouth struct {
	BaseFeature
	Width   int
	Height  int
	Spacing int
	Count   int
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

// Speech(height, corner_radius...); filled in rectangle with rounded corners.
type SpeechMouth struct {
	BaseFeature
	Width  int
	Height int
}

func (m *SpeechMouth) Draw(img *image.RGBA) {
	w := int(float64(m.Width) * m.Scale)
	h := int(float64(m.Height) * m.Scale)
	DrawRotatedRect(img, m.Position, w, h, m.Rotation, m.Color)
}

// Soft(corner_radius); an outline similar to AND or OR in set theory.
type SoftMouth struct {
	BaseFeature
	Width  int
	Height int
}

func (m *SoftMouth) Draw(img *image.RGBA) {
	w := int(float64(m.Width) * m.Scale)
	DrawArc(img, m.Position, w/2, m.LineWidth, m.Rotation+3.14, 0.5, m.Color)
}

// Cutie(); uwu-styled w.
type CutieMouth struct {
	BaseFeature
	Size int
}

func (m *CutieMouth) Draw(img *image.RGBA) {
	s := int(float64(m.Size) * m.Scale)
	center := m.Position
	rot := m.Rotation - 1.57 // -PI/2

	leftCenter := RotatePoint(image.Pt(center.X-s/2, center.Y), center, m.Rotation)
	rightCenter := RotatePoint(image.Pt(center.X+s/2, center.Y), center, m.Rotation)

	DrawArc(img, leftCenter, s/2, m.LineWidth, rot, 3.14, m.Color)
	DrawArc(img, rightCenter, s/2, m.LineWidth, rot, 3.14, m.Color)
}
