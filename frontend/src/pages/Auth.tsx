import { useState } from "react";
import { toast } from "sonner";
import { useNavigate } from "react-router";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import AppService from "../service/AppService";
import Header from "../components/Header";
import Layout from "../components/Layout";

const Auth = () => {
  const navigate = useNavigate();
  const [signin, setSignin] = useState<boolean>(false);

  const [userName, setUserName] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  const handleLogin = async () => {
    try {
      await AppService.login(userName, password);
      toast.success("Login successful");
      navigate("/files");
    } catch (error) {
      console.error("Failed to login: ", error);
      toast.error("Login failed. Please check your credentials.");
    }
  };

  const handleSignIn = () => {
    toast.success("Successfully created account", {
      description: "Please login to your account to continue",
    });
    setSignin(true);
  };

  return (
    <Layout header={<Header />}>
      <div className="h-[94vh] w-screen flex justify-center items-center px-4">
        <div className="p-8 w-full max-w-[400px] rounded-lg shadow-lg bg-card border">
          <h2 className="text-2xl font-bold mb-6">
            {signin ? "Login" : "Create Account"}
          </h2>

          <div className="space-y-4">
            <div className="space-y-2">
              <Label htmlFor="username">Username</Label>
              <Input
                id="username"
                placeholder="Enter username"
                className="bg-muted border-primary"
                value={userName}
                onChange={(e) => setUserName(e.target.value)}
              />
            </div>

            <div className="space-y-2">
              <Label htmlFor="password">Password</Label>
              <Input
                id="password"
                type="password"
                placeholder="Enter password"
                className="bg-muted border-primary"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </div>

            {signin ? (
              <Button className="w-full h-11 text-lg" onClick={handleLogin}>
                Login
              </Button>
            ) : (
              <Button className="w-full h-11 text-lg" onClick={handleSignIn}>
                Register
              </Button>
            )}

            <p className="text-sm text-center">
              {signin ? "Don't have an account? " : "Already have an account? "}
              <span
                className="text-primary hover:underline cursor-pointer font-semibold"
                onClick={() => setSignin((preValue) => !preValue)}
              >
                {signin ? "Register" : "Login"}
              </span>
            </p>
          </div>
        </div>
      </div>
    </Layout>
  );
};

export default Auth;
