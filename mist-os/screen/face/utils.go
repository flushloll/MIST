package face

import (
	"image"
	"image/color"
	"math"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

var mainFont font.Face

func LoadFont(path string, size float64) error {
	fontBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	f, err := opentype.Parse(fontBytes)
	if err != nil {
		return err
	}
	mainFont, err = opentype.NewFace(f, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	return err
}

func DrawTTFChar(img *image.RGBA, center image.Point, char string, angle float64, c color.Color) {
	if mainFont == nil {
		DrawChar(img, center, char, 15, angle, c)
		return
	}

	temp := image.NewRGBA(image.Rect(0, 0, 512, 512))
	d := &font.Drawer{
		Dst:  temp,
		Src:  image.NewUniform(c),
		Face: mainFont,
		Dot:  fixed.P(256, 384),
	}
	d.DrawString(char)

	minX, minY, maxX, maxY := 512, 512, 0, 0
	found := false
	for y := 0; y < 512; y++ {
		for x := 0; x < 512; x++ {
			if _, _, _, a := temp.At(x, y).RGBA(); a > 0 {
				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
				found = true
			}
		}
	}
	if !found {
		return
	}

	charCenter := image.Pt((minX+maxX)/2, (minY+maxY)/2)
	halfW := float64(maxX-minX)/2.0 + 2.0
	halfH := float64(maxY-minY)/2.0 + 2.0
	radius := int(math.Sqrt(halfW*halfW+halfH*halfH)) + 4
	cosA, sinA := math.Sincos(-angle)

	for ty := -radius; ty <= radius; ty++ {
		for tx := -radius; tx <= radius; tx++ {
			sx := float64(tx)*cosA - float64(ty)*sinA + float64(charCenter.X)
			sy := float64(tx)*sinA + float64(ty)*cosA + float64(charCenter.Y)

			x0, y0 := int(math.Floor(sx)), int(math.Floor(sy))
			if x0 < 0 || y0 < 0 || x0 >= 511 || y0 >= 511 {
				continue
			}

			p00 := temp.RGBAAt(x0, y0)
			p10 := temp.RGBAAt(x0+1, y0)
			p01 := temp.RGBAAt(x0, y0+1)
			p11 := temp.RGBAAt(x0+1, y0+1)

			fx, fy := sx-float64(x0), sy-float64(y0)
			a0 := float64(p00.A)*(1-fx) + float64(p10.A)*fx
			a1 := float64(p01.A)*(1-fx) + float64(p11.A)*fx
			finalA := a0*(1-fy) + a1*fy

			if finalA > 0.5 {
				r, g, b, _ := c.RGBA()
				finalColor := color.RGBA{
					R: uint8(r >> 8),
					G: uint8(g >> 8),
					B: uint8(b >> 8),
					A: uint8(math.Round(finalA)),
				}
				img.Set(center.X+tx, center.Y+ty, finalColor)
			}
		}
	}
}

func DrawCircle(img *image.RGBA, center image.Point, radius int, c color.Color) {
	r2 := radius * radius
	bounds := img.Bounds()
	xMin, xMax := center.X-radius, center.X+radius
	yMin, yMax := center.Y-radius, center.Y+radius

	if xMin < bounds.Min.X {
		xMin = bounds.Min.X
	}
	if xMax >= bounds.Max.X {
		xMax = bounds.Max.X - 1
	}
	if yMin < bounds.Min.Y {
		yMin = bounds.Min.Y
	}
	if yMax >= bounds.Max.Y {
		yMax = bounds.Max.Y - 1
	}

	for y := yMin; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {
			dx, dy := x-center.X, y-center.Y
			if dx*dx+dy*dy <= r2 {
				img.Set(x, y, c)
			}
		}
	}
}

func DrawArc(img *image.RGBA, center image.Point, radius int, thickness int, rotation float64, gap float64, c color.Color) {
	rInner, rOuter := float64(radius-thickness/2), float64(radius+thickness/2)
	rInner2, rOuter2 := rInner*rInner, rOuter*rOuter
	normRot := math.Mod(rotation, 2*math.Pi)
	if normRot > math.Pi {
		normRot -= 2 * math.Pi
	} else if normRot < -math.Pi {
		normRot += 2 * math.Pi
	}
	for y := -radius - thickness; y <= radius+thickness; y++ {
		for x := -radius - thickness; x <= radius+thickness; x++ {
			distSq := float64(x*x + y*y)
			if distSq >= rInner2 && distSq <= rOuter2 {
				angle := math.Atan2(float64(y), float64(x))
				diff := math.Abs(angle - normRot)
				if diff > math.Pi {
					diff = 2*math.Pi - diff
				}
				if diff > gap/2.0 {
					img.Set(center.X+x, center.Y+y, c)
				}
			}
		}
	}

	capRadius := thickness / 2
	for _, side := range []float64{-1, 1} {
		angle := normRot + (gap/2.0)*side
		tipX := float64(center.X) + float64(radius)*math.Cos(angle)
		tipY := float64(center.Y) + float64(radius)*math.Sin(angle)
		DrawCircle(img, image.Pt(int(tipX), int(tipY)), capRadius, c)
	}
}

func DrawEllipticalArc(img *image.RGBA, center image.Point, rx, ry int, thickness int, rotation float64, gap float64, c color.Color) {
	cosR, sinR := math.Sincos(rotation)
	bx, by := int(math.Max(float64(rx), float64(ry)))+thickness, int(math.Max(float64(rx), float64(ry)))+thickness

	for y := -by; y <= by; y++ {
		for x := -bx; x <= bx; x++ {
			tx := float64(x)*cosR + float64(y)*sinR
			ty := -float64(x)*sinR + float64(y)*cosR
			dist := (tx/float64(rx))*(tx/float64(rx)) + (ty/float64(ry))*(ty/float64(ry))
			if math.Abs(math.Sqrt(dist)-1.0) < (float64(thickness) / (2.0 * math.Min(float64(rx), float64(ry)))) {
				angle := math.Atan2(ty, tx)
				if math.Abs(angle) > gap/2.0 {
					img.Set(center.X+x, center.Y+y, c)
				}
			}
		}
	}
}

func DrawCross(img *image.RGBA, center image.Point, size int, thickness int, rotation float64, c color.Color) {
	s := size / 2
	p1 := RotatePoint(image.Pt(center.X-s, center.Y-s), center, rotation)
	p2 := RotatePoint(image.Pt(center.X+s, center.Y+s), center, rotation)
	p3 := RotatePoint(image.Pt(center.X+s, center.Y-s), center, rotation)
	p4 := RotatePoint(image.Pt(center.X-s, center.Y+s), center, rotation)

	DrawLine(img, p1, p2, thickness, c)
	DrawLine(img, p3, p4, thickness, c)
}

func DrawRoundedRotatedRect(img *image.RGBA, center image.Point, w, h int, angle float64, r1, r2, r3, r4 float64, c color.Color) {
	cosA, sinA := math.Sincos(angle)
	halfW, halfH := float64(w)/2.0, float64(h)/2.0
	maxR := math.Min(halfW, halfH)

	// Convert normalized radii to pixels
	rr1 := math.Max(0, math.Min(r1, 1.0)) * maxR
	rr2 := math.Max(0, math.Min(r2, 1.0)) * maxR
	rr3 := math.Max(0, math.Min(r3, 1.0)) * maxR
	rr4 := math.Max(0, math.Min(r4, 1.0)) * maxR

	// Determine bounding box for iteration
	bound := int(math.Sqrt(float64(w*w+h*h)))/2 + 4
	for y := -bound; y <= bound; y++ {
		for x := -bound; x <= bound; x++ {
			// Transform to local rect-aligned space
			tx := float64(x)*cosA + float64(y)*sinA
			ty := -float64(x)*sinA + float64(y)*cosA

			// Check if within the main rectangle bounds
			if math.Abs(tx) > halfW || math.Abs(ty) > halfH {
				continue
			}

			// Check corners using robust distance method
			inside := true
			// Top-Left (rr1)
			if tx < -halfW+rr1 && ty < -halfH+rr1 {
				if math.Hypot(tx-(-halfW+rr1), ty-(-halfH+rr1)) > rr1 {
					inside = false
				}
			} else if tx > halfW-rr2 && ty < -halfH+rr2 {
				// Top-Right (rr2)
				if math.Hypot(tx-(halfW-rr2), ty-(-halfH+rr2)) > rr2 {
					inside = false
				}
			} else if tx > halfW-rr3 && ty > halfH-rr3 {
				// Bottom-Right (rr3)
				if math.Hypot(tx-(halfW-rr3), ty-(halfH-rr3)) > rr3 {
					inside = false
				}
			} else if tx < -halfW+rr4 && ty > halfH-rr4 {
				// Bottom-Left (rr4)
				if math.Hypot(tx-(-halfW+rr4), ty-(halfH-rr4)) > rr4 {
					inside = false
				}
			}

			if inside {
				img.Set(center.X+x, center.Y+y, c)
			}
		}
	}
}

func DrawRotatedRect(img *image.RGBA, center image.Point, w, h int, angle float64, c color.Color) {
	DrawRoundedRotatedRect(img, center, w, h, angle, 0, 0, 0, 0, c)
}

func DrawLine(img *image.RGBA, p1, p2 image.Point, thickness int, c color.Color) {
	halfT := float64(thickness) / 2.0
	r2 := halfT * halfT

	xMin := int(math.Min(float64(p1.X), float64(p2.X))) - thickness
	xMax := int(math.Max(float64(p1.X), float64(p2.X))) + thickness
	yMin := int(math.Min(float64(p1.Y), float64(p2.Y))) - thickness
	yMax := int(math.Max(float64(p1.Y), float64(p2.Y))) + thickness

	// Clip to image bounds
	bounds := img.Bounds()
	if xMin < bounds.Min.X {
		xMin = bounds.Min.X
	}
	if xMax >= bounds.Max.X {
		xMax = bounds.Max.X - 1
	}
	if yMin < bounds.Min.Y {
		yMin = bounds.Min.Y
	}
	if yMax >= bounds.Max.Y {
		yMax = bounds.Max.Y - 1
	}

	dx, dy := float64(p2.X-p1.X), float64(p2.Y-p1.Y)
	ab2 := dx*dx + dy*dy

	for y := yMin; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {
			px, py := float64(x), float64(y)
			var t float64
			if ab2 > 0 {
				apx, apy := px-float64(p1.X), py-float64(p1.Y)
				t = (apx*dx + apy*dy) / ab2
				if t < 0 {
					t = 0
				}
				if t > 1 {
					t = 1
				}
			}
			cx := float64(p1.X) + t*dx
			cy := float64(p1.Y) + t*dy
			dist2 := (px-cx)*(px-cx) + (py-cy)*(py-cy)
			if dist2 <= r2 {
				img.Set(x, y, c)
			}
		}
	}
}

