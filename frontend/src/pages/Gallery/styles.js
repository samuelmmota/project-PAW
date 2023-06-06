import styled from "styled-components";
import { primaryColor } from "../../resources/constants.js";

export const ContainerSubmissions = styled.div`
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
`;

export const Title = styled.h1`
text-align: center;
margin-top: 20px;
font-size: 2rem;
color: #333;
text-transform: uppercase;
`;

export const FilterContainer = styled.div`
  background-color: ${primaryColor};
  color: white;
  padding: 10px;
  margin-bottom: 10px;
  display: flex;
  gap: 40px;
`;

export const FilterItem = styled.div`
  display: flex;
  flex-direction: column;
`;

export const FilterLabel = styled.label`
  font-weight: bold;
  text-align: center;
`;

export const PageContainer = styled.div`
  margin-bottom: 90px; /* Adjust the margin as needed */
`;