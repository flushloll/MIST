"""
# Brain / Singleton Loop

MIST's Behavioural suggestion:
While Disney's robots are impressive in performance, the primary goal in creating MIST is authenticity and usefulness. She isn't merely an object of entertainment but rather an alive-like creature that posesses personality, full control over their body and natural procedural non-pre-defined movements. Thus, the problem of defining the robot's AI becomes harder than just putting her CAD model into NVIDIA Isaac Lab and letting her learn how to walk (though that could be a part of it).

As a human being, my movements are action-based. I do not care what my posture would be and how I hold fingers in my hands in walking. Similarly, I do not care how my legs or eyebrows are positioned when typing. I control only the part of my body I care about: face+hands+torso when reacting, legs+head while crossing the road... Our lives are objective-based, with brain thinking of controlling a part of the whole body that is connected to the objective the most in the time of execution, with it being able to switch which part of the body to focus on immiediately. Asking an LLM for each action would be lazy, slow, error-prone, and unoriginal.

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