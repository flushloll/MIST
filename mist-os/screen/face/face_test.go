package face_test

import (
	"image"
	"image/color"
	"math"
	"testing"

	"mist-os/screen/face"
	"mist-os/screen/face/eyes"
	"mist-os/screen/face/mouths"
)

func TestFaceFeatures(t *testing.T) {
	// Base Feature settings
	colorCyan := color.RGBA{0, 255, 255, 255}
	base := face.BaseFeature{
		Position: image.Pt(25, 25), AnchorPosition: image.Pt(25, 25), TargetPosition: image.Pt(25, 25),
		Scale: 1.0, TargetScale: 1.0, Rotation: 0, TargetRotation: 0,
		LineWidth: 4, TargetLineWidth: 4, Color: colorCyan, TargetColor: colorCyan,
		TransitionRate: 0.15,
	}

	// Create test image to render on
	img := image.NewRGBA(image.Rect(0, 0, 50, 50))

	// Define all eyes to test
	testEyes := []face.Feature{
		&eyes.IdleEye{BaseFeature: base, Radius: 10, GapSize: 0.2, TargetRadius: 10, TargetGapSize: 0.2},
		&eyes.EnergeticEye{BaseFeature: base, Size: 15, TargetSize: 15},
		&eyes.SoftEye{BaseFeature: base, Width: 20, TargetWidth: 20},
		&eyes.SoftSadEye{BaseFeature: base, Width: 20, TargetWidth: 20},
		&eyes.CrossEye{BaseFeature: base, Size: 20},
		&eyes.FancyEye{BaseFeature: base, Char: "E"},
		&eyes.BlinkEye{BaseFeature: base, Width: 20, TargetWidth: 20},
	}

	// Define all mouths to test
	testMouths := []face.Feature{
		&mouths.NoneMouth{BaseFeature: base},
		&mouths.SilentMouth{BaseFeature: base, Width: 20, TargetWidth: 20},
		&mouths.SpeechMouth{BaseFeature: base, Width: 20, Height: 10, TargetWidth: 20, TargetHeight: 10},
		&mouths.SpeechNerdyMouth{BaseFeature: base, Width: 10, Height: 20, TargetWidth: 10, TargetHeight: 20},
		&mouths.SoftMouth{BaseFeature: base, Width: 20, TargetWidth: 20},
		&mouths.CutieMouth{BaseFeature: base, Size: 10, TargetSize: 10},
		&mouths.DotMouth{BaseFeature: base, Radius: 5, TargetRadius: 5},
		&mouths.ThreeDotMouth{BaseFeature: base, Radius: 3, Spacing: 8, TargetRadius: 3, TargetSpacing: 8},
		&mouths.SpeechCurveMouth{BaseFeature: base, Width: 20, TargetWidth: 20},
		&mouths.SpeechHappyMouth{BaseFeature: base, Width: 20, TargetWidth: 20},
		&mouths.SpeechSadMouth{BaseFeature: base, Width: 20, TargetWidth: 20},
	}

	// Helper function to test feature operations
	testFeature := func(f face.Feature, name string) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("%s panicked: %v", name, r)
			}
		}()

		// Test Base Getter
		if f.GetBase() == nil {
			t.Errorf("%s returned nil base feature", name)
		}

		// Test Update
		f.Update(0.016)

		// Test Draw
		f.Draw(img)
	}

	// Run tests on eyes
	for _, eye := range testEyes {
		testFeature(eye, "Eye type")
	}

	// Run tests on mouths
	for _, mouth := range testMouths {
		testFeature(mouth, "Mouth type")
	}
}

func TestGazeUpdate(t *testing.T) {
	// Test face look direction limits and sensitivity logic
	lookT := image.Pt(100, 100)
	f := &face.Face{
		LookTarget:      &lookT,
		LookLimit:       30,
		LookSensitivity: 0.5,
		LeftEye: &eyes.IdleEye{
			BaseFeature: face.BaseFeature{
				Position: image.Pt(200, 200), AnchorPosition: image.Pt(200, 200), TargetPosition: image.Pt(200, 200),
				Scale: 1.0, TargetScale: 1.0, Rotation: math.Pi, TargetRotation: math.Pi,
				LineWidth: 10, TargetLineWidth: 10, Color: color.RGBA{255, 255, 255, 255}, TargetColor: color.RGBA{255, 255, 255, 255},
				TransitionRate: 0.15,
			},
		},
		RightEye: &eyes.IdleEye{
			BaseFeature: face.BaseFeature{
				Position: image.Pt(400, 200), AnchorPosition: image.Pt(400, 200), TargetPosition: image.Pt(400, 200),
				Scale: 1.0, TargetScale: 1.0, Rotation: 0, TargetRotation: 0,
				LineWidth: 10, TargetLineWidth: 10, Color: color.RGBA{255, 255, 255, 255}, TargetColor: color.RGBA{255, 255, 255, 255},
				TransitionRate: 0.15,
			},
		},
		Mouth: &mouths.NoneMouth{},
	}

	// Update face to invoke updateGaze
	f.Update(0.016)

	// Verify target position shifted towards the look target
	lBase := f.LeftEye.GetBase()
	if lBase.TargetPosition == lBase.AnchorPosition {
		t.Error("expected left eye target position to shift towards look target")
	}
}
