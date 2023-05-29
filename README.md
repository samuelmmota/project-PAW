# project-PAW
Project (PAW) Programação de Aplicações Web 2022-2023


PAW 

1. User Registration: The users will be able to register themselves on the website using their email address and password. Once registered, the user will be authenticated and can access the features of the website.
   [x] a. Users will navigate to the website and select the "Register" button.
   [x] b. Users will be directed to a registration form where they will enter their email address and password.
   [x] c. Once the user submits the registration form, the server will verify the email address is valid and not already registered.
   [x] d. If the email address is valid and not already registered, the server will create a new user account and authenticate the user.
   [x] e. The user will be redirected to the homepage of the website.

2. Image Submission: Authenticated users will be able to upload pictures of their skin affected by psoriasis. The images will be stored on the server-side.
   [] a. Authenticated users will navigate to the "Upload Image" page.
   [x] b. Users will select an image file from their local device and provide a description and the body positions of the image (e.g., head, stomach, leg, arm, etc.).
   [] c. The server will encrypt the image using a secure encryption mechanism.
   [] d. The encrypted image and the image description will be stored in the database. 
   [] e. The user will be redirected to the "My Images" page (i.e., gallery of images).

3. Image Management: Authenticated users will be able to view a list of their uploaded images, update and delete them. They will also be able to filter the images based on the date and body positions.
   [] a. Authenticated users will navigate to the "My Images" page. 
   [x] b. Users will be able to view a list of their uploaded images.
Web Application Design & Project in Systems and Networks 2
   [] c. Users will be able to select an image and view a larger version of the image along with its description.
   [] d. Users will be able to update an image's description.
   [] e. Users will be able to delete an image.

4. Clinical Feedback: Users can authorize clinicians to view their uploaded images. Authorized clinicians can access the images and provide feedback to the patients. Multiple clinicians can be authorized to view the same patient's images. Clinicians can leave textual feedback for the patient.
   [] a. Users will navigate to the "Clinical Feedback" page.
   [] b. Users will select which clinicians they want to authorize to view their uploaded images.
   [] c. Clinicians will navigate to the "Clinical Feedback" page and view the authorized patient's uploaded images.
   [] d. Clinicians will be able to provide textual feedback for each image.
   [] e. Clinicians will be able to view their own feedback and feedback provided by other authorized clinicians.

5. Image Security: All images will be encrypted using a secure mechanism so that they can only be viewed by the patient and their authorized clinicians.
   [] a. All images will be encrypted using a secure encryption mechanism before being stored in the database.
   [] b. Only the patient and authorized clinicians will be able to view the images by logging into the web application.
   [] c. The web application will use appropriate security measures to protect user data, such as, SSL/TLS encryption, secure password storage, and two-factor authentication.

6. Future Enhancements:
   [] a. Add support for video submissions to allow users to capture the progression of their psoriasis over time.
   [] b. Implement a feature to allow users to share their images with others for research or educational purposes, with their consent
   [] c. Be able to relate a new image to a previous one in order to follow the clinical evolution of a given patient (timeline)