import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import Axios from "axios";
import Footer from "../../components/Footer";
import Header from "../../components/Header";
import User from "../../components/User";
import styled from "styled-components";
import { FilterItem, FilterLabel,FilterContainer,ContainerSubmissions} from "./../Gallery/styles";

import PageLayout from "../../components/PageLayout";
import { evaluateUrl, loginUrl, refreshTokenUrl, patientUrl } from "../../resources/constants.js";
import { Title, Form } from "./styles";
import Tabs from "@mui/material/Tabs";
import Tab from "@mui/material/Tab";
import Submission from "../../components/Submission";

const Evaluate = () => {
  const token = sessionStorage.getItem("token");
  const isLoggedIn = token !== null;
  const navigate = useNavigate();
  const [currentTab, setCurrentTab] = React.useState(0);
  const [patients, setPatients] = React.useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [selectedBodyPart, setSelectedBodyPart] = useState("");
  const [descriptionFilter, setDescriptionFilter] = useState("");
  const [startDate, setStartDate] = useState("");
  const [endDate, setEndDate] = useState("");

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
    console.log("fetch Patietns");
    const decodedToken = JSON.parse(atob(token.split(".")[1]));
    const userID = decodedToken.user_id;
    const url = patientUrl + userID;

    try {
      const response = await Axios.get(url, {
        headers: { "Content-Type": "application/json", Authorization: token },
      });
        setPatients(response.data.clinicals);
        console.log(response.data.clinicals);
      console.log(response);
    } catch (error) {
      console.log("Error fetching patients:", error);
    }
  }

  const handleTabChange = (event, newValue) => {
    setCurrentTab(newValue);
  };

  return (
    <PageLayout>
      <ToastContainer />
      <Tabs value={currentTab} onChange={handleTabChange}>
        {patients.map((patient, index) => (
          <Tab key={index} label={patient.patient_email} />
        ))}
      </Tabs>
      {patients.map((patient, index) => (
        <div key={index} hidden={currentTab !== index}>
          <h3>{patient.email}</h3>
          {patient.submission.map((submission, index) => (
      <ContainerSubmissions>
            <Submission
              image={submission.image}
              body_part={submission.body_part}
              media_type={submission.media_type}
              media={submission.media}
              date={submission.date}
              description={submission.description}
              id={submission.id}
              isClinicalViewing={true}
              key={index}
            />
      </ContainerSubmissions>
          ))}
        </div>
      ))}
    </PageLayout>
  );
};

export default Evaluate;
