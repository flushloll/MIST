package eyes

import (
	"image"
	"math"

	"mist-os/screen/face"
)

type IdleEye struct {
	face.BaseFeature
	Radius        int
	GapSize       float64 // 0.0 (full circle) to 1.0 (empty)
	TargetRadius  int
	TargetGapSize float64
}

func (e *IdleEye) Draw(img *image.RGBA) {
	gapRadians := e.GapSize * 2 * math.Pi
	face.DrawArc(img, e.Position, int(float64(e.Radius)*e.Scale), e.LineWidth, e.Rotation, gapRadians, e.Color)
}

func (e *IdleEye) IsClosed() bool { return false }

func (e *IdleEye) Update(dt float64) {
	e.BaseFeature.Update(dt)
	if e.TransitionRate <= 0 {
		return
	}
	t := e.TransitionRate * dt * 60.0
	if t > 1.0 {
		t = 1.0
	}
	e.Radius = face.LerpInt(e.Radius, e.TargetRadius, t)
	e.GapSize = face.Lerp(e.GapSize, e.TargetGapSize, t)
}
