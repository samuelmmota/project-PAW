import React, { useEffect, useState } from "react";

import User from "../../components/User";
import Axios from "axios";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { useNavigate } from "react-router-dom";
import { userUrl, clinicalUrl, refreshTokenUrl } from "../../resources/constants.js";
import { Form, ClinicalList, ClinicalItem, RemoveButton, PageContainer } from "./styles";
import PageLayout from "../../components/PageLayout";
const Profile = () => {
  const navigate = useNavigate();
  const token = sessionStorage.getItem("token");
  const isLoggedIn = token !== null;

  const [user, setUser] = useState({
  });

  const [clinical, setClinical] = useState("");

  useEffect(() => {
    RefreshToken();
    getMyUser();
    getClinicals();
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


  async function getMyUser() {
  const decodedToken = JSON.parse(atob(token.split(".")[1]));
  const userID = decodedToken.user_id;
    const url = userUrl + userID;
    try {
      const response = await Axios.get(url, {
        headers: { "Content-Type": "application/json", Authorization: token },
      });
      const userData = response.data.user;
      setUser({
        user_email: userData.email,
        user_is_clinical: userData.isClinical,
      });
      console.log(response);
    } catch (err) {
      console.log(err);
      toast.error(err.response.data.message, {
        position: toast.POSITION.TOP_RIGHT,
      });
      if (err.response.data.message === "token is not valid") {
        navigate("/login");
      }
    }
  }

  async function getClinicals() {
    const decodedToken = JSON.parse(atob(token.split(".")[1]));
  const userID = decodedToken.user_id;
    const url = clinicalUrl + userID;
    try {
      const response = await Axios.get(url, {
        headers: { "Content-Type": "application/json", Authorization: token },
      });
      setUser((prevState) => ({
        ...prevState,
        clinicals: response.data.clinicals,
      }));
    } catch (err) {
      console.log(err);
      toast.error(err.response.data.message, {
        position: toast.POSITION.TOP_RIGHT,
      });
    }
  }

  async function addClinical(e) {
    e.preventDefault();
    const decodedToken = JSON.parse(atob(token.split(".")[1]));
  const userID = decodedToken.user_id;
    const url = clinicalUrl + userID;
    try {
      const response = await Axios.post(
        url,
        {
          clinical_email: clinical,
        },
        {
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
          },
        }
      );
      const updatedClinicals = [...user.clinicals, response.data.clinical];
      setUser((prevState) => ({
        ...prevState,
        clinicals: updatedClinicals,
      }));
      toast.success("Clinical added successfully", {
        position: toast.POSITION.TOP_RIGHT,
      });
      setClinical(""); // Clear the input field after adding clinical
      window.location.reload(); // Refresh the page
    } catch (err) {
      console.log(err);
      toast.error(err.response.data.message, {
        position: toast.POSITION.TOP_RIGHT,
      });
    }
  }
  

  async function removeClinical(clinicalID) {
    const url = userUrl + clinicalID;
    try {
      await Axios.delete(url, {
        headers: { "Content-Type": "application/json", Authorization: token },
      });
      const updatedClinicals = user.clinicals.filter(
        (clinical) => clinical.id !== clinicalID
      );
      setUser((prevState) => ({
        ...prevState,
        clinicals: updatedClinicals,
      }));
      toast.success("Clinical removed successfully", {
        position: toast.POSITION.TOP_RIGHT,
      });
    } catch (err) {
      console.log(err);
      toast.error(err.response.data.message, {
        position: toast.POSITION.TOP_RIGHT,
      });
    }
  }

  return (
    <>
<PageLayout>
      <ToastContainer />
      {isLoggedIn && (
        <User userName={user.user_name} userEmail={user.user_email} isClinical= {user.user_is_clinical} />
      )}
       {!user.user_is_clinical && (
              <> 
              <Form onSubmit={addClinical}>
              <label>Add Clinicals to Evaluate your submissions:</label>
              <input
                type="text"
                name="clinical"
                value={clinical}
                onChange={(e) => setClinical(e.target.value)}
              />
              <button type="submit">Add</button>
            </Form>
            <ClinicalList>
        {user.clinicals && user.clinicals.length !== 0 && (
          user.clinicals.map((clinical) => (
            <ClinicalItem key={clinical.clinical.id}>
              {clinical.clinical.email}
              <RemoveButton onClick={() => removeClinical(clinical.clinical.id)}>
                X
              </RemoveButton>
            </ClinicalItem>
          ))
        )}
      </ClinicalList>
      </>
              )}

</PageLayout>
    </>
  );
};

export default Profile;
