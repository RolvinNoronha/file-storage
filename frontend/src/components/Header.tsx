import { Link, useParams } from "react-router";
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
import AppService from "@/service/AppService";
import { toast } from "sonner";
import { useQueryClient } from "@tanstack/react-query";
import type { CompletedPart } from "@/store/interfaces";

const Header = () => {
  const { isAuthenticated, logout } = useAuth();
  const { theme, toggleTheme } = useAppTheme();
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [folderName, setFolderName] = useState("");
  const [isCreatingFolder, setIsCreatingFolder] = useState(false);
  const params = useParams();
  const queryClient = useQueryClient();

  const getFolderId = () => {
    const folderIds = params["*"]?.split("/");
    if (folderIds && folderIds.length > 0) {
      const id = folderIds[folderIds.length - 1];
      return id ? parseInt(id) : undefined;
    }
    return undefined;
  };

  const handleCreateFolder = async () => {
    if (!folderName.trim()) return;

    setIsCreatingFolder(true);
    const folderId = getFolderId();

    try {
      await AppService.createFolder(folderName, folderId);
      toast.success("Folder created successfully");
      setIsModalOpen(false);
      setFolderName("");
      queryClient.invalidateQueries({ queryKey: ["folderId"] });
    } catch (error) {
      console.error("Failed to create folder:", error);
      toast.error("Failed to create folder");
    } finally {
      setIsCreatingFolder(false);
    }
  };

  const handleFileUpload = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;

    const folderId = getFolderId();
    const toastId = toast.loading("Initiating upload...");

    try {
      const initRes = await AppService.initiateUpload({
        fileName: file.name,
        fileType: file.type,
        fileSize: file.size,
        folderId: folderId,
      });

      if (!initRes.success) {
        throw new Error(initRes.message);
      }

      const { uploadId, key, parts } = initRes.data;
      const completedParts: CompletedPart[] = [];
      const chunkSize = 5 * 1024 * 1024; // 5MB

      for (const part of parts) {
        const start = (part.partNumber - 1) * chunkSize;
        const end = Math.min(start + chunkSize, file.size);
        const chunk = file.slice(start, end);

        toast.loading(`Uploading part ${part.partNumber}/${parts.length}...`, {
          id: toastId,
        });

        const uploadRes = await AppService.uploadPart(part.url, chunk);

        const etag = uploadRes.headers["etag"];
        if (!etag) throw new Error(`No ETag for part ${part.partNumber}`);

        completedParts.push({
          partNumber: part.partNumber,
          etag: etag,
        });
      }

      toast.loading("Finalizing upload...", { id: toastId });
      const completeRes = await AppService.completeUpload({
        uploadId,
        key,
        parts: completedParts,
        fileName: file.name,
        fileSize: file.size,
        fileType: file.type,
        folderId,
      });

      if (completeRes.success) {
        toast.success("File uploaded successfully", { id: toastId });
        queryClient.invalidateQueries({ queryKey: ["files"] });
      } else {
        throw new Error(completeRes.message);
      }
    } catch (err: any) {
      console.error(err);
      toast.error(err.message || "Upload failed", { id: toastId });
    } finally {
      e.target.value = "";
    }
  };

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
                    onChange={handleFileUpload}
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
                value={folderName}
                onChange={(e) => setFolderName(e.target.value)}
                disabled={isCreatingFolder}
              />
            </div>
            <DialogFooter className="flex justify-center sm:justify-center gap-2">
              <Button
                variant="outline"
                onClick={() => setIsModalOpen(false)}
                className="border-primary text-primary"
                disabled={isCreatingFolder}
              >
                Cancel
              </Button>
              <Button onClick={handleCreateFolder} disabled={isCreatingFolder}>
                {isCreatingFolder ? "Creating..." : "Create"}
              </Button>
            </DialogFooter>
          </DialogContent>
        </Dialog>
      </div>

      <div className="h-[1px] w-full bg-border"></div>
    </>
  );
};

export default Header;
