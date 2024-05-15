import {Route, Routes} from "react-router-dom";
import HomeContent from "./Content/Home";
import AboutContent from "./Content/About";
import PageNotFound from "../Base/PageNotFound";
import LoginForm from "../Base/Forms/Login";

function ContentBox() {
    return (
        <main className="container md:mx-auto text-justify shadow-md flex-grow bg-slate-200 dark:bg-slate-500 mb-5">
            <Routes>
                <Route exact path="/" element={<HomeContent />} />
                <Route path="/aboutme" element={<AboutContent />} />
                <Route path="/login" element={<LoginForm />} /> 
                <Route path="*" element={<PageNotFound />}/>
            </Routes>
        </main>
    )
}

export default ContentBox