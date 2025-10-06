import { FaShieldAlt, FaUpload, FaUsers } from "react-icons/fa";
import { LuZap } from "react-icons/lu";
import { useAppTheme } from "./context/ThemeContext";
import { Button, Card, Text, Title } from "@mantine/core";
import { useNavigate } from "react-router";
import heroImage from "./assets/hero.jpg";
import Header from "./components/Header";
import Layout from "./components/Layout";

function App() {
  const { colors } = useAppTheme();
  const navigate = useNavigate();

  const features = [
    {
      icon: <FaUpload size={30} color={colors.primary} />,
      title: "Easy Upload",
      description:
        "Drag and drop your files or click to upload. Support for all file types.",
    },
    {
      icon: <FaShieldAlt size={30} color={colors.primary} />,
      title: "Secure Storage",
      description:
        "Your files are encrypted and stored securely with enterprise-grade protection.",
    },
    {
      icon: <LuZap size={30} color={colors.primary} />,
      title: "Lightning Fast",
      description:
        "Access your files instantly with our optimized cloud infrastructure.",
    },
    {
      icon: <FaUsers size={30} color={colors.primary} />,
      title: "Easy Sharing",
      description:
        "Share files with teammates and control access permissions effortlessly.",
    },
  ];

  return (
    <Layout header={<Header />}>
      {/* Hero Section */}
      <main className="relative">
        <div className="container mx-auto px-6 py-20">
          <div className="grid lg:grid-cols-2 gap-12 items-center">
            <div className="space-y-8">
              <div className="space-y-4">
                <h1
                  style={{ backgroundColor: colors.primary }}
                  className="text-5xl lg:text-7xl font-bold  bg-clip-text text-transparent leading-tight"
                >
                  Your Files, Everywhere
                </h1>
                <Text c={colors.text} size="xl">
                  Store, organize, and share your files with our modern cloud
                  platform. Access your documents from anywhere, anytime.
                </Text>
              </div>

              <div className="flex flex-col sm:flex-row gap-4">
                <Button
                  color={colors.primary}
                  size="lg"
                  className="hover:shadow-glow transition-all duration-300 text-lg px-8 py-6"
                  onClick={() => navigate("/auth")}
                >
                  Get Started Free
                </Button>
                <Button
                  color={colors.primary}
                  variant="outline"
                  size="lg"
                  className="text-lg px-8 py-6 border-primary/20 hover:bg-primary/5"
                >
                  View Demo
                </Button>
              </div>
            </div>

            <div className="relative">
              <div className="absolute inset-0 bg-gradient-primary opacity-20 blur-3xl"></div>
              <img
                src={heroImage}
                alt="File management interface"
                className="relative rounded-2xl shadow-card w-full h-auto"
              />
            </div>
          </div>
        </div>

        {/* Features Section */}
        <div className="container mx-auto px-6 py-20">
          <div className="text-center mb-16">
            <Title order={1} c={colors.text}>
              Why Choose Our Platform?
            </Title>
            <Text size="xl" c={colors.text}>
              Experience the next generation of file management with powerful
              features designed for modern workflows.
            </Text>
          </div>

          <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-8">
            {features.map((feature, index) => (
              <Card
                bg={colors.background2}
                key={index}
                className="border-0 shadow-soft hover:shadow-card transition-all duration-300 group"
              >
                <Card.Section className="p-8 text-center">
                  <div className="mb-6 relative">
                    <div className="w-16 h-16 mx-auto bg-gradient-primary rounded-2xl flex items-center justify-center group-hover:scale-110 transition-transform duration-300">
                      {feature.icon}
                    </div>
                  </div>
                  <Title c={colors.text} order={3}>
                    {feature.title}
                  </Title>
                  <Text size="lg" c={colors.text}>
                    {feature.description}
                  </Text>
                </Card.Section>
              </Card>
            ))}
          </div>
        </div>

        {/* CTA Section */}
        <div className="border-y border-primary/10">
          <div className="container mx-auto px-6 py-20">
            <div className="text-center max-w-3xl mx-auto">
              <Title order={2} c={colors.text}>
                Ready to Get Started?
              </Title>
              <Text my={10} c={colors.text} size="lg">
                Join thousands of users who trust us with their files. Sign up
                today and get 10GB free storage.
              </Text>
              <Button
                color={colors.primary}
                size="md"
                className="hover:shadow-glow transition-all duration-300 text-lg px-12 py-6"
                onClick={() => navigate("/auth")}
              >
                Start Your Journey
              </Button>
            </div>
          </div>
        </div>
      </main>
    </Layout>
  );
}

export default App;
