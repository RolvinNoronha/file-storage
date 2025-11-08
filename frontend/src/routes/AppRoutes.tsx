import { createBrowserRouter, RouterProvider } from "react-router";
import App from "../App";
import ProtectedRoutes from "./ProtectedRoutes";
import Auth from "../pages/Auth";
import Dashboard from "../pages/Dashboard";
import NotFound from "../pages/NotFound";

const AppRoutes = () => {
  const router = createBrowserRouter([
    {
      path: "/",
      children: [
        { index: true, element: <App /> },
        { path: "/auth", element: <Auth /> },
        {
          element: <ProtectedRoutes />,
          children: [
            { path: "/dashboard", element: <Dashboard /> },
            { path: "/dashboard/*", element: <Dashboard /> },
          ],
        },
      ],
    },
    {
      path: "*",
      element: <NotFound />,
    },
  ]);

  return <RouterProvider router={router} />;
};

export default AppRoutes;
