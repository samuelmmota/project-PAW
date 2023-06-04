import React, { useEffect, useState } from "react";
import Axios from "axios";
import { useParams, useNavigate } from "react-router-dom";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import PageLayout from "../../components/PageLayout";
import { Container, Typography, TextField, Button, List, ListItem, ListItemText, Divider } from "@material-ui/core";
import { submissionUrl, messageUrl } from "../../resources/constants.js";
import { ContainerSubmission, TitleSubmission, ContainerImage, ImageSubmission, VideoSubmission } from "./../../components/Submission/styles.js";

const EvaluateSubmission = () => {
    const navigate = useNavigate();
    const { id } = useParams();

    const [submission, setSubmission] = useState({});
    const [messages, setMessages] = useState([]);
    const [comment, setComment] = useState("");

    useEffect(() => {
        fetchSubmission();
        fetchMessages();
    }, []);

    async function fetchSubmission() {
        const token = sessionStorage.getItem("token");

        try {
            const response = await Axios.get(submissionUrl + id, {
                headers: { Authorization: `Bearer ${token}` },
            });
            setSubmission(response.data.submission);
        } catch (error) {
            console.log(error);
            toast.error(error.response.data.message, {
                position: toast.POSITION.TOP_RIGHT,
            });
        }
    }

    async function fetchMessages() {
        const token = sessionStorage.getItem("token");

        try {
            const response = await Axios.get(messageUrl + id, {
                headers: { Authorization: `Bearer ${token}` },
            });
            setMessages(response.data.messages);
        } catch (error) {
            console.log(error);
            toast.error(error.message, {
                position: toast.POSITION.TOP_RIGHT,
            });
        }
    }

    const handleMessageChange = (event) => {
        setComment(event.target.value);
    };

    const handleSubmitMessage = async (event) => {
        event.preventDefault();

        const token = sessionStorage.getItem("token");
        const decodedToken = JSON.parse(atob(token.split(".")[1]));
        const userID = decodedToken.user_id;
        const url = messageUrl;

        const body = {
            date: new Date().toString(), // Update with the desired date value
            message_content: comment,
            submission_id: parseInt(id), // Update with the desired submission_id value
            clinical_id: parseInt(userID), // Update with the desired clinical_id value
        };

        try {
            const response = await Axios.post(url, body, {
                headers: { "Content-Type": "application/json", Authorization: `Bearer ${token}` },
            });
            const newMessage = response.data;
            setMessages([...messages, newMessage]);
            setComment("");
            toast.success("Message added successfully", {
                position: toast.POSITION.TOP_RIGHT,
            });
        } catch (error) {
            console.error("Error adding message:", error);
            toast.error(error.response.data.error, {
                position: toast.POSITION.TOP_RIGHT,
            });
        }
    };

    return (
        <PageLayout>
            <ToastContainer />
            <Container>
                <Typography variant="h4">Evaluate Submission</Typography>

                <ContainerSubmission>
                    <TitleSubmission>{submission.body_part}</TitleSubmission>
                    <ContainerImage>
                        {submission.media_type === "image" && (
                            <ImageSubmission key={id} src={`data:image/jpeg;base64,${submission.media}`} alt="Image" />
                        )}
                        {submission.media_type === "video" && (
                            <VideoSubmission url={`data:video/mp4;base64,${submission.media}`} controls />
                        )}
                    </ContainerImage>
                    <Typography variant="body1">{submission.description}</Typography>
                </ContainerSubmission>

                <Typography variant="h5">Clinical Feedback:</Typography>

                <List>
                    {messages.map((message, index) => (
                        <React.Fragment key={message.id}>
                            <ListItem alignItems="flex-start">
                                <ListItemText
                                    primary={`Date: ${message.date}`}
                                    secondary={
                                        <>
                                            <Typography component="span" variant="subtitle2" color="textPrimary">
                                                Clinical Email:
                                            </Typography>{" "}
                                            {message.clinical_email}
                                        </>
                                    }
                                />
                            </ListItem>
                            <ListItem>
                                <ListItemText primary={`Message: ${message.message_content}`} />
                            </ListItem>
                            {index !== messages.length - 1 && <Divider />}
                        </React.Fragment>
                    ))}
                </List>

                <Typography variant="h5">Add Comment</Typography>
                <form onSubmit={handleSubmitMessage}>
                    <TextField
                        label="Comment"
                        multiline
                        rows={4}
                        variant="outlined"
                        fullWidth
                        value={comment}
                        onChange={handleMessageChange}
                    />
                    <Button variant="contained" color="primary" type="submit">
                        Add Comment
                    </Button>
                </form>
            </Container>
            <br /><br /><br /><br />
        </PageLayout>
    );
};

export default EvaluateSubmission;
