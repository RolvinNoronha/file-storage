import { useNavigate } from "react-router";
import { Button } from "@/components/ui/button";

const NotFound = () => {
  const navigate = useNavigate();

  return (
    <div className="h-screen w-screen flex justify-center items-center bg-background text-foreground">
      <div className="max-w-md w-full p-8 rounded-3xl flex flex-col items-center justify-around shadow-xl bg-card border space-y-6 text-center">
        <h2 className="text-3xl font-bold">404 - Page Not Found</h2>
        <p className="text-muted-foreground text-lg">
          The page you are looking for does not exist
        </p>
        <Button
          size="lg"
          className="w-full h-12"
          onClick={() => navigate("/files")}
        >
          Go To Dashboard
        </Button>
      </div>
    </div>
  );
};

export default NotFound;