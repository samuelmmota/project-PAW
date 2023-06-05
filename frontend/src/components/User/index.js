import React from "react";
import { useNavigate } from "react-router-dom";
import {
  ProfileSection,
  Container,
  ProfileImage,
  Name,
  Email,
  ButtonContainer,
  Button
} from "./styles";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { userUrl } from "../../resources/constants.js";
import Axios from "axios";
const User = ({ userName, userEmail, isClinical, isExportToResearchSet }) => {

  // variavel usada pra fazer a navegação pelas paginas
  const navigate = useNavigate();
  const token = sessionStorage.getItem("token");
  const isLoggedIn = token !== null;
  const decodedToken = JSON.parse(atob(token.split(".")[1]));
  const userID = decodedToken.user_id;

  function deleteUser(element) {
    console.log("delete user");
    console.log(element.target);
  }
  function editUser(element) {
    console.log("edit user button");
    console.log(element.target);
    navigate("/edituser");
  }

  async function imaDoctor(element) {
    const url = userUrl + userID;
    try {
      const response = await Axios.put(
        url,
        {
          isClinical: "true",
        },
        {
          headers: {
            "Content-Type": "application/json",
            Authorization: token,
          },
        }
      );
      window.location.reload();
    } catch (error) {
      toast.error(error.response?.data?.message, {
        position: toast.POSITION.TOP_RIGHT,
      });
    }
  }


  async function allowToExportData(element) {
    const url = userUrl + userID;
    try {
      const response = await Axios.put(
        url,
        {
          exportToReasearcher: "true",
        },
        {
          headers: {
            "Content-Type": "application/json",
            Authorization: token,
          },
        }
      );
    } catch (error) {
      toast.error(error.response?.data?.message, {
        position: toast.POSITION.TOP_RIGHT,
      });
    }

  
    element.target.remove();
  }

  return (
    <>
      <ToastContainer />
      <Container>
        <ProfileSection>
          <ProfileImage src="https://thumbs.dreamstime.com/b/default-avatar-profile-vector-user-profile-default-avatar-profile-vector-user-profile-profile-179376714.jpg" alt="User profile Image" />
          {isClinical && (
            <>
              <Name> User is a doctor!</Name>
            </>
          )}
          <Email>{userEmail}</Email>
        </ProfileSection>
        <center>
          {isLoggedIn && (
            <ButtonContainer>
              <Button type="button" onClick={editUser}>
                Update
              </Button>
              <Button type="button" onClick={deleteUser}>
                Delete
              </Button>
              {!isClinical && (
                <>
                  <Button type="button" onClick={imaDoctor}>
                    I'm a Doctor!
                  </Button>
                  {!isExportToResearchSet && (
                    <>
                      <Button type="button" onClick={allowToExportData}>
                        Allow to export data to research
                      </Button>
                    </>)}
                </>
              )}
            </ButtonContainer>
          )}
        </center>
      </Container>

    </>
  );
};

export default User;
