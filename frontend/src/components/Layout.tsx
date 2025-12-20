import React from "react";

const Layout = ({
  header,
  children,
}: {
  header: React.ReactNode;
  children: React.ReactNode;
}) => {
  return (
    <div className="h-screen w-screen overflow-auto transition-colors ease-out duration-100 bg-background text-foreground">
      {header}
      {children}
    </div>
  );
};

export default Layout;
