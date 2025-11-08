import { Alert, Grid, Loader, Text } from "@mantine/core";
import { useEffect, useState } from "react";
import { FaFolder } from "react-icons/fa";
import { useAppTheme } from "../context/ThemeContext";
import { useNavigate, useParams } from "react-router";
import { useFetchFolders } from "../hooks/hooks";
import { IoIosAlert } from "react-icons/io";
import { type FolderType } from "../store/interfaces";

const Folders = () => {
  const params = useParams();
  const navigate = useNavigate();
  const { colors } = useAppTheme();
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
      <div className="h-full w-full flex flex-col justify-center items-center">
        <Loader color={colors.primary} />
        <Text c={colors.text}>Loading folders...</Text>
      </div>
    );
  }

  if (isError) {
    return (
      <div className="h-full w-full flex flex-col justify-center items-center">
        <Alert
          variant="light"
          color="red"
          title="Error Fetching Folders"
          styles={{
            message: {
              color: colors.text,
            },
          }}
          icon={<IoIosAlert />}
        >
          {error.message}
        </Alert>
      </div>
    );
  }

  return (
    <Grid mt={10} gutter={"lg"} h={"100%"} w={"90%"} mx={"auto"}>
      {folders.map((f) => {
        return (
          <Grid.Col span={2} key={f.id}>
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
                {f.name + " - 12/12/2024"}
              </Text>
            </div>
          </Grid.Col>
        );
      })}
    </Grid>
  );
};

export default Folders;
