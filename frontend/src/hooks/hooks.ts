import { useQuery } from "@tanstack/react-query";
import AppService from "../service/AppService";

export const useFetchFiles = (folderId?: string) =>
  useQuery({
    queryKey: ["files", folderId],
    queryFn: () => AppService.getFiles(folderId),
  });

export const useFetchFolders = (folderId?: string) =>
  useQuery({
    queryKey: ["folderId", folderId],
    queryFn: () => AppService.getFolders(folderId),
  });
