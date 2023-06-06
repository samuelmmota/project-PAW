import React, { useEffect, useState } from "react";
import Footer from "../../components/Footer";
import Header from "../../components/Header";
import Axios from "axios";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { ContainerSubmissions } from "./styles";
import { submissionUrl,images_test} from "../../resources/constants.js";


const Gallery = () => {
//TESTING code
const [media, setMedia] = useState([]);

useEffect(() => {
  fetchMedia();
}, []);

const fetchMedia = async () => {
  try {
    const response = await Axios.get(images_test); 
    setMedia(response.data);
  } catch (error) {
    toast.error("Failed to fetch media");
  }
};
// END TESTING code

  const [submissions, setSubmissions] = useState([]);
  const token = sessionStorage.getItem("token");
  const isLoggedIn = token !== null;


  useEffect(() => {
    getSubmissions();
  }, []);

  
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

if (!Array.isArray(media)) {
  return null; // or display an error message
}

return (
  <>
    <Header />
    <ToastContainer />
    <ContainerSubmissions>
      {/* Render the media */}
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
export default Gallery;
