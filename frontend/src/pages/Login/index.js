import React, { useRef, useState } from "react";
import Footer from "../../components/Footer";
import Header from "../../components/Header";
import { Navigate } from "react-router-dom";

import Axios from "axios";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { auth_loginUrl } from "../../resources/constants.js";

import {
  ContainerFormLogin,
  FormLogin,
  InputProfile,
  ContainerAllButtons,
  ContainerButton,
  ButtonLogin,
  ButtonRegister,
} from "./styles";
const Login = () => {
  const emailRef = useRef();
  const passwordRef = useRef();

  const [login, setLogin] = useState(false);

  function handleSubmit(event) {
    event.preventDefault();
    console.log("handle submit");

    loginUser({
      email: emailRef.current.value,
      password: passwordRef.current.value,
    });
  }
  async function loginUser(credentials) {
    console.log("login user");
    console.log(credentials);
    const url = auth_loginUrl;
    try {
      const response = await Axios.post(url, JSON.stringify(credentials), {
        headers: { "Content-Type": "application/json" },
      });
      console.log("user is valid!");
      console.log(response.data.token);

      setLogin(true);

      sessionStorage.setItem("token", response.data.token);
    } catch (error) {
      sessionStorage.removeItem("token");
      setLogin(false);

      console.log(error);
      toast.error(error.response.data.message, {
        position: toast.POSITION.TOP_RIGHT,
      });
      console.log("user is invalid!");
    }
  }

  function registerButton() {}
  return (
    <>
      <Header />
      {login ? (
        <>
          <Navigate to="/gallery" replace={true} />
        </>
      ) : (
        <ContainerFormLogin>
         <ToastContainer />

          <FormLogin onSubmit={handleSubmit}>
            <InputProfile
              placeholder="Email"
              name="email"
              type="text"
              ref={emailRef}
              required
            ></InputProfile>
            <InputProfile
              placeholder="Password"
              name="passwUser"
              type="password"
              ref={passwordRef}
              required
            ></InputProfile>
            <ContainerAllButtons>
              <ContainerButton>
                <ButtonLogin type="submit">Login</ButtonLogin>
              </ContainerButton>
              <ContainerButton>
                <ButtonRegister type="button" onClick={registerButton}>
                  Register
                </ButtonRegister>
              </ContainerButton>
            </ContainerAllButtons>
          </FormLogin>
        </ContainerFormLogin>
      )}
      <Footer />
    </>
  );
};

export default Login;