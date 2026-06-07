package mouths

import (
	"image"

	"mist-os/screen/face"
)

type NoneMouth struct {
	face.BaseFeature
}

func (m *NoneMouth) Draw(img *image.RGBA) {}
func (m *NoneMouth) Update(dt float64)    { m.BaseFeature.Update(dt) }
func (m *NoneMouth) IsClosed() bool      { return true }
