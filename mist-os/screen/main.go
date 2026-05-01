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
				Position:       image.Pt(width/3, height/2),
				Scale:          1.0,
				Rotation:       0.0,
				LineWidth:      10,
				Color:          cyan,
				TargetPosition: image.Pt(width/3, height/2),
				TargetScale:    1.0,
				TargetRotation: 0.0,
				TargetColor:    cyan,
				TransitionRate: 0.05,
			},
			Radius:        60,
			GapSize:       0.2,
			TargetRadius:  60,
			TargetGapSize: 0.2,
		},
		RightEye: &face.IdleEye{
			BaseFeature: face.BaseFeature{
				Position:       image.Pt(2*width/3, height/2),
				Scale:          1.0,
				Rotation:       0.0,
				LineWidth:      10,
				Color:          cyan,
				TargetPosition: image.Pt(2*width/3, height/2),
				TargetScale:    1.0,
				TargetRotation: 0.0,
				TargetColor:    cyan,
				TransitionRate: 0.05,
			},
			Radius:        60,
			GapSize:       0.2,
			TargetRadius:  60,
			TargetGapSize: 0.2,
		},
		Mouth: &face.SpeechMouth{
			BaseFeature: face.BaseFeature{
				Position:       image.Pt(width/2, 3*height/4),
				Scale:          1.0,
				Rotation:       1.57079632679,
				LineWidth:      8,
				Color:          cyan,
				TargetPosition: image.Pt(width/2, 3*height/4),
				TargetScale:    1.0,
				TargetRotation: 1.57079632679,
				TargetColor:    cyan,
				TransitionRate: 0.05,
			},
			Width:          80,
			Height:         20,
			TLRadius:       0.5,
			TRRadius:       0.5,
			BRRadius:       0.5,
			BLRadius:       0.5,
			TargetWidth:    80,
			TargetHeight:   20,
			TargetTLRadius: 0.5,
			TargetTRRadius: 0.5,
			TargetBRRadius: 0.5,
			TargetBLRadius: 0.5,
		},
	}

	lastSwap := time.Now()
	isExcited := false

	for {
		// Toggle state every 3 seconds
		if time.Since(lastSwap) > 3*time.Second {
			isExcited = !isExcited
			lastSwap = time.Now()

			leftEye := f.LeftEye.(*face.IdleEye)
			rightEye := f.RightEye.(*face.IdleEye)
			mouth := f.Mouth.(*face.SpeechMouth)

			if isExcited {
				// Excited state
				leftEye.TargetPosition = image.Pt(width/3, height/2-40)
				leftEye.TargetScale = 1.3
				leftEye.TargetRadius = 80
				leftEye.TargetGapSize = 0.1 // Smaller gap when excited

				rightEye.TargetPosition = image.Pt(2*width/3, height/2-40)
				rightEye.TargetScale = 1.3
				rightEye.TargetRadius = 80
				rightEye.TargetGapSize = 0.1

				mouth.TargetWidth = 150
				mouth.TargetHeight = 60
				mouth.TargetTLRadius = 0.2
				mouth.TargetTRRadius = 0.2
				mouth.TargetBRRadius = 1.0
				mouth.TargetBLRadius = 1.0
			} else {
				// Neutral state
				leftEye.TargetPosition = image.Pt(width/3, height/2)
				leftEye.TargetScale = 1.0
				leftEye.TargetRadius = 60
				leftEye.TargetGapSize = 0.2

				rightEye.TargetPosition = image.Pt(2*width/3, height/2)
				rightEye.TargetScale = 1.0
				rightEye.TargetRadius = 60
				rightEye.TargetGapSize = 0.2

				mouth.TargetWidth = 80
				mouth.TargetHeight = 20
				mouth.TargetTLRadius = 0.5
				mouth.TargetTRRadius = 0.5
				mouth.TargetBRRadius = 0.5
				mouth.TargetBLRadius = 0.5
			}
		}

		img := image.NewRGBA(image.Rect(0, 0, width, height))

		// 1. Clear with solid black
		draw.Draw(img, img.Bounds(), &image.Uniform{black}, image.Point{}, draw.Src)

		// 2. Update and draw the face
		f.Update(0.033) // ~30fps
		f.Draw(img)

		// 3. Present
		if err := s.Present(img); err != nil {
			log.Printf("failed to present: %v", err)
			break
		}

		time.Sleep(33 * time.Millisecond)
	}
}
