import axios, { type AxiosRequestConfig } from "axios";
import type {
  FileType,
  FolderType,
  RequestMethod,
  Response,
  User,
} from "../store/interfaces";

interface LoginResponse extends Response {
  data: {
    user: User;
  };
}

interface GetFilesResponse extends Response {
  data: {
    files: FileType[];
  };
}

interface GetFolderResponse extends Response {
  data: {
    folders: FolderType[];
  };
}

const baseApiUrlv1 = import.meta.env.VITE_PUBLIC_API_URL_V1;

class AppServiceClass {
  private static instance: AppServiceClass;
  private baseUrl: string;
  private token: string | undefined;

  private getTokenFromCookie() {
    const cookies = document.cookie.split("; ");
    const token = cookies.find((cookie) => cookie.startsWith("token="));
    return token ? token.split("=")[1] : undefined;
  }

  private constructor() {
    this.token = this.getTokenFromCookie();
    this.baseUrl = baseApiUrlv1;
  }

  private async request<T>(
    url: string,
    headers: Record<string, string> | null,
    data: string | null,
    method: RequestMethod,
    formData?: FormData
  ): Promise<T> {
    const baseHeaders: Record<string, string> = {
      Authorization: `Bearer ${this.token}`,
      ...(headers || {}),
    };

    if (!formData) {
      baseHeaders["Content-Type"] = "application/json";
    }

    const reqOptions: AxiosRequestConfig = {
      url: `${this.baseUrl}${url}`,
      method: method,
      headers: baseHeaders,
    };

    if (method !== "GET" && data) {
      if (formData) {
        reqOptions.data = formData;
      } else {
        reqOptions.data = data;
      }
    }

    console.log(reqOptions);

    try {
      const response = await axios.request(reqOptions);
      return response.data;
    } catch (error) {
      console.error("Failed to make an API request: ", error);
      throw error;
    }
  }

  public static getInstance(): AppServiceClass {
    if (!AppServiceClass.instance) {
      AppServiceClass.instance = new AppServiceClass();
    }
    return AppServiceClass.instance;
  }

  register = async (username: string, password: string) => {
    return this.request<Response>(
      "/user/register",
      null,
      JSON.stringify({ username, password }),
      "POST"
    );
  };

  login = async (username: string, password: string) => {
    return this.request<LoginResponse>(
      "/user/login",
      null,
      JSON.stringify({ username, password }),
      "POST"
    );
  };

  getFiles = async (folderId?: string) => {
    return this.request<GetFilesResponse>(
      !folderId ? "/file/files" : `/file/files/${folderId}`,
      null,
      null,
      "GET"
    );
  };

  createFile = async (formData: FormData) => {
    return this.request<Response>("/file/create", null, null, "POST", formData);
  };

  getFolders = async (folderId?: string) => {
    return this.request<GetFolderResponse>(
      !folderId ? "/folder/folders" : `/folder/folders/${folderId}`,
      null,
      null,
      "GET"
    );
  };

  createFolder = async (folderName: string, folderId?: number) => {
    return this.request<Response>(
      "/folder/create",
      null,
      JSON.stringify({ folderName, folderId }),
      "POST"
    );
  };
}

const AppService = AppServiceClass.getInstance();

export default AppService;
