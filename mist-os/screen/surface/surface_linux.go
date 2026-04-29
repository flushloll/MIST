package surface

import (
	"image"
	"os"
)

type FBSurface struct {
	file *os.File
	w, h int
}

func NewSurface(w, h int) Surface {
	f, err := os.OpenFile("/dev/fb0", os.O_RDWR, 0)
	if err != nil {
		return nil
	}
	return &FBSurface{
		file: f,
		w:    w,
		h:    h,
	}
}

func (s *FBSurface) Present(img *image.RGBA) error {
	if s.file == nil {
		return nil
	}

	// RGB565: 2 bytes per pixel
	hwBuf := make([]byte, s.w*s.h*2)
	for i := 0; i < s.w*s.h; i++ {
		r := uint16(img.Pix[i*4])
		g := uint16(img.Pix[i*4+1])
		b := uint16(img.Pix[i*4+2])

		// Pack into RGB565: RRRRRGGGGGGBBBBB
		// R: 5 bits, G: 6 bits, B: 5 bits
		rgb := ((r >> 3) << 11) | ((g >> 2) << 5) | (b >> 3)

		hwBuf[i*2] = byte(rgb & 0xFF)
		hwBuf[i*2+1] = byte(rgb >> 8)
	}

	s.file.Seek(0, 0)
	_, err := s.file.Write(hwBuf)
	return err
}

func (s *FBSurface) Close() {
	if s.file != nil {
		s.file.Close()
	}
}
