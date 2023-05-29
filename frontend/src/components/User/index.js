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

const User = ({userName, userEmail }) => {
  const isLoggedIn = true;
  // variavel usada pra fazer a navegação pelas paginas
  const navigate = useNavigate();

  function deleteUser(element) {
    console.log("delete user");
    console.log(element.target);
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
          <ProfileImage src="https://thumbs.dreamstime.com/b/default-avatar-profile-vector-user-profile-default-avatar-profile-vector-user-profile-profile-179376714.jpg" alt="User profile Image" />
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
