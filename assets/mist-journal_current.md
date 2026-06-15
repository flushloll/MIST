# MIST — Journal Export

- Exported at: 2026-06-15T22:18:43Z
- Project ID: 2463
- Entries: 71

## Entry 1
- ID: 3694
- Author: Esia
- Created At: 2026-04-24T10:45:09Z

### Content

We've begun working on laying out the materials needed (motors, electrical components, physical framing) to make MIST (the rice cooker robot lol) a reality, as well as starting on the CAD of the robot's base whilst planning on eventually developing semi-humanoid legs to add onto the base and a functional lid mechanism to emulate an actual rice cooker.

![Screenshot 2026-04-24 034224.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6NzYxNCwicHVyIjoiYmxvYl9pZCJ9fQ==--29e0d2853e63831ec1cb308ab7814d829c3d9322/Screenshot 2026-04-24 034224.png)



### Recording Links

- https://lookout.hackclub.com/api/media/73c03f60-d6fb-4926-9cfd-aa727cfd0269/video.mp4
- https://lookout.hackclub.com/api/media/e32ee504-f344-4883-b765-bb6c7ec15c7f/video.mp4

## Entry 2
- ID: 3721
- Author: Umarbek
- Created At: 2026-04-24T15:03:42Z

### Content

In planning creation of MIST with Esia earlier today, we've planned a complicated system for the robot, with a lot of resource intensive elements such as face-recognition, object-detection and autonomous navigation and balancing. 

