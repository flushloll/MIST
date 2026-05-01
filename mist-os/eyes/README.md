# Eyes

The aim of this system is to analyse the camera feed and return insightful data for the main brain to process. It is quite hard to make an elegant system which doesn't consume computer's all resources, requiring thousands of packages. We might have to run the model as a docker container and streaming the camera input to the container. In the end, I got it working! Now I must re-evaluate the choices and compare using docker vs native.

What we'll have:
- Full-spectrum camera feed

We want:
- Object Detection
- Facial Recognition
- Point-of-attention