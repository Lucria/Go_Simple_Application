import {createTheme, ThemeProvider} from "@mui/material/styles";
import * as React from 'react';
import {useState} from 'react';
import dayjs, {Dayjs} from 'dayjs';
import TextField, {TextFieldProps} from '@mui/material/TextField';
import {AdapterDayjs} from '@mui/x-date-pickers/AdapterDayjs';
import {LocalizationProvider} from '@mui/x-date-pickers/LocalizationProvider';
import {Card, CardContent, Stack} from "@mui/material";
import {DateTimePicker} from "@mui/x-date-pickers";
import Box from "@mui/material/Box";
import {Calendar, Event, momentLocalizer} from 'react-big-calendar';
import moment from 'moment';
import Button from "@mui/material/Button";
import {useNavigate} from "react-router-dom";

import 'react-big-calendar/lib/addons/dragAndDrop/styles.css'
import 'react-big-calendar/lib/css/react-big-calendar.css'

export default function DashboardPage() {
  const navigate = useNavigate();

  const presentDate = new Date(Date.now()).setMinutes(0, 0, 0);
  const [startDate, setStartDate] = React.useState<Dayjs | null>(dayjs(presentDate));
  const [endDate, setEndDate] = React.useState<Dayjs | null>(dayjs(presentDate));

  const calendarLocalizer = momentLocalizer(moment);
  const [events, setEvents] = useState<Event[]>([
    {
      title: 'Learn cool stuff',
      start: new Date(presentDate),
      end: moment(presentDate).add(1, 'days').toDate(),
    },
  ])

  const theme = createTheme();

  return (
    <ThemeProvider theme={theme}>
      <LocalizationProvider dateAdapter={AdapterDayjs}>
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
            <CardContent>
              <Stack spacing={3}>
                <DateTimePicker
                  minutesStep={60}
                  label="Start"
                  value={startDate}
                  onChange={(newValue: any) => {
                    console.log(newValue);
                    setStartDate(newValue);
                  }}
                  renderInput={(params: TextFieldProps) => <TextField {...params} />}
                />
                <DateTimePicker
                  minutesStep={60}
                  label="End"
                  value={endDate}
                  onChange={(newValue: any) => {
                    console.log(newValue);
                    setEndDate(newValue);
                  }}
                  renderInput={(params: TextFieldProps) => <TextField {...params} />}
                />
              </Stack>
            </CardContent>
          </Card>
        </Box>

        <Box
          sx={{
            paddingX: 4
          }}
        >
          <Calendar
            localizer={calendarLocalizer}
            events={events}
            startAccessor="start"
            endAccessor="end"
            style={{height: 500}}>
          </Calendar>
        </Box>

        <Button
          onClick={() => {
            document.cookie.split(";").forEach((c) => {
              document.cookie = c
                .replace(/^ +/, "")
                .replace(/=.*/, "=;expires=" + new Date().toUTCString() + ";path=/");
            });
            navigate("/login");
          }}
          type="submit"
          fullWidth
          variant="contained"
          sx={{mt: 8, mb: 2}}
        >
          Logout
        </Button>
      </LocalizationProvider>
    </ThemeProvider>
  )
}
