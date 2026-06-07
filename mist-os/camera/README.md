# Camera
For testing purposes, there are both Mac's and actual MIST's camera implementations.
This package exposes a function `Shoot()`, which when called returns an `image.Image` from the camera module.

MIST's camera module: Raspberry Pi 5 Camera 5MP 1080P Optional Night Vision Wide Angle IR-CUT FF AF Cameras for RPI 5 Pi5 Zero | FF 69

Usage guide:
```go
package main

import (
 "mist-os/screen"
)

func main() {
 cam, _ := camera.NewCamera(0)
 defer cam.Close()
 imageBytes, _ := cam.Shoot()
 os.WriteFile("test_capture.png", imageBytes, 0644)
}
```