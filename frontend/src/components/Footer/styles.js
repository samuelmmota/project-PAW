import styled from "styled-components";
import { primaryColor} from "../../resources/constants.js";
/**
Style components CSS servem para ser usados como copmpoentes de estilos em jss
*/
export const ContainerFooter = styled.div`
  background-color: ${primaryColor};
  color: white;
  font-family: "Anton", sans-serif;
  font-size: 18px;
  font-weight: bold;
  display: flex;
  justify-content: space-between;
  padding: 16px 60px;
  margin-top: 60px;
  bottom: 0;
  position: fixed;
  width: 100%;
`;