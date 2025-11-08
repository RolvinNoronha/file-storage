import { Alert, Grid, Loader, Text } from "@mantine/core";
import { useEffect, useState } from "react";
import {
  FaFileExcel,
  FaFileImage,
  FaFilePdf,
  FaFilePowerpoint,
  FaFileVideo,
  FaFileWord,
} from "react-icons/fa";
import { useAppTheme } from "../context/ThemeContext";
import { useParams } from "react-router";
import { useFetchFiles } from "../hooks/hooks";
import { IoIosAlert } from "react-icons/io";
import { type FileType } from "../store/interfaces";

const Files = () => {
  const params = useParams();
  const { colors } = useAppTheme();
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

  if (isLoading) {
    return (
      <div className="h-full w-full flex flex-col justify-center items-center">
        <Loader color={colors.primary} />
        <Text c={colors.text}>Loading files...</Text>
      </div>
    );
  }

  if (isError) {
    return (
      <div className="h-full w-full flex flex-col justify-center items-center">
        <Alert
          variant="light"
          color="red"
          styles={{
            message: {
              color: colors.text,
            },
          }}
          title="Error Fetching Files"
          icon={<IoIosAlert />}
        >
          {error.message}
        </Alert>
      </div>
    );
  }

  return (
    <Grid mt={10} gutter={"lg"} w={"90%"} mx={"auto"}>
      {files.map((f) => {
        return (
          <Grid.Col span={2} key={f.id}>
            <div
              style={{ backgroundColor: colors.background2 }}
              className="flex flex-col justify-center items-center p-2 rounded-lg hover:cursor-pointer"
            >
              {getIcon(f.type)}
              <Text fw={500} my={10} c={colors.text}>
                {f.name + " - 12/12/2024"}
              </Text>
            </div>
          </Grid.Col>
        );
      })}
    </Grid>
  );
};

export default Files;
