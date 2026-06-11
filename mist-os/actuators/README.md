# Actuators
This is a module responsible for all of the servos and brushless motors that make up MIST.

!TODO: Insert Servo (part of MIST) to Index Mapping Diagram

Usage guide:
```go
package main

import (
	"fmt"
	"log"
	"mist-os/actuators"
	"time"
)

func main() {
	// setup
	cntr, err := actuators.NewController()
	if err != nil {
		log.Fatal(err)
	}
	defer cntr.Close()

	// Servos:
	// move servo 0 to center position
	cntr.SetServo(0, 90)

	// smoothly move servo 1 from 0 to 180
	for angle := 0.0; angle <= 180; angle += 10 {
		cntr.SetServo(1, angle)
		time.Sleep(50 * time.Millisecond)
	}

	// Motors:
	// warming esc's
	for ch := 12; ch <= 15; ch++ {
		if err := cntr.ArmESC(ch); err != nil {
			log.Fatalf("failed to arm channel %d: %w", ch, err)
		}
	}
	fmt.Println("ESCs armed.")

	// spin motor on channel 12 up slowly then back down
	for throttle := 0.0; throttle <= 0.5; throttle += 0.1 {
		cntr.SetThrottle(12, throttle)
		time.Sleep(200 * time.Millisecond)
	}
	cntr.SetThrottle(12, 0)
}
```

## Connections:
PCA9685 to RP5:
- 1GND to 25Ground
- 3SCL to 5SCL(GPIO3)
- 4SDA to 3SDA(GPIO2)
- 5VCC to 17 3V3 power
PCA9685 to Actuators:
- 12 slots of 9,10,11 for TowerPro MG996R 180 degree servos.
- 4 slots of 9.10,11 for DYS D2822 Brushless Outrunner Motor 2-3S for Multicopter RC Aircraft Fixed-Wing UAV（D2822 1800KV）controlled by 5A/5A PRO/10A/20A/30A/40A/50A/60A/120A Dual Way Bidirectional Brushed ESC Control For Rc Model Boat/tank 180-795 Brushed Motor. 