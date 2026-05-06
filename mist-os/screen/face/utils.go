package face

import (
	"image"
	"image/color"
	"math"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

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

func DrawRoundedRotatedRect(img *image.RGBA, center image.Point, w, h int, angle float64, r1, r2, r3, r4 float64, c color.Color) {
	cosA, sinA := math.Sincos(angle)
	halfW, halfH := float64(w)/2.0, float64(h)/2.0
	maxR := math.Min(halfW, halfH)
	r1, r2, r3, r4 = math.Min(r1, maxR), math.Min(r2, maxR), math.Min(r3, maxR), math.Min(r4, maxR)
	bound := int(math.Sqrt(float64(w*w+h*h)))/2 + 1
	for y := -bound; y <= bound; y++ {
		for x := -bound; x <= bound; x++ {
			tx := float64(x)*cosA + float64(y)*sinA
			ty := -float64(x)*sinA + float64(y)*cosA
			if math.Abs(tx) <= halfW && math.Abs(ty) <= halfH {
				var cx, cy, r float64
				ic := false
				if tx < -halfW+r1 && ty < -halfH+r1 {
					cx, cy, r, ic = -halfW+r1, -halfH+r1, r1, true
				} else if tx > halfW-r2 && ty < -halfH+r2 {
					cx, cy, r, ic = halfW-r2, -halfH+r2, r2, true
				} else if tx > halfW-r3 && ty > halfH-r3 {
					cx, cy, r, ic = halfW-r3, halfH-r3, r3, true
				} else if tx < -halfW+r4 && ty > halfH-r4 {
					cx, cy, r, ic = -halfW+r4, halfH+r4, r4, true
				}
				if !ic || (tx-cx)*(tx-cx)+(ty-cy)*(ty-cy) <= r*r {
					img.Set(center.X+x, center.Y+y, c)
				}
			}
		}
	}
}

func DrawRotatedRect(img *image.RGBA, center image.Point, w, h int, angle float64, c color.Color) {
	DrawRoundedRotatedRect(img, center, w, h, angle, 0, 0, 0, 0, c)
}

func DrawLine(img *image.RGBA, p1, p2 image.Point, thickness int, c color.Color) {
	dx, dy := float64(p2.X-p1.X), float64(p2.Y-p1.Y)
	length := math.Sqrt(dx*dx + dy*dy)
	if length == 0 {
		DrawCircle(img, p1, thickness/2, c)
		return
	}
	for i := 0.0; i <= length; i += 1.0 {
		px, py := float64(p1.X)+(i/length)*dx, float64(p1.Y)+(i/length)*dy
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
				tx, ty := (float64(sx)-float64(charW)/2.0)*scale, (float64(sy)-float64(charH)/2.0)*scale
				rx, ry := tx*cosA-ty*sinA, tx*sinA+ty*cosA
				DrawCircle(img, image.Pt(int(rx)+center.X, int(ry)+center.Y), int(scale/2.0)+1, c)
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
