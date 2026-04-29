package face

import (
	"image"
	"math"
)

// IDLE(circle_angle); a circle outline with a cut part of it denoted by rotation parameter.
type IdleEye struct {
	BaseFeature
	Radius  int
	GapSize float64 // 0.0 (full circle) to 1.0 (empty)
}

func (e *IdleEye) Draw(img *image.RGBA) {
	gapRadians := e.GapSize * 2 * math.Pi
	DrawArc(img, e.Position, int(float64(e.Radius)*e.Scale), e.LineWidth, e.Rotation, gapRadians, e.Color)
}

// Energetic(); like "> <".
type EnergeticEye struct {
	BaseFeature
	Size int
}

func (e *EnergeticEye) Draw(img *image.RGBA) {
	s := int(float64(e.Size) * e.Scale)
	center := e.Position

	p1 := RotatePoint(image.Pt(center.X-s/2, center.Y-s/2), center, e.Rotation)
	p2 := RotatePoint(image.Pt(center.X+s/2, center.Y), center, e.Rotation)
	p3 := RotatePoint(image.Pt(center.X-s/2, center.Y+s/2), center, e.Rotation)

	DrawLine(img, p1, p2, e.LineWidth, e.Color)
	DrawLine(img, p2, p3, e.LineWidth, e.Color)
}

// Soft(height, corner_radius1, corner_radius2); rectangle with one side missing.
type SoftEye struct {
	BaseFeature
	Width  int
	Height int
}

func (e *SoftEye) Draw(img *image.RGBA) {
	w := int(float64(e.Width) * e.Scale)
	h := int(float64(e.Height) * e.Scale)
	center := e.Position

	p1 := RotatePoint(image.Pt(center.X-w/2, center.Y+h/2), center, e.Rotation)
	p2 := RotatePoint(image.Pt(center.X-w/2, center.Y-h/2), center, e.Rotation)
	p3 := RotatePoint(image.Pt(center.X+w/2, center.Y-h/2), center, e.Rotation)
	p4 := RotatePoint(image.Pt(center.X+w/2, center.Y+h/2), center, e.Rotation)

	DrawLine(img, p1, p2, e.LineWidth, e.Color)
	DrawLine(img, p2, p3, e.LineWidth, e.Color)
	DrawLine(img, p3, p4, e.LineWidth, e.Color)
}

// Fancy(char, mirrored[true|false]); any character from english alphabet.
type FancyEye struct {
	BaseFeature
	Char     string
	Mirrored bool
}

func (e *FancyEye) Draw(img *image.RGBA) {
	DrawChar(img, e.Position, e.Char, e.Scale, e.Rotation, e.Color)
}
