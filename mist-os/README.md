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

### Eyes
- IDLE(circle_angle); a circle outline with a cut part of it denoted by rotation parameter.
- Energetic(); like "> <".
- Soft(height, corner_radius1, corner_radius2); rectangle with one side msising. (height is how big the gap between two parallel lines are.)
- Fancy(char, mirrored[true|false]); any character from english alphabet.

### Mouths:
- None; in a case of only-eyes.
- Silent(count); line(s) or dot(s).
- Speech(height, corner_radius1, corner_radius2, corner_radius3, corner_radius4); filled in rectangle with rounded corners.
- Soft(corner_radius); an outline similar to AND or OR in set theory.
- Cutie(); uwu-styled w.

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