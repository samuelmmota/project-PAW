import React from "react";
import { useNavigate } from "react-router-dom";
import {
  ContainerSubmission,
  ContainerImage,
  ImageSubmission,
  TitleSubmission,
  Button,
  ButtonContainer,
  DateText,
  VideoSubmission,
  ButtonFeedback,
} from "./styles";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import Axios from "axios";
import { submissionUrl } from "../../resources/constants.js";
import styled from "styled-components";
import { primaryColor } from "../../resources/constants.js";
import { element } from "prop-types";
import Modal from "react-modal";
import { useState } from "react";

const Submission = ({ body_part, media, media_type, date, id, description, refreshSubmissions, isClinicalViewing }) => {
  const token = sessionStorage.getItem("token");
  const isLoggedIn = token !== null;
  // variavel usada pra fazer a navegação pelas paginas
  const navigate = useNavigate();
  const [isModalOpen, setIsModalOpen] = useState(false);

  // Modal configuration
  Modal.setAppElement("#root");
  const customModalStyles = {
    content: {
      maxWidth: "100%",
      maxHeight: "100%",
      margin: "auto",
      display: "flex",
      alignItems: "center",
      justifyContent: "center",
    },
    overlay: {
      backgroundColor: "rgba(0, 0, 0, 0.8)",
    },
  };

  let displayDate = date;
  let displayMonth = "";
  let displayDay = "";
  let displayYear = "";

  const openModal = () => {


    setIsModalOpen(true);

  };

  const closeModal = () => {
    setIsModalOpen(false);
  };

  if (date !== undefined && date !== null && date !== "") {
    const regex = /(?<=\b)\w+\s(\w+)\s(\d+)\s(\d+)\b/;
    const matchResult = date.match(regex);

    if (matchResult !== null) {
      const [, month, day, year] = matchResult;
      displayMonth = month;
      displayDay = day;
      displayYear = year;
    } else {
      console.log("Invalid date format");
    }
  } else {
    console.log("Invalid date");
  }

  async function deleteSubmission(element) {
    console.log("delete Submission" + id);
    // console.log(element.target);
    // element.target.parentNode.parentNode.remove();

    const url = submissionUrl + id;
    try {
      await Axios.delete(url, {
        headers: {
          Authorization: "Bearer " + token,
        },
      });
      refreshSubmissions(" from delete");
    } catch (err) {
      console.log(err);
      toast.error(err.response.data.message, {
        position: toast.POSITION.TOP_RIGHT,
      });
    }
  }

  const handleButtonClick = () => {

  };

  /*
  First we perform a get to know if the user is the ownner of the Submission and have the right to edit it
  */
  async function editSubmission(element) {
    console.log("edit Submission button");

    const url = submissionUrl + id;
    try {
      await Axios.get(url, {
        headers: {
          Authorization: "Bearer " + token,
        },
      });
      console.log(element.target);
      navigate(`/editsubmission/${id}`);
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

  function evaluateSubmission(element) {
    navigate(`/evaluatesubmission/${id}`);
  }

  function viewSubmission(element) {
    navigate(`/viewsubmission/${id}`);
  }

  return (
    <ContainerSubmission>
      <DateText>
        {displayMonth !== "" && displayDay !== "" && displayYear !== "" ? (
          <><b>Date: </b> {displayMonth}, {displayDay}, {displayYear}</>
        ) : (
          <><b>Date: </b> {displayDate}</>
        )}
      </DateText>
      <ContainerImage>
        {media_type === "image" && (
          <>
            <ImageSubmission
              key={id}
              src={`data:image/jpeg;base64,${media}`}
              alt="Image"
              onClick={openModal} // Open the modal on image click
            />
            <Modal
              isOpen={isModalOpen}
              onRequestClose={closeModal}
              style={customModalStyles}
            >
              <ImageSubmission
                key={id}
                src={`data:image/jpeg;base64,${media}`}
                alt="Large Image"
              />
              <Button onClick={closeModal}>Close</Button>
            </Modal>
          </>
        )}
        {media_type === "video" && (
          <VideoSubmission
            url={`data:video/mp4;base64,${media}`}
            controls
          />
        )}
      </ContainerImage>
      <TitleSubmission><b>Description: </b>{description}</TitleSubmission>
      <TitleSubmission><b>Body Part: </b>{body_part}</TitleSubmission>
      <center>
        <ButtonContainer>
          {isClinicalViewing == null || isClinicalViewing == false ? (
            <>
              <Button onClick={editSubmission}>Update</Button>
              <Button onClick={deleteSubmission}>Delete</Button>
              <ButtonFeedback onClick={viewSubmission}> View Feedbacks</ButtonFeedback>
            </>
          ) : (
            <Button onClick={evaluateSubmission}>Evaluate</Button>
          )}

        </ButtonContainer>
      </center>
    </ContainerSubmission>
  );
};

export default Submission;