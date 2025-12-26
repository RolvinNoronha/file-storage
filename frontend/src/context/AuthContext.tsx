import AppService from "@/service/AppService";
import {
  createContext,
  useContext,
  useState,
  useEffect,
  type ReactNode,
} from "react";

type AuthContextType = {
  authToken: string | null | undefined;
  isAuthenticated: boolean;
  isLoading: boolean;
  login: (userName: string, password: string) => void;
  logout: () => void;
};
const AuthContext = createContext<AuthContextType | null>(null);

const getCookie = (name: string) => {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);

  if (parts.length === 2) {
    const part = parts.pop();
    if (part !== undefined) {
      return part.split(";").shift();
    }
    return null;
  }

  return null;
};

export const AuthProvider = ({ children }: { children: ReactNode }) => {
  const [authToken, setAuthToken] = useState<string | null | undefined>(null);
  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(false);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  useEffect(() => {
    const token = getCookie("jwt_token");
    if (token) {
      setAuthToken(token);
      setIsAuthenticated(true);
    } else {
      setIsAuthenticated(false);
    }
    setIsLoading(false);
  }, []);

  const login = async (userName: string, password: string) => {
    // document.cookie = `jwt_token=${token}; path=/; secure; samesite=strict`;
    try {
      await AppService.login(userName, password);
      setIsAuthenticated(true);
    } catch (error) {
      throw error;
    }
  };

  const logout = () => {
    document.cookie = "jwt_token=; path=/; max-age=-1";
    setAuthToken(null);
    setIsAuthenticated(false);
  };

  return (
    <AuthContext.Provider
      value={{ authToken, isAuthenticated, isLoading, login, logout }}
    >
      {children}
    </AuthContext.Provider>
  );
};

// Custom hook to use the Auth context
export const useAuth = () => {
  const authContext = useContext(AuthContext);
  if (!authContext) {
    throw new Error("useAuth has to be used within <AuthContext.Provider>");
  }

  return authContext;
};
