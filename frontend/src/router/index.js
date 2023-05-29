import React from "react";
import { Routes, Route, BrowserRouter } from "react-router-dom";

import Home from "../pages/Home";
import Gallery from "../pages/Gallery";
import AddSubmission from "../pages/AddSubmission";
import EditSubmission from "../pages/EditSubmission";
import Login from "../pages/Login";
import Profile from "../pages/Profile";
import Edituser from "../pages/EditUser";
import Evaluate from "../pages/Evaluate";
import { Navigate } from "react-router-dom";

const Router = () => {
  const token = sessionStorage.getItem("token");
  const isLoggedIn = token !== null;

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="gallery" element={<Gallery />} />
        <Route path="login" element={<Login />} />
        <Route path="editsubmission/:id" element={<EditSubmission />} />
        <Route path="addsubmission" element={<AddSubmission />} />
        <Route path="edituser" element={<Edituser />} />
        <Route path="evaluate" element={<Evaluate />} />
        <Route path="editsubmission" element={<EditSubmission />} />
        <Route path="profile" element={<Profile />} />
      </Routes>
    </BrowserRouter>
  );
};

export default Router;