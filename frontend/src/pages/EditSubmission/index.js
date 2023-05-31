import React, { useRef, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import Footer from "../../components/Footer";
import Header from "../../components/Header";
import Axios from "axios";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { submissionUrl } from "../../resources/constants.js";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";

import {
    ContainerSubmission,
    ContainerInfosSubmission,
    ContainerInputs,
    InputEditSubmissionDescription,
    ContainerButtonAdd,
    ButtonAddSubmission,
    Title,
} from "./styles";

const EditSubmission = () => {
    const { id } = useParams();
    const navigate = useNavigate();

    const description = useRef();
    const [selectedDate, setSelectedDate] = useState(null);
    const body_part = useRef();

    async function editSubmission(e) {
        e.preventDefault();
        console.log("edit Submission");

        const editedSubmission = {
            description: description.current.value,
            date: selectedDate,
            body_part: body_part.current.value,
        };

        const url = submissionUrl + id;
        const token = sessionStorage.getItem("token");

        try {
            const response = await Axios.put(url, editedSubmission, {
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
            <Title>Edit Submission {id}</Title>
            <ContainerSubmission>
                <ContainerInfosSubmission onSubmit={editSubmission}>
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
                        <ButtonAddSubmission>Edit</ButtonAddSubmission>
                    </ContainerButtonAdd>
                </ContainerInfosSubmission>
            </ContainerSubmission>
            <Footer />
        </>
    );
};

export default EditSubmission;
