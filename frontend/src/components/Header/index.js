import React from "react";
import {
  ContainerLinks,
  Logo,
  ProfileIcon,
  Row,
  ContainerMenu,
  LinkHome,
  ButtonLogout,
  ButtonLogin,
  Button,
} from "./styles";
import { useNavigate } from "react-router-dom";
import profile from "../../assets/profile.svg";
import home from "../../assets/home.svg";

const Header = () => {
  const token = sessionStorage.getItem("token");
  const isLoggedIn = token !== null;
  // variavel usada pra fazer a navegação pelas paginas
  const navigate = useNavigate();
  
  function logout() {
    sessionStorage.removeItem("token");
    console.log("logout");
    navigate("/login");
  }

  function login() {
    console.log("login");
    navigate("/login");
  }

  function register() {
    console.log("register");
    navigate("/register");
  }

  function gallery() {
    console.log("gallery");
    navigate("/gallery");
  }

  function evaluate() {
    console.log("evaluate");
    navigate("/evaluate");
  }
  /*
   onClick={gallery}>Gallery  quando se clica chama a função de logout
  é semelhante ao fazer a navegação diretamente:
         <Button onClick={() => navigate("/addbook")}>Add Book</Button>
  */
  return (
    <>
      <header>
        <ContainerLinks>
          <LinkHome title="Home Page" to="/">
            <Logo src={home} alt="logo" />
          </LinkHome>
          <ContainerMenu>
            {isLoggedIn ? (
              <>
                <ButtonLogout onClick={evaluate}>Evaluate</ButtonLogout>
                <Button onClick={() => navigate("/addsubmission")}>Add Submission</Button>
                <ButtonLogin onClick={gallery}>Gallery</ButtonLogin>
                <ButtonLogout onClick={logout}>Logout</ButtonLogout>
                <a title="Your Profile" href="/profile">
                  <ProfileIcon src={profile} alt="Profile icon" />
                </a>
              </>
            ) : (
              <>
                <ButtonLogin onClick={login}>Login</ButtonLogin>
              </>
            )}
          </ContainerMenu>
        </ContainerLinks>
      </header>
      <Row />
    </>
  );
};

export default Header;
