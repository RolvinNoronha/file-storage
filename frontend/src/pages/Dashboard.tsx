import { Button, Group, Input, Menu, Stack, Tabs } from "@mantine/core";
import { useAppTheme } from "../context/ThemeContext";
import { IoIosSearch } from "react-icons/io";
import { LuFile, LuFilter, LuFolder } from "react-icons/lu";
import Header from "../components/Header";
import Layout from "../components/Layout";
import Folders from "../components/Folders";
import Files from "../components/Files";

const Dashboard = () => {
  const { colors } = useAppTheme();

  return (
    <Layout header={<Header />}>
      <Stack mt={10} h={"94vh"} w={"100vw"}>
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

        <Tabs
          styles={{
            tab: {
              display: "flex",
              justifyContent: "center",
              color: colors.text,
              fontWeight: "bold",
            },
            tabLabel: {
              flex: 0,
            },
          }}
          h={"100%"}
          color={colors.primary}
          defaultValue="files"
        >
          <Tabs.List>
            <Tabs.Tab
              className="tab-button"
              w={"50%"}
              value="files"
              leftSection={<LuFile size={18} />}
            >
              Files
            </Tabs.Tab>
            <Tabs.Tab
              className="tab-button"
              w={"50%"}
              value="folders"
              leftSection={<LuFolder size={18} />}
            >
              Folders
            </Tabs.Tab>
          </Tabs.List>

          <Tabs.Panel h={"100%"} value="files">
            <Files />
          </Tabs.Panel>

          <Tabs.Panel h={"100%"} value="folders">
            <Folders />
          </Tabs.Panel>
        </Tabs>
      </Stack>
    </Layout>
  );
};

export default Dashboard;
