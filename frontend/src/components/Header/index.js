import React from "react";
import { AppBar, Toolbar, IconButton, Typography, Button } from "@mui/material";
import { styled } from "@mui/system";
import { useNavigate } from "react-router-dom";
import profileIcon from "../../assets/profile.svg";
import { useState, useEffect } from "react";
import Axios from "axios";
import { userUrl } from "../../resources/constants.js";

const CustomAppBar = styled(AppBar)({
  zIndex: (theme) => theme.zIndex.drawer + 1,
});

const CustomTitle = styled(Typography)({
  flexGrow: 1,
});

const CustomProfileIcon = styled("img")({
  marginLeft: (theme) => theme.spacing(2),
});

const Header = () => {
  const navigate = useNavigate();
  const token = sessionStorage.getItem("token");
  const isLoggedIn = token !== null;

  var userID = 0;

  if(token !== null){
    const decodedToken = JSON.parse(atob(token.split(".")[1]));
    userID = decodedToken.user_id;
  }
  
  const [user, setUser] = useState({
  });

  useEffect(() => {
    getMyUser();
  }, []);

  async function getMyUser() {
    if(token === null){
      return;
    }
    const url = userUrl + userID;
    try {
      const response = await Axios.get(url, {
        headers: { "Content-Type": "application/json", Authorization: token },
      });
      const userData = response.data.user;
      setUser({
        user_email: userData.email,
        user_is_clinical: userData.isClinical,
      });
      console.log(response);
    } catch (err) {
      console.log(err);}
  }



  const handleLogout = () => {
    sessionStorage.removeItem("token");
    navigate("/login");
  };

  const handleLogin = () => {
    navigate("/login");
  };

  const handleGallery = () => {
    navigate("/gallery");
  };

  const handleEvaluate = () => {
    navigate("/evaluate");
  };


  return (
    <CustomAppBar position="fixed">
      <Toolbar>
        <CustomTitle variant="h6" onClick={() => navigate("/")}>Clinical Feedback</CustomTitle>
        {isLoggedIn ? (
          <>
            <Button color="inherit" onClick={handleEvaluate}>
              Evaluate
            </Button>

            {!user.user_is_clinical && (
              <>
            <Button color="inherit" onClick={() => navigate("/addsubmission")}>
              Add Submission
            </Button>
            <Button color="inherit" onClick={handleGallery}>
              Gallery
            </Button>
            </>
            )}
            <Button color="inherit" onClick={handleLogout}>
              Logout
            </Button>
            <IconButton edge="end" color="inherit" href="/profile">
              <CustomProfileIcon src={profileIcon} alt="Profile icon" />
            </IconButton>
          </>
        ) : (
          <>
            <Button color="inherit" onClick={handleLogin}>
              Login
            </Button>
          </>
        )}
      </Toolbar>
    </CustomAppBar>
  );
};

export default Header;
