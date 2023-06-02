import React, { useRef, useState, useEffect} from "react";
import Footer from "../../components/Footer";
import Header from "../../components/Header";
import Axios from "axios";
import { useNavigate } from "react-router-dom";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import { evaluateUrl, loginUrl, refreshTokenUrl } from "../../resources/constants.js";

import {
  ContainerSubmission,
  ContainerInfosSubmission,
  InputChangeImage,
  ContainerInputs,
  InputEditSubmissionDescription,
  InputEditSubmission,
  ContainerButtonAdd,
  ButtonAddSubmission,
} from "./styles";
import { submissionUrl } from "../../resources/constants.js";


const AddSubmission = () => {
  const navigate = useNavigate();

  const body_part = useRef();
  const description = useRef();
  const date = useRef();
  const media = useRef();

  const [selectedDate, setSelectedDate] = useState(null);

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

  async function addSubmission(e) {
    e.preventDefault();
    console.log("add submission");

    const formData = new FormData();
    formData.append("body_part", body_part.current.value);
    formData.append("description", description.current.value);
    formData.append("media", media.current.files[0]);


   /* const newSubmission = {
      body_part: body_part.current.value,
      description: description.current.value,
  
      media: media.current.value,
      //media_type: ""
    };

    if (body_part.current.value !== "") {
      newSubmission.body_part = "Unknown";
    }*/

    const file = media.current.files[0];
    var media_type = "";

    if (file) {
      // Check the file type
    if (file.type.startsWith("image/")) {
      media_type = "image";
     
    } else if (file.type.startsWith("video/")) {
      //newSubmission.media_type = "video";
      media_type = "video";
    } else {
      toast.error("Invalid file type. Discarding file.", {
        position: toast.POSITION.TOP_RIGHT,
      });
      return;
    }

      formData.append("media_type", media_type);
      formData.append("file", file);
      formData.append("date", selectedDate); 

      //newSubmission.file = formData;
    }


    const url = submissionUrl;
    const token = sessionStorage.getItem("token");

    try {
      const response = await Axios.post(url, formData, {
        headers: { Authorization: token },
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
      <ContainerSubmission>
        <ContainerInfosSubmission onSubmit={addSubmission}>
          <ContainerInputs>
            <InputEditSubmissionDescription
              name="descriptionSubmission"
              placeholder="Submission description"
              type="text"
              id="add_submission_description"
              required
              ref={description}
            />
          <DatePicker
              selected={selectedDate}
              onChange={(date) => setSelectedDate(date)}
              placeholderText="Submission date"
              required
            />
            {/* Add a file input for uploading images or videos */}
            <input
              type="file"
              name="file"
              accept="image/*, video/*"
              ref={media}
            />
            <select
              name="bodyPartSubmission"
              id="add_submission_body_part"
              required
              ref={body_part}
            >
              <option value="">Select Body Part</option>
              <option value="Face">Face</option>
              <option value="Head">Head</option>
              <option value="Hand">Hand</option>
              <option value="Feet">Feet</option>
              <option value="Leg">Leg</option>
              <option value="Body">Body</option>
              <option value="Chest">Chest</option>
              <option value="Back">Back</option>
              <option value="Arm">Arm</option>
              <option value="Belly">Belly</option>
            </select>
          </ContainerInputs>
          <ContainerButtonAdd>
            <ButtonAddSubmission>Add</ButtonAddSubmission>
          </ContainerButtonAdd>
        </ContainerInfosSubmission>
      </ContainerSubmission>
      <Footer />
    </>
  );
};

export default AddSubmission;
