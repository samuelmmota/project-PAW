import React, { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import Axios from "axios";
import Footer from "../../components/Footer";
import Header from "../../components/Header";
import User from "../../components/User";
import styled from "styled-components";

import PageLayout from "../../components/PageLayout";
import { evaluateUrl, loginUrl, refreshTokenUrl } from "../../resources/constants.js";
import { Title, Form } from "./styles";
import Tabs from '@mui/material/Tabs';
import Tab from '@mui/material/Tab';

const Evaluate = () => {
  const token = sessionStorage.getItem("token");
  const isLoggedIn = token !== null;
  const navigate = useNavigate();
  const [currentTab, setCurrentTab] = React.useState(0);
  const [patients, setPatients] = React.useState([]);

  useEffect(() => {
    RefreshToken();
    fetchPatients(); // Fetch the patients data
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

  async function fetchPatients() {
    // Fetch the patients data from the API
    try {
      const response = await Axios.get("your_api_endpoint_here");
      setPatients(response.data);
    } catch (error) {
      console.log("Error fetching patients:", error);
    }
  }

  const handleTabChange = (event, newValue) => {
    setCurrentTab(newValue);
  };

  return (
    <>
     <PageLayout>
      <ToastContainer />
      <Tabs value={currentTab} onChange={handleTabChange}>
        <Tab label="Clinicals" />
        <Tab label="Patients" />
      </Tabs>
      {currentTab === 0 && (
        <div>
          {/* Render clinicals component */}
        </div>
      )}
      {currentTab === 1 && (
        <div>
          {patients.map((patient) => (
            <div key={patient.id}>
              <h3>{patient.email}</h3>
              {/* Render patient's information */}
            </div>
          ))}
        </div>
      )}
    </PageLayout>
    </>
  );
};

export default Evaluate;
