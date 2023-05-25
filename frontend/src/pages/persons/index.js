import React, { useEffect, useState } from "react";
import Footer from "../../components/Footer";
import Header from "../../components/Header";
import User from "../../components/User";
import Axios from "axios";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { useNavigate } from "react-router-dom";

import { 
  ContainerUser,
  Title
 } from "./styles";

 const Persons = () => {

  const [users, setUsers] = useState([]);
  const token = sessionStorage.getItem("token");
  const isLoggedIn = token !== null;

  useEffect(() => {
    getUsers();
  }, []);

 // Não usa token de auth
 async function getUsers() {
  const url = "http://localhost:3000/api/v1/user/";
  try {
    const response = await Axios.get(url, {
      headers: { "Content-Type": "application/json" },
    });
    console.log(response.data.users);
    setUsers(response.data.users);
  } catch (error) {
    console.log(error);
  }
}
  return (
    <>
      <Header />
      <ToastContainer />
      <ContainerUser>
        {users.map((user, index) => (
          <User
          userImage={user.profile_picture}
          userName={user.name}
          userEmail={user.email}
          />
        ))}
      </ContainerUser>
      <Footer />
    </>
  );
};

export default Persons;
