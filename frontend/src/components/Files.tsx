import { useEffect, useState } from "react";
import {
  FileImage,
  FileVideo,
  FileText,
  FileSpreadsheet,
  FileArchive,
  File as FileIcon,
  AlertCircle,
  Loader2,
} from "lucide-react";
import { useParams } from "react-router";
import { useFetchFiles } from "../hooks/hooks";
import { type FileType } from "../store/interfaces";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";

const Files = () => {
  const params = useParams();
  const [folderId, setFolderId] = useState<string | undefined>(undefined);
  const [files, setFiles] = useState<FileType[]>([]);
  const { data, isLoading, isError, error } = useFetchFiles(folderId);

  useEffect(() => {
    const folderIds = params["*"]?.split("/");
    if (folderIds && folderIds.length > 0) {
      const folderId = folderIds[folderIds.length - 1];
      setFolderId(folderId);
    }
  }, [params]);

  useEffect(() => {
    if (data && data.data) {
      setFiles(data.data.files);
    }
  }, [data]);

  const getIcon = (fileType: string) => {
    const size = 80;
    if (fileType.includes("image")) {
      return <FileImage size={size} className="text-blue-500" />;
    } else if (fileType.includes("pdf")) {
      return <FileText size={size} className="text-red-500" />;
    } else if (fileType.includes("video")) {
      return <FileVideo size={size} className="text-purple-500" />;
    } else if (fileType.includes("sheet") || fileType.includes("excel")) {
      return <FileSpreadsheet size={size} className="text-green-600" />;
    } else if (fileType.includes("presentation") || fileType.includes("powerpoint")) {
      return <FileArchive size={size} className="text-orange-600" />;
    }

    return <FileIcon size={size} className="text-gray-500" />;
  };

  if (isLoading) {
    return (
      <div className="h-full w-full flex flex-col justify-center items-center gap-2">
        <Loader2 className="h-10 w-10 animate-spin text-primary" />
        <p className="font-medium">Loading files...</p>
      </div>
    );
  }

  if (isError) {
    return (
      <div className="h-full w-full flex flex-col justify-center items-center p-4">
        <Alert variant="destructive" className="max-w-md">
          <AlertCircle className="h-4 w-4" />
          <AlertTitle>Error Fetching Files</AlertTitle>
          <AlertDescription>
            {error instanceof Error ? error.message : "Something went wrong"}
          </AlertDescription>
        </Alert>
      </div>
    );
  }

  return (
    <div className="mt-4 grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-6 gap-6 w-[90%] mx-auto pb-8">
      {files.map((f) => {
        return (
          <div
            key={f.id}
            className="flex flex-col justify-center items-center p-4 rounded-lg bg-card hover:bg-accent/50 transition-colors cursor-pointer border border-transparent hover:border-border"
          >
            {getIcon(f.type)}
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

export default Files;