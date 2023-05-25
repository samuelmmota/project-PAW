import React from "react";
import { Routes, Route, BrowserRouter } from "react-router-dom";

import Home from "../pages/Home";
import Gallery from "../pages/Gallery";
import AddSubmission from "../pages/AddSubmission";
import EditSubmission from "../pages/EditSubmission";
import Login from "../pages/Login";
import Profile from "../pages/Profile";
import Register from "../pages/Register";
import Edituser from "../pages/EditUser";
import Persons from "../pages/persons";
import { Navigate } from "react-router-dom";

const Router = () => {
  const token = sessionStorage.getItem("token");
  const isLoggedIn = token !== null;
     
  /*
        <Route path="addbook" element={<Addbook />} />
        <Route path="edituser" element={<Edituser />} />
        <Route path="persons" element={<Persons />} />
        <Route path="editbook" element={<Editbook />} />
        <Route path="profile" element={<Profile />} />
                <Route path="register" element={<Register />} />

       <Route path="editbook/:id" element={
         isLoggedIn ? (
            <Editbook />
         ) : (
            <Navigate to="/login" replace={true} />
         )
       }
       />

  */

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="gallery" element={<Gallery />} />
        <Route path="login" element={<Login />} />
        <Route path="editsubmission/:id" element={<EditSubmission />} />
        <Route path="addsubmission" element={<AddSubmission />} />
        <Route path="edituser" element={<Edituser />} />
        <Route path="persons" element={<Persons />} />
        <Route path="editsubmission" element={<EditSubmission />} />
        <Route path="profile" element={<Profile />} />
        <Route path="register" element={<Register />} />
      </Routes>
    </BrowserRouter>
  );
};

export default Router;
/*
Adicionar rotas e todos os componentes que vão ser carregados (regra de negocio)  5566#
*/

