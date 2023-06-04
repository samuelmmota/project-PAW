import React, { useEffect, useState } from "react";
import Axios from "axios";
import { useParams } from "react-router-dom";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import PageLayout from "../../components/PageLayout";
import styled from "styled-components";
import { submissionUrl } from "../../resources/constants.js";
import {
    ContainerSubmission, TitleSubmission
    , ContainerImage, ImageSubmission, VideoSubmission, Title
} from "./../../components/Submission/styles.js";

import { useNavigate } from "react-router-dom";
import { sub } from "date-fns";
const Description = styled.p`
  margin-bottom: 10px;
`;

const EvaluateSubmission = ({ submissionId }) => {
    const navigate = useNavigate();

    console.log("submissionId", submissionId);

    if (submissionId === undefined || submissionId === null) {
        navigate(`/`);
    }

    const [comment, setComment] = useState("");
    const [submission, setSubmission] = useState({});

    useEffect(() => {
        fetchSubmission();
    }, []);


async function fetchSubmission() {
    const token = sessionStorage.getItem("token");
    const isLoggedIn = token !== null;

    const url = submissionUrl + submissionId ;
    try {
      const response = await Axios.get(url, {
        headers: { "Content-Type": "application/json", Authorization: `Bearer ${token}`},
      });
      console.log(response.data.submissions);
      setSubmission(response.data.submissions);
    } catch (error) {
      console.log(error);
    }
  }

    const handleCommentChange = (event) => {
        setComment(event.target.value);
    };

    const handleSubmitComment = async (event) => {
        event.preventDefault();

        try {
            const response = await Axios.post(`/api/submissions/${submissionId}/comments`, {
                comment: comment,
            });
            const newComment = response.data;
            /*setSubmission((prevSubmission) => ({
              ...prevSubmission,
              comments: [...prevSubmission.comments, newComment],
            }));*/
            setComment("");
            toast.success("Comment added successfully", {
                position: toast.POSITION.TOP_RIGHT,
            });
        } catch (error) {
            console.error("Error adding comment:", error);
            toast.error("Failed to add comment", {
                position: toast.POSITION.TOP_RIGHT,
            });
        }
    };



    return (
        <PageLayout>
            <ToastContainer />
            <div>
                <h1>Evaluate Submission</h1>
                media_type
                {
/*
 <ContainerSubmission>

                    <TitleSubmission>{body_part}</TitleSubmission>
                    <ContainerImage>
                        {media_type === "image" && (
                            <ImageSubmission
                                key={id}
                                src={`data:image/jpeg;base64,${media}`}
                                alt="Image"
                            />
                        )}
                        {media_type === "video" && (
                            <VideoSubmission
                                url={`data:video/mp4;base64,${media}`}
                                controls
                            />
                        )}
                    </ContainerImage>
                    <Description>{description}</Description>
                </ContainerSubmission>
*/
                }
               
                <h2>Comments</h2>

                <ul>

                    <li key={comment.id}>
                        <p>Comment ID: {comment.id}</p>
                        <p>Author: {comment.author}</p>
                        <p>Message: {comment.message}</p>
                    </li>

                </ul>

                <h2>Add Comment</h2>
                <form onSubmit={handleSubmitComment}>
                    <textarea value={comment} onChange={handleCommentChange} />
                    <button type="submit">Add Comment</button>
                </form>



            </div>
        </PageLayout>
    );
};

export default EvaluateSubmission;
