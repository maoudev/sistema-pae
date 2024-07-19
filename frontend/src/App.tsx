import Home from "./pages/Home";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import AddStudent from "./pages/AddStudent";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home/>
  },
  {
    path: "/add",
    element: <AddStudent />
  }

])


function App() {
  return <RouterProvider router={router} />;
}

export default App;
