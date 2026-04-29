package main

import (
	"image"
	"image/color"
	"image/draw"
	"log"
	"time"

	"mist-screen/face"
	"mist-screen/surface"
)

const width = 800
const height = 480

func main() {
	s := surface.NewSurface(width, height)
	if s == nil {
		log.Fatal("failed to create surface")
	}
	defer s.Close()

	cyan := color.RGBA{0, 255, 255, 255}
	black := color.RGBA{0, 0, 0, 255}

	f := &face.Face{
		LeftEye: &face.IdleEye{
			BaseFeature: face.BaseFeature{
				Position:  image.Pt(width/3, height/2),
				Scale:     1.0,
				Rotation:  0.0,
				LineWidth: 10,
				Color:     cyan,
			},
			Radius:  60,
			GapSize: 0.2,
		},
		RightEye: &face.IdleEye{
			BaseFeature: face.BaseFeature{
				Position:  image.Pt(2*width/3, height/2),
				Scale:     1.0,
				Rotation:  0.0,
				LineWidth: 10,
				Color:     cyan,
			},
			Radius:  60,
			GapSize: 0.2,
		},
		Mouth: &face.SilentMouth{
			BaseFeature: face.BaseFeature{
				Position:  image.Pt(width/2, 3*height/4),
				Scale:     1.0,
				Rotation:  0.0,
				LineWidth: 8,
				Color:     cyan,
			},
			Count:   3,
			Width:   40,
			Height:  8,
			Spacing: 30,
		},
	}

	for {
		img := image.NewRGBA(image.Rect(0, 0, width, height))

		// 1. Clear with solid black
		draw.Draw(img, img.Bounds(), &image.Uniform{black}, image.Point{}, draw.Src)

		// 2. Draw the static face
		f.Draw(img)

		// 3. Present
		if err := s.Present(img); err != nil {
			log.Printf("failed to present: %v", err)
			break
		}

		time.Sleep(33 * time.Millisecond)
	}
}
