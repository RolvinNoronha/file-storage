import { Search, Filter, File, Folder } from "lucide-react";
import Header from "../components/Header";
import Layout from "../components/Layout";
import Folders from "../components/Folders";
import Files from "../components/Files";
import { Input } from "@/components/ui/input";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Button } from "@/components/ui/button";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

const Dashboard = () => {
  return (
    <Layout header={<Header />}>
      <div className="mt-2.5 h-[94vh] w-screen flex flex-col gap-4">
        <div className="w-full flex flex-row justify-center gap-4 px-4">
          <div className="relative w-[60%]">
            <Input
              className="w-full h-12 text-lg bg-card border-primary pr-12"
              placeholder="Search Files"
            />
            <Search className="absolute right-4 top-1/2 -translate-y-1/2 text-muted-foreground" size={24} />
          </div>
          
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button
                variant="outline"
                size="lg"
                className="h-12 border-primary text-primary hover:text-primary"
              >
                Filter
                <Filter className="ml-2 h-4 w-4" />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end" className="w-[200px]">
              <DropdownMenuItem>Alphabetically (A-Z)</DropdownMenuItem>
              <DropdownMenuItem>Alphabetically (Z-A)</DropdownMenuItem>
              <DropdownMenuItem>Created At (asc)</DropdownMenuItem>
              <DropdownMenuItem>Created At (desc)</DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        </div>

        <Tabs defaultValue="files" className="w-full h-full">
          <TabsList className="w-full flex justify-center bg-transparent border-b rounded-none h-auto p-0">
            <TabsTrigger 
              value="files" 
              className="flex-1 py-3 font-bold flex items-center justify-center gap-2 data-[state=active]:border-b-2 data-[state=active]:border-primary rounded-none bg-transparent"
            >
              <File size={18} />
              Files
            </TabsTrigger>
            <TabsTrigger 
              value="folders" 
              className="flex-1 py-3 font-bold flex items-center justify-center gap-2 data-[state=active]:border-b-2 data-[state=active]:border-primary rounded-none bg-transparent"
            >
              <Folder size={18} />
              Folders
            </TabsTrigger>
          </TabsList>

          <TabsContent value="files" className="h-full mt-0">
            <Files />
          </TabsContent>

          <TabsContent value="folders" className="h-full mt-0">
            <Folders />
          </TabsContent>
        </Tabs>
      </div>
    </Layout>
  );
};

export default Dashboard;