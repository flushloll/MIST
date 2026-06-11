package imu

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"

	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

const (
	gyroSensitivity  = 131.0
	accelSensitivity = 16384.0
	tempSensitivity  = 340.0
	tempOffset       = 36.53
)

type Vector3 struct {
	X, Y, Z float64
}

func (v Vector3) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

type Reading struct {
	Gyroscope     Vector3
	Accelerometer Vector3
	Temperature   float64
	Timestamp     time.Time
}

// pitch angle in degrees (might be noisy when robot is accelerating?)
func (r Reading) Pitch() float64 {
	return math.Atan2(
		r.Accelerometer.X,
		math.Sqrt(
			r.Accelerometer.Y*r.Accelerometer.Y+
				r.Accelerometer.Z*r.Accelerometer.Z,
		)) * 180 / math.Pi
}

// angle in degrees from accelerometer data
func (r Reading) Roll() float64 {
	return math.Atan2(
		r.Accelerometer.Y,
		math.Sqrt(
			r.Accelerometer.X*r.Accelerometer.X+
				r.Accelerometer.Z*r.Accelerometer.Z),
	) * 180 / math.Pi
}

// checks if robot parallel to the ground plane
func (r Reading) IsLevel(toleranceDeg float64) bool {
	return math.Abs(r.Pitch()) <= toleranceDeg && math.Abs(r.Roll()) <= toleranceDeg
}

type IMU struct {
	driver  *i2c.MPU6050Driver
	adaptor *raspi.Adaptor
	mu      sync.RWMutex
	latest  Reading
}

func New() (*IMU, error) {
	adaptor := raspi.NewAdaptor()
	if err := adaptor.Connect(); err != nil {
		return nil, fmt.Errorf("imu connecting adaptor error: %w", err)
	}

	driver := i2c.NewMPU6050Driver(adaptor)
	if err := driver.Start(); err != nil {
		return nil, fmt.Errorf("imu starting driver error: %w", err)
	}

	time.Sleep(50 * time.Millisecond) // to warm up after power on
	return &IMU{driver: driver, adaptor: adaptor}, nil
}

func (imu *IMU) Read() (Reading, error) {
	if err := imu.driver.GetData(); err != nil {
		return Reading{}, fmt.Errorf("imu reading error: %w", err)
	}

	r := Reading{
		Gyroscope: Vector3{
			X: float64(imu.driver.Gyroscope.X) / gyroSensitivity,
			Y: float64(imu.driver.Gyroscope.Y) / gyroSensitivity,
			Z: float64(imu.driver.Gyroscope.Z) / gyroSensitivity,
		},
		Accelerometer: Vector3{
			X: float64(imu.driver.Accelerometer.X) / accelSensitivity * 9.80665,
			Y: float64(imu.driver.Accelerometer.Y) / accelSensitivity * 9.80665,
			Z: float64(imu.driver.Accelerometer.Z) / accelSensitivity * 9.80665,
		},
		Temperature: float64(imu.driver.Temperature)/tempSensitivity + tempOffset,
		Timestamp:   time.Now(),
	}

	imu.mu.Lock()
	imu.latest = r
	imu.mu.Unlock()

	return r, nil
}

func (imu *IMU) Latest() Reading {
	imu.mu.RLock()
	defer imu.mu.Unlock()
	return imu.latest
}

// Unsure on this one yet
func (imu *IMU) Stream(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if _, err := imu.Read(); err != nil {
				fmt.Printf("imu: %v\n", err)
			}
		}
	}
}

func (imu *IMU) Close() {
	imu.driver.Halt()
	imu.adaptor.Finalize()
}

// Raw Readings
type RawVector3 struct {
	X, Y, Z int16
}

type RawReading struct {
	Gyroscope     RawVector3
	Accelerometer RawVector3
	Temperature   int16
	Timestamp     time.Time
}

func (imu *IMU) ReadRaw() (RawReading, error) {
	if err := imu.driver.GetData(); err != nil {
		return RawReading{}, fmt.Errorf("imu reading error: %w", err)
	}
	return RawReading{
		Gyroscope:     RawVector3(imu.driver.Gyroscope),
		Accelerometer: RawVector3(imu.driver.Accelerometer),
		Temperature:   imu.driver.Temperature,
		Timestamp:     time.Now(),
	}, nil
}
