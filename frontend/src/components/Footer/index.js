import { ContainerFooter } from "./styles";

const Footer = () => {
  const year = new Date().getFullYear();
  return (
    <ContainerFooter>
      <p>Psoriasis Image Submission and Clinical Feedback Web Application using React © {year}</p>
    </ContainerFooter>
  );
};

export default Footer;
