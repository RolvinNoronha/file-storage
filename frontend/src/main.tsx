import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { createTheme, MantineProvider } from "@mantine/core";
import { ThemeProvider } from "./context/ThemeContext.tsx";

import "./index.css";
import "@mantine/core/styles.css";
import "@mantine/notifications/styles.css";

import AppRoutes from "./routes/AppRoutes.tsx";
import { AuthProvider } from "./context/AuthContext.tsx";
import { Notifications } from "@mantine/notifications";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

const Main = () => {
  const theme = createTheme({
    fontFamily: "Inter, sans-serif",
  });

  const queryClient = new QueryClient();

  return (
    <QueryClientProvider client={queryClient}>
      <MantineProvider theme={theme}>
        <Notifications />
        <AppRoutes />
      </MantineProvider>
    </QueryClientProvider>
  );
};

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <AuthProvider>
      <ThemeProvider>
        <Main />
      </ThemeProvider>
    </AuthProvider>
  </StrictMode>
);
