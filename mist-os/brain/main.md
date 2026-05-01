"""
# Brain / Singleton Loop

MIST's Behavioural suggestion:
While Disney's robots are impressive in performance, the primary goal in creating MIST is authenticity and usefulness. She isn't merely an object of entertainment but rather an alive-like creature that posesses personality, full control over their body and natural procedural non-pre-defined movements. Thus, the problem of defining the robot's AI becomes harder than just putting her CAD model into NVIDIA Isaac Lab and letting her learn how to walk (though that could be a part of it).

As a human being, my movements are action-based. I do not care what my posture would be and how I hold fingers in my hands in walking. Similarly, I do not care how my legs or eyebrows are positioned when typing. I control only the part of my body I care about: face+hands+torso when reacting, legs+head while crossing the road... Our lives are objective-based, with brain thinking of controlling a part of the whole body that is connected to the objective the most in the time of execution, with it being able to switch which part of the body to focus on immiediately. Asking an LLM for each action would be lazy, slow, error-prone, and unoriginal.

Maybe we could embed the robot's emotional state as a vector/matrix which will be crafted so each axis is a range of certain emotion. I'll definitely have to read on the proposed psychological models of human mind and emotion or something on those lines. Nonetheless, the point is that we could potentially use this emotion vector as a point as an input which acts as a filter on all of the robot's movements, as curve modifiers on the frame interpolation. (I really like Disney's RobotKeyframing idea and this approach would most likely be partially adopted by our robot.)

It might sound cliche and I don't like what I'm saying, but Maslow's Hierarchy of Needs /  Herzberg's Two-Factor Theory from business management course seem quite fitting. Separating each of the robot's possible actions into layers, starting from physical needs to "self-actualisation" (whatever it would be in this case), seems like a good initial model to start with. Robot's actions will be determined by things like: "Does it have enough charge? Are all of the body parts responding and functioning?" and if it is all satisfied, only then we move to the more complex actions of exploring surrounding, meeting people, answering questions, and more.

Layer 3: Intelligence-long; tries to understand intent from the user or explore to choose what is the current goal.
Layer 2: Intelligence-short; performing of the current goal.
Layer 1: Personality/Mood; defining the current emotional state of teh robot, how fast/energetically does the robot move.
Layer 0: Instinct; balance, battery, mapping the environment, not-bumping into things.

For the recognition and tracing of the environment, I always thought that 3D-environmental maps like ones from DJI would be the best option, but that is computationally expensive, and not fitting for the first MIST model. However, as we're already planning to put a camera, with facial and object recognition system, maybe, as we recognise objects, we could store them in the order of discovery or groups if seen simultaneously. Combining it with how much we've moved motors and in which direction, we could potentially make an approximate map of the surroundings, which could be useful in back-tracking or even search of objects in some cases? This will be explored upon later.

As we're attaching a camera to the robot, why not go all in with a full-spectrum camera? It would not only allow us to just take pretty pictures of what we don't see but also see in the dark, see alive beaings, and more. The seeing in the dark especially would be way more helpful for turned off lights in the room or night.

Social battery, boredom? Mapping of People and their personalities?

Perhaps for the memory we could use an embedding model and SQLite? 

Should we allow her to write functions/skils for herself?

We could and probably should add an [UWB](https://ja.aliexpress.com/item/1005001575673574.html?spm=a2g0o.wiki_detail.0.0.28ee65f4OaKNx0&gatewayAdapt=glo2jpn)-styled module for positioning of MIST. It would help maintaining the "follow" functionality, especially in larger crows like where we're going to, fallout. Gladly, they seem to be super-cheap. How should the "charm" / "legendary item which grands you the status of the MIST's owner" look like?

## mini-study: Welch Labs
Author tells the story of [how physical intelligence](https://youtu.be/2mrGMMmrVNE?si=-nkhhYY6tyNthH1P) has developed over time.

## mini-study: Iterator
## mini-study: Larry
## mini-study: Cozmo & Vector

## mini-study: Disney's Droid
[Disney's Droid Example](https://youtu.be/7_LW7u-nk6Q?si=kZ3XEBfBi1F7m4WY) shows a method on they've introduced different ways of robot movement. It features both freeform animations and RL-based movement which takes inputs of walking velocity and robot pose and outputs actuator commands. PPO algorithm is used. It has 5 actuators in each leg and 4 in the neck. Speakers, antenas, eyes. 1 hour battery life.

All motions are separated into 3 categories:
- Perpetual: indefinite such as standing.
- Periodic: cycling motions like walking.
- Episodic: pre-defined duration like animatoin clips.
During control, animatoin engine looks at the databse of pre-recorded animations and chooses the optimal one depending on the button pressed.

Robustness is obtained through RL and random terrain optimisation and introduction of disturbances during training.

Animation states are managed by layering animations. Background animation + Predifined Sequences + Modulation with Joysticks (such as head turn). Posture and Gaze are separately controlled. The key flaw seems to be that all of the animations are pre-recorded and pre-testied, with the robot just selecting the correct memory and replaying it during performance.

They've also invented method of keyframing movements; it looks promising in achieving multi-purple steps: [RobotKeyframing](https://youtu.be/YpOABpwdxko?si=oiyq4JZLEhh3ghAP)

"""