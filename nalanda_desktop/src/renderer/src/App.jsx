import { useEffect, useState } from "react";
import { library } from '@fortawesome/fontawesome-svg-core'
import { fab } from '@fortawesome/free-brands-svg-icons'
import { fas } from '@fortawesome/free-solid-svg-icons'
import { far } from '@fortawesome/free-regular-svg-icons'
import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

import LoginPage from "./Components/LoginPage/Main";
import DashboardPage from "./Components/Dashboard/Main";

function App() {
  const [userLoginComplete, setUserLoginComplete] = useState(false);

  // Set the title of the document.
  useEffect(() => {
    document.title = "Nalanda"
  }, [])

  // Check whether the user has logged in previously or not.
  useEffect(() => {
    // The algorithm.
    // Call store and check whether teh user login complete flag is set or not.
    // setUserLoginComplete(true);
  });

  return (
    <div className="bg-[#FFFFF8] dark:bg-slate-800" >
      <ToastContainer />
      { userLoginComplete ? <DashboardPage /> : <LoginPage /> }
    </div>
  );
}

export default App;
library.add(fab, fas, far)