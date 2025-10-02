import Header from "./components/Header";
import Layout from "./components/Layout";

function App() {
  return (
    <Layout header={<Header />}>
      <div>home content</div>
    </Layout>
  );
}

export default App;
