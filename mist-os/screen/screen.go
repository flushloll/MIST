package screen

import (
	"image"
	"image/color"
	"image/draw"
	"math"
	"math/rand"
	"time"

	"mist-os/screen/face"
	"mist-os/screen/surface"
)

type Screen struct {
	surface    surface.Surface
	Face       *face.Face
	LookTarget *image.Point
	Width      int
	Height     int
	Color      color.RGBA

	EyeSpacing int
	EyeY       int
	MouthY     int
	Radius     int
	LineWidth  int
	MouthW     int
	MouthH     int

	// Animation State
	Mode      string
	startTime time.Time
	duration  float64
	animType  int

	// Blinking State
	nextBlink    time.Time
	isBlinking   bool
	currentEyes  string
	currentMouth string

	dirty bool
}

func NewScreen(w, h int) *Screen {
	s := surface.NewSurface(w, h)
	if s == nil {
		return nil
	}

	c := color.RGBA{0, 255, 255, 255}
	sc := &Screen{
		surface: s, Width: w, Height: h, Color: c,
		EyeSpacing: 350, EyeY: 205, MouthY: 365,
		Radius: 75, LineWidth: 18, MouthW: 165, MouthH: 20,
		Mode:        "neutral",
		currentEyes: "idle",
		nextBlink:   time.Now().Add(3 * time.Second),
		dirty:       true,
	}

	lx, rx := w/2-sc.EyeSpacing/2, w/2+sc.EyeSpacing/2
	sc.Face = &face.Face{
		LookLimit: 60, LookSensitivity: 0.3,
		LeftEye:  sc.newIdleEye(lx, sc.EyeY, math.Pi),
		RightEye: sc.newIdleEye(rx, sc.EyeY, 0),
		Mouth:    &face.NoneMouth{},
	}

	return sc
}

func (s *Screen) newIdleEye(x, y int, rot float64) *face.IdleEye {
	return &face.IdleEye{
		BaseFeature: face.BaseFeature{
			Position: image.Pt(x, y), AnchorPosition: image.Pt(x, y), TargetPosition: image.Pt(x, y),
			Scale: 1.0, TargetScale: 1.0, Rotation: rot, TargetRotation: rot,
			LineWidth: s.LineWidth, TargetLineWidth: s.LineWidth, Color: s.Color, TargetColor: s.Color,
			TransitionRate: 0.15,
		},
		Radius: s.Radius, TargetRadius: s.Radius, GapSize: 0.2, TargetGapSize: 0.2,
	}
}

func (s *Screen) NewEyes(style string) {
	if style == "blink" {
		s.applyEyeStyle("blink")
		return
	}
	s.currentEyes = style
	if !s.isBlinking {
		s.applyEyeStyle(style)
	}
}

func (s *Screen) applyEyeStyle(style string) {
	lx, rx := s.Width/2-s.EyeSpacing/2, s.Width/2+s.EyeSpacing/2
	create := func(x int, rot float64, old face.Feature) face.Feature {
		base := face.BaseFeature{
			Position: image.Pt(x, s.EyeY), AnchorPosition: image.Pt(x, s.EyeY), TargetPosition: image.Pt(x, s.EyeY),
			Scale: 1.0, TargetScale: 1.0, Rotation: rot, TargetRotation: rot,
			LineWidth: s.LineWidth, TargetLineWidth: s.LineWidth, Color: s.Color, TargetColor: s.Color,
			TransitionRate: 0.25,
		}
		if old != nil {
			b := old.GetBase()
			base.Position, base.Scale, base.Rotation = b.Position, b.Scale, b.Rotation
		}

		switch style {
		case "energetic":
			return &face.EnergeticEye{BaseFeature: base, Size: s.Radius, TargetSize: s.Radius}
		case "soft":
			return &face.SoftEye{BaseFeature: base, Width: s.Radius * 2, Height: s.Radius, TargetWidth: s.Radius * 2, TargetHeight: s.Radius}
		case "fancy":
			return &face.FancyEye{BaseFeature: base, Char: "?"}
		case "blink":
			return &face.SoftEye{BaseFeature: base, Width: s.Radius * 2, Height: 0, TargetWidth: s.Radius * 2, TargetHeight: 0}
		default:
			return &face.IdleEye{BaseFeature: base, Radius: s.Radius, TargetRadius: s.Radius, GapSize: 0.2, TargetGapSize: 0.2}
		}
	}
	s.Face.LeftEye = create(lx, math.Pi, s.Face.LeftEye)
	s.Face.RightEye = create(rx, 0, s.Face.RightEye)
	s.dirty = true
}

