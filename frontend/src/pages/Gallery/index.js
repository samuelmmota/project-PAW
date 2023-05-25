import React, { useEffect, useState } from "react";
import Footer from "../../components/Footer";
import Header from "../../components/Header";
import Submission from "../../components/Submission";
import Axios from "axios";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { ContainerSubmissions } from "./styles";
import { submissionUrl} from "../../resources/constants.js";


const Gallery = () => {
const [media, setMedia] = useState([]);

useEffect(() => {
  fetchMedia();
}, []);

const fetchMedia = async () => {
  try {
    const response = await Axios.get(submissionUrl);
    setMedia(response.data);
  } catch (error) {
    toast.error("Failed to fetch media");
  }
};

  const [submissions, setSubmissions] = useState([]);
  const token = sessionStorage.getItem("token");
  const isLoggedIn = token !== null;


  useEffect(() => {
    getSubmissions();
  }, []);

  // Não usa token de auth
  async function getSubmissions() {
    const url = submissionUrl;
    try {
      const response = await Axios.get(url, {
        headers: { "Content-Type": "application/json" },
      });
      console.log(response.data.submissions);
      setSubmissions(response.data.submissions);
    } catch (error) {
      console.log(error);
    }
  }

  /*async function isOwner(id) {
    const url = submissionUrl + id;
    try {
      const response = await Axios.get(url, {
        headers: {
          Authorization: "Bearer " + token,
        },
      });
      console.log(response);
      return true
    } catch (error) {
      return false
      //console.log(error);
    }
  }*/

if (!Array.isArray(submissions)) {
  toast.error("No Submissions found", {
    position: toast.POSITION.TOP_RIGHT,
  });
};
  return (
    <>
      <Header />
        <ToastContainer />
      <ContainerSubmissions>
        {submissions.map((submissions, index) => (
          <Submission
            image={submissions.image}
            body_part={submissions.body_part}
            media_type={submissions.media_type}
            media= {submissions.media}
            date={submissions.date}
            description={submissions.description}
            id={submissions.id}
            refreshSubmissions={getSubmissions}
            key={index}
            owner={{}/*isOwner() TODO: fix this*/}
          />
        ))}
      </ContainerSubmissions>
      <Footer />
    </>
  );
};
export default Gallery;