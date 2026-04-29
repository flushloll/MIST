package face

import (
	"image"
	"image/color"
	"math"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

type Feature interface {
	Draw(img *image.RGBA)
	GetBase() *BaseFeature
}

type BaseFeature struct {
	Position  image.Point
	Scale     float64
	Rotation  float64
	LineWidth int
	Color     color.Color
}

func (b *BaseFeature) GetBase() *BaseFeature {
	return b
}

type Face struct {
	LeftEye  Feature
	RightEye Feature
	Mouth    Feature
}

func (f *Face) Draw(img *image.RGBA) {
	if f.LeftEye != nil { f.LeftEye.Draw(img) }
	if f.RightEye != nil { f.RightEye.Draw(img) }
	if f.Mouth != nil { f.Mouth.Draw(img) }
}

// --- SOLID HELPERS ---

func DrawCircle(img *image.RGBA, center image.Point, radius int, c color.Color) {
	r2 := radius * radius
	for y := -radius; y <= radius; y++ {
		for x := -radius; x <= radius; x++ {
			if x*x+y*y <= r2 {
				img.Set(center.X+x, center.Y+y, c)
			}
		}
	}
}

func DrawArc(img *image.RGBA, center image.Point, radius int, thickness int, rotation float64, gap float64, c color.Color) {
	rInner := float64(radius - thickness/2)
	rOuter := float64(radius + thickness/2)
	rInner2 := rInner * rInner
	rOuter2 := rOuter * rOuter

	for y := -radius - thickness; y <= radius+thickness; y++ {
		for x := -radius - thickness; x <= radius+thickness; x++ {
			distSq := float64(x*x + y*y)
			if distSq >= rInner2 && distSq <= rOuter2 {
				angle := math.Atan2(float64(y), float64(x))
				diff := math.Abs(angle - rotation)
				if diff > math.Pi { diff = 2*math.Pi - diff }

				if diff > gap/2.0 {
					img.Set(center.X+x, center.Y+y, c)
				}
			}
		}
	}
}

func DrawRotatedRect(img *image.RGBA, center image.Point, w, h int, angle float64, c color.Color) {
	cosA, sinA := math.Sincos(angle)
	halfW, halfH := float64(w)/2.0, float64(h)/2.0
	bound := int(math.Sqrt(float64(w*w + h*h))) / 2 + 1
	for y := -bound; y <= bound; y++ {
		for x := -bound; x <= bound; x++ {
			tx := float64(x)*cosA + float64(y)*sinA
			ty := -float64(x)*sinA + float64(y)*cosA
			if math.Abs(tx) <= halfW && math.Abs(ty) <= halfH {
				img.Set(center.X+x, center.Y+y, c)
			}
		}
	}
}

func RotatePoint(p, center image.Point, angle float64) image.Point {
	sinA, cosA := math.Sincos(angle)
	x, y := float64(p.X-center.X), float64(p.Y-center.Y)
	return image.Pt(int(x*cosA-y*sinA)+center.X, int(x*sinA+y*cosA)+center.Y)
}

func DrawLine(img *image.RGBA, p1, p2 image.Point, thickness int, c color.Color) {
	dx, dy := float64(p2.X-p1.X), float64(p2.Y-p1.Y)
	length := math.Sqrt(dx*dx + dy*dy)
	if length == 0 {
		DrawCircle(img, p1, thickness/2, c)
		return
	}
	for i := 0.0; i <= length; i += 1.0 {
		px := float64(p1.X) + (i/length)*dx
		py := float64(p1.Y) + (i/length)*dy
		DrawCircle(img, image.Pt(int(px), int(py)), thickness/2, c)
	}
}

func DrawChar(img *image.RGBA, center image.Point, char string, scale float64, angle float64, c color.Color) {
	charW, charH := 7*len(char), 13
	srcImg := image.NewRGBA(image.Rect(0, 0, charW, charH))
	d := &font.Drawer{Dst: srcImg, Src: image.NewUniform(color.White), Face: basicfont.Face7x13}
	d.Dot = fixed.P(0, 10)
	d.DrawString(char)

	cosA, sinA := math.Sincos(angle)
	for sy := 0; sy < charH; sy++ {
		for sx := 0; sx < charW; sx++ {
			if _, _, _, a := srcImg.At(sx, sy).RGBA(); a > 0 {
				tx := (float64(sx) - float64(charW)/2.0) * scale
				ty := (float64(sy) - float64(charH)/2.0) * scale
				rx := tx*cosA - ty*sinA
				ry := tx*sinA + ty*cosA
				DrawCircle(img, image.Pt(int(rx)+center.X, int(ry)+center.Y), int(scale/2.0)+1, c)
			}
		}
	}
}
