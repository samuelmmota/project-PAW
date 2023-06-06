import React, { useEffect, useState } from "react";
import Submission from "../../components/Submission";
import Axios from "axios";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { ContainerSubmissions, FilterContainer, FilterItem, FilterLabel } from "./styles";
import { submissionUrl } from "../../resources/constants.js";
import { useNavigate } from "react-router-dom";
import { refreshTokenUrl } from "../../resources/constants.js";
import PageLayout from "../../components/PageLayout";

const Gallery = () => {
  const [media, setMedia] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [selectedBodyPart, setSelectedBodyPart] = useState("");
  const [descriptionFilter, setDescriptionFilter] = useState("");
  const [startDate, setStartDate] = useState("");
  const [endDate, setEndDate] = useState("");



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
      <PageLayout>
      <ToastContainer />
      </PageLayout>
      </>
    )
  }

  return (
    <>
    <PageLayout>
      <ToastContainer />
        <FilterContainer>
          <FilterItem>
            <FilterLabel>Body Part:</FilterLabel>{
          <select
            value={selectedBodyPart}
            onChange={(event) => setSelectedBodyPart(event.target.value)}
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
          }
          </FilterItem>
          <FilterItem>
    <FilterLabel>Description:</FilterLabel>
    {
          <input
            type="text"
            value={descriptionFilter}
            onChange={(event) => setDescriptionFilter(event.target.value)}
            placeholder="Filter by description"
          />}
          </FilterItem>
          <FilterItem>
    <FilterLabel>Start Date:</FilterLabel>
    {
          <input
            type="date"
            value={startDate}
            onChange={(event) => setStartDate(event.target.value)}
          />}
          </FilterItem>
          <FilterItem>
    <FilterLabel>End Date:</FilterLabel>
    {
          <input
            type="date"
            value={endDate}
            onChange={(event) => setEndDate(event.target.value)}
          />}
          </FilterItem>
        </FilterContainer>  
      <ContainerSubmissions>
        {submissions.length > 0 ? (
          submissions
          .filter((submission) => {
            // Filtrar pelo body_part
            if (selectedBodyPart !== "" && submission.body_part !== selectedBodyPart) {
              return false;
            }
            // Filtrar pela descrição
            if (
              descriptionFilter !== "" &&
              !submission.description.toLowerCase().includes(descriptionFilter.toLowerCase())
            ) {
              return false;
            }
            // Filtrar pela data de início
            if (startDate !== "" && submission.date <= startDate) {
              return false;
            }
            // Filtrar pela data de fim
            if (endDate !== "" && submission.date >= endDate) {
              return false;
            }
            return true;
          })
          .map((submission, index) => (
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
      </PageLayout>
    </>
  );
};

export default Gallery;
