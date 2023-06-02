import React, { useEffect, useState } from "react";
import Footer from "../../components/Footer";
import Header from "../../components/Header";
import Submission from "../../components/Submission";
import Axios from "axios";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { ContainerSubmissions, PageContainer } from "./styles";
import { submissionUrl } from "../../resources/constants.js";
import { useNavigate } from "react-router-dom";
import { evaluateUrl, loginUrl, refreshTokenUrl } from "../../resources/constants.js";

const Gallery = () => {
  const [media, setMedia] = useState([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    fetchMedia();
  }, []);

  const fetchMedia = async () => {
    try {
      const response = await Axios.get(submissionUrl);
      setMedia(response.data);
    } catch (error) {
      if(!isLoading)
        toast.error("Failed to fetch media");
    } finally {
      setIsLoading(false);
    }
  };

  const [submissions, setSubmissions] = useState([]);
  const token = sessionStorage.getItem("token");
  const isLoggedIn = token !== null;
  const navigate = useNavigate();

  useEffect(() => {
    RefreshToken();
    getSubmissions();
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
      toast.error(error.response?.data?.message || "Failed to refresh token", {
        position: toast.POSITION.TOP_RIGHT,
      });
    }
  }

  // Não usa token de auth
  async function getSubmissions() {
    const url = submissionUrl;
    try {
      const response = await Axios.get(url, {
        headers: { "Content-Type": "application/json", Authorization: `Bearer ${token}`},
      });
      console.log(response.data.submissions);
      setSubmissions(response.data.submissions);
    } catch (error) {
      console.log(error);
    }
  }

  if (!isLoading && !submissions) {
    return (
      <>
      <Header />
      <ToastContainer />
      <Footer />
      </>
    )
  }

  return (
    <>
    <PageContainer>
      <Header />
      <ToastContainer />
      <ContainerSubmissions>
        {submissions.length > 0 ? (
          submissions.map((submission, index) => (
            <Submission
              image={submission.image}
              body_part={submission.body_part}
              media_type={submission.media_type}
              media={submission.media}
              date={submission.date}
              description={submission.description}
              id={submission.id}
              refreshSubmissions={getSubmissions}
              key={index}
            />
          ))
        ) : (
          <p>No submissions found.</p>
        )}
      </ContainerSubmissions>
      <Footer />
      </PageContainer>
    </>
  );
};

export default Gallery;
