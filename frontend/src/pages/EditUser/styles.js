import styled from "styled-components";

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
  background-color: #fdc544;
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