func (s *Screen) NewMouth(style string) {
	s.currentMouth = style
	base := face.BaseFeature{
		Position: image.Pt(s.Width/2, s.MouthY), AnchorPosition: image.Pt(s.Width/2, s.MouthY), TargetPosition: image.Pt(s.Width/2, s.MouthY),
		Scale: 1.0, TargetScale: 1.0, Rotation: 0, TargetRotation: 0,
		LineWidth: s.LineWidth, TargetLineWidth: s.LineWidth, Color: s.Color, TargetColor: s.Color, TransitionRate: 0.15,
	}
	if s.Face.Mouth != nil {
		b := s.Face.Mouth.GetBase()
		base.Position, base.Scale, base.Rotation = b.Position, b.Scale, b.Rotation
	}

	switch style {
	case "speech":
		base.Rotation, base.TargetRotation = math.Pi/2, math.Pi/2
		s.Face.Mouth = &face.SpeechMouth{BaseFeature: base, Width: s.MouthW, Height: s.MouthH, TLRadius: 0.5, TRRadius: 0.5, BRRadius: 0.5, BLRadius: 0.5, TargetWidth: s.MouthW, TargetHeight: s.MouthH, TargetTLRadius: 0.5, TargetTRRadius: 0.5, TargetBRRadius: 0.5, TargetBLRadius: 0.5}
	case "silent":
		s.Face.Mouth = &face.SilentMouth{BaseFeature: base, Count: 1, Width: s.MouthW, Height: 0, Spacing: 0, TargetCount: 1, TargetWidth: s.MouthW, TargetHeight: 0, TargetSpacing: 0}
	case "soft":
		s.Face.Mouth = &face.SoftMouth{BaseFeature: base, Width: s.MouthW, Height: s.MouthH, TargetWidth: s.MouthW, TargetHeight: s.MouthH}
	case "cutie":
		s.Face.Mouth = &face.CutieMouth{BaseFeature: base, Size: s.MouthW / 2, TargetSize: s.MouthW / 2}
	default:
		s.Face.Mouth = &face.NoneMouth{BaseFeature: base}
	}
	s.dirty = true
}

func (s *Screen) StartLoading(animType int, duration float64) {
	s.Mode = "loading"
	s.animType = animType
	s.startTime = time.Now()
	s.duration = duration
	s.NewMouth("none")
	s.Face.LookTarget = nil
	s.dirty = true
}