func DrawChar(img *image.RGBA, center image.Point, char string, scale float64, angle float64, c color.Color) {
	charW, charH := 7*len(char), 13
	srcImg := image.NewRGBA(image.Rect(0, 0, charW, charH))
	d := &font.Drawer{Dst: srcImg, Src: image.NewUniform(color.White), Face: basicfont.Face7x13}
	d.Dot = fixed.P(0, 10)
	d.DrawString(char)

	cosA, sinA := math.Sincos(angle)
	s := int(scale)
	if s < 1 {
		s = 1
	}

	for sy := 0; sy < charH; sy++ {
		for sx := 0; sx < charW; sx++ {
			if _, _, _, a := srcImg.At(sx, sy).RGBA(); a > 0 {
				lx := (float64(sx) - float64(charW)/2.0) * scale
				ly := (float64(sy) - 6.5) * scale
				halfS := scale / 2.0
				for dy := -halfS; dy <= halfS; dy++ {
					for dx := -halfS; dx <= halfS; dx++ {
						tx, ty := lx+dx, ly+dy
						rx, ry := tx*cosA-ty*sinA, tx*sinA+ty*cosA
						img.Set(int(rx)+center.X, int(ry)+center.Y, c)
					}
				}
			}
		}
	}
}

func RotatePoint(p, center image.Point, angle float64) image.Point {
	sinA, cosA := math.Sincos(angle)
	x, y := float64(p.X-center.X), float64(p.Y-center.Y)
	return image.Pt(int(x*cosA-y*sinA)+center.X, int(x*sinA+y*cosA)+center.Y)
}

func Lerp(a, b, t float64) float64    { return a + (b-a)*t }
func LerpInt(a, b int, t float64) int { return int(math.Round(float64(a) + float64(b-a)*t)) }
func LerpPoint(a, b image.Point, t float64) image.Point {
	return image.Pt(LerpInt(a.X, b.X, t), LerpInt(a.Y, b.Y, t))
}
func LerpColor(a, b color.RGBA, t float64) color.RGBA {
	return color.RGBA{R: uint8(LerpInt(int(a.R), int(b.R), t)), G: uint8(LerpInt(int(a.G), int(b.G), t)), B: uint8(LerpInt(int(a.B), int(b.B), t)), A: uint8(LerpInt(int(a.A), int(b.A), t))}
}
