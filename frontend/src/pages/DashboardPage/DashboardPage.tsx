import {createTheme, ThemeProvider} from "@mui/material/styles";
import * as React from 'react';
import {useEffect, useState} from 'react';
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
  const [startDate, setStartDate] = React.useState<Dayjs>(dayjs(presentDate));
  const [endDate, setEndDate] = React.useState<Dayjs>(dayjs(presentDate));

  const calendarLocalizer = momentLocalizer(moment);
  const [events, setEvents] = useState<Event[]>([]);
  const [loadEvent, setLoadEventState] = useState<boolean>(true);
  useEffect(() => {
    if (loadEvent) {
      fetch('http://localhost:8080/appointments', {
        method: "GET",
        headers: {
          'Accept': 'application/json',
          'Content-type': 'application/json; charset=UTF-8',
        },
        credentials: "include",
        mode: "cors"
      })
        .then((res) => res.json())
        .then(data => {
          console.log(data);
          const appointments = data.appointments.map((appointment: any) => mapAppointmentToCalendarEvent(appointment));
          setEvents(appointments);
          setLoadEventState(false);
        })
        .catch(err => {
          console.log(err);
          navigate("/login");
        })
    }
  }, [loadEvent]);

  const theme = createTheme();

  const bookAppointment = (startDate: Dayjs, endDate: Dayjs) => {
    console.log("Handle appointment");
    fetch("http://localhost:8080/appointment", {
      method: "POST",
      body: JSON.stringify({
        title: "Test Title", // TODO new form field for title
        owner: "Admin", // TODO retrieve owner's name from login response
        startDateTime: startDate.valueOf(),
        endDateTime: endDate.valueOf()
      }),
      headers: {
        "Content-type": "application/json; charset=UTF-8",
      },
      credentials: "include",
      mode: "cors"
    })
      .then(response => response.json())
      .then(() => {
        setLoadEventState(true);
      })
      .catch(err => console.warn(err))
  }

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
                    setStartDate(newValue);
                  }}
                  renderInput={(params: TextFieldProps) => <TextField {...params} />}
                />
                <DateTimePicker
                  minutesStep={60}
                  label="End"
                  value={endDate}
                  onChange={(newValue: any) => {
                    setEndDate(newValue);
                  }}
                  renderInput={(params: TextFieldProps) => <TextField {...params} />}
                />
                <Button
                  onClick={() => bookAppointment(startDate, endDate)}
                  type="submit"
                  fullWidth
                  variant="contained"
                >
                  Book Appointment
                </Button>
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

        <Box
          sx={{
            marginX: 50
          }}
        >
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
            sx={{mt: 4}}
          >
            Logout
          </Button>
        </Box>

      </LocalizationProvider>
    </ThemeProvider>
  )
}

function mapAppointmentToCalendarEvent(appointment: any): Event {
  return {
    title: appointment.title,
    start: new Date(appointment.startDateTime),
    end: new Date(appointment.endDateTime)
  }
}
