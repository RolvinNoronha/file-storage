import { useEffect, useState } from "react";
import { Folder } from "lucide-react";
import { useNavigate, useParams } from "react-router";
import { useFetchFolders } from "../hooks/hooks";
import { AlertCircle, Loader2 } from "lucide-react";
import { type FolderType } from "../store/interfaces";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";

const Folders = () => {
  const params = useParams();
  const navigate = useNavigate();
  const [folders, setFolders] = useState<FolderType[]>([]);
  const [folderId, setFolderId] = useState<string | undefined>(undefined);

  const { data, isLoading, error, isError } = useFetchFolders(folderId);

  useEffect(() => {
    const folderIds = params["*"]?.split("/");
    if (folderIds && folderIds.length > 0) {
      const folderId = folderIds[folderIds.length - 1];
      setFolderId(folderId);
    }
  }, [params]);

  useEffect(() => {
    if (data && data.data) {
      setFolders(data.data.folders);
    }
  }, [data]);

  if (isLoading) {
    return (
      <div className="h-full w-full flex flex-col justify-center items-center gap-2">
        <Loader2 className="h-10 w-10 animate-spin text-primary" />
        <p className="font-medium">Loading folders...</p>
      </div>
    );
  }

  if (isError) {
    return (
      <div className="h-full w-full flex flex-col justify-center items-center p-4">
        <Alert variant="destructive" className="max-w-md">
          <AlertCircle className="h-4 w-4" />
          <AlertTitle>Error Fetching Folders</AlertTitle>
          <AlertDescription>
            {error instanceof Error ? error.message : "Something went wrong"}
          </AlertDescription>
        </Alert>
      </div>
    );
  }

  return (
    <div className="mt-4 grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-6 gap-6 w-[90%] mx-auto pb-8">
      {folders.map((f) => {
        return (
          <div
            key={f.id}
            className="flex flex-col justify-center items-center p-4 rounded-lg bg-card hover:bg-accent/50 transition-colors cursor-pointer border border-transparent hover:border-border"
            onClick={() => {
              const path = params["*"];
              if (!path) {
                navigate(`/files/${f.id}`);
              } else {
                navigate(`/files/${path}/${f.id}`);
              }
            }}
          >
            <Folder size={80} className="text-primary fill-primary/20" />
            <p className="font-medium mt-3 text-center truncate w-full">
              {f.name}
            </p>
            <p className="text-xs text-muted-foreground mt-1">
              12/12/2024
            </p>
          </div>
        );
      })}
    </div>
  );
};

export default Folders;