import styled from "styled-components";
import { primaryColor, secondaryColor } from "../../resources/constants.js";

export const Button = styled.button`
  background-color: ${primaryColor};
  border: none;
  border-radius: 5px;
  color: #fff;
  cursor: pointer;
  font-family: "Anton", sans-serif;
  font-size: 18px;
  padding: 8px 20px;
  margin: 8px;
  transition: 0.3s;
`;

export const ContainerSubmission = styled.div`
  padding: 20px;
  padding-bottom: 0;
  width: fit-content;
`;

export const TitleSubmission = styled.p`
  text-align: center;
  margin-bottom: 0;
`;

export const ContainerImage = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  width: fit-content;
`;

export const ImageSubmission = styled.img`
  margin: 8px;
  max-width: 60%;
  width: 300px;
  height: 225px;
`;

export const ButtonContainer = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
`;

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

export const VideoSubmission = styled.video`
  margin: 8px;
  max-width: 60%;
  width: 300px;
  height: 225px;
`;