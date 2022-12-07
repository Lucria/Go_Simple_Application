import React from 'react';
import {useRoutes} from "react-router-dom";
import router from "./Router";
import CssBaseline from "@mui/material/CssBaseline";
import {createTheme, ThemeProvider} from "@mui/material/styles";

export default function App() {
  const content = useRoutes(router);
  const theme = createTheme();

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline/>
      {content}
    </ThemeProvider>
  );
}
