import { Button, Group, Input, Menu, Stack } from "@mantine/core";
import { useAppTheme } from "../context/ThemeContext";
import { IoIosSearch } from "react-icons/io";
import { LuFilter } from "react-icons/lu";
import Header from "../components/Header";
import Layout from "../components/Layout";
import FilesAndFolders from "../components/FIlesAndFolders";

const Dashboard = () => {
  const { colors } = useAppTheme();

  return (
    <Layout header={<Header />}>
      <Stack h={"90vh"} w={"100vw"}>
        <Group w={"100%"} justify="center">
          <Input
            w={"60%"}
            size="lg"
            variant="filled"
            placeholder="Search Files"
            styles={{
              input: {
                backgroundColor: colors.background2,
                color: colors.text,
                borderColor: colors.primary,
              },
            }}
            rightSection={<IoIosSearch size={30} />}
          />
          <Menu shadow="md" width={200}>
            <Menu.Target>
              <Button
                size="lg"
                rightSection={<LuFilter />}
                variant="outline"
                color={colors.primary}
              >
                Filter
              </Button>
            </Menu.Target>

            <Menu.Dropdown bg={colors.background2} bd={0}>
              <Menu.Item
                styles={{
                  item: {
                    backgroundColor: colors.background2,
                    color: colors.text,
                  },
                }}
              >
                {"Alphabetically (A-Z)"}
              </Menu.Item>
              <Menu.Item
                styles={{
                  item: {
                    backgroundColor: colors.background2,
                    color: colors.text,
                  },
                }}
              >
                {"Alphabetically (Z-A)"}
              </Menu.Item>
              <Menu.Item
                styles={{
                  item: {
                    backgroundColor: colors.background2,
                    color: colors.text,
                  },
                }}
              >
                {"Created At (asc)"}
              </Menu.Item>
              <Menu.Item
                styles={{
                  item: {
                    backgroundColor: colors.background2,
                    color: colors.text,
                  },
                }}
              >
                {"Created At (desc)"}
              </Menu.Item>
            </Menu.Dropdown>
          </Menu>
        </Group>
        <FilesAndFolders />
      </Stack>
    </Layout>
  );
};

export default Dashboard;
