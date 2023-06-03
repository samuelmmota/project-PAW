import { Box, Typography, Link } from "@mui/material";
import { styled } from "@mui/system";

const CustomFooter = styled(Box)`
  background-color: #f5f5f5;
  padding: 16px;
  text-align: center;
  position: fixed;
  bottom: 0;
  width: 100%;
`;

const Footer = () => {
  const year = new Date().getFullYear();

  return (
    <CustomFooter component="footer">
      <Typography variant="body2" color="textSecondary">
        Psoriasis Image Submission and Clinical Feedback Web Application using React © {year}
      </Typography>
      <Typography variant="body2" color="textSecondary">
        Created by Your Name
      </Typography>
      <Typography variant="body2" color="textSecondary">
        Contact: <Link href="mailto:yourname@example.com">yourname@example.com</Link>
      </Typography>
    </CustomFooter>
  );
};

export default Footer;
