import React from "react";
import { useAppTheme } from "../context/ThemeContext";

const Layout = ({
  header,
  children,
}: {
  header: React.ReactNode;
  children: React.ReactNode;
}) => {
  const { colors } = useAppTheme();

  return (
    <div
      className="h-screen w-screen  transition-colors ease-out duration-100"
      style={{ backgroundColor: colors.background1 }}
    >
      {header}
      {children}
    </div>
  );
};

export default Layout;
