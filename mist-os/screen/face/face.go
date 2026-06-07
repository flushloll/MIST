package face

import (
	"image"
	"image/color"
	"math"
)

type Feature interface {
	Draw(img *image.RGBA)
	Update(dt float64)
	GetBase() *BaseFeature
	IsClosed() bool
}

type BaseFeature struct {
	Position  image.Point
	Scale     float64
	Rotation  float64
	LineWidth int
	Color     color.RGBA

	AnchorPosition  image.Point
	TargetPosition  image.Point
	TargetScale     float64
	TargetRotation  float64
	TargetLineWidth int
	TargetColor     color.RGBA

	RotationSpeed  float64
	TransitionRate float64
}

func (b *BaseFeature) GetBase() *BaseFeature {
	return b
}

func (b *BaseFeature) Update(dt float64) {
	if b.TransitionRate <= 0 {
		return
	}
	t := b.TransitionRate * dt * 60.0
	if t > 1.0 {
		t = 1.0
	}

	b.Position = LerpPoint(b.Position, b.TargetPosition, t)
	b.Scale = Lerp(b.Scale, b.TargetScale, t)
	if b.RotationSpeed == 0 {
		b.Rotation = Lerp(b.Rotation, b.TargetRotation, t)
	}
	b.LineWidth = LerpInt(b.LineWidth, b.TargetLineWidth, t)
	b.Color = LerpColor(b.Color, b.TargetColor, t)

	b.Rotation += b.RotationSpeed * dt
}

type Face struct {
	LeftEye  Feature
	RightEye Feature
	Mouth    Feature

	LookTarget      *image.Point
	LookLimit       int
	LookSensitivity float64
}

func (f *Face) Draw(img *image.RGBA) {
	if f.LeftEye != nil {
		f.LeftEye.Draw(img)
	}
	if f.RightEye != nil {
		f.RightEye.Draw(img)
	}
	if f.Mouth != nil {
		f.Mouth.Draw(img)
	}
}

func (f *Face) Update(dt float64) {
	if f.LookTarget != nil {
		f.updateGaze()
	}
	if f.LeftEye != nil {
		f.LeftEye.Update(dt)
	}
	if f.RightEye != nil {
		f.RightEye.Update(dt)
	}
	if f.Mouth != nil {
		f.Mouth.Update(dt)
	}
}

func (f *Face) updateGaze() {
	target := *f.LookTarget
	limit := float64(f.LookLimit)
	sensitivity := f.LookSensitivity
	if sensitivity <= 0 {
		sensitivity = 0.1
	}

	var cx, cy, count int
	for _, eye := range []Feature{f.LeftEye, f.RightEye} {
		if eye != nil {
			b := eye.GetBase()
			cx += b.AnchorPosition.X
			cy += b.AnchorPosition.Y
			count++
		}
	}
	if count == 0 {
		return
	}

	faceCenter := image.Pt(cx/count, cy/count)
	dx := float64(target.X - faceCenter.X)
	dy := float64(target.Y - faceCenter.Y)
	dist := math.Sqrt(dx*dx + dy*dy)

	if dist > 0 {
		mag := math.Min(dist*sensitivity, limit)
		offX := int((dx / dist) * mag)
		offY := int((dy / dist) * mag)

		for _, feat := range []Feature{f.LeftEye, f.RightEye, f.Mouth} {
			if feat != nil {
				b := feat.GetBase()
				b.TargetPosition = image.Pt(b.AnchorPosition.X+offX, b.AnchorPosition.Y+offY)
			}
		}
	}
}
