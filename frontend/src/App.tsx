import { Shield, Upload, Zap, Users } from "lucide-react";
import { useNavigate } from "react-router";
import heroImage from "./assets/hero.jpg";
import Header from "./components/Header";
import Layout from "./components/Layout";
import { Button } from "@/components/ui/button";

function App() {
  const navigate = useNavigate();

  const features = [
    {
      icon: <Upload size={30} className="text-primary" />,
      title: "Easy Upload",
      description:
        "Drag and drop your files or click to upload. Support for all file types.",
    },
    {
      icon: <Shield size={30} className="text-primary" />,
      title: "Secure Storage",
      description:
        "Your files are encrypted and stored securely with enterprise-grade protection.",
    },
    {
      icon: <Zap size={30} className="text-primary" />,
      title: "Lightning Fast",
      description:
        "Access your files instantly with our optimized cloud infrastructure.",
    },
    {
      icon: <Users size={30} className="text-primary" />,
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
                <h1 className="text-5xl lg:text-7xl font-bold leading-tight bg-gradient-to-r from-primary to-primary/60 bg-clip-text text-transparent">
                  Your Files, Everywhere
                </h1>
                <p className="text-xl text-muted-foreground">
                  Store, organize, and share your files with our modern cloud
                  platform. Access your documents from anywhere, anytime.
                </p>
              </div>

              <div className="flex flex-col sm:flex-row gap-4">
                <Button
                  size="lg"
                  className="text-lg px-8 py-6 rounded-xl"
                  onClick={() => navigate("/auth")}
                >
                  Get Started Free
                </Button>
                <Button
                  variant="outline"
                  size="lg"
                  className="text-lg px-8 py-6 rounded-xl border-primary/20"
                >
                  View Demo
                </Button>
              </div>
            </div>

            <div className="relative">
              <div className="absolute inset-0 bg-primary/20 blur-3xl rounded-full"></div>
              <img
                src={heroImage}
                alt="File management interface"
                className="relative rounded-2xl shadow-2xl w-full h-auto border"
              />
            </div>
          </div>
        </div>

        {/* Features Section */}
        <div className="container mx-auto px-6 py-20">
          <div className="text-center mb-16 space-y-4">
            <h2 className="text-3xl font-bold">Why Choose Our Platform?</h2>
            <p className="text-xl text-muted-foreground max-w-2xl mx-auto">
              Experience the next generation of file management with powerful
              features designed for modern workflows.
            </p>
          </div>

          <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-8">
            {features.map((feature, index) => (
              <div
                key={index}
                className="bg-card p-8 rounded-2xl border shadow-sm hover:shadow-md transition-all duration-300 group"
              >
                <div className="mb-6 relative">
                  <div className="w-16 h-16 bg-primary/10 rounded-2xl flex items-center justify-center group-hover:scale-110 transition-transform duration-300">
                    {feature.icon}
                  </div>
                </div>
                <h3 className="text-xl font-bold mb-3">{feature.title}</h3>
                <p className="text-muted-foreground leading-relaxed">
                  {feature.description}
                </p>
              </div>
            ))}
          </div>
        </div>

        {/* CTA Section */}
        <div className="border-y bg-muted/30">
          <div className="container mx-auto px-6 py-20">
            <div className="text-center max-w-3xl mx-auto space-y-6">
              <h2 className="text-3xl font-bold">Ready to Get Started?</h2>
              <p className="text-lg text-muted-foreground">
                Join thousands of users who trust us with their files. Sign up
                today and get 10GB free storage.
              </p>
              <Button
                size="lg"
                className="text-lg px-12 py-6 rounded-xl"
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