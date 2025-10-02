import React, { createContext, useState, useEffect, useContext } from "react";
import COLORS from "../constants/colors";

type ThemeType = "light" | "dark";
type ThemeContextType = {
  theme: ThemeType;
  toggleTheme: () => void;
};

export const ThemeContext = createContext<ThemeContextType | null>(null);

export function ThemeProvider({ children }: { children: React.ReactNode }) {
  const [theme, setTheme] = useState<ThemeType>("light");

  useEffect(() => {
    loadSavedTheme();
  }, []);

  const loadSavedTheme = async () => {
    try {
      const savedTheme = localStorage.getItem("theme");
      if (savedTheme) {
        setTheme(savedTheme as ThemeType);
      } else {
        setTheme("light");
      }
    } catch (error) {
      console.error("Failed to load theme:", error);
    }
  };

  const toggleTheme = async () => {
    const newTheme = theme === "light" ? "dark" : "light";
    setTheme(newTheme);
    try {
      localStorage.setItem("theme", newTheme);
    } catch (error) {
      console.error("Failed to save theme:", error);
    }
  };

  return (
    <ThemeContext.Provider value={{ theme, toggleTheme }}>
      {children}
    </ThemeContext.Provider>
  );
}

export const useAppTheme = () => {
  const themeContext = useContext(ThemeContext);

  if (!themeContext) {
    throw new Error(
      "useAppTheme has to be used within <ThemeContext.Provider>"
    );
  }

  const { theme, toggleTheme } = themeContext;
  const colors = COLORS[theme];

  return { colors, theme, toggleTheme };
};
