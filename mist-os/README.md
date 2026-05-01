# MIST-OS
This is a program responsible for functioning of MIST-a fancy ricecooker robot inspired from Pantheon show. Esia is handling the hardware, leaving this part a small "mystery" for me. However, it is certain that we're running this project on Raspberry Pi 5 8gb edition with a DietPi on it. DietPi was chosen due to its stability and efficiency, especially that it doesn't have a GUI. The components I must take into account are as follows:
- Singleton Loop: main codebase which will allow us to modularise the robot's functions and pass around the arguments between modules (like that one loop in the game engines).
- Movement of actuators + smoothing curves: as the motors are wired correctly and identified in the software, I must create an interface for them to be remote-controlled. It would be a milestone to achieve a smooth lively animation similar to one portrayed in this [video](https://youtu.be/KPoeNZZ6H4s?si=CQ2NKyPPLrJgOUS9).
- Procedural animation for the eyes: I might have to create a double-interface for it to be able to be displayed on both mac and linux screens for testing and production respectively. The current screen code writes the generated photo into a frame-buffer, which isn't available on mac. 
- Emergent behaviour + AI: be it linkage of a smaller LLM such as Gemma or just a fancy if else statement, for it to function properly, we must figure out how to automate the robot. The current most prioritised features would be holding conversations with body language and following a selected person autonomously in a case of traveling. 

## Screen
There are different combinations of either Eyes + Mouths or just Eyes. In total, there are unique 7 eyes and 12 mouths found so far. We could minimise this by a lot if we give parameters such as rotation, scale, and width. Thus, we are able to simplify to 4 eyes and 5 mouth types. Important to remember that change in parameters changes the emotion, not change in type.

```bash
# Build for DietPi
GOOS=linux GOARCH=arm64 go build -o mist-screen .
./mist-screen
# If you need to change permissions before running it:
chmod +x mist-screen # make it executable by all users

# Run on MacOS:
go run .
```

Both eyes and the mouth are handled as three separate objects. Each part has a general parameter of position, scale, rotation, line_width which can be animated as well as internal parameters for each of the body type.

Now that nearly everything is prooved possible, I must figure out on how to simplify process of using the face. Perhpas some mix of pre-build animations and simpler face-changing method would be optimal?

We must figure out how to make the eyes follow the interest of the camera (what is the key feature we're starting into / paying attention to in the camera's object detection) and then look into the same direction. We could set a point towards which the eyes would point (or lack of the point) towards which the face will be animated to look.

Should we add blinking?

### Eyes
```go
// IdleEye: A circle outline with a gap.
LeftEye: &face.IdleEye{
    BaseFeature: face.BaseFeature{
        Position:  image.Pt(width/3, height/2),
        Scale:     1.0,
        Rotation:  0.0,
        LineWidth: 10,
        Color:     cyan,
    },
    Radius:  60,
    GapSize: 0.2, // 0.0 (full circle) to 1.0 (empty)
},

// EnergeticEye: A "> <" styled eye.
LeftEye: &face.EnergeticEye{
    BaseFeature: face.BaseFeature{
        Position:  image.Pt(width/3, height/2),
        Scale:     1.0,
        Rotation:  0.0,
        LineWidth: 10,
        Color:     cyan,
    },
    Size: 60,
},

// SoftEye: A rectangle with one side missing (U-shape).
LeftEye: &face.SoftEye{
    BaseFeature: face.BaseFeature{
        Position:  image.Pt(width/3, height/2),
        Scale:     1.0,
        Rotation:  0.0,
        LineWidth: 10,
        Color:     cyan,
    },
    Width:  60,
    Height: 40,
},

// FancyEye: Any character from the English alphabet.
LeftEye: &face.FancyEye{
    BaseFeature: face.BaseFeature{
        Position:  image.Pt(width/3, height/2),
        Scale:     1.0,
        Rotation:  0.0,
        LineWidth: 10,
        Color:     cyan,
    },
    Char:     "A",
    Mirrored: false,
},
```

### Mouths
```go
// NoneMouth: No mouth (useful for eyes-only expressions).
Mouth: &face.NoneMouth{},

// SilentMouth: One or more horizontal lines or dots.
Mouth: &face.SilentMouth{
    BaseFeature: face.BaseFeature{
        Position:  image.Pt(width/2, 3*height/4),
        Scale:     1.0,
        Rotation:  0.0,
        LineWidth: 8,
        Color:     cyan,
    },
    Count:   3,
    Width:   40,
    Height:  8,
    Spacing: 100,
},

// SpeechMouth: A filled rectangle.
Mouth: &face.SpeechMouth{
    BaseFeature: face.BaseFeature{
        Position:  image.Pt(width/2, 3*height/4),
        Scale:     1.0,
        Rotation:  0.0,
        LineWidth: 8,
        Color:     cyan,
    },
    Width:  100,
    Height: 40,
},

// SoftMouth: An arc outline (smile/frown).
Mouth: &face.SoftMouth{
    BaseFeature: face.BaseFeature{
        Position:  image.Pt(width/2, 3*height/4),
        Scale:     1.0,
        Rotation:  0.0,
        LineWidth: 8,
        Color:     cyan,
    },
    Width:  100,
    Height: 40,
},

// CutieMouth: A "w" styled mouth (uwu).
Mouth: &face.CutieMouth{
    BaseFeature: face.BaseFeature{
        Position:  image.Pt(width/2, 3*height/4),
        Scale:     1.0,
        Rotation:  0.0,
        LineWidth: 8,
        Color:     cyan,
    },
    Size: 40,
},
```

### Animation
All features (Eyes and Mouths) support smooth animation for both their base parameters and their specific properties. To animate, set a `TransitionRate` (0.0 to 1.0) and specify `Target` values.

```go
f := &face.Face{
    LeftEye: &face.IdleEye{
        BaseFeature: face.BaseFeature{
            Position:       image.Pt(width/3, height/2),
            Scale:          1.0,
            Rotation:       0.0,
            LineWidth:      10,
            Color:          cyan,
            TargetPosition: image.Pt(width/3, height/2 - 20), // Move up
            TargetScale:    1.2,                             // Scale up
            TransitionRate: 0.1,                             // Speed
        },
        Radius:       60,
        TargetRadius: 80, // Expand radius
    },
    // ...
}

// In your main loop:
f.Update(0.033) // ~30fps
f.Draw(img)
```

# Mainboard setup
For the purposes of saving money and resources, for now, we're using a raspberry pi 5, 8gb version for controlling fo the robot. LANE transcievers will output display connection, Gobot will work with actuators through PCIO interface, power supply will come from usb-c connected to power-bank, and bluethooth/wifi will be used for further connection to other devices, game controller, and maybe additional compute for now.

## Setting up OS
We've chosen DietPi due to its stability and in order to maximise available resources by ditching the GUI. It would be preferarble to flash some ssd instead of a flash card for reliability. In installation of the DietPi on the raspberry pi, we must choose to install these programs: tailscale, openssh, go.

We'll setup Tailscale through a token using these commands:
```bash
# Authentication
tailscale up

# Port Forwarding
echo 'net.ipv4.ip_forward = 1' | sudo tee -a /etc/sysctl.conf 
echo 'net.ipv6.conf.all.forwarding = 1' | sudo tee -a /etc/sysctl.conf
sudo sysctl -p
```

Now, we can access Raspberry pi through ssh by simply `ssh root@mist`. Notice that the face animation on screen won't be available through ssh, and will only work when a physical display is connected to the board.

# Controller
We've successfully prooved connected PS5 controller to both Mac and RaspberryPi, allowing us to rely on it as means of control in future for our robot. We've written proof of concept code to get the data from the controller, which includes every button press/release and movement of joysticks. Next stage for this would be to decide on how we want MIST to be controlled and implement as planned.

# Actuators
(Gobot)[https://gobot.io] seems like a good framework for controlling PID pins. As tehre are so many motors and sensors, we'll need to look into PWM Driver or something similar to wire motors and separate power supply for the motors and all.