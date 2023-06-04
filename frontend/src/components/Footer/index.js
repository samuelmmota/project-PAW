import { Box, Typography, Link, Button } from "@mui/material";
import { styled } from "@mui/system";
import Axios from "axios";
import { exportToReaserchURL } from "../../resources/constants.js";
import { saveAs } from "file-saver";
import { tertiaryColor } from "../../resources/constants.js";

const CustomFooter = styled(Box)`
  background-color: ${tertiaryColor};
  padding: 16px;
  text-align: center;
  position: fixed;
  bottom: 0;
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
`;

const ButtonContainer = styled(Box)`
  display: flex;
  align-items: center;
`;


const Footer = () => {
  const year = new Date().getFullYear();



  async function exportSubmissions() {
    const url = exportToReaserchURL;
    try {
      const response = await Axios.get(url, {
        headers: { "Content-Type": "application/json" },
      });
  
      const contentDispositionHeader = response.headers["content-disposition"];
      const fileName = contentDispositionHeader
        ? contentDispositionHeader.split("filename=")[1]
        : "submissions.csv";
  
      const submissions = response.data.submissions;
      console.log(submissions);
  
      // Create the CSV content
      let csvContent = "date,body_part,media,media_type\n";
      submissions.forEach((submission) => {
        const { date, body_part, media, media_type } = submission;
        csvContent += `${date},${body_part},${media},${media_type}\n`;
      });
  
      // Create a Blob object from the CSV content
      const blob = new Blob([csvContent], { type: "text/csv" });
  
      // Save the file using the FileSaver.js library
      saveAs(blob, fileName);
    } catch (err) {
      console.log(err);
    }
  }

  return (
    <CustomFooter component="footer">
        <Button color="inherit" onClick={exportSubmissions}>Export</Button>
      <Typography variant="body2" color="textSecondary">
        Psoriasis Image Submission and Clinical Feedback 
      </Typography>
      <Typography variant="body2" color="textSecondary">
        Contact: <Link href="mailto:36726@ufp.edu.pt">36726@ufp.edu.pt</Link>
     / <Link href="mailto:37146@ufp.edu.pt">37146@ufp.edu.pt</Link>
      </Typography>
      <Typography variant="body2" color="textSecondary">
        Created by Samuel Mota & João Piedade
      </Typography>
    
      
    </CustomFooter>
  );
};

export default Footer;
