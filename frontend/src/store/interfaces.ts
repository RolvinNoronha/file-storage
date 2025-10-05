export interface Response {
  data: any;
  message: string;
  error: any;
}

export interface FileType {
  name: string;
  path: string;
  type: string;
  size: number;
  userId: number;
  folderId: number;
  createdAt: Date;
}

export interface FileUrlType {
  fileUrl: string;
  fildId: number;
}

export interface FolderType {
  name: string;
  userId: number;
  createdAt: Date;
}

export interface User {
  userName: string;
  userId: string;
}

export type RequestMethod = "GET" | "POST" | "PATCH" | "PUT" | "DELTETE";
