import styled from "styled-components";
import { primaryColor, secondaryColor } from "../../resources/constants.js";

export const ContainerRegister = styled.div`
  display: flex;
  justify-content: center;
`;

export const ContainerInfosUser = styled.form`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  max-width: 50%;
`;

export const ContainerInputs = styled.div`
  margin: 20px 0px;
  display: flex;
  justify-content: center;
  flex-direction: column;
`;

export const InputChangeImage = styled.input`
  color: #000;
  font-size: 16px;
  margin: 4px 0;
  padding: 2px 8px;
`;

export const InputEditUser = styled.input`
  color: #000;
  font-size: 16px;
  margin: 4px 0;
  padding: 2px 8px;
`;

export const InputEditUserDescription = styled.input`
  color: #000;
  font-size: 16px;
  margin: 4px 0;
  padding: 2px 8px;
  height: 200px;
`;

export const ContainerButtonRegister = styled.div`
  display: flex;
  justify-content: flex-end;
`;

export const ButtonSubmitRegister = styled.button`
  background-color: ${primaryColor};
  color: #fff;
  border-radius: 16px;
  border: none;
  padding: 8px 20px;
  text-transform: uppercase;
  font-weight: bold;

  &:hover {
    opacity: 100%;
    -webkit-transform: scale(1.4);
    transform: scale(1.1);
  }
`;

export const Title = styled.h1`
text-align: center;
margin-top: 20px;
font-size: 2rem;
color: #333;
text-transform: uppercase;
`;


export const ContainerFormRegister = styled.div`
  display: flex;
  justify-content: center;
`;


export const FormRegister = styled.form`
  margin: 40px 0px;
  background: #ffffff;
  box-shadow: 0px 0px 15px 1px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
  padding: 60px 80px;
`;

export const InputProfile = styled.input`
  filter: drop-shadow(0px 0px 15px rgba(0, 0, 0, 0.1));
  border: none;
  margin-bottom: 10px;
  outline: none;
  box-sizing: border-box;
  font-size: 18px;
  padding: 8px 16px;
  display: flex;
  justify-content: center;

  &::placeholder {
    color: #000;
    opacity: 30%;
    justify-content: center;
  }
`;

export const ContainerAllButtons = styled.div`
  margin-top: 15%;
`;

export const ContainerButton = styled.div`
  display: flex;
  justify-content: center;
`;

export const ButtonRegister = styled.button`
  margin-top: 8px;
  border: none;
  color: #000;
  padding: 8px 54px;
  background: ${secondaryColor};
  border-radius: 4px;
  font-size: 18px;

  &:hover {
    -webkit-transform: scale(1.4);
    transform: scale(1.1);
  }
`;
