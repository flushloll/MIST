package surface

import "image"

type Surface interface {
	Present(img *image.RGBA) error
	Close()
}
