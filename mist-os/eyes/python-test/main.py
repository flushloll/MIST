from ultralytics import YOLO
import cv2

# Load YOLOv11 Nano (Detects 80+ objects like persons, laptops, bottles, etc.)
model = YOLO('yolo11n.pt') 

# Load a dedicated Face Detector (Haar Cascade is fast for Pi)
face_cascade = cv2.CascadeClassifier(cv2.data.haarcascades + 'haarcascade_frontalface_default.xml')

cap = cv2.VideoCapture(0)

while cap.isOpened():
    success, frame = cap.read()
    if not success: break

    # 1. Detect 80+ Objects with YOLO
    results = model(frame, verbose=False)
    for result in results:
        for box in result.boxes:
            coords = box.xyxy[0].tolist() # [xmin, ymin, xmax, ymax]
            label = model.names[int(box.cls)]
            print(f"Detected: {label} at {coords}")

    # 2. Detect Faces Specifically
    gray = cv2.cvtColor(frame, cv2.COLOR_BGR2GRAY)
    faces = face_cascade.detectMultiScale(gray, 1.1, 4)
    for (x, y, w, h) in faces:
        print(f"Face found at: [{x}, {y}, {x+w}, {y+h}]")
        cv2.rectangle(frame, (x, y), (x+w, y+h), (255, 0, 0), 2)

    # Display results
    cv2.imshow("Robot Vision", results[0].plot()) # Plot YOLO boxes
    if cv2.waitKey(1) & 0xFF == ord('q'): break

cap.release()
