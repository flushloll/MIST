package knees

import (
	"fmt"
	"math"
	"mist-os/actuators"
)

const (
	L1 = 0.0293 // distance between 2 servo centers
	L2 = 0.126  // distance between servo and motor center

	CloseRest = 90.0
	CloseMin  = 0.0
	CloseMax  = 180.0

	FarRest = 90.0
	FarMin  = 0.0
	FarMax  = 180.0

	deg2rad = math.Pi / 180
)

// returns foot/wheel position (x(horizontal), y(vertical) meters) relative to the hip
func footFK(closeDeg, farDeg float64) (x, y float64) {
	theta1 := (closeDeg - CloseRest) * deg2rad
	theta2 := theta1 - math.Pi/2 + (farDeg-FarRest)*deg2rad

	x = L1*math.Cos(theta1) + L2*math.Cos(theta2)
	y = L1*math.Sin(theta1) + L2*math.Sin(theta2)
	return
}

// returns servo angles (degress)  that place the foot at (x,y) relative to hip (error if out of reach)
func footIK(x, y float64) (closeDeg, farDeg float64, err error) {
	d2 := x*x + y*y
	d := math.Sqrt(d2)

	if d > L1+L2 || d < math.Abs(L1-L2) {
		return 0, 0, fmt.Errorf("knee error: foot target is unreachable")
	}

	cosRel := (d2 - L1*L1 - L2*L2) / (2 * L1 * L2)
	cosRel = math.Max(-1, math.Min(1, cosRel))
	thetaRel := math.Acos(cosRel)

	theta1 := math.Atan2(y, x) - math.Atan2(L2*math.Sin(thetaRel), L1+L2*math.Cos(thetaRel))
	theta2 := theta1 + thetaRel

	closeDeg = theta1/deg2rad + CloseRest
	farDeg = (theta2-theta1+math.Pi/2)/deg2rad + FarRest
	return
}

var restX, restY = footFK(CloseRest, FarRest)

type Leg struct {
	Close *actuators.Servo
	Far   *actuators.Servo
}

// moves the foot to x, y relative to hip axis.
func (l *Leg) SetFoot(x, y float64) error {
	closeDeg, farDeg, err := footIK(x, y)
	if err != nil {
		return err
	}

	closeDeg = clamp(closeDeg, CloseMin, CloseMax)
	farDeg = clamp(farDeg, FarMin, FarMax)

	if err := l.Close.SetAngle(closeDeg); err != nil {
		return err
	}
	return l.Far.SetAngle(farDeg)
}

func clamp(v, lo, hi float64) float64 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
