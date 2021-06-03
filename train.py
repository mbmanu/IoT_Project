import os           #importing modules
import cv2 as cv
import numpy as np

people = ['rahul', 'meghana', 'nevil','shujan', 'mustafa', 'manu']      # people who's face will be detected
DIR = r'/home/manu/Documents/IoT_Project/assets/train'      #location varies

haar_cascade = cv.CascadeClassifier('haar_face.xml')        #face detection xml file

features = []
labels = []

def train_model():      #training function
    for i in people:
        path = os.path.join(DIR, i)
        label = people.index(i)

        for j in os.listdir(path):
            imgPath = os.path.join(path,j)

            img_array = cv.imread(imgPath)
            if img_array is None:
                continue

            gray = cv.cvtColor(img_array, cv.COLOR_BGR2GRAY)        #BGR to gray colour conversion

            faces = haar_cascade.detectMultiScale(gray, scaleFactor=1.1, minNeighbors=4)    #face detection

            for (x,y,w,h) in faces:
                faces_roi = gray[y:y+h, x:x+w]
                features.append(faces_roi)
                labels.append(label)

train_model()           #function call
print('Training complete!')

features = np.array(features, dtype='object')
labels = np.array(labels)

face_recognizer = cv.face.LBPHFaceRecognizer_create()

face_recognizer.train(features, labels)
face_recognizer.save('trained.yml')         #saving the trained model