import { useEffect } from "react";
import { library } from '@fortawesome/fontawesome-svg-core'
import { fab } from '@fortawesome/free-brands-svg-icons'
import { fas } from '@fortawesome/free-solid-svg-icons'
import { far } from '@fortawesome/free-regular-svg-icons'
import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";


import Header from './components/Commons/Header';
import Footer from "./components/Commons/Footer";
import ContentBox from "./components/Content/ContentBox";


function App() {
  // Set the title of the document.
  useEffect(() => {
    document.title = "Serene Link"
  }, [])

  return (
    <div className="bg-[#FFFFF8] dark:bg-slate-800" >
      <ToastContainer />
      <Header/>
      <ContentBox />
      <Footer/>
    </div>
  );
}

export default App;
library.add(fab, fas, far)