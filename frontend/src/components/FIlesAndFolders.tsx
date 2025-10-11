// @ts-nocheck
import { Card, Grid, Loader, Text } from "@mantine/core";
import { useEffect, useState } from "react";
import {
  FaFileExcel,
  FaFileImage,
  FaFilePdf,
  FaFilePowerpoint,
  FaFileVideo,
  FaFileWord,
  FaFolder,
} from "react-icons/fa";
import { useAppTheme } from "../context/ThemeContext";
import { useNavigate, useParams } from "react-router";

const getFilesAndFolders = (folderId = null) => {
  return new Promise((resolve) => {
    setTimeout(() => {
      // Data (same as provided in the original question)
      const data = [
        {
          id: 1,
          type: "file",
          fileType: "application/pdf",
          fileName: "pdf.pdf",
        },
        {
          id: 2,
          type: "file",
          fileType: "application/vnd.ms-powerpoint",
          fileName: "presentation.pptx",
        },
        {
          id: 3,
          type: "folder",
          folderName: "Folder 1",
          files: [
            {
              id: 1,
              type: "file",
              fileType:
                "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
              fileName: "sheet.xlsx",
            },
            {
              id: 2,
              type: "file",
              fileType:
                "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
              fileName: "word.docx",
            },
            {
              id: 3,
              type: "folder",
              folderName: "Folder 2",
              files: [
                {
                  id: 1,
                  type: "file",
                  fileType: "application/pdf",
                  fileName: "pdf.pdf",
                },
                {
                  id: 2,
                  type: "folder",
                  folderName: "Folder 3",
                  files: [
                    {
                      id: 1,
                      type: "file",
                      fileType: "audio/mpeg",
                      fileName: "audio.mp3",
                    },
                    {
                      id: 2,
                      type: "folder",
                      folderName: "Folder 4",
                      files: [
                        {
                          id: 1,
                          type: "file",
                          fileType: "application/json",
                          fileName: "data.json",
                        },
                        {
                          id: 2,
                          type: "file",
                          fileType: "text/plain",
                          fileName: "readme.txt",
                        },
                      ],
                    },
                  ],
                },
              ],
            },
            {
              id: 4,
              type: "file",
              fileType: "application/zip",
              fileName: "archive.zip",
            },
          ],
        },
        {
          id: 4,
          type: "file",
          fileType: "video/mp4",
          fileName: "video.mp4",
        },
        {
          id: 5,
          type: "folder",
          folderName: "Folder 5",
          files: [
            {
              id: 1,
              type: "file",
              fileType: "image/jpeg",
              fileName: "image.jpg",
            },
            {
              id: 2,
              type: "folder",
              folderName: "Folder 6",
              files: [
                {
                  id: 1,
                  type: "file",
                  fileType: "application/msword",
                  fileName: "document.doc",
                },
                {
                  id: 2,
                  type: "file",
                  fileType: "audio/wav",
                  fileName: "sound.wav",
                },
                {
                  id: 3,
                  type: "folder",
                  folderName: "Folder 7",
                  files: [
                    {
                      id: 1,
                      type: "file",
                      fileType: "application/vnd.ms-excel",
                      fileName: "spreadsheet.xls",
                    },
                    {
                      id: 2,
                      type: "file",
                      fileType: "video/avi",
                      fileName: "video.avi",
                    },
                  ],
                },
              ],
            },
            {
              id: 3,
              type: "file",
              fileType: "application/vnd.ms-powerpoint",
              fileName: "presentation.pptx",
            },
          ],
        },
        {
          id: 6,
          type: "file",
          fileType: "image/gif",
          fileName: "animation.gif",
        },
        {
          id: 7,
          type: "file",
          fileType:
            "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
          fileName: "sheet.xlsx",
        },
        {
          id: 8,
          type: "file",
          fileType:
            "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
          fileName: "word.docx",
        },
        {
          id: 9,
          type: "file",
          fileType: "image/gif",
          fileName: "animation.gif",
        },
        {
          id: 10,
          type: "file",
          fileType: "image/gif",
          fileName: "animation.gif",
        },
      ];

      // If a folder ID is provided, find the folder and return its contents
      if (folderId) {
        const findFolder = (id: number, items: any) => {
          for (const item of items) {
            if (item.id === id && item.type === "folder") {
              return item.files;
            }
            if (item.type === "folder" && item.files) {
              const folder: any = findFolder(id, item.files);
              if (folder) return folder;
            }
          }
          return null;
        };

        // Fetch the folder data using the ID
        const folderContents = findFolder(folderId, data);
        resolve(folderContents || []);
      } else {
        // If no folder ID is provided, return all files and folders on the first level
        const firstLevelFilesAndFolders = data.filter(
          (item) => item.type === "file" || item.type === "folder"
        );
        resolve(firstLevelFilesAndFolders);
      }
    }, 2000); // Simulate a delay of 500ms
  });
};

const FilesAndFolders = () => {
  const params = useParams();
  const navigate = useNavigate();
  const { colors } = useAppTheme();
  const [filesFolders, setFilesFolders] = useState([]);
  const [loading, setLoading] = useState<boolean>(false);

  const getData = async () => {
    setLoading(true);
    const result = await getFilesAndFolders();
    setFilesFolders(result);
    setLoading(false);
  };

  useEffect(() => {
    const folderIds = params["*"]?.split("/");
    if (folderIds && folderIds.length > 0) {
      const folderId = folderIds[folderIds.length - 1];
      getData(Number(folderId));
      console.log(params);
    } else {
      getData();
    }
  }, [params]);

  const getIcon = (fileType) => {
    if (fileType.includes("image")) {
      return <FaFileImage size={100} color={colors.secondary} />;
    } else if (fileType.includes("pdf")) {
      return <FaFilePdf size={100} color="#E60012" />;
    } else if (fileType.includes("video")) {
      return <FaFileVideo size={100} color={colors.secondary} />;
    } else if (fileType.includes("sheet")) {
      return <FaFileExcel size={100} color="#217346" />;
    } else if (fileType.includes("powerpoint")) {
      return <FaFilePowerpoint size={100} color="#D24B00" />;
    }

    return <FaFileWord size={100} color="#2B579A" />;
  };

  if (loading) {
    return (
      <div className="h-full w-full flex flex-col justify-center items-center">
        <Loader color={colors.primary} />
        <Text c={colors.text}>Loading data...</Text>
      </div>
    );
  }

  return (
    <Grid mt={10} gutter={"lg"} w={"90%"} mx={"auto"}>
      {filesFolders.map((f) => {
        return (
          <Grid.Col span={2} key={f.id}>
            {f.type === "file" ? (
              <div
                style={{ backgroundColor: colors.background2 }}
                className="flex flex-col justify-center items-center p-2 rounded-lg hover:cursor-pointer"
              >
                {getIcon(f.fileType)}
                <Text fw={500} my={10} c={colors.text}>
                  {f.fileName + " - 12/12/2024"}
                </Text>
              </div>
            ) : (
              <div
                style={{ backgroundColor: colors.background2 }}
                className="flex flex-col justify-center items-center p-2 rounded-lg hover:cursor-pointer"
                onClick={() => {
                  const path = params["*"];
                  if (!path) {
                    navigate(`/files/${f.id}`);
                  } else {
                    navigate(`/files/${path}/${f.id}`);
                  }
                }}
              >
                <FaFolder size={100} color={colors.primary} />
                <Text fw={500} my={10} c={colors.text}>
                  {f.folderName + " - 12/12/2024"}
                </Text>
              </div>
            )}
          </Grid.Col>
        );
      })}
    </Grid>
  );
};

export default FilesAndFolders;
