package eyes

import (
	"image"

	"mist-os/screen/face"
)

type FancyEye struct {
	face.BaseFeature
	Char       string
	TargetChar string
}

func (e *FancyEye) Draw(img *image.RGBA) {
	face.DrawTTFChar(img, e.Position, e.Char, e.Rotation, e.Color)
}

func (e *FancyEye) IsClosed() bool { return false }

func (e *FancyEye) Update(dt float64) {
	e.BaseFeature.Update(dt)
	if e.TargetChar != "" {
		e.Char = e.TargetChar
	}
}
