import React from "react";
import Footer from "../../components/Footer";
import Header from "../../components/Header";
import { 
  Title,
  Layout,
  FeatureContainer,
  FeatureTitle,
  FeatureDescription,
      
} from "./style";
import PageLayout from "../../components/PageLayout";
const Home = () => {
  return (
    <>
    <PageLayout>
    <Layout>
      <Title>Enhancing Psoriasis Awareness and Support</Title>

      <FeatureContainer>
        <FeatureTitle>User Registration</FeatureTitle>
        <FeatureDescription>
          The users will be able to register themselves on the website using their email address and password.
          Once registered, the user will be authenticated and can access the features of the website.
        </FeatureDescription>
      </FeatureContainer>

      <FeatureContainer>
        <FeatureTitle>Image Submission</FeatureTitle>
        <FeatureDescription>
          Authenticated users will be able to upload pictures of their skin affected by psoriasis. The images will be stored on the server-side.
        </FeatureDescription>
      </FeatureContainer>

      <FeatureContainer>
        <FeatureTitle>Image Management</FeatureTitle>
        <FeatureDescription>
          Authenticated users will be able to view a list of their uploaded images, update and delete them.
          They will also be able to filter the images based on the date and body positions.
        </FeatureDescription>
      </FeatureContainer>

      <FeatureContainer>
        <FeatureTitle>Clinical Feedback</FeatureTitle>
        <FeatureDescription>
          Users can authorize clinicians to view their uploaded images.
          Authorized clinicians can access the images and provide feedback to the patients.
          Multiple clinicians can be authorized to view the same patient's images.
          Clinicians can leave textual feedback for the patient.
        </FeatureDescription>
      </FeatureContainer>

      <FeatureContainer>
        <FeatureTitle>Image Security</FeatureTitle>
        <FeatureDescription>
          All images will be encrypted using a secure mechanism so that they can only be viewed by the patient and their authorized clinicians.
        </FeatureDescription>
      </FeatureContainer>

      <FeatureContainer>
        <FeatureTitle>Future Enhancements</FeatureTitle>
        <FeatureDescription>
          <ul>
            <li>Implement a feature to allow users to share their images with others for research or educational purposes, with their consent</li>
            <li>Be able to relate a new image to a previous one in order to follow the clinical evolution of a given patient (timeline)</li>
          </ul>
        </FeatureDescription>
      </FeatureContainer>
    </Layout>
    </PageLayout>
    </>
  );
};


// ter acesso ao ficheiro dentro do componente
export default Home;