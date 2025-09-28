import { Button, Group, Modal, Stack, TextInput, Title } from "@mantine/core";
import { Link } from "react-router";
import { useAppTheme } from "../context/ThemeContext";
import { MdDarkMode, MdLightMode } from "react-icons/md";
import { useAuth } from "../context/AuthContext";
import { useDisclosure } from "@mantine/hooks";

const Header = () => {
  const { isAuthenticated, logout } = useAuth();
  const { theme, toggleTheme, colors } = useAppTheme();
  const [opened, { open, close }] = useDisclosure(false);

  return (
    <Stack h={"10vh"} justify="center">
      <Group px={"md"} justify="space-between">
        <Title c={colors.text}>File Uploader</Title>
        <Group>
          {isAuthenticated ? (
            <>
              <div className="flex items-center justify-center">
                <input
                  type="file"
                  id="fileInput"
                  className="hidden"
                  accept="*/*"
                  onChange={() => {}}
                />
                <label
                  htmlFor="fileInput"
                  className="cursor-pointer text-white font-semibold px-6 py-2 rounded-sm focus:outline-none focus:ring-2"
                  style={{ background: colors.primary }}
                >
                  Choose File
                </label>
              </div>
              <Button size="md" color={colors.primary} onClick={open}>
                Add Folder
              </Button>
              <Button
                size="md"
                variant="outline"
                color={"red"}
                onClick={logout}
              >
                Log Out
              </Button>
            </>
          ) : (
            <Link to={"/auth"}>
              <Button size="md" color={colors.primary}>
                Get Started
              </Button>
            </Link>
          )}
          {theme === "dark" ? (
            <MdDarkMode
              color={colors.text}
              size={30}
              className="hover:cursor-pointer"
              onClick={toggleTheme}
            />
          ) : (
            <MdLightMode
              color={colors.text}
              size={30}
              className="hover:cursor-pointer"
              onClick={toggleTheme}
            />
          )}
        </Group>
      </Group>

      <div
        style={{
          backgroundColor:
            theme === "light"
              ? "rgba(0, 0, 0, 0.2)"
              : "rgba(255, 255, 255, 0.2)",
        }}
        className="h-[1px] w-full"
      ></div>
      <Modal
        opened={opened}
        centered
        onClose={close}
        withCloseButton={false}
        closeOnClickOutside={false}
        title="Folder Name"
        styles={{
          title: {
            color: colors.text,
          },
          content: {
            backgroundColor: colors.background1,
          },
          header: {
            backgroundColor: colors.background1,
          },
        }}
      >
        <TextInput
          my={10}
          c={colors.text}
          placeholder="enter folder name"
          variant="filled"
          styles={{
            input: {
              backgroundColor: colors.background3,
              color: colors.text,
            },
          }}
        />
        <Group justify="center">
          <Button
            variant="outline"
            my={10}
            size="md"
            color={colors.primary}
            onClick={close}
          >
            Cancel
          </Button>
          <Button my={10} size="md" color={colors.primary}>
            Create
          </Button>
        </Group>
      </Modal>
    </Stack>
  );
};

export default Header;
