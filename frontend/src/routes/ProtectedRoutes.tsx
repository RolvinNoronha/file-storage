import { Outlet } from "react-router";
import { useAuth } from "../context/AuthContext";
import { Loader2 } from "lucide-react";

const ProtectedRoutes = () => {
  const { isLoading } = useAuth();

  if (isLoading) {
    return (
      <div className="h-screen w-screen flex flex-col justify-center items-center gap-2 bg-background text-foreground">
        <Loader2 className="h-10 w-10 animate-spin text-primary" />
        <p className="font-medium">Please Wait...</p>
      </div>
    );
  }

  // if (!isAuthenticated || !authToken) {
  //   return <Navigate to={"/auth"} />;
  // }

  return <Outlet />;
};

export default ProtectedRoutes;