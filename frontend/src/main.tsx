import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { MantineProvider } from "@mantine/core";
import { ThemeProvider } from "./context/ThemeContext.tsx";

import "./index.css";
import "@mantine/core/styles.css";
import "@mantine/notifications/styles.css";

import AppRoutes from "./routes/AppRoutes.tsx";
import { AuthProvider } from "./context/AuthContext.tsx";
import { Notifications } from "@mantine/notifications";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <AuthProvider>
      <MantineProvider>
        <Notifications />
        <ThemeProvider>
          <AppRoutes />
        </ThemeProvider>
      </MantineProvider>
    </AuthProvider>
  </StrictMode>
);
