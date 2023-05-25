import React, { useRef } from "react";
import { useParams } from "react-router-dom";
import Footer from "../../components/Footer";
import Header from "../../components/Header";
import Axios from "axios";
import { useNavigate } from "react-router-dom";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { submissionUrl } from "../../resources/constants.js";

import {
  ContainerSubmission,
  ContainerInfosSubmission,
  InputChangeImage,
  ContainerInputs,
  InputEditSubmissionDescription,
  InputEditSubmission,
  ContainerButtonAdd,
  ButtonAddSubmission,
  Title,
} from "./styles";

const EditSubmission = () => {
  const { id } = useParams();
  const navigate = useNavigate();

  const title = useRef();
  const description = useRef();
  const year = useRef();
  const submissionCover = useRef();

 async function editSubmission(e) {
    //previne reload à pagina
    e.preventDefault();
    console.log("edit Submission");

    const editedSubmission = {
      title: title.current.value,
      description: description.current.value,
      year: year.current.value,
    };

    const url = submissionUrl + id;
    const token = sessionStorage.getItem("token");

    try {
     const response = await Axios.put(url, JSON.stringify(editedSubmission), {
        headers: { "Content-Type": "application/json", Authorization: token },
      });

      console.log(response);
      navigate("/gallery");
    } catch (error) {
      console.log(error);
      toast.error(error.response.data.message, {
        position: toast.POSITION.TOP_RIGHT,
      });
    }
  }

  return (
    <>
      <Header />
      <ToastContainer />
      <Title>Edit Submission {id}</Title>
      <ContainerSubmission>
        <ContainerInfosSubmission onSubmit={editSubmission}>
          <ContainerInputs>
            <InputChangeImage
              name="imageSubmission"
              type="text"
              placeholder="Submission cover url"
              id="add_submission_cover"
              ref={submissionCover}
            />
            <InputEditSubmission
              name="titleSubmission"
              placeholder="Submission title"
              type="text"
              id="add_submission_title"
              ref={title}
            />
            <InputEditSubmissionDescription
              name="descriptionSubmission"
              placeholder="Submission description"
              type="text"
              id="add_submission_description"
              ref={description}
            />
            <InputEditSubmission
              name="yearSubmission"
              placeholder="Submission year"
              type="number"
              id="add_submission_number"
              ref={year}
            />
          </ContainerInputs>
          <ContainerButtonAdd>
            <ButtonAddSubmission>Edit</ButtonAddSubmission>
          </ContainerButtonAdd>
        </ContainerInfosSubmission>
      </ContainerSubmission>
      <Footer />
    </>
  );
};

export default EditSubmission;
