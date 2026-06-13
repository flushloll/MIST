# Body
The purpose of this module is to provide a higher-level control over the actuators in MIST's body rather than individual motor control of her body. The modes offered: **Ride**, which treats MIST as a fancy car, and **Stand**, which allows for precise control of each limb.
Whether this will be the interface provided to the robot ai, is still undecided.
The controls are robot-relative. 
The ride-drone movement style will be used. (below for more.)

Roadmap:
1. Head rotation - singular servo?
2. Motors only - car-styled movement
3. Servos - manual control
3. Servos - automatic height 
4. Servos - automatic tilt
3. Servos - automatic Ride
5. Hand - Forward Kinematics
6. Hand - Inverse Kinematics

What we have:
Currently, I have established the protocol for controlling each motor in MIST (a physical twin of the cute rice-cooker robot from pantheon). The main issue/uncertainty comes from how should the commands from ps5 controller be interpreted in what ways by the robot.
MIST has:
- head rotation with a servo 180 degrees (pan)
- 4 motors on each of the legs
- 8 servos (2 per leg) in the same axis, that act more as emotional points, and can tilt head I suppose?
- There is some sort of pan mechanism for legs of which I'm still unsure, so it should be decided upon on tomorrow.
- For now, assumption is, we can rotate the front legs by 30-40 degrees each.

Proposed type of controls:
- Ride-car: There's one axis on a controller responsible for the speed and rotation of controller 
- Ride-rocket: As if MIST is in space and at the end of her back legs, there are thrusters of which intensity we can control. The R2 and L2 control how much thrust is in right and left sides respectedly. The possible downside might be instability (wobbling right and left), we might need to add thretholds of difference between the two sides.
- Ride-tank: R3 and L3 control MIST is similar way, but now we can also do things like 360 rotation on spot by moving one joystick forward and second backward. Same downside as ride-rocket?
- Ride-drone: L3 controls the direction of the robot. (There will be some blind spots on the right and left sides but everything except that should be covered.) R3 controls the rotation in place (right/left) and height of the head. The question would be on how to handle input from both controls simultaneously. 