package surface

import (
	"image"
	"log"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

type SDLSurface struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	texture  *sdl.Texture
	w, h     int
}

func NewSurface(w, h int) Surface {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		log.Printf("failed to init SDL: %v", err)
		return nil
	}

	window, err := sdl.CreateWindow("Mist Screen Simulator", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(w), int32(h), sdl.WINDOW_SHOWN)
	if err != nil {
		log.Printf("failed to create window: %v", err)
		return nil
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Printf("failed to create renderer: %v", err)
		return nil
	}

	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(w), int32(h))
	if err != nil {
		log.Printf("failed to create texture: %v", err)
		return nil
	}

	return &SDLSurface{
		window:   window,
		renderer: renderer,
		texture:  texture,
		w:        w,
		h:        h,
	}
}

func (s *SDLSurface) Present(img *image.RGBA) error {
	if s.texture == nil {
		return nil
	}

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			return nil
		}
	}

	s.texture.Update(nil, unsafe.Pointer(&img.Pix[0]), img.Stride)
	s.renderer.Clear()
	s.renderer.Copy(s.texture, nil, nil)
	s.renderer.Present()
	return nil
}

func (s *SDLSurface) Close() {
	if s.texture != nil {
		s.texture.Destroy()
	}
	if s.renderer != nil {
		s.renderer.Destroy()
	}
	if s.window != nil {
		s.window.Destroy()
	}
	sdl.Quit()
}
