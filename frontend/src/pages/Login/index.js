import React, { useRef, useState } from "react";
import { Navigate } from "react-router-dom";

import Axios from "axios";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { auth_loginUrl, userUrl } from "../../resources/constants.js";

import {
  ContainerFormLogin,
  FormLogin,
  InputProfile,
  ContainerAllButtons,
  ContainerButton,
  ButtonLogin,
  ButtonRegister,
} from "./styles";
import PageLayout from "../../components/PageLayout";
const Login = () => {
  const emailRef = useRef();
  const passwordRef = useRef();

  const [login, setLogin] = useState(false);

  function handleSubmit(event, type) {
    event.preventDefault();
    console.log("handle submit");
    const email = emailRef.current.value;
    const password = passwordRef.current.value;
  
    if (!email || !password) {
      toast.warning("Please enter both email and password", {
        position: toast.POSITION.TOP_RIGHT,
      });
      return;
    }

    if (type === "login") {
      loginUser({
        email: emailRef.current.value,
        password: passwordRef.current.value,
      });
    } else if (type === "register") {
      registerButton({
        email: emailRef.current.value,
        password: passwordRef.current.value,
      });
    }
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

  async function registerButton(credentials) {
    try {
      const response = await Axios.post(userUrl, JSON.stringify(credentials), {
        headers: { "Content-Type": "application/json" },
      });
      toast.success("User Registered Successfully", {
        position: toast.POSITION.TOP_RIGHT,
      });
    } catch (error) {
      sessionStorage.removeItem("token");
      console.log(error);
      toast.error(error.response.data.error, {
        position: toast.POSITION.TOP_RIGHT,
      });
    }

  }

  return (
    <>
    <PageLayout>
      {login ? (
        <>
          <Navigate to="/" replace={true} />
        </>
      ) : (
        <ContainerFormLogin>
         <ToastContainer />
          <FormLogin onSubmit={handleSubmit}>
            <InputProfile
              placeholder="Email"
              name="email"
              type="email"
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
              <ButtonLogin
                  type="submit"
                  onClick={(event) => handleSubmit(event, "login")}>
                  Login
                </ButtonLogin>
              </ContainerButton>
              <ContainerButton>
              <ButtonRegister
                  type="submit"
                  onClick={(event) => handleSubmit(event, "register")}>
                  Register
                </ButtonRegister>
              </ContainerButton>
            </ContainerAllButtons>
          </FormLogin>
        </ContainerFormLogin>
      )}
      </PageLayout>
    </>
  );
};

export default Login;