import { Navigate, Outlet } from "react-router";
import { useAuth } from "../context/AuthContext";
import { Loader, Text } from "@mantine/core";
import { useAppTheme } from "../context/ThemeContext";

const ProtectedRoutes = () => {
  const { isLoading, isAuthenticated, authToken } = useAuth();
  const { colors } = useAppTheme();

  if (isLoading) {
    return (
      <div className="h-screen w-screen flex flex-col justify-center items-center">
        <Loader color={colors.primary} />
        <Text c={colors.text}>Please Wait...</Text>
      </div>
    );
  }

  // if (!isAuthenticated || !authToken) {
  //   return <Navigate to={"/auth"} />;
  // }

  return <Outlet />;
};

export default ProtectedRoutes;
