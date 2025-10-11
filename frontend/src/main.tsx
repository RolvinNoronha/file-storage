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

const theme = createTheme({
  fontFamily: "Inter, sans-serif",
});

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <AuthProvider>
      <MantineProvider theme={theme}>
        <Notifications />
        <ThemeProvider>
          <AppRoutes />
        </ThemeProvider>
      </MantineProvider>
    </AuthProvider>
  </StrictMode>
);
