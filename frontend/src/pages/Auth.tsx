import { Button, PasswordInput, Text, TextInput, Title } from "@mantine/core";
import { useAppTheme } from "../context/ThemeContext";
import { useState } from "react";
import { useAuth } from "../context/AuthContext";

import Header from "../components/Header";
import Layout from "../components/Layout";
import { notifications } from "@mantine/notifications";
import { useNavigate } from "react-router";

const Auth = () => {
  const navigate = useNavigate();
  const { colors } = useAppTheme();
  const { login } = useAuth();
  const [signin, setSignin] = useState<boolean>(false);

  const handleLogin = () => {
    login("sometoken");
    navigate("/files");
  };

  const handleSignIn = () => {
    notifications.show({
      color: "green",
      title: <Text c={colors.text}>Successfully created account</Text>,
      position: "top-right",
      message: "Please login to your account to continue",
      style: {
        color: colors.text,
        backgroundColor: colors.background1,
        border: "1px solid green",
      },
    });
    setSignin(true);
  };

  return (
    <Layout header={<Header />}>
      <div className="h-[90vh] w-[100vw] flex justify-center items-center">
        <div
          className="p-8 w-[30rem] rounded-lg shadow-md"
          style={{ backgroundColor: colors.background2 }}
        >
          <Title c={colors.text} order={2} my={10}>
            {signin ? "Login" : "Create Account"}
          </Title>
          <TextInput
            my={10}
            c={colors.text}
            label="Username"
            placeholder="enter username"
            variant="filled"
            styles={{
              input: {
                backgroundColor: colors.background3,
                color: colors.text,
              },
            }}
          />
          <PasswordInput
            my={10}
            c={colors.text}
            label="Password"
            placeholder="enter password"
            variant="filled"
            styles={{
              input: {
                backgroundColor: colors.background3,
                color: colors.text,
              },
            }}
          />

          {signin ? (
            <Button
              my={10}
              size="md"
              color={colors.primary}
              onClick={handleLogin}
            >
              Login
            </Button>
          ) : (
            <Button
              my={10}
              size="md"
              color={colors.primary}
              onClick={handleSignIn}
            >
              Register
            </Button>
          )}

          <Text size="sm" c={colors.text}>
            {signin ? "Don't have an account? " : "Already have an account? "}
            <span
              className="text-blue-400 hover:cursor-pointer"
              onClick={() => setSignin((preValue) => !preValue)}
            >
              {signin ? "Register" : "Login"}
            </span>
          </Text>
        </div>
      </div>
    </Layout>
  );
};

export default Auth;