Here, I've tested the possibility of using DietPiOS on the raspberry pi as it doesn't have GUI, relying solely on terminal. However, we will have to animate eyes for the robot on a screen, and the choices we had were either to use ASCII art (we'll definitely add this if we have extra time) or render from zero using lower-level programming language (golang was chosen in this scenario as it will most likely be the singleton which will control the entire robot mechanics later on due to its speed and concurrency and all).

The result of this was simple animated (bouncing left and right) eyes for the screen. I couldn't figure out how to make it automatically detect the screen size due to complexities with low-level kernel architecture so I hard-coded whatever looked right to me for now 😅. Now that the idea works, I will next continue on defining the function loop for the robot and then continue with designing all the eyes states & implementing them in golang.

Final Eyes moving:
![Screenshot 2026-04-24 at 22.44.31.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6NzYzNywicHVyIjoiYmxvYl9pZCJ9fQ==--9076328d290de73c690e147a8f1609fb5f4ce131/Screenshot 2026-04-24 at 22.44.31.png)

Original Glorious Cube that worked, bringing hope in this idea:
![IMG_7268.JPG](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6NzYzNiwicHVyIjoiYmxvYl9pZCJ9fQ==--d1d94f6ac98dbc19599983afabd05c13e135d248/IMG_7268.JPG)

### Recording Links

- https://lookout.hackclub.com/api/media/026a0bd4-ef73-40e5-b696-359a6614e63e/video.mp4
- https://lookout.hackclub.com/api/media/16481bfa-b237-43ed-a15b-0fd930c1bbcb/video.mp4

## Entry 3
- ID: 3986
- Author: Umarbek
- Created At: 2026-04-26T03:25:57Z

### Content

Following yesterday's success of running animated eyes in DietPi, operating system without graphical operating system, today I've updated my setup to run on MacOS. Now it automatically compiles to the correct operating system depending on the host, creating a window for Mac and wiring to the frame buffer on Linux. Here's the new (now blue) ping-pong eyes running on Mac:
![screenshot2.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6ODIzNCwicHVyIjoiYmxvYl9pZCJ9fQ==--a3912b985488b44bb3337aa001c0a6c0dd3e9dc9/screenshot2.png)
(I know it might look intimidating but next thing I'll do is ensure correct colours and start implementing all the eye & mouth variations, making it more accurate to the source.)

I've also researched on all the possible faces for MIST, categorising all the various the ones I found into few categories by introducing parameters such as line-width, corner-radius, size, and rotation. 

![screenshot1.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6ODIzMywicHVyIjoiYmxvYl9pZCJ9fQ==--b219985c080da544c097c1a1a3d34d57679fa3fd/screenshot1.png)


### Recording Links

- https://lookout.hackclub.com/api/media/5870a76a-7e8f-4c3b-add2-98061efec9b3/video.mp4

## Entry 4
- ID: 4160
- Author: Umarbek
- Created At: 2026-04-27T02:33:07Z

### Content

I have continued the process of classifying the face types and figured out a way how to further break down each eye and mouth into basic shapes, allowing me to implement the least amount of facial features while achieving as much as possible. Next, I will start implementing them.

![screenshot.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6ODYyNywicHVyIjoiYmxvYl9pZCJ9fQ==--29c7c3a8b48573b44b856f901fa9cfbfcd8d98a6/screenshot.png)



### Recording Links

- https://lookout.hackclub.com/api/media/683e05c7-87ec-437c-8347-b03e6e2e362e/video.mp4

## Entry 5
- ID: 4177
- Author: Umarbek
- Created At: 2026-04-27T04:38:06Z

### Content

I have finally implemented some of the faces and figured out a way how to animate them. While it functions, each change is still hand-written and further steps for this would be to finish implementing all the possible facial features, create a framework to properly animate them in a simpler way (maybe using an http server or something on those lines) and creating pre-configured animations such as "default-state", "curious", "loading", "charging", "text", and more. Here are some of the faces we got so far:

![screenshot3.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6ODY2MCwicHVyIjoiYmxvYl9pZCJ9fQ==--917a9db76b62569613427030aa3dbf1b3e4041ed/screenshot3.png)


### Recording Links

- https://lookout.hackclub.com/api/media/a6880284-f1fe-44db-858e-74e1314de928/video.mp4

## Entry 6
- ID: 4294
- Author: Umarbek
- Created At: 2026-04-28T01:57:43Z

### Content

We've successfully prooved connected PS5 controller to both Mac and RaspberryPi, allowing us to rely on it as means of control in future for our robot. We've written proof of concept code to get the data from the controller, which includes every button press/release and movement of joysticks. Next stage for this would be to decide on how we want MIST to be controlled and implement as planned. (It functions on both MacOS and Raspberry Pi!)

![screenjot.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6ODkwNiwicHVyIjoiYmxvYl9pZCJ9fQ==--04c856477610e37232d4aec8d2d9add6c73a2a6a/screenjot.png)


### Recording Links

- https://lookout.hackclub.com/api/media/93c8a870-bbe5-4fd5-aa43-449a9fa7a26a/video.mp4
- https://lookout.hackclub.com/api/media/3302e526-94e0-4ada-bf5a-e6cae6ba62f2/video.mp4

## Entry 7
- ID: 4460
- Author: Umarbek
- Created At: 2026-04-29T07:39:54Z

### Content

In this session, I've analysed Disney's approaches in their Droid and attempted to write a pseudo-code algorithm for our robot. While it isn't finished, now there are some ideas on how movement and control should be handled and it is more clear on how to continue.
![shcreen.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6OTI1MywicHVyIjoiYmxvYl9pZCJ9fQ==--ef7a29d1d8999b00845ec765b5c93e46f34bdde4/shcreen.png)


### Recording Links

- https://lookout.hackclub.com/api/media/3e3b9fc0-bfd3-4016-b8ca-d296ae6b9f8a/video.mp4

## Entry 8
- ID: 4462
- Author: Umarbek
- Created At: 2026-04-29T08:44:30Z

### Content

With Esia, we've caught up on the progress. He've shared on the parts he've selected for the robot and we've reviewed the motor options, approximating battery life, power, and more.

![screenshot.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6OTI1OCwicHVyIjoiYmxvYl9pZCJ9fQ==--4dca8c05275e118111d082f3274e42dedb5c0d65/screenshot.png)


### Recording Links

- https://lookout.hackclub.com/api/media/aa841e17-7862-493c-bfeb-3292589eeeff/video.mp4

## Entry 9
- ID: 4624
- Author: Umarbek
- Created At: 2026-04-30T14:07:58Z

### Content

Today I've fixed some bugs on the facial software and finished implement all the eye/mouth types. I've also expanded upon what can be animated. The next step would be to clean up the codebase and figure out final way to let the higher-level brain to control the states of the face. (The demo is now animated)

![screenshot.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6OTY4NCwicHVyIjoiYmxvYl9pZCJ9fQ==--50a52c44fd3b6c1fe9484aa7cc35d07f41fc277e/screenshot.png)


### Recording Links

- https://lookout.hackclub.com/api/media/5c981180-3157-4423-be4e-4dcb309924af/video.mp4

## Entry 10
- ID: 4649
- Author: Esia
- Created At: 2026-04-30T17:57:49Z

### Content

Umarbek and I have now figured out the core of the electronics system together and I've continued CADDING whilst having a better vision of what the project will end up looking like.

A brief list of the electronics we'll be using:

12 servos and 4 DC motors 
A 3 cell LIPO battery.
A motor distribution board (PCA9685 - Adafruit)
Custom motor controller PCBs for the DC motors
Raspberry Pi5 (sourcing it ourselves) as the computer
A power bank for the Pi5
A PS5 controller (sourcing it ourselves probably)
12v to 6v DC converter for the servos
voltage sensor module to read battery life

![Screenshot 2026-04-30 10.55.50 AM.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6OTc1MiwicHVyIjoiYmxvYl9pZCJ9fQ==--6936e5181770040d2bde7d00e9a85d7806cfa74a/Screenshot 2026-04-30 10.55.50 AM.png)


### Recording Links

- https://lookout.hackclub.com/api/media/7813575c-b4e9-4d54-8d99-accfd3d6d907/video.mp4
- https://lookout.hackclub.com/api/media/8d835758-2ae4-44bc-bafb-96150c5e8d85/video.mp4

## Entry 11
- ID: 4730
- Author: Umarbek
- Created At: 2026-05-01T08:10:48Z

### Content

The hardest challenge for me is to create a behaviour that makes the robot feel real/alive. Therefore, I've written more on how should the robot's decisions/actions/movements be handled, coming up with what seems like an elegant "filtering" solution to its actions, and hierarchy of movements.
![screenshot.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6OTk5MSwicHVyIjoiYmxvYl9pZCJ9fQ==--bfd0d4dee2ea41d651d501bb7a2d848da99e1ab4/screenshot.png)


### Recording Links

- https://lookout.hackclub.com/api/media/f5a1cfda-4ccd-4e5a-9609-e60a23b87ef4/video.mp4

## Entry 12
- ID: 4829
- Author: Esia
- Created At: 2026-05-01T23:26:35Z

### Content

I've completed the CAD for one leg segment out of a total six that I will need to CAD. We've decided on using a belt drive beginning inside the robot's base and ending at the wheels, accelerating them.

We'll be using GT2 belts as they are the cheapest option available and using 0.375" hex shafts everywhere we decide to place an axle of rotation. 

Here's a front view of what the inside of one leg segment looks like:

![Screenshot 2026-05-01 053201.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MTAyNDcsInB1ciI6ImJsb2JfaWQifX0=--f37c7f15bce891cefd97fb791a004c299d5e01b0/Screenshot 2026-05-01 053201.png)


### Recording Links

- https://lookout.hackclub.com/api/media/8c873caa-98bf-4587-99f5-c955e238557a/video.mp4

## Entry 13
- ID: 4852
- Author: Umarbek
- Created At: 2026-05-02T03:21:42Z

### Content

I've explored upon the emotional models of human mind and wrote a little of pseudo-code to plan out how MIST will function. Additionally, together with Esia, we've thought of how autonomous-following functionality could work, and attempted running whisper locally for Speech to Text, and checked if Gemini would be a good "brain" for the MIST (it turned out to be surprisingly good, I now understand why Google wants to just put their models on robots)

![Screenshot 2026-05-02 at 12.18.46.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MTAzMDAsInB1ciI6ImJsb2JfaWQifX0=--b5b1c29ffeaee79c39572147f198d98dbea4393a/Screenshot 2026-05-02 at 12.18.46.png)


### Recording Links

- https://lookout.hackclub.com/api/media/6a2de2ed-2a25-4845-b4c1-bda2144b6f36/video.mp4

## Entry 14
- ID: 5227
- Author: Umarbek
- Created At: 2026-05-03T05:03:03Z

### Content

We've decided that we'll add a camera to MIST, which will inference OpenCV or some other object detection algorithm as well as facial recognition. I've successfully set up OpenCV face detection, ready to be integrated to the main singleton function. I must, however, still optimise the code and figure out a more elegant approach to object and facial recognition.

![Screenshot 2026-05-03 at 14.02.08.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MTA4MTksInB1ciI6ImJsb2JfaWQifX0=--883376f596beedca99b7729ef7fedf239f9babce/Screenshot 2026-05-03 at 14.02.08.png)


### Recording Links

- https://lookout.hackclub.com/api/media/c4f07891-65a2-4095-a853-8836b233251a/video.mp4

## Entry 15
- ID: 5410
- Author: Umarbek
- Created At: 2026-05-04T10:07:50Z

### Content

Further developed the idea of emotional vector and mapped out which emotions lead to what changes in faces. Finally have a functioning face with a fancy boot animation. Next I'll start working on coding for the servos and figuring out how to simulate cushioning.

![updated_mist_face.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MTEyODksInB1ciI6ImJsb2JfaWQifX0=--d047d1a98b63363ad242e982e8526a35699d324c/updated_mist_face.png)


### Recording Links

- https://lookout.hackclub.com/api/media/d23e75e9-98e4-41c1-bcaa-438a7c2405cd/video.mp4

## Entry 16
- ID: 5580
- Author: Umarbek
- Created At: 2026-05-05T10:45:16Z

### Content

I've run Yolo11n object detection model using python instead of go, which resulted in a higher accuracy and quantity of objects detected. However, it was still very low as we were only getting "person", "smartphone", "fridge" for any object I've tried. Object detection seems to be a sub-optimal method for eyes of our robot. 

Next, I've tested usage of Gemma models for the robot, and they blew all my expectations. After such an experience, you start understanding why Google just puts Gemini into all of their robots and why competitors follow–it's just disgustingly simple while performing better than any other method available. Trying various sizes ranging from Gemma3:270m to Gemma4:31b, the latter one provided sufficient intelligence to comfortably operate a prompt+picture simulated robot, bringing hopes in its autonomy. While this goes against original idea of full offline robot controlled with raspberry pi 5, it brings the end goal of intelligent robot to the line.

Next, I will look at how other robots or AI-apps create harness and how could our one look like for MIST to integrate Gemma while keeping it 'alive'.

![aqy5ff.jpg](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MTE2MzUsInB1ciI6ImJsb2JfaWQifX0=--8903abe0969e9a3cbd68be9848e3e2f7b03ce702/aqy5ff.jpg)


### Recording Links

- https://lookout.hackclub.com/api/media/32385c7d-ebe9-4f3e-9a0c-dcfdfa072252/video.mp4

## Entry 17
- ID: 5700
- Author: Umarbek
- Created At: 2026-05-06T07:36:59Z

### Content

We've encountered some issues with motors, which after some discussion were solved. I've quickly tested smaller models for the brain, made a brief map of the current code repository for presentation to the team this weekend.

![image.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MTE5NTIsInB1ciI6ImJsb2JfaWQifX0=--1b7fbdbfa9b6fc6d538ff1b22fa1386aa181cceb/image.png)


### Recording Links

- https://lookout.hackclub.com/api/media/1720a2b1-344a-4b0e-8d7f-aae7b555a3d2/video.mp4

## Entry 18
- ID: 5869
- Author: Umarbek
- Created At: 2026-05-07T09:52:28Z

### Content

I've tested few more smaller AI models, successfully connected to local instance of Ollama through golang and implemented particular AI functionality. 

Picked a camera for the robot and more...
![Screenshot 2026-05-07 at 18.52.16.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MTIzNTcsInB1ciI6ImJsb2JfaWQifX0=--b436766558871dab752e17c130947917fb4e4a2b/Screenshot 2026-05-07 at 18.52.16.png)


### Recording Links

- https://lookout.hackclub.com/api/media/acbfd201-6021-4357-a483-f68888b5722e/video.mp4

## Entry 19
- ID: 6044
- Author: Esia
- Created At: 2026-05-08T06:15:08Z

### Content

I spent a significant portion of time iterating through joint designs to manage actuating each leg's rotation separately. Originally, I went into pushing myself to try to create an impossibly compact design to remain accurate to the show all the while in our miniature formfactor.

Eventually I had to compromise and choose functionality over accuracy but after receiving second opinions from Umarbek, William, and outside sources, I determined this would be the best way to approach the problem and so I began working on a modular joint that we can use across the entire robot.

![Screenshot 2026-05-07 231436.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MTI3NDAsInB1ciI6ImJsb2JfaWQifX0=--aed257fd612f5df45be05acd19388e58bdd4bc01/Screenshot 2026-05-07 231436.png)

### Recording Links

- https://lookout.hackclub.com/api/media/7d2b6ae1-12c9-4942-9c66-09ea27bdfc14/video.mp4
- https://lookout.hackclub.com/api/media/e2a33ce6-6770-4ac2-82e1-5c67b1775964/video.mp4
- https://lookout.hackclub.com/api/media/b828f226-6e9d-4ea1-919d-cc6e571f1d0c/video.mp4

## Entry 20
- ID: 6098
- Author: Umarbek
- Created At: 2026-05-08T15:13:21Z

### Content

I have written the logic for the camera module, which works on Mac and (should) work on MIST as well. We'll test it on the robot during her assembly.

I understand why they do so. This techniques works too well to ignore it or invent drastically something new. An interesting observation from their [blog](https://deepmind.google/blog/gemini-robotics-er-1-6/) is that the core fundamental idea behind this model is its ability to identify objects, their center-points, and boxes. This information is then used in a standard way by LLM creating a script to calculate the distances/relationship between dots/boxes to then decide on what to do. 

![log.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MTI4NDUsInB1ciI6ImJsb2JfaWQifX0=--2232fe9d4aca0610a41eead0abfcd7a8ff9170dc/log.png)


### Recording Links

- https://lookout.hackclub.com/api/media/67de4cb7-f098-4924-8d4a-205e41f94336/video.mp4

## Entry 21
- ID: 6215
- Author: Esia
- Created At: 2026-05-09T06:33:20Z

### Content

I continued to work on revamping the joint into becoming modular and usable across the entire robot. I essentially finished it, and I've left to work on integrating it where needed in the robot.

The joint's design relies heavily on coaxial features to manage stability and mounting. There are a couple mounted bearings for stability purposes and one on either side to mount each leg directly onto.

![Screenshot 2026-05-07 230949.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MTMxNjQsInB1ciI6ImJsb2JfaWQifX0=--34ed4b79797c4611338354266d03e9c1ed298899/Screenshot 2026-05-07 230949.png)
![Screenshot 2026-05-08 223710.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MTMxNjUsInB1ciI6ImJsb2JfaWQifX0=--03827a36cfbef972383ae376ad9f81ecd54db338/Screenshot 2026-05-08 223710.png)


### Recording Links

- https://lookout.hackclub.com/api/media/70515fef-0fa0-42f9-a5ac-2fa92412ee17/video.mp4
- https://lookout.hackclub.com/api/media/89d895ee-1b3a-4481-94d4-5ff8573f50a8/video.mp4

## Entry 22
- ID: 6716
- Author: Esia
- Created At: 2026-05-12T05:01:50Z

### Content

I finished the first two leg segments of MIST entirely, and began to work on the third leg segment before I realized how impractical the design overall was and succumbed to dissatisfaction.

![Screenshot 2026-05-09 035626.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MTQ2MTEsInB1ciI6ImJsb2JfaWQifX0=--ce4d1cb6d9a70ae19440acd2311dae2a0b1cf6d6/Screenshot 2026-05-09 035626.png)

So, I thought to myself why am I making such a CHUNKY joint in the first place? Of which, I deduced two main reasons along with a bad design choice I had made previously.

These two reasons being:
1. I was designing too much for stability's sake, using dead axles and bearing combos wherever I could to support the robot's weight on all sides of the legs.
2. I was insistent on the fact that a motor could not possibly fit within the leg's volume and that a belt drive originating from MIST's body would be necessary.

For the first reason, I now say: it'll be fineeee...
In response to the second reason, I now say: small motors exist, don't be so stubborn and find one that'll work. (ehem: 

![newvsold.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MTQ2MzIsInB1ciI6ImJsb2JfaWQifX0=--2e3f0b005688cc8fb2ffc527b3cf018f57d8e5b5/newvsold.png)

The bad design decision I made being: I placed the servo in a vertical position, independent of any leg segment, rather than storing it within the former leg segment of each rotational piece. Such a setup took up too much space and led to intersecting conflicts.

So, with a new resolve and an even BETTER (and simpler) design in mind, I now embarked on the journey toward redesigning the joints for a third time.

### Recording Links

- https://lookout.hackclub.com/api/media/90f5964f-3105-4445-8038-a84a89962956/video.mp4
- https://lookout.hackclub.com/api/media/2b33c7e6-fea8-4aab-8c45-d09612915324/video.mp4
- https://lookout.hackclub.com/api/media/5247c1c3-9c12-434f-8bd1-3f1582cd37f0/video.mp4
- https://lookout.hackclub.com/api/media/019000b2-f068-4050-8a70-c0c3383134dc/video.mp4

## Entry 23
- ID: 6772
- Author: Umarbek
- Created At: 2026-05-12T14:03:02Z

### Content

Since I've implemented all the eye and mouth types at the beginning, now with all the screen transformations, some of them having changed significantly to the point I must rewrite the logic behind their rendering, thus here it is more time to drawing MIST's face :sob:

![MIST-excited.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MTQ4MTksInB1ciI6ImJsb2JfaWQifX0=--6ba0445d3d9bc8ee7ce9ca8494e4f1e50d1c23b2/MIST-excited.png)


### Recording Links

- https://lookout.hackclub.com/api/media/71add299-3c4f-4b92-8f96-fe99e061460d/video.mp4

## Entry 24
- ID: 6773
- Author: Umarbek
- Created At: 2026-05-12T14:03:38Z

### Content

Mujoco's interface while intimidating at first, is growing on me. I've been learning their XML formatting and simulation logic; while we wanted to import the CAD models from CREO to Mujoco right away, I've learned that it is too much, and we should rather aim to recreate a simplified copy of MIST, where only moving parts matter, and focus is rather on weight distribution, moment of inertia and other physical parameters rather than assembly. Thus, here I am, manually typing into XML to see if it works at all

![Mujoco.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MTQ4MjAsInB1ciI6ImJsb2JfaWQifX0=--157ff6d66e8071815a80354765ff5e5ab31b5fbd/Mujoco.png)


### Recording Links

- https://lookout.hackclub.com/api/media/e31cea6b-1676-4520-a262-2d67383b6250/video.mp4

## Entry 25
- ID: 6781
- Author: Umarbek
- Created At: 2026-05-12T15:16:22Z

### Content

Nothing too revolutionary, just iterated over all the eyes and fixed them one by one–finally made the non-pixelated versions of fancy-eyes (character eyes); below you can see the one with X's. Next, I will work on mouths. 

I also developed way cleaner approach of animating the eyes, greatly reducing the complexity in the main controller. Now it only requires this: 

```
package main

import (
	"mist-os/screen"
)

func main() {
	sc := screen.NewScreen(800, 480)
	if sc == nil {
		return
	}
	defer sc.Close()
	testFaces(sc)

	sc.StartLoading(0, 0.5)
	sc.SetFace("fancy-x", "none")
	sc.Run()
}
```

![dead-mist.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MTQ4MzksInB1ciI6ImJsb2JfaWQifX0=--f7c098cb71323e098ae29ec20b6d47c0049ee142/dead-mist.png)


### Recording Links

- https://lookout.hackclub.com/api/media/1637479e-8703-4b37-9683-d3b7a0701c18/video.mp4

## Entry 26
- ID: 6866
- Author: Umarbek
- Created At: 2026-05-13T02:35:29Z

### Content

Addition to my previous journal: I was able to animate in Mujoco, now I am trying to figure out on how to simulate it. Here, you can see an animation of a cube rotating:

![Screenshot 2026-05-13 at 11.32.54.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MTUwMTksInB1ciI6ImJsb2JfaWQifX0=--43a777e8551973f1aa6c5b5159989ec630aca2df/Screenshot 2026-05-13 at 11.32.54.png)


### Recording Links

- https://lookout.hackclub.com/api/media/e727a20c-720b-4e75-855e-f45028ae6761/video.mp4

## Entry 27
- ID: 7247
- Author: Umarbek
- Created At: 2026-05-15T12:44:39Z

### Content

I've been looking through the current schema for electronics, understanding how to connect to each component and checking for alternatives.

![screne.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MTU4MzMsInB1ciI6ImJsb2JfaWQifX0=--d92b2b8263c22ea10327779912ecdf734f75e9cf/screne.png)


### Recording Links

- https://lookout.hackclub.com/api/media/e812e053-596b-400d-94f4-4ea94474e825/video.mp4

## Entry 28
- ID: 7253
- Author: Umarbek
- Created At: 2026-05-15T13:51:28Z

### Content

I've continued the auditing of the components, and weighted the choice between bluetooth controllers such as ones of consoles and FPV-styled radio-controllers.

Then, I've written pseudo-code for self-positioning and ground angle approximation for a robot whose base is variably tilted. 

There are many many degrees of freedom. I've mapped some possible movements to a controller for now. (still unfinished)

![screen.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MTU4NDEsInB1ciI6ImJsb2JfaWQifX0=--2d15f627501fbc7fbe9ac119127afcee1fb556cd/screen.png)


### Recording Links

- https://lookout.hackclub.com/api/media/8fb666bd-ad79-4d44-b459-3b3fb31a7595/video.mp4

## Entry 29
- ID: 7379
- Author: Umarbek
- Created At: 2026-05-16T07:41:39Z

### Content

Expanded upon yesterday's work on figuring out robot's movement controls. I've mapped our controller's inputs to SDL, and attempted assigning them the correct motors for our robot. It is a challenge, as there are so many degrees of freedom in MIST, but doing this, it is now way clearer on how the code for motors should look like.

![sorew.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MTYxMjMsInB1ciI6ImJsb2JfaWQifX0=--85fd95a1532bcfce2aea4bcd24b97d21b2cb1607/sorew.png)


### Recording Links

- https://lookout.hackclub.com/api/media/739c8c01-57ae-49ed-a001-2444ed00942a/video.mp4

## Entry 30
- ID: 7740
- Author: Umarbek
- Created At: 2026-05-18T04:42:34Z

### Content

I've been working on Mujoco simulations, experimenting with the XML files and trying to find a way to properly see what I'm typing without having to run a python simulation and export it as a video every time...
![Screenshot 2026-05-18 at 13.41.41.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MTY5MTEsInB1ciI6ImJsb2JfaWQifX0=--0cdea57ae6fd29a482599c16d81d034e6ca8edf1/Screenshot 2026-05-18 at 13.41.41.png)


### Recording Links

- https://lookout.hackclub.com/api/media/bedeef45-fcb5-4023-8220-34a03268dba7/video.mp4
- https://lookout.hackclub.com/api/media/97053490-86f1-434e-9b40-7756d39ffbb4/video.mp4

## Entry 31
- ID: 7741
- Author: Umarbek
- Created At: 2026-05-18T04:46:18Z

### Content

I've learned of existence of visual studio code Mujoco viewer extension, which turned out to be a game-changer! It allowed to simulate the model in real time and gave an already set up beautiful environment with good lightning. Next, I'll be figuring out how to control the digital robot with external feedback during simulation. 

Btw, this model is whatever I could scramble by myself and isn't the model being used for CAD or the model we'll actually simulate later, it is purely for testing purposes.

![Screenshot 2026-05-16 at 22.53.01.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MTY5MTIsInB1ciI6ImJsb2JfaWQifX0=--2e31ff4d336e26ceb4227f81f34baacf97c80c98/Screenshot 2026-05-16 at 22.53.01.png)


### Recording Links

- https://lookout.hackclub.com/api/media/4c3ea4c0-9171-4803-a020-533e36372e9a/video.mp4

## Entry 32
- ID: 7898
- Author: Umarbek
- Created At: 2026-05-19T03:36:02Z

### Content

In this session, together with Esia, we've been revising and evaluating the choices of motors and their positioning. Several different approaches were discussed and we've come at more elegant design which will (hopefully) help simplify building of front legs, bringing it closer to original design.

![fjaf.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MTcyNjgsInB1ciI6ImJsb2JfaWQifX0=--763e905301b4860cd5b0ea6c187f9741f6baa4c8/fjaf.png)


### Recording Links

- https://lookout.hackclub.com/api/media/1b6a3850-e437-45c6-9e6d-fcdefe390879/video.mp4

## Entry 33
- ID: 9258
- Author: Umarbek
- Created At: 2026-05-26T02:59:59Z

### Content

Before writing Gobot code, I've realised I must know which pints must I connect to in order get data and control the actuators. This marked the beginning of series of great circuitry sessions which consists of a continuous loop of me searching for non-existent documentation of electrical components, getting humbled by power distribution calculations in order to avoid frying the precious raspberry pi, and (hopefully) correctly wiring the components.

![Screenshot 2026-05-26 at 11.56.27.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjA1NTcsInB1ciI6ImJsb2JfaWQifX0=--7b678d0dc08d86f0212a741f4d73a6ad0e059bd5/Screenshot 2026-05-26 at 11.56.27.png)


### Recording Links

- https://lookout.hackclub.com/api/media/70f7bbd9-917b-4036-aa4c-962078268ec3/video.mp4

## Entry 34
- ID: 9260
- Author: Umarbek
- Created At: 2026-05-26T03:03:34Z

### Content

Continuation of the previous session, where I've been working on circuitry for MIST. This time, I've only concentrated on wiring, putting Gobot and programming aside (I'll come back to it once the wiring will be completed.) All that's left is the hardest parts of this process—wiring 8 servos and 4 motors, figuring out how to distribute power across all of these things without having to get dozen separate batteries.

![Screenshot 2026-05-26 at 11.56.39.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjA1NjMsInB1ciI6ImJsb2JfaWQifX0=--3333c0e0f1797e7bd8834e2cefcd7edd41b8db43/Screenshot 2026-05-26 at 11.56.39.png)


### Recording Links

- https://lookout.hackclub.com/api/media/d1a65c28-dd19-4d14-9f14-397c3b01358b/video.mp4

## Entry 35
- ID: 9262
- Author: Umarbek
- Created At: 2026-05-26T03:05:46Z

### Content

Continuing from the previous session, I've finally finished connecting the servos! All that's left is power, motors, and peripherals (they'll just connect through USB's so it won't be as much of hurdle as actual circuitry components.)

![hga.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjA1NjQsInB1ciI6ImJsb2JfaWQifX0=--8d663c6e6419ed9e01237a8bb48db17590203dc0/hga.png)


### Recording Links

- https://lookout.hackclub.com/api/media/b056a568-a73e-42de-a9a5-8292f88df3ba/video.mp4

## Entry 36
- ID: 9298
- Author: Umarbek
- Created At: 2026-05-26T05:55:11Z

### Content

Understood how wiring and controlling of brushless motors work, yet now there are questions whether it will be strong enough for our heavy robot. Will continue later on it.

![operw.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjA2MTUsInB1ciI6ImJsb2JfaWQifX0=--c599ed5a28da7163840b0f0e59fd0ca13bc62f6c/operw.png)


### Recording Links

- https://lookout.hackclub.com/api/media/f6620127-5ee7-4e61-97a9-5f0446cfd1da/video.mp4

## Entry 37
- ID: 10080
- Author: Umarbek
- Created At: 2026-05-29T06:49:53Z

### Content

Until now, I've been creating separate modules for each part of the robot: camera, AI, screen, imu, and more. After reading an article on how all these parts are combined in robotics, I've created a basic diagram for the data flow between different packages. I've been trying to figure out how to actually move the robot (which turns out is a part of Control theory), that I've been studying after this session at my own time.

![Screenshot 2026-05-29 at 15.47.00.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjI3NTIsInB1ciI6ImJsb2JfaWQifX0=--2344ef0ff3871bd1966a7172280b2d4feb860655/Screenshot 2026-05-29 at 15.47.00.png)


### Recording Links

- https://lookout.hackclub.com/api/media/234dcb60-f5c7-443d-b1ae-30f7e6c7564f/video.mp4

## Entry 38
- ID: 10092
- Author: Esia
- Created At: 2026-05-29T09:23:55Z

### Content

I began working on the 3rd version of the joint around this time. In this new version of the joint, there's a few differences:

1. The servo is mounted directly to the joint itself and the leg that precedes it. This is just a better way of coupling rotational motion rather than what I was doing before where the servo was connected only to itself.

![image.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjI3NzIsInB1ciI6ImJsb2JfaWQifX0=--60c0ad873a8b7adeafcfc36ba0a88e7a9227e57d/image.png)
![image.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjI3NzgsInB1ciI6ImJsb2JfaWQifX0=--105bf90312a3c3c3c826a099e2d82d1b4a0287f6/image.png)

2. The motor is a smaller alternative that can be mounted directly in the leg. This removes the need for a chain of belts connecting the chunky mini cim motor we were previously using:

![IMG_8289.jpg](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjI3ODEsInB1ciI6ImJsb2JfaWQifX0=--2edfbcf3c4ecf25c2df1fc3ff35218ca81e999b0/IMG_8289.jpg)

Thus making our joints more compact overall.

3. There is simply less things to mount overall, so axial stability is less of a concern. This lets us use just one bearing for support rather than the bearing stack that was being used before.

![image.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjI3ODQsInB1ciI6ImJsb2JfaWQifX0=--001d879be97afc649d38d4a2df0e0134858be6f8/image.png)


### Recording Links

- https://lookout.hackclub.com/api/media/2cde9235-83ac-4cb1-a8d2-085cfc29ac14/video.mp4
- https://lookout.hackclub.com/api/media/9706a848-2cf6-469c-b5f4-1f129247040a/video.mp4

## Entry 39
- ID: 10098
- Author: Esia
- Created At: 2026-05-29T09:48:47Z

### Content

With the 3rd version of the joint done, I returned back to finishing the design of the legs. This part was rather simple, I just finished the shells of each leg segment that remained for the robot and attached a joint to each respectively. This allowed us to go from this:

![image.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjI3ODgsInB1ciI6ImJsb2JfaWQifX0=--42dd81203d2d6579b7e909070da1c82d17da1972/image.png)

To this:

![image.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjI3ODksInB1ciI6ImJsb2JfaWQifX0=--874ee872d54c557b903c7205d65a3c04884a04a9/image.png)

And by shell, I mean literally that. These leg segments are just hollow and empty, only existing so that our project remains loyal to the show we are deriving it from.

![image.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjI3OTAsInB1ciI6ImJsb2JfaWQifX0=--c2c93aaca0c911cc1e58f8c9acb8747ac3598a0b/image.png)
![image.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjI3OTIsInB1ciI6ImJsb2JfaWQifX0=--b5dddbaaf092a51d314f4ac4aa48f968ef732ea2/image.png)

With an entire leg done, I mounted it onto the robot, did some mirroring magic to receive the model of the second version of this leg we'll be using and voila: 

![Screenshot 2026-05-29 023823.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjI3OTYsInB1ciI6ImJsb2JfaWQifX0=--c70a709c88ac16ce01841ea64bdca531e6d2573d/Screenshot 2026-05-29 023823.png)
![Screenshot 2026-05-29 023823.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjI3OTcsInB1ciI6ImJsb2JfaWQifX0=--320359e68bacf9ee1dcee9ca7a811b7a7ecc6c35/Screenshot 2026-05-29 023823.png)
![Screenshot 2026-05-29 023752.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjI3OTgsInB1ciI6ImJsb2JfaWQifX0=--de39e453bc34a0c4f8f41e2d62adee3483c0bf92/Screenshot 2026-05-29 023752.png)


### Recording Links

- https://lookout.hackclub.com/api/media/4dee5033-5493-44d3-8d79-0c1e83e73d37/video.mp4
- https://lookout.hackclub.com/api/media/99fadd5c-3786-4258-b535-f6de5c78851d/video.mp4

## Entry 40
- ID: 10099
- Author: Esia
- Created At: 2026-05-29T10:04:48Z

### Content

I basically took some time to further develop the wiring of our robot and make sure everything checks out:

![image.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjI4MDMsInB1ciI6ImJsb2JfaWQifX0=--1bc29189f3b02364ed3ce4ef8056d15756b4001c/image.png)

It was rather informal and I just did it on the figma board with marker and no pin layouts, but I just established a baseline for Umarbek to eventually build upon with specific pin connections and more detail.

Afterwards Umarbek left some comments which you can see on the sticky notes.

At the same time, I also very minimally began on the cad for the other 2 legs remaining on the robot (being the skinny legs).

### Recording Links

- https://lookout.hackclub.com/api/media/e37accac-9d07-4eaf-8377-e5b673488e2b/video.mp4

## Entry 41
- ID: 10108
- Author: Esia
- Created At: 2026-05-29T10:56:27Z

### Content

So I 3D printed some of the stuff, including the cover for the servo horns which our robot's joints use.. take a look at it:

![AF432733-6F86-4B02-A83D-2554A9AC9522.jpg](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjI4MDcsInB1ciI6ImJsb2JfaWQifX0=--96fdb40e5980c1a8ceffc670eb3524c99679b61b/AF432733-6F86-4B02-A83D-2554A9AC9522.jpg)

Issue. The holes are incredibly small. Now, I have never seen screws that small other than in the assembly of extremely delicate electronics. In our case the weight of a robot leg is resting upon those screws.

Screws in general tend to be stronger than needed for all purposes I've been exposed to, so to follow that trend I chose to also upsize our screws as well, mostly for the peace of mind. In order to do so, I'd have needed to design a custom servo horn and ensure it can actually fit onto the tooth profile of a servo.

So I did some research on previous attempts:

![Screenshot 2026-05-29 033022.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjI4MTUsInB1ciI6ImJsb2JfaWQifX0=--99273a7d0f4d2ac67cbab8944f054121b9ec6f4e/Screenshot 2026-05-29 033022.png)
![Screenshot 2026-05-29 033100.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjI4MTYsInB1ciI6ImJsb2JfaWQifX0=--ef8aafa9fbc7586f7be398077fd03bf1ef5ed55a/Screenshot 2026-05-29 033100.png)

Printed a few variations with different tooth profiles for mounting:

![image.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjI4MTksInB1ciI6ImJsb2JfaWQifX0=--cc00608ef18436624bc094f3b15dff056216b141/image.png)

And had some moderate success with fit at the very least.

![IMG_8392.jpg](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjI4MjEsInB1ciI6ImJsb2JfaWQifX0=--f6f46a79802a7ab14af62ff4bde137d95ee9d356/IMG_8392.jpg)

What I found is that generally 3D printers tend to struggle with printing details finer than the nozzle size (in my case 0.4 mm):

![image.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjI4MjgsInB1ciI6ImJsb2JfaWQifX0=--1ee0f04c88c0e6f11cffdad985a90cfb5b023a6e/image.png)

Even the slicer sometimes turns the details into goop:

![image.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjI4MzAsInB1ciI6ImJsb2JfaWQifX0=--f8abd162d7ed25eab754c0d16c2fe2736278e3e2/image.png)

So you have to work around that. I did so by just undersizing the tooth profile (scaling it down by 2%) and relying on a press fit.

I also started working on doing some math on what drive ratios I'd need to accelerate with maximum efficiency. It involved a significant chunk of research and honestly most of it was just physics. Here are some of the constants I defined before getting into the math:

![image.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjI4MzMsInB1ciI6ImJsb2JfaWQifX0=--d399301ddedd0fea76ca99695c00793348e8c54f/image.png)


### Recording Links

- https://lookout.hackclub.com/api/media/20d9913f-4804-4faa-8477-c5a1e76b8ce0/video.mp4

## Entry 42
- ID: 10116
- Author: Esia
- Created At: 2026-05-29T11:28:12Z

### Content

To decide the drive gear ratio we'd be using for our wheels, I had to do the math with the constants I had given myself.

![image.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjI4MzksInB1ciI6ImJsb2JfaWQifX0=--d8f80097ad852fdaa4dab2b44f49d6a0beb3f468/image.png)

I basically calculated the maximum force the ground could theoretically exert against our wheels before they started slipping and decided that it would only make sense to go up to that fast and geared such that we could go at that speed within our motors' efficient range of operation (~80%)

The reason we decided it was necessary to be capable of going so fast (18.14 m/s) is because of this clip from Pantheon: [https://youtu.be/1tU2ZMSHPR4?t=954](url)

We want to stay accurate to the show, and sometimes it posed challenges like having to push your robot to go _theoretically_ as fast as a car on a road.

And with the gear/pulley ratio now calculated, I modified the model to match it appropriately:

![image.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjI4NDIsInB1ciI6ImJsb2JfaWQifX0=--3d4310083ba4586a8e8b4e3afd216086d03cfe42/image.png)

### Recording Links

- https://lookout.hackclub.com/api/media/318c4d66-f633-489e-8f08-5504fa8e8c16/video.mp4

## Entry 43
- ID: 10343
- Author: Umarbek
- Created At: 2026-05-30T08:05:08Z

### Content

With the Burnout pitch deadline approaching, I've started documenting our total progress on MIST in a singular document. In writing it all into document, I've realised how far we've come. While I've written most of the pitch, it is still to be completed.

![Screenshot 2026-05-30 at 16.55.33.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjM0MzgsInB1ciI6ImJsb2JfaWQifX0=--0719c17461908590b33f2cc2db703fda9f4b2d4d/Screenshot 2026-05-30 at 16.55.33.png)


### Recording Links

- https://lookout.hackclub.com/api/media/d79effb6-2c98-4544-8783-e790e19756d2/video.mp4

## Entry 44
- ID: 10344
- Author: Umarbek
- Created At: 2026-05-30T08:07:11Z

### Content

Continuation from previous time, I've read through other people's failed and successful burnout pitches, styling ours in a similar manner to ensure we'll pass it. I've completed all I could, leaving some space for my teammate to fill out (parts specific to his knowledge and skill that I don't know how to answer.) We're planning to submit it by Saturday evening.

![Screenshot 2026-05-30 at 17.06.09.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjM0MzksInB1ciI6ImJsb2JfaWQifX0=--5f8e03a8573d1248f56f11dbd496db90525e0c65/Screenshot 2026-05-30 at 17.06.09.png)


### Recording Links

- https://lookout.hackclub.com/api/media/afc3c53c-6b00-4251-b3e9-0f3310a0f7b5/video.mp4

## Entry 45
- ID: 10346
- Author: Umarbek
- Created At: 2026-05-30T08:13:17Z

### Content

I have started the Bill of Materials and completed it to the extent to which I could. There are many smaller materials/parts that still have to be accounted for (I've only put their names), but I am hoping that they could be better approximated by my teammate. I've reviewed the electrical components and their delivery dates, cheaper listings on other marketplaces, and updated the list to reflect the latest decisions on MIST. So far, it sums up to $371.94 USD, but it is expected to be higher. (I have also came to a realisation that the power supply for Raspberry Pi 5 we've been using is too weak, as RP5 required non-standard current. For now, current converter seemed to be the most elegant choice.)



![Screenshot 2026-05-30 at 17.11.47.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjM0NDIsInB1ciI6ImJsb2JfaWQifX0=--dbe4c907828cf2a35575d065cd0f1f30b065b1b2/Screenshot 2026-05-30 at 17.11.47.png)


### Recording Links

- https://lookout.hackclub.com/api/media/a5833e07-74eb-45b6-acf5-84540a2b7a6e/video.mp4

## Entry 46
- ID: 10843
- Author: William
- Created At: 2026-06-01T06:38:27Z

### Content

### Mist Robot Arm: Development Log



I made the robot arm for Mist. Here is a breakdown of the design iterations:



**Version 1**

The v1 was way too bulky and ended up looking pretty bad. It used an 11kg servo motor at each joint. While powerful, this ended up making the arm almost as wide as the leg, which looked terrible.



**Version 2**

For the v2 (which I never fully finished), I decided to try to move all of the joints out of the arm itself by controlling it with a rope drive system. This worked great in theory, but it just concentrated the size issue further down into the arm, as it would interfere with the internal electronics if I tried to conceal it in the body.



**Current Iteration**

I went back to a servo at each joint, but this time I used micro servos. The reason I avoided this initially is that micro servos didn't seem able to hit the torque requirements without being absurdly expensive. After much searching, I managed to find a half-decent price, although they are still much more expensive and not as strong as the full-sized servos. 



**Next Steps**

* **Remaining Work:** The only things that remain are aesthetic and weight optimization, as well as sorting out which gripper to use. 

* **Gripper:** I currently have it set up as a modular system, so anything could be put there.
![image.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MjQ3NDEsInB1ciI6ImJsb2JfaWQifX0=--72ff3892e25eaf5011cf3ea6a9daa5bc8e7c3df3/image.png)


### Recording Links

- https://lookout.hackclub.com/api/media/b6f9eda6-308f-4d0e-a838-ee51ae80b7e2/video.mp4
- https://lookout.hackclub.com/api/media/ff0bda29-f026-4195-839b-19ed1f5cf061/video.mp4
- https://lookout.hackclub.com/api/media/d802a121-baac-4944-a09e-d6cbeafa0fc3/video.mp4
- https://lookout.hackclub.com/api/media/3b65ee98-45fb-4659-80c2-427ac2cd1a8b/video.mp4

## Entry 47
- ID: 12177
- Author: Umarbek
- Created At: 2026-06-07T02:21:40Z

### Content

It's been a while since we've last updated guide for setting up software on our robot. Therefore, I've written all of the instructions to get software part of the robot ready and running. We will test it out in the upcoming meeting with my teammate during build process. 

![Screenshot 2026-06-07 at 11.19.46.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6Mjg1OTgsInB1ciI6ImJsb2JfaWQifX0=--20d71b1887c99bad7d7dee84be3f77a52f2ab64a/Screenshot 2026-06-07 at 11.19.46.png)


### Recording Links

- https://lookout.hackclub.com/api/media/56e95f4c-89ac-445a-9889-17e6c9ad37d7/video.mp4

## Entry 48
- ID: 12201
- Author: Umarbek
- Created At: 2026-06-07T04:11:02Z

### Content

I have written and cleaned the code for controller and camera modules for MIST. Additionally, I've written some documentation for both with an example usage. Now, I can easily initialise and use controllers and cameras through simple go interface.

Below, there are screenshots of README's that capture simplicity and abstraction of using these packages now.

Controller:
![controller.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6Mjg2NDAsInB1ciI6ImJsb2JfaWQifX0=--b2532f1c182d596b10a5da772a9845825329aa4f/controller.png)

Camera:
![camera.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6Mjg2NDEsInB1ciI6ImJsb2JfaWQifX0=--08ad1b9e5449b55ac8fa28fd4f7e0a9625875879/camera.png)

I will now move on to next modules. The goal for today is to have all of mist-os ready for review.

### Recording Links

- https://lookout.hackclub.com/api/media/0201fbac-ea3e-451e-a138-d805214f6223/video.mp4

## Entry 49
- ID: 12236
- Author: Umarbek
- Created At: 2026-06-07T09:19:23Z

### Content

While the screen module had a clean interface already, it for certain wasn't clean nor modular. I've modularised code for each eye and mouth types into separate folders and optimised speed of rendering for many features. Additionally, I've removed a lot of dead code. Here's MIST's happy face from changes I've made:

![face.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6Mjg3MTgsInB1ciI6ImJsb2JfaWQifX0=--6455975afb369ee81d1ecd987f153f7a3e19654d/face.png)


### Recording Links

- https://lookout.hackclub.com/api/media/e2b8680c-09e3-4473-96d8-31710879b633/video.mp4
- https://lookout.hackclub.com/api/media/c37768de-433a-4a1c-a737-5babb4a0f069/video.mp4

## Entry 50
- ID: 12265
- Author: Umarbek
- Created At: 2026-06-07T12:35:24Z

### Content

I have created a new version of text to speech module, which utilises the OS's built-in program in order to create and play an audio.

With a simple bash script, I have tested all of the sounds available on Mac on a short phrase. It worked out perfectly, and 

Testing on Raspberry Pi's Dietpi was of more trouble, but by creating audio files and exporting them into a folder which my own alternative of google drive shares to the internet, I was able to get the voice lines onto my device. Unfortunately, the voices in linux configuration were too basic. I have chosen one of the voices for now, but now there are plans on using neural-network based models or simple vocaloid of some sort.

The module is now ready and working, the package is usable and production ready.



![voice.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6Mjg3NzMsInB1ciI6ImJsb2JfaWQifX0=--c4b6d3ce50f53ab04a503cf580b5c4ff4e1f15df/voice.png)


### Recording Links

- https://lookout.hackclub.com/api/media/f5b67658-44a1-4395-bb0d-810dccae58e4/video.mp4
- https://lookout.hackclub.com/api/media/0068e534-4daa-42ce-b2f7-d08ad6805467/video.mp4
- https://lookout.hackclub.com/api/media/ae29fc2e-f836-425e-8ad5-391f9d409690/video.mp4

## Entry 51
- ID: 12283
- Author: Umarbek
- Created At: 2026-06-07T13:53:06Z

### Content

I've tested running LLMs on edge using Ollama. The hoped-for Gemma4:2eb turned out to be too large, nearly bricking my raspberry pi 5:8gb by filling out its ram and heating the CPU to the level of boiling water. As last reading before disconnect showed critical temperature, I dived into the box of cables to find my raspberry pi, and reboot it. Whilst waiting for the red light to turn green, I ran to the kitchen for ice, which I've used to cool the cpu down. (just realised that it's quite similar to pantheon.)

![gemma.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6Mjg4MDQsInB1ciI6ImJsb2JfaWQifX0=--d551760bad5e5931b015d118280a87333b43ce4e/gemma.png)

Gemma4:2eb was a total disaster, as I've started searching for other models. Qwen3.5:0.8b looked promising as it was the smallest model supporting images. Running it on raspberry pi, I got stable 74 degrees celcius and quite a few words per minute. It is promising. In order to use it, we will limit or disable thinking, and add cooling to the raspberry pi. 

![qwen.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6Mjg4MDUsInB1ciI6ImJsb2JfaWQifX0=--2c6ff4ffd7ec32fdad8f591936093fa2f1d57dcf/qwen.png)

All that's concerning is for the LLM not to take all the resources of the raspberry pi, as there are way more critical components to the system, such as sensor readings and actuator control. I should find a way to ensure that physical component control is prioritised over LLM.
I am also considering if getting something adjacent to Raspberry Pi AI HAT+ 2 or Coral AI accelerator would be a good idea. If budget allows, this could be the solution.

I have added guide for installation of Ollama and correct model for diet-pi. It is now updated and up on GitHub. Next, I will start creating harness/brain for our robot.

### Recording Links

- https://lookout.hackclub.com/api/media/13c74f3d-05cf-4dea-bc6a-d35280aff4a2/video.mp4
- https://lookout.hackclub.com/api/media/7bfc6e0c-6d38-4681-a48a-18864ea3f6f2/video.mp4

## Entry 52
- ID: 12542
- Author: William
- Created At: 2026-06-08T17:10:30Z

### Content

Alright, so I finished up the gripper and arm assembly as a whole. I still need to do some weight reduction, although a lot of that can be accomplished with slicer tweaks. I didn't record it, but I got my second 3d printer up and running with some new firmware, which will let it go a whole lot faster. Still working on tuning it. For the gripper, I went with a worm drive to a normal 6mm geared motor; even the smallest servo was too big to fit, so I went with the smallest motor I could find. I'm hoping that I can use the current draw to detect force so it can know how hard it's holding something. The worm drive also makes it not backdriveable, so that's pretty cool. The grippers are compliant, so they can grip a multitude of shapes. I'm going to keep working on maintaining and upgrading the printers cause theyll need it. Also working on making a phone stand to record IRL.
![image.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6Mjk1MDgsInB1ciI6ImJsb2JfaWQifX0=--7417a2652f91f460dfd6c500dc0706da06f1d4e3/image.png)
![image.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6Mjk1MDksInB1ciI6ImJsb2JfaWQifX0=--1add0d790aed72499437386aa03ee65bdb57f379/image.png)
![image.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6Mjk1MTAsInB1ciI6ImJsb2JfaWQifX0=--f6a8d5b933b2cc0495bd3183a3318735ba0c26d3/image.png)


### Recording Links

- https://lookout.hackclub.com/api/media/a7d1f5be-a928-474f-a2b8-96d8435a8d1d/video.mp4

## Entry 53
- ID: 12759
- Author: Umarbek
- Created At: 2026-06-09T13:59:20Z

### Content

To be completely honest, I had my wisdom tooth taken out this afternoon after school, and pain has greatly reduced my capacity to do work ("Energy" as by physics definition, lol?). Due to anesthesia, my brain is foggy, but here's what we got:
1. A beautiful orchestrator/harness for the local LLM.
1. A confirmation that local models are too dumb (yes, I'm referring to you, Qwen2.5vl:3b, Qwen3.5:0.8b, and moondream.) As my teammate has proposed yesterday, we will (unless some genius will invent a way to make smaller models substantially smarter, helping me with this project and crippling Ai industry in the background) have to fine tune it on our tool-calling.


![criticalig.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6Mjk5NzEsInB1ciI6ImJsb2JfaWQifX0=--ea0454bfddae674d6e5ccfb24657c6527b5ee9de/criticalig.png)


### Recording Links

- https://lookout.hackclub.com/api/media/b7faaf59-ce58-4164-af3a-543abe37c29d/video.mp4
- https://lookout.hackclub.com/api/media/f534824b-13ac-4661-99e2-a847e2e1e7a5/video.mp4

## Entry 54
- ID: 12981
- Author: Umarbek
- Created At: 2026-06-10T10:05:45Z

### Content

I've written documentation on how to activate I2C protocol on raspberry pi with dietpi. (I am now realising how many steps does MIST take to configure, writing an automatic script later in time would be a good choice.)

I have encountered many many issues with our current chosen model of the IMU, which includes driver issues, protocol issues, wiring issues and more.
Here are the options I've considered:
We could do one of the following:
1. Limit amount of data we get, and stream more primitive readings to some part of the OS as a peripheral, which will be read by our singleton.
2. Have a contained python script that will stream all the data to our Go singleton
3. Just get another IMU (MPU-6050 seems like a great choice tbg) (way cheaper as well)

And honestly, since we haven't started building her yet, we can at no loss choose option 3, saving us a lot of pain and headache in the future. So far, MPU6050 seems like a good option with great support (and surprisingly lower cost). I will now try adjust the wiring (I wrote most of the updated script). 

![imu.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzA1NTUsInB1ciI6ImJsb2JfaWQifX0=--3342fc73fdd7c51a1a9fe3e6dc198e3c55dc7acd/imu.png)

Now, in writing this, I realise that the IMU data I'm working with are all random numbers. I should also add translator capabilities to turning these numbers into m/s and other SI units. + I'll have to write streaming option rather than function calling for real-time robot body processing later.

![imu2.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzA1NTYsInB1ciI6ImJsb2JfaWQifX0=--2ec58ce6b16c95367091692ad03f7da1031c2595/imu2.png)


### Recording Links

- https://lookout.hackclub.com/api/media/d91175b3-a419-483f-afd4-6bdba9b7a778/video.mp4

## Entry 55
- ID: 13138
- Author: Umarbek
- Created At: 2026-06-11T03:08:45Z

### Content

I have finished working on the IMU package. I have written clean, modular code which (should) provide an easy way to work with the physical sensor.

![imu4.webp](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzEwMjAsInB1ciI6ImJsb2JfaWQifX0=--d38edc8f8b10e880c7c419d5c9fa0ca6013df225/imu4.webp)

It is important to say, however, due to not having physical hardware on myself right now, it is uncertain whether this code would work. I've added configuration information on the main README and few methods of testing if everything is correctly connected and functioning. 

I have also updated the wiring guide on Figma, replacing the old IMU to the new one. (after copying from an online guide I found, at the end, I've double checked-with Gemini, and it said everything is correct, so...)

So far, this is the list of packages in the repository: 
  + camera
  + controller
  + screen
  + voice
  + imu
  ~ ai
  - body
  - hand
  - brain
  - eyes
  - maths
  - main
, and I will probably now go to the body module, working on controlling of the servos and brushless motors as well as finishing the wiring I've started a while ago.

![imu4.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzEwMjEsInB1ciI6ImJsb2JfaWQifX0=--3b28408355ba2b5d445b594c978f4bd7d2889334/imu4.png)


### Recording Links

- https://lookout.hackclub.com/api/media/99ee39db-fea8-4c5a-a9ab-adf71ecee4c2/video.mp4
- https://lookout.hackclub.com/api/media/cb45e654-df09-42f0-baec-a365229295a0/video.mp4
- https://lookout.hackclub.com/api/media/ff738864-4dea-4f67-98ef-abfa4447e220/video.mp4
- https://lookout.hackclub.com/api/media/49097bdb-58df-40e6-8e99-43562d5627e1/video.mp4

## Entry 56
- ID: 13200
- Author: Umarbek
- Created At: 2026-06-11T10:46:41Z

### Content

I have created a module for servos and motors! This was surprisingly easy to be honest, and controlling them is rather fun than tedious. 


![wfikxmg6r6a51.jpg](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzExODgsInB1ciI6ImJsb2JfaWQifX0=--db3df86a15e3de2da19dbf19e373ae4528621b96/wfikxmg6r6a51.jpg)

I think it is mainly due to Esia choosing a popular variant which has a good documentation and community support.

The PCA9685 provides a clear minimal interface, making development enjoyable!

![servo.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzExNTEsInB1ciI6ImJsb2JfaWQifX0=--8640c89ca5dfb267eba189ee49eaedcf232cd01c/servo.png)

### Recording Links

- https://lookout.hackclub.com/api/media/7259157b-d89a-4b31-b6e3-c608d9174642/video.mp4
- https://lookout.hackclub.com/api/media/9a7b8910-b7c3-4ef3-878a-d0689df26387/video.mp4
- https://lookout.hackclub.com/api/media/cce0d27e-6592-46a8-b883-2cebafdc7ebe/video.mp4

## Entry 57
- ID: 13226
- Author: Umarbek
- Created At: 2026-06-11T13:56:05Z

### Content

I have worked further on the wiring for MIST, specifically focusing on power and its distribution; it was a challenge. It is uncertain what specific hardware to use for some parts such as Battery, Power Distribution Board, Buck Converters.

Additionally, it seems like we've been using wrong motor type for the job, and I only now came to a realisation on the difference between aerial motors and ground motors, with aerial motors having nearly no torque. Remembering my previous conversations and looking into CAD model (why navigation in CAD with a trackpad is so challenging for no reason? is it really hard figure it out like Blender did decades back?), my friend and teammate has implemented gear reduction mechanisms, but the question is: is this complexity and uncertainty justified or should we just use a ground motor?

(Btw, I'm unsure if it applies to everything but I've noticed that harder it is to wire something, the easier it is to program it and vice-versa.)

It is still unfinished, and the wiring must be reviewed, but the progress is honestly great! Here's the progress so far:

![wiring.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzEyNjQsInB1ciI6ImJsb2JfaWQifX0=--c19e1daab9817f38b56843ba663cac368ad5b3a6/wiring.png)


### Recording Links

- https://lookout.hackclub.com/api/media/e03a89be-4320-4bd8-8cf8-6a84bc3276c6/video.mp4

## Entry 58
- ID: 13674
- Author: Umarbek
- Created At: 2026-06-13T08:14:19Z

### Content

In earlier designs of methods of controlling MIST, there were too many different options. Thus, there was a proposition to use a singular button in a morse-code style in order to expand upon available inputs. (Currently, it is to be used in controlling screen types.)

![meme.jpeg](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzIzOTksInB1ciI6ImJsb2JfaWQifX0=--1a457dfa01a1a57761ef07ae7eb76866a9704f9a/meme.jpeg)

Creation of telegraph module in golang was quite engaging, and seeing the results felt magical? Turns out, after fine-tuning, hard-coded time durations for difference between dot and the line as well as idle time until the character is recorded it works flawlessly!

![proof.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzI0MDAsInB1ciI6ImJsb2JfaWQifX0=--243aeec5b5ac5c564ba3ef2f35042f725955f786/proof.png)


### Recording Links

- https://lookout.hackclub.com/api/media/d6c60a16-7392-4168-9568-b9d8c83229c3/video.mp4

## Entry 59
- ID: 13720
- Author: Umarbek
- Created At: 2026-06-13T13:36:01Z

### Content

I have finally embarked on a journey of actually figuring out MIST's higher-level physical controls. But before that... of course, I have decided to rewrite the structure of actuators' code. 

Now, instead of having a large object where one must have all of the robot's motors at once in order to control a singular servo on the head, it is possible to create a singular actuator controller, which can then be passed as an argument to NewServo(), NewMotor() functions to have a pointer and methods for a specific motor. 

This significantly simplifies the process of working with the hardware, keeping everything modular and safe: no servo functions available on motors anymore!)

I have also thought more on how to map controls from a gamepad onto movements in MIST. I was struggling with mediocre-limiting methods until I remembered how much I love drones! And how closely MIST resembles one! Therefore, it is now decided. 

For the higher level body functions, I have implemented one for the head, and partially for wheels (the actual way of controlling mist was important here, so I steered off a little as you see.) I have also realised that we could get more data from the controller from its gyroscope and accelerometers! 

Next, I will forget about the body for a little, implementing the functionality of gyro+accel read on a controller, after which I'll come back to the body, finishing the head off, and continuing with wheels, knees, and hand.

![meme.jpeg](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzI1MzIsInB1ciI6ImJsb2JfaWQifX0=--d881cb1893083385934109e6f2bdc8fe1a1ef7a3/meme.jpeg)

![proof.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzI1MzMsInB1ciI6ImJsb2JfaWQifX0=--0c2ff5aabe5d4dcffbd46f06553db4f0283610fd/proof.png)


### Recording Links

- https://lookout.hackclub.com/api/media/280d6d38-8387-4936-9191-2e566fb74599/video.mp4
- https://lookout.hackclub.com/api/media/80d52b06-6a14-4ccf-9572-3f7f793cf8a3/video.mp4

## Entry 60
- ID: 13733
- Author: Umarbek
- Created At: 2026-06-13T14:36:51Z

### Content

I have worked on getting the accelerometer and gyroscope data from the controller. The stable version of sdl I was using turned out to be too old, requiring me to update to the new alpha version, which stalled. In a dubious debugging session, I've figured out that the new version has switched from iterating pointers to values directly. Thus, updating the switch statement by just removing the *'s has fixed all of the problems. 

![proof2.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzI1NzQsInB1ciI6ImJsb2JfaWQifX0=--1f7c9091dedbb71e70e923046ed27f8a4f12214f/proof2.png)


### Recording Links

- https://lookout.hackclub.com/api/media/f3f1865a-eff5-4367-a71a-118005a6befc/video.mp4

## Entry 61
- ID: 13878
- Author: Umarbek
- Created At: 2026-06-14T03:39:34Z

### Content

I have finished upon figuring out the controls and writing code for wheels! It will be controlled in this way: 

R3:
- forward-backward: movement, riding in that direction by reversing motor direction
- right/left tilt: rotation on place (left back right forward and wise-versa)
L3:
- forward-backward: up-down (controlled from the servos on knees)
- right/left: turning (forward legs turn up to 30 degrees right/left, like in a car)
(Gemini wrote a simulation from which my brother and I have tested the controls, and it seems to be perfect!)

![proof1.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzI5NTMsInB1ciI6ImJsb2JfaWQifX0=--0859ec20c591e50152b0250cd1cf4da1baab1cff/proof1.png)

I have also updated the motor class from the actuator module to support bidirectional movement (it only worked for movement in one direction before.) Without physical hardware on my hands, it is hard to confirm if the codebase actually works, but we'll be hoping for the best. 

Now, when choosing a new ESC Control, we must ensure it supports bidirectional throttle. 

![proof2.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzI5NTQsInB1ciI6ImJsb2JfaWQifX0=--9ae2992756bfbf7311e49d3c8736b6f086579f14/proof2.png)


### Recording Links

- https://lookout.hackclub.com/api/media/d9819b1d-a102-4f9f-b75d-0a7de2064f19/video.mp4

## Entry 62
- ID: 13887
- Author: Umarbek
- Created At: 2026-06-14T04:25:05Z

### Content

In working on the knees (which is a headache of its own 8 servos worth), I got distracted a little, and worked a little on what MIST's shot for the zine might look like. I've exported the FreeCAD model to blender (unfortunately the mesh was ugly and unusable, leaving me with the straight standing pose unable to rig or even simply move it.)

I applied a glass shader, created a simple 3 point light setup and rendered it, getting this masterpiece:

![700830dd31cd71a6043b59b7ab4980c86fc773206602666de0ff170cb3e90d8b.jpeg](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzI5NjksInB1ciI6ImJsb2JfaWQifX0=--4e6e280d25b01e1b016f76736cef26c33f3af96b/700830dd31cd71a6043b59b7ab4980c86fc773206602666de0ff170cb3e90d8b.jpeg)

I've put settings on low in order to see approximately how will it look; I will later render it at better quality with more passes (in order to avoid using denoise and making it blurry.). Additionally, this is an outdated model, and I will replace it with articulated final model from Esia (hopefully) tomorrow.

### Recording Links

- https://lookout.hackclub.com/api/media/d80e6c78-49a5-45c5-8771-ec745aeba7b4/video.mp4

## Entry 63
- ID: 13914
- Author: Umarbek
- Created At: 2026-06-14T08:15:47Z

### Content

The hardest part of coding MIST is now finished!!

The 8 servos responsible for the knees/legs seemed as an impossible challenge where I will struggle a lot. Thanks to the YouTube and couple great creators, I have understood the mathematics behind forward, and backward kinematics sufficient enough for a joint of 2 degrees of freedom.

![proof1.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzMwNDYsInB1ciI6ImJsb2JfaWQifX0=--7f9aaf97eb40bc97d7528f23006489aefa249770/proof1.png)


![proof2.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzMwNDcsInB1ciI6ImJsb2JfaWQifX0=--a8614956173314b6b03d40e962252e10dfa6ef1c/proof2.png)


Now, MIST tilts forward and backward, and can go up and down in height. I have now remembered that lastly, I need mechanics to turn right and left. I will work on it afterwards.

![meme.jpg.webp](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzMwNDgsInB1ciI6ImJsb2JfaWQifX0=--19677de607c2baaf9d052eb006d776579df3046c/meme.jpg.webp)


### Recording Links

- https://lookout.hackclub.com/api/media/6b293bb0-c594-4091-bba1-9f254aef85d2/video.mp4
- https://lookout.hackclub.com/api/media/cb53d225-e268-41be-a431-104429ae88ca/video.mp4
- https://lookout.hackclub.com/api/media/c4534e6f-52d1-4d09-95dd-e5d77fecacbe/video.mp4

## Entry 64
- ID: 13917
- Author: Umarbek
- Created At: 2026-06-14T08:37:23Z

### Content

After working with all of the complex modules, animating hips was a little too simple. It just required rotation of two servos from a singular vector input, which I've implemented. Nothing special, just a hip module finished.

Next, I will implement the hand module, after which, I will assemble everything into body class, which will then be controlled by the brain, which finally connects controller inputs to the actuators, making MIST's code complete!

I have also noticed how we've ran out of ports on PCA9685,  we should resolve it later, while making the wiring/assembly guides. For now, I've added it to an invalid port number:

![proof_hips.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzMwNTQsInB1ciI6ImJsb2JfaWQifX0=--ad3581a421300688e81074de995d304fe97abeb8/proof_hips.png)

![Screenshot 2026-06-14 at 17.32.31.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzMwNTUsInB1ciI6ImJsb2JfaWQifX0=--dde06b1898f7656098134293b361ededb0a1a98a/Screenshot 2026-06-14 at 17.32.31.png)


### Recording Links

- https://lookout.hackclub.com/api/media/346da6c9-8b0a-4f8e-ad0f-b579eb5a3252/video.mp4

## Entry 65
- ID: 13959
- Author: Umarbek
- Created At: 2026-06-14T12:20:06Z

### Content

Initially, I was planning to work upon the last, hand module, but it is quite hard to understand what William has designed with black magic. Therefore, I skip this step for now, skipping to body module right away.

![proof.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzMxODUsInB1ciI6ImJsb2JfaWQifX0=--33b0de334cdb18e9368a51b68a2e116c23111546/proof.png)


I have created a body module, that combines controller inputs and all of the actuators in order to finally make MIST move! It's been a long way until here and there's a little more left, mainly combining the screen inputs, and camera to the body module.

Further, if sufficient time is given, I should add intelligence layer on top. I completely give up on local LLM models: they are too dumb (though I will probably add whisper for speech to text and embedding Gemma through Ollama for tiny local database and facial recognition.) These three things should allow to have autonomous MIST, though who knows how much tokens will this setup burn through.

### Recording Links

- https://lookout.hackclub.com/api/media/b5b0914e-1c50-4ca3-95fd-2266c3870ef9/video.mp4
- https://lookout.hackclub.com/api/media/0fd0d475-25bc-49f8-b997-21012fdee4af/video.mp4

## Entry 66
- ID: 13989
- Author: Umarbek
- Created At: 2026-06-14T15:05:26Z

### Content

The face module is pain module! 
I am tired of wasting so much time on designing each facial feature one by one. I have come to understanding that I could use an SVG library to design these on my own and just displaying them in screen module. I am taking a rest from the screen module for now...

With creating of a zine! This is a prototype of a zine for now, which uses the earlier concept art of a transparent glass MIST. I've designed it along other peers in a slack huddle, getting some feedback in the process.

![1a34522628508fb483fb3f8e6fba3728600a4098bd7a64154f42e42f2a68b1a8.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzMyNDUsInB1ciI6ImJsb2JfaWQifX0=--c7a094d30c1fcf92e510bfbf93716889e92c0aba/1a34522628508fb483fb3f8e6fba3728600a4098bd7a64154f42e42f2a68b1a8.png)


### Recording Links

- https://lookout.hackclub.com/api/media/4c7b7e78-86d3-49ac-85eb-08d4f0c67e25/video.mp4
- https://lookout.hackclub.com/api/media/33043589-afa3-423f-891b-8b3dd3c9966d/video.mp4

## Entry 67
- ID: 14153
- Author: Umarbek
- Created At: 2026-06-15T03:00:25Z

### Content

The codebase for MIST is 100% complete, and (should be) functional!! The satisfaction of opening dozens of modules in the main.go and brain.go is indescribable.

![5c759513190b9e4919195ea7fe958517.jpg](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzM2MjMsInB1ciI6ImJsb2JfaWQifX0=--24e53f870667d6025e2703407d31b5e568b46c78/5c759513190b9e4919195ea7fe958517.jpg)

In fact, everything is so elegantly modularised that this is all that's needed from main.go:

![proofx.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzM2MjQsInB1ciI6ImJsb2JfaWQifX0=--9310b21dd9b857b7ad220b45dc3cded5b735095a/proofx.png)


### Recording Links

- https://lookout.hackclub.com/api/media/8ba2a7b3-8ee7-487d-ad0d-4bdd8d6f8f81/video.mp4

## Entry 68
- ID: 14176
- Author: Umarbek
- Created At: 2026-06-15T05:35:58Z

### Content

This is the continuation of creating a brain module for MIST. It is a continuation of the previous journal. For the photo, I will just add a screenshot of brain module. 

![proof_b.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzM2NzYsInB1ciI6ImJsb2JfaWQifX0=--bcebc4f03b9c759a30a93ad154daab43309b359d/proof_b.png)


### Recording Links

- https://lookout.hackclub.com/api/media/43fe7301-b290-4b3f-9af7-63f3ac0ac494/video.mp4

## Entry 69
- ID: 14179
- Author: Umarbek
- Created At: 2026-06-15T05:37:54Z

### Content

I have completed wiring schematics for the base robot (all excluding the arm). We already had a quite good scheme but it was all over the place; now, everything is colour-coded and available in a singular image with port allocations.

![Circuit Guide.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzM2NzcsInB1ciI6ImJsb2JfaWQifX0=--5ca4f888a2578fcd06ca9646bf8b01ab034b1a01/Circuit Guide.png)


### Recording Links

- https://lookout.hackclub.com/api/media/cb6ad4ac-8ec8-426e-8631-93b50476c625/video.mp4

## Entry 70
- ID: 14273
- Author: Umarbek
- Created At: 2026-06-15T16:49:29Z

### Content

This was supposed to be part of the previous recording, where I have finished the wiring. I forgot to stop the recording before typing in and submitting the previous journal. I will add the same diagram as the last time: 

![proop.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzM5NjksInB1ciI6ImJsb2JfaWQifX0=--f5536956ad38eabc4e5837617d5cd4e61d89fef2/proop.png)


### Recording Links

- https://lookout.hackclub.com/api/media/3bec42fa-f9be-48d5-a74e-e7a8c68433b7/video.mp4

## Entry 71
- ID: 14300
- Author: Umarbek
- Created At: 2026-06-15T19:05:32Z

### Content

I have been doing last touch up changes to the MIST before submitting it for review. I've touched upon few small details in wiring, did the whole exporting, file conversion black magic, and rendered the last version of MIST in Blender, now with updated colour  palette. 

I also made the last iteration of zine!

![1.png](/user-attachments/blobs/redirect/eyJfcmFpbHMiOnsiZGF0YSI6MzQwMzMsInB1ciI6ImJsb2JfaWQifX0=--d4cdfbd597a5118178ac5617eee38ce50ea72268/1.png)

### Recording Links

- https://lookout.hackclub.com/api/media/8088257b-fb5a-4517-8af7-9f07ef3d676d/video.mp4
