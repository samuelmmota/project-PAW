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
export const primaryColor = "#1A76D2"
export const secondaryColor = "#009DE7";
export const tertiaryColor = "#00BED8";
export const quaternaryColor = "#00D9B1";
//Integers