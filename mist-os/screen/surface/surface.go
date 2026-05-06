package surface

import (
	"image"
)

type Surface interface {
	Present(img *image.RGBA) error
	GetMousePos() (x, y int, available bool)
	PollEvent() any
	IsQuitEvent(event any) bool
	Close()
}
