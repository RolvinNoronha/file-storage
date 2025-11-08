import { Button, Text, Title } from "@mantine/core";
import { useAppTheme } from "../context/ThemeContext";
import { useNavigate } from "react-router";

const NotFound = () => {
  const { colors } = useAppTheme();
  const navigate = useNavigate();

  return (
    <div
      style={{ backgroundColor: colors.background1 }}
      className="h-screen w-screen flex justify-center items-center"
    >
      <div
        style={{ backgroundColor: colors.background2 }}
        className="h-[30%] w-[30%] px-4 py-8 rounded-3xl flex flex-col items-center justify-around shadow-xl"
      >
        <Title c={colors.text} order={2}>
          404 - Page Not Found
        </Title>
        <Text c={colors.text}>The page you are looking for does not exist</Text>
        <Button
          size="md"
          color={colors.primary}
          onClick={() => navigate("/dashboard")}
        >
          Go To Dashboard
        </Button>
      </div>
    </div>
  );
};

export default NotFound;
