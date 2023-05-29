import React, { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import Axios from "axios";
import Footer from "../../components/Footer";
import Header from "../../components/Header";
import User from "../../components/User";
import { evaluateUrl, loginUrl, refreshTokenUrl } from "../../resources/constants.js";
import {
  Title
} from "./styles";

const Evaluate = () => {
  const token = sessionStorage.getItem("token");
  const isLoggedIn = token !== null;
  const navigate = useNavigate();

  useEffect(() => {
    RefreshToken();
  }, []);

  async function RefreshToken() {
    const token = sessionStorage.getItem("token");

    try {
      const response = await Axios.post(refreshTokenUrl, null, {
        headers: { Authorization: `Bearer ${token}` },
      });
      console.log("Token is valid!");
    } catch (error) {
      console.log("Token is expired or not existent!");
      sessionStorage.removeItem("token");
      navigate("/login");
      toast.error(error.response.data.message, {
        position: toast.POSITION.TOP_RIGHT,
      });
    }
  }

  return (
    <>
      <Header />
      <ToastContainer />
      <Footer />
    </>
  );
};

export default Evaluate;
