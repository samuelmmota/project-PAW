import React from "react";
import { useNavigate } from "react-router-dom";
import {
  ProfileSection,
  Container,
  ProfileImage,
  Name,
  Email,
  ButtonContainer,
  Button
} from "./styles";


const User = ({ userImage, userName, userEmail }) => {
  const isLoggedIn = true;
  // variavel usada pra fazer a navegação pelas paginas
  const navigate = useNavigate();

  function deleteUser(element) {
    console.log("delete user");
    console.log(element.target);
    element.target.parentNode.parentNode.remove();
  }
  function editUser(element) {
    console.log("edit user button");
    console.log(element.target);
    navigate("/edituser");
  }
  return (
    <>
        <Container>
          <ProfileSection>
          <ProfileImage src={userImage} alt="User profile Image" />
          <Name>{userName}</Name>
          <Email>{userEmail}</Email>
        </ProfileSection>
        <center>
          {isLoggedIn && (
            <ButtonContainer>
              <Button type="button" onClick={editUser}>
              Update
            </Button>
            <Button type="button" onClick={deleteUser}>
              Delete
            </Button>
            </ButtonContainer>
          )}
        </center>
       </Container>

    </>
  );
};

export default User;
