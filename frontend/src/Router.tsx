import {RouteObject} from "react-router-dom";
import ErrorPage from "./pages/ErrorPage/ErrorPage";
import RegisterPage from "./pages/RegisterPage/RegisterPage";
import LoginPage from "./pages/LoginPage/LoginPage";
import React from "react";
import BaseLayout from "./layouts/BaseLayout/BaseLayout";
import WelcomePage from "./pages/WelcomePage/WelcomePage";
import DashboardPage from "./pages/DashboardPage/DashboardPage";

const routes: RouteObject[] = ([
  {
    path: "",
    element: <BaseLayout/>,
    children: [
      {
        path: "/",
        element: <WelcomePage/>
      },
      {
        path: "/register",
        element: <RegisterPage/>
      },
      {
        path: "/login",
        element: <LoginPage/>
      },
      {
        path: "/dashboard",
        element: <DashboardPage/>
      },
      {
        path: "*",
        element: <ErrorPage/>
      }
    ]
  },
]);

export default routes;