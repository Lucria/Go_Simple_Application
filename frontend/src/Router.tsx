import {RouteObject} from "react-router-dom";
import ErrorPage from "./pages/ErrorPage/ErrorPage";
import RegisterPage from "./pages/RegisterPage/RegisterPage";
import LoginPage from "./pages/LoginPage/LoginPage";
import React from "react";
import BaseLayout from "./layouts/BaseLayout/BaseLayout";

const routes: RouteObject[] = ([
  {
    path: "",
    element: <BaseLayout/>,
    children: [
      {
        path: "/register",
        element: <RegisterPage/>
      },
      {
        path: "/login",
        element: <LoginPage/>
      },
      {
        path: "*",
        element: <ErrorPage/>
      }
    ]
  },
]);

export default routes;