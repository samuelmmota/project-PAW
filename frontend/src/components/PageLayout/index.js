import React from "react";
import Box from "@mui/material/Box";
import Header from "../../components/Header";
import Footer from "../../components/Footer";

const PageLayout = ({ children }) => {
  return (
    <>
      <Header />
      <Box sx={{ py: 8 }}>{children}</Box>
      <Footer />
    </>
  );
};

export default PageLayout;
