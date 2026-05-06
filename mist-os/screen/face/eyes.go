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

	TargetRadius  int
	TargetGapSize float64
}

func (e *IdleEye) Draw(img *image.RGBA) {
	gapRadians := e.GapSize * 2 * math.Pi
	DrawArc(img, e.Position, int(float64(e.Radius)*e.Scale), e.LineWidth, e.Rotation, gapRadians, e.Color)
}

func (e *IdleEye) Update(dt float64) {
	e.BaseFeature.Update(dt)
	if e.TransitionRate <= 0 {
		return
	}
	t := e.TransitionRate * dt * 60.0
	if t > 1.0 {
		t = 1.0
	}
	e.Radius = LerpInt(e.Radius, e.TargetRadius, t)
	e.GapSize = Lerp(e.GapSize, e.TargetGapSize, t)
}

// Energetic(); like "> <".
type EnergeticEye struct {
	BaseFeature
	Size int

	TargetSize int
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

func (e *EnergeticEye) Update(dt float64) {
	e.BaseFeature.Update(dt)
	if e.TransitionRate <= 0 {
		return
	}
	t := e.TransitionRate * dt * 60.0
	if t > 1.0 {
		t = 1.0
	}
	e.Size = LerpInt(e.Size, e.TargetSize, t)
}

// Soft(height, corner_radius1, corner_radius2); rectangle with one side missing.
type SoftEye struct {
	BaseFeature
	Width  int
	Height int

	TargetWidth  int
	TargetHeight int
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

func (e *SoftEye) Update(dt float64) {
	e.BaseFeature.Update(dt)
	if e.TransitionRate <= 0 {
		return
	}
	t := e.TransitionRate * dt * 60.0
	if t > 1.0 {
		t = 1.0
	}
	e.Width = LerpInt(e.Width, e.TargetWidth, t)
	e.Height = LerpInt(e.Height, e.TargetHeight, t)
}

// Fancy(char, mirrored[true|false]); any character from english alphabet.
type FancyEye struct {
	BaseFeature
	Char     string
	Mirrored bool

	TargetChar string
}

func (e *FancyEye) Draw(img *image.RGBA) {
	DrawChar(img, e.Position, e.Char, e.Scale, e.Rotation, e.Color)
}

func (e *FancyEye) Update(dt float64) {
	e.BaseFeature.Update(dt)
	if e.TargetChar != "" {
		e.Char = e.TargetChar
	}
}
