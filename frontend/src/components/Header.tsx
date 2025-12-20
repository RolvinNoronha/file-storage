import { Link } from "react-router";
import { useAppTheme } from "../context/ThemeContext";
import { Moon, Sun, LogOut } from "lucide-react";
import { useAuth } from "../context/AuthContext";
import { useState } from "react";
import filesIcons from "../assets/files-icon.png";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";

const Header = () => {
  const { isAuthenticated, logout } = useAuth();
  const { theme, toggleTheme } = useAppTheme();
  const [isModalOpen, setIsModalOpen] = useState(false);

  return (
    <>
      <div className="h-[6vh] flex flex-col justify-center sticky top-0 bg-background/80 backdrop-blur-3xl z-50 px-4">
        <div className="flex flex-row justify-between items-center w-full">
          <div className="flex items-center gap-2">
            <img className="h-8 w-8" src={filesIcons} alt="File Icon" />
            <h1 className="text-2xl font-bold">File Uploader</h1>
          </div>
          <div className="flex items-center gap-4">
            {isAuthenticated ? (
              <>
                <div className="flex items-center justify-center">
                  <input
                    type="file"
                    id="fileInput"
                    className="hidden"
                    accept="*/*"
                    onChange={() => {}}
                  />
                  <label
                    htmlFor="fileInput"
                    className="cursor-pointer bg-primary text-primary-foreground font-medium px-4 py-[6px] rounded-md hover:opacity-90 transition-opacity"
                  >
                    Choose File
                  </label>
                </div>
                <Button size="default" onClick={() => setIsModalOpen(true)}>
                  Add Folder
                </Button>
                <Button
                  size="default"
                  variant="outline"
                  className="text-destructive border-destructive hover:text-destructive-foreground"
                  onClick={logout}
                  // onClick={() => {}}
                >
                  <LogOut className="mr-2 h-4 w-4" />
                  Log Out
                </Button>
              </>
            ) : (
              <Link to={"/auth"}>
                <Button size="default">Get Started</Button>
              </Link>
            )}
            <Button
              variant="ghost"
              size="icon"
              onClick={toggleTheme}
              className="rounded-full"
            >
              {theme === "dark" ? (
                <Moon className="h-[1.5rem] w-[1.5rem]" />
              ) : (
                <Sun className="h-[1.5rem] w-[1.5rem]" />
              )}
            </Button>
          </div>
        </div>

        <Dialog open={isModalOpen} onOpenChange={setIsModalOpen}>
          <DialogContent className="sm:max-w-[425px]">
            <DialogHeader>
              <DialogTitle>Folder Name</DialogTitle>
            </DialogHeader>
            <div className="grid gap-4 py-4">
              <Input
                placeholder="Enter folder name"
                className="col-span-3 bg-muted border-primary"
              />
            </div>
            <DialogFooter className="flex justify-center sm:justify-center gap-2">
              <Button
                variant="outline"
                onClick={() => setIsModalOpen(false)}
                className="border-primary text-primary"
              >
                Cancel
              </Button>
              <Button onClick={() => setIsModalOpen(false)}>Create</Button>
            </DialogFooter>
          </DialogContent>
        </Dialog>
      </div>

      <div className="h-[1px] w-full bg-border"></div>
    </>
  );
};

export default Header;
