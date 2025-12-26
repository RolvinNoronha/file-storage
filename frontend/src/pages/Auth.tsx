import { useState } from "react";
import { toast } from "sonner";
import { useNavigate } from "react-router";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import AppService from "../service/AppService";
import Header from "../components/Header";
import Layout from "../components/Layout";
import { useAuth } from "@/context/AuthContext";

const Auth = () => {
  const navigate = useNavigate();
  const { login } = useAuth();
  const [signin, setSignin] = useState<boolean>(false);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const [userName, setUserName] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  const handleLogin = async () => {
    setIsLoading(true);
    try {
      login(userName, password);
      toast.success("Login successful");
      navigate("/files");
    } catch (error) {
      console.error("Failed to login: ", error);
      toast.error("Login failed. Please check your credentials.");
    } finally {
      setIsLoading(false);
    }
  };

  const handleSignIn = async () => {
    setIsLoading(true);
    try {
      await AppService.register(userName, password);
      toast.success("Successfully created account", {
        description: "Please login to your account to continue",
      });
      setSignin(true);
      setUserName("");
      setPassword("");
    } catch (error) {
      console.error("Failed to login: ", error);
      toast.error("Login failed. Please check your credentials.");
    } finally {
      setIsLoading(false);
    }
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
                disabled={isLoading}
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
                disabled={isLoading}
              />
            </div>

            {signin ? (
              <Button
                className="w-full h-11 text-lg"
                onClick={handleLogin}
                disabled={isLoading}
              >
                {isLoading ? "Please wait..." : "Login"}
              </Button>
            ) : (
              <Button
                className="w-full h-11 text-lg"
                onClick={handleSignIn}
                disabled={isLoading}
              >
                {isLoading ? "Please wait..." : "Register"}
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
