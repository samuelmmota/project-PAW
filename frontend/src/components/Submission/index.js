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
} from "./styles";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import Axios from "axios";
import { submissionUrl } from "../../resources/constants.js";


const Submission = ({ body_part, media, media_type, date, id, description, refreshSubmissions, owner }) => {
  const token = sessionStorage.getItem("token");
  const isLoggedIn = token !== null;
  // variavel usada pra fazer a navegação pelas paginas
  const navigate = useNavigate();

  let displayDate = date;
  let displayMonth = "";
  let displayDay = "";
  let displayYear = "";

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
 /* return (
    <>
     <ToastContainer />
      <ContainerSubmission>
        <TitleSubmission>{bodyPart}</TitleSubmission>
        <ContainerImage>
          <ImageSubmission src={image} />
        </ContainerImage>
        <center>
          {description}
          {isLoggedIn && owner && (
            <ButtonContainer>
              <Button type="button" onClick={editSubmission}>
              Update
            </Button>
            <Button type="button" onClick={deleteSubmission}>
              Delete
            </Button>
            </ButtonContainer>
          )}
        </center>
      </ContainerSubmission>
    </>
  );*/
  return (
    <>
      <ToastContainer />
      <ContainerSubmission>
      <DateText>
  {displayMonth !== "" && displayDay !== "" && displayYear !== "" ? (
    <>Date: {displayMonth}, {displayDay}, {displayYear}</>
  ) : (
    <>Date: {displayDate}</>
  )}
</DateText>

      <ContainerImage>
        {media_type === "image" && (
           <ImageSubmission
             key={id}
            src={`data:image/jpeg;base64,${media}`}
             alt="Image"
          />
                )}
        {media_type === "video" && (
            <VideoSubmission controls>
       <         source src={`data:video/mp4;base64,${media}`} type="video/mp4" />
        </VideoSubmission>
  )}
</ContainerImage>
        <p>{description}</p>
        <TitleSubmission>{body_part}</TitleSubmission>
        <center>
          {owner && (
            <ButtonContainer>
              <Button type="button" onClick={editSubmission}>
                Update
              </Button>
              <Button type="button" onClick={deleteSubmission}>
                Delete
              </Button>
            </ButtonContainer>
          )}
        </center>
      </ContainerSubmission>
    </>
  );
};



export default Submission;


/*
return (
  <>
    <Header />
    <ToastContainer />
    <ContainerSubmissions>
      {media.map((item) => (
    
         <div key={item.id}>
           {item.mediaType === "image" && (
             <img
                key={item.id}
               src={`data:image/jpeg;base64,${item.image}`}
               alt="Image"
             />
           )}
           {item.mediaType === "video" && (
             <video controls>
               <source
                 src={`data:video/mp4;base64,${item.image}`}
                 type="video/mp4"
               />
             </video>
           )}
         </div>
      ))}
    </ContainerSubmissions>
    <Footer />
  </>
);
};
/*
<img
key={item.id}
src={`data:image/jpeg;base64,${item.image}`}
alt="Image"
/>*/

