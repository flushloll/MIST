package eyes

import (
	"image"

	"mist-os/screen/face"
)

type CrossEye struct {
	face.BaseFeature
	Size int
}

func (e *CrossEye) Draw(img *image.RGBA) {
	s := int(float64(e.Size) * e.Scale)
	face.DrawCross(img, e.Position, s, e.LineWidth, e.Rotation, e.Color)
}

func (e *CrossEye) IsClosed() bool { return true }

func (e *CrossEye) Update(dt float64) {
	e.BaseFeature.Update(dt)
}
