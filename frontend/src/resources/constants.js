import { useNavigate } from "react-router-dom";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import Axios from "axios";
// Constants for the frontend

//Strings
const apiUrl = "http://localhost:3000/paw/api/v1";
export const submissionUrl = `${apiUrl}/submission/`;
export const userUrl = `${apiUrl}/user/`;
export const auth_loginUrl = `${apiUrl}/auth/login`; 
export const auth_registerUrl = `${apiUrl}/auth/register`; 
export const refreshTokenUrl = `${apiUrl}/auth/refreshtoken`;
export const images_test = `${apiUrl}/image`;
export const evaluateUrl = `${apiUrl}/evaluate`;
export const clinicalUrl = `${apiUrl}/user/clinical/`;

//Colors
export const primaryColor = "#4169E1"
export const secondaryColor = "#8A2BE2";
//Integers