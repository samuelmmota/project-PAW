import React from "react";
import Footer from "../../components/Footer";
import Header from "../../components/Header";

import {
  ContainerRegister,
  ContainerInfosUser,
  InputChangeImage,
  ContainerInputs,
  ContainerButtonRegister,
  InputEditUser,
  ButtonSubmitRegister,
  Title,
} from "./styles";

const Register = () => {

  function register(e) {
    //previne reload à pagina
    e.preventDefault();
    console.log("Register User");
  }

  return (
    <>
      <Header />
      <Title>Register</Title>
      <ContainerRegister>
        <ContainerInfosUser onSubmit={register}>
          <ContainerInputs>
            <InputChangeImage
              name="registerEmail"
              type="email"
              placeholder="example@email.com"
              id="login_input_email"
              required
            />
            <InputEditUser
              name="registerPassword"
              placeholder="Insert password"
              type="password"
              id="login_input_password"
              required
            />
            <InputEditUser
              name="registerRepeatedPassword"
              placeholder="Insert repeated password"
              type="password"
              id="login_input_repeated_password"
              required
            />
             <InputChangeImage
              name="registerName"
              type="text"
              placeholder="Insert your name"
              id="login_input_name"
              required
            />
          </ContainerInputs>
          <ContainerButtonRegister>
            <ButtonSubmitRegister>Register</ButtonSubmitRegister>
          </ContainerButtonRegister>
        </ContainerInfosUser>
      </ContainerRegister>
      <Footer />
    </>
  );
};

export default Register;