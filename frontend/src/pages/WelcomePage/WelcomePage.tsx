import Box from "@mui/material/Box";
import {createTheme, ThemeProvider} from "@mui/material/styles";
import {Card, CardActions, CardContent, CardHeader, IconButton} from "@mui/material";
import Typography from "@mui/material/Typography";
import LoginIcon from '@mui/icons-material/Login';
import AppRegistrationIcon from '@mui/icons-material/AppRegistration';
import {useNavigate} from "react-router-dom";

export default function WelcomePage() {
  const theme = createTheme();
  const navigate = useNavigate();

  return (
    <ThemeProvider theme={theme}>

      <Box component="span"
           sx={{
             overflow: 'auto',
             display: 'flex',
             overflowX: 'hidden',
             alignItems: 'center',
             justifyContent: 'center'
           }}>
        <Card variant="outlined"
              sx={{p: 5, mb: 10, borderRadius: 16, borderColor: 'primary.main¸'}}>
          <CardHeader
            title={"Welcome to Dental System"}
            subheader={"I love Foodpanda!"}
          ></CardHeader>
          <CardContent>
            <Typography variant="body1" color="text.primary" gutterBottom>
              Click the icons below to login or register for an account.
            </Typography>
          </CardContent>
          <CardActions>
            <IconButton aria-label="login" onClick={() => {
              navigate("/login")
            }}>
              <LoginIcon/>
            </IconButton>
            <IconButton aria-label="register" onClick={() => {
              navigate("/register")
            }}>
              <AppRegistrationIcon/>
            </IconButton>
          </CardActions>
        </Card>
      </Box>
    </ThemeProvider>
  )
}

