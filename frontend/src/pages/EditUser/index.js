import React from "react";
import Footer from "../../components/Footer";
import Header from "../../components/Header";

import {
  ContainerRegister,
  ContainerInfosUser,
  InputChangeImage,
  ContainerInputs,
  ContainerButtonRegister,
  ButtonSubmitRegister,
  Title,
} from "./styles";

const Edituser = () => {

  function editUser(e) {
    //previne reload à pagina
    e.preventDefault();
    console.log("edit user");
  }

  return (
    <>
      <Header />
      <Title>Edit User</Title>
      <ContainerRegister>
        <ContainerInfosUser onSubmit={editUser}>
          <ContainerInputs>
            <InputChangeImage
              name="registerEmail"
              type="email"
              placeholder="editedemail@email.com"
              id="login_input_email"
            />
             <InputChangeImage
              name="registerName"
              type="text"
              placeholder="Edit your name"
              id="login_input_name"
            />
          </ContainerInputs>
          <ContainerButtonRegister>
            <ButtonSubmitRegister>Update</ButtonSubmitRegister>
          </ContainerButtonRegister>
        </ContainerInfosUser>
      </ContainerRegister>
      <Footer />
    </>
  );
};

export default Edituser;
