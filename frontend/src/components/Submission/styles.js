import styled from "styled-components";
import { primaryColor, secondaryColor, tertiaryColor,quaternaryColor } from "../../resources/constants.js";
import ReactPlayer from "react-player";
import { Container } from "@mui/material";

export const Title = styled.h1`
text-align: center;
margin-top: 20px;
font-size: 2rem;
color: #333;
text-transform: uppercase;
`;

export const DateText = styled.p`
  text-align: center;
  font-size: 14px;
  color: #777;
  margin-top: 10px;
`;

export const ContainerSubmission = styled.div`
  width: 300px; /* Set the desired fixed width */
  height: 400px; /* Set the desired fixed height */
  background-color: #f5f5f5; /* Replace with your desired shade of gray */
  padding: 20px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
`;

export const TitleSubmission = styled.p`
  text-align: center;
  margin-bottom: 0;
`;

export const ContainerImage = styled(Container)`
display: flex;
align-items: center;
justify-content: center;
height: 200px; /* Set the desired height for the media container */
`;

export const ImageSubmission = styled.img`
  max-width: 100%;
  max-height: 100%;
`;

export const VideoSubmission = styled(ReactPlayer)`
  max-width: 100%;
  max-height: 100%;
`;

export const Description = styled.p`
  margin-bottom: 10px;
`;

export const ButtonContainer = styled.div`
  display: flex;
  flex-direction: column;
`;

export const Button = styled.button`
  background-color: ${primaryColor};
  border: none;
  border-radius: 5px;
  color: #fff;
  cursor: pointer;
  font-family: "Anton", sans-serif;
  font-size: 18px;
  padding: 8px 20px;
  margin-bottom: 10px;
`;

export const ButtonFeedback = styled.button`
  background-color: ${tertiaryColor};
  border: none;
  border-radius: 5px;
  color: #fff;
  cursor: pointer;
  font-family: "Anton", sans-serif;
  font-size: 18px;
  padding: 8px 20px;
  margin-bottom: 10px;
`;