func (s *Screen) Update(dt float64) {
	// 1. Handle loading logic
	elapsed := time.Since(s.startTime).Seconds()
	if s.Mode == "loading" {
		progress := math.Min(1.0, elapsed/s.duration)

		speed := 2.0 + progress*28.0
		if s.animType == 0 {
			// Old: Synchronous rotation
			if le := s.Face.LeftEye.GetBase(); le != nil {
				le.RotationSpeed = speed
			}
			if re := s.Face.RightEye.GetBase(); re != nil {
				re.RotationSpeed = speed
			}
		} else {
			// New: Mirrored rotation
			if le := s.Face.LeftEye.GetBase(); le != nil {
				le.RotationSpeed = speed
			}
			if re := s.Face.RightEye.GetBase(); re != nil {
				re.RotationSpeed = -speed
			}
		}

		gap := 1.0 - progress
		s.setGapSize(gap)
		if progress >= 1.0 {
			s.Mode = "transitioning"
			s.startTime = time.Now()

			// Reset rotations and stop spinning
			if le := s.Face.LeftEye.GetBase(); le != nil {
				le.RotationSpeed = 0
				le.Rotation = math.Pi
				le.TargetRotation = math.Pi
			}
			if re := s.Face.RightEye.GetBase(); re != nil {
				re.RotationSpeed = 0
				re.Rotation = 0
				re.TargetRotation = 0
			}

			s.setTargetGapSize(0.2)
			s.NewMouth("speech")
			s.Face.Mouth.GetBase().Scale = 0
			s.Face.Mouth.GetBase().TargetScale = 1.0
		}
		s.dirty = true
	} else if s.Mode == "transitioning" && elapsed > 1.5 {
		s.Mode = "looking"
	}

	// 2. Handle blinking logic
	if s.Mode != "loading" && s.Mode != "transitioning" && time.Now().After(s.nextBlink) && !s.isBlinking {
		s.performBlink()
	}

	// 3. Sync positions
	lx, rx := s.Width/2-s.EyeSpacing/2, s.Width/2+s.EyeSpacing/2
	s.syncFeature(s.Face.LeftEye, lx, s.EyeY, s.Radius)
	s.syncFeature(s.Face.RightEye, rx, s.EyeY, s.Radius)
	s.syncFeature(s.Face.Mouth, s.Width/2, s.MouthY, 0)

	s.Face.Update(dt)
	s.dirty = true
}

func (s *Screen) performBlink() {
	s.isBlinking = true
	s.applyEyeStyle("blink")

	dur := time.Duration(150+rand.Intn(150)) * time.Millisecond
	time.AfterFunc(dur, func() {
		s.isBlinking = false
		s.applyEyeStyle(s.currentEyes)
		s.nextBlink = time.Now().Add(time.Duration(3000+rand.Intn(3000)) * time.Millisecond)
	})
}

func (s *Screen) syncFeature(feat face.Feature, x, y, radius int) {
	if feat == nil {
		return
	}
	b := feat.GetBase()
	b.AnchorPosition = image.Pt(x, y)
	if s.Face.LookTarget == nil {
		b.TargetPosition = b.AnchorPosition
	}
	if eye, ok := feat.(*face.IdleEye); ok {
		eye.TargetRadius = radius
	}
}

func (s *Screen) setGapSize(g float64) {
	if le, ok := s.Face.LeftEye.(*face.IdleEye); ok {
		le.GapSize, le.TargetGapSize = g, g
	}
	if re, ok := s.Face.RightEye.(*face.IdleEye); ok {
		re.GapSize, re.TargetGapSize = g, g
	}
}

func (s *Screen) setTargetGapSize(g float64) {
	if le, ok := s.Face.LeftEye.(*face.IdleEye); ok {
		le.TargetGapSize = g
	}
	if re, ok := s.Face.RightEye.(*face.IdleEye); ok {
		re.TargetGapSize = g
	}
}

func (s *Screen) DrawAndPresent() {
	if !s.dirty {
		return
	}
	img := image.NewRGBA(image.Rect(0, 0, s.Width, s.Height))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{0, 0, 0, 255}}, image.Point{}, draw.Src)
	s.Face.Draw(img)
	s.surface.Present(img)
	s.dirty = false
}

func (s *Screen) Run() {
	for {
		for event := s.surface.PollEvent(); event != nil; event = s.surface.PollEvent() {
			if s.surface.IsQuitEvent(event) {
				return
			}
		}
		s.Update(0.033)
		s.DrawAndPresent()
		time.Sleep(33 * time.Millisecond)
	}
}

func (s *Screen) Close() { s.surface.Close() }
