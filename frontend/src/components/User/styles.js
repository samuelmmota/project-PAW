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

export const ContainerUser = styled.div`
  padding: 20px;
  padding-bottom: 0;
  width: fit-content;
`;

export const ContainerImage = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  width: fit-content;
`;

export const ImageUser = styled.img`
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

export const Container = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100vh;
`;

export const ProfileSection = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: #f1f1f1;
  padding: 20px;
  border-radius: 10px;
  width: 50%;
  box-shadow: 0px 2px 10px rgba(0, 0, 0, 0.2);
`;

export const ProfileImage = styled.img`
  width: 200px;
  height: 200px;
  object-fit: cover;
  border-radius: 50%;
`;

export const Name = styled.h1`
  margin-top: 20px;
  font-size: 2rem;
  color: #333;
`;

export const Email = styled.p`
  margin-top: 10px;
  font-size: 1.2rem;
  color: #666;
`;
