export interface Response {
  data: any;
  message: string;
  error: any;
  success: boolean;
}

export interface FileType {
  id: number;
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
  id: number;
  name: string;
  userId: number;
  createdAt: Date;
}

export interface User {
  userName: string;
  userId: string;
}

export type RequestMethod = "GET" | "POST" | "PATCH" | "PUT" | "DELTETE";

export interface InitiateUploadRequest {
  fileName: string;
  fileType: string;
  fileSize: number;
  folderId?: number;
}

export interface PresignedPart {
  partNumber: number;
  url: string;
}

export interface InitiateUploadResponse extends Response {
  data: {
    uploadId: string;
    key: string;
    parts: PresignedPart[];
  };
}

export interface CompletedPart {
  partNumber: number;
  etag: string;
}

export interface CompleteUploadRequest {
  uploadId: string;
  key: string;
  parts: CompletedPart[];
  fileName: string;
  fileSize: number;
  fileType: string;
  folderId?: number;
}
