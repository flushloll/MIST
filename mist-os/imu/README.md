# IMU
Comprehensive package to work with the IMU module. It can both give one-time readings as well as stream data. It operates in SI units, but you can get arbitary raw sensor readings as well.

MIST's IMU module: MPU-6050

Usage guide:
```go
package main

import (
	"fmt"
	"log"
	"mist-os/imu"
	"time"
)

func main() {
	sensor, err := imu.New()
	if err != nil {
		log.Fatal("imu initialisation error: %w", err)
	}
	defer sensor.Close()

	for {
		data, err := sensor.Read()
		if err != nil {
			fmt.Print("imu read error: %w", err)
			continue
		}

		fmt.Printf("Pitch: %.1f | Roll: %1.f | Temp: %.1f",
			data.Pitch(), data.Roll(), data.Temperature)
		time.Sleep(10 * time.Millisecond)
	}
}
```