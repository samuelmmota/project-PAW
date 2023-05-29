import React, { useEffect, useState } from "react";
import Footer from "../../components/Footer";
import Header from "../../components/Header";
import User from "../../components/User";
import Axios from "axios";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { useNavigate } from "react-router-dom";
import { userUrl } from "../../resources/constants.js";

const Profile = () => {
  const navigate = useNavigate();
  const token = sessionStorage.getItem("token");
  const isLoggedIn = token !== null;
  const decodedToken = JSON.parse(atob(token.split(".")[1]));
  const userID = decodedToken.user_id;

  const [user, setUser] = useState({
    user_name: "",
    user_email: "",
  });

  useEffect(() => {
    getMyUser();
  }, []);

  async function getMyUser() {
    const url = userUrl + userID;
    try {
      const response = await Axios.get(url, {
        headers: { "Content-Type": "application/json", Authorization: token },
      });
      const userData = response.data.user;
      setUser({
        user_name: userData.name,
        user_email: userData.email,
      });
    } catch (err) {
      console.log(err);
      toast.error(err.response.data.message, {
        position: toast.POSITION.TOP_RIGHT,
      });
    if (err.response.data.message == "token is not valid") {
       navigate("/login");
      }
    }
  }

  return (
    <>
      <Header />
      <ToastContainer />
      {isLoggedIn && (
        <User
          userName={user.user_name}
          userEmail={user.user_email}
        />
      )}
      <Footer />
    </>
  );
};

export default Profile;
