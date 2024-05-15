import {Route, Routes} from "react-router-dom";
import HomeContent from "./Home";
import AboutContent from "./About";
import RecognitionContent from "./Recognition";
import ProductsContent from "./Products"
import PageNotFound from "../Commons/PageNotFound";
import LoginForm from "../Commons/Forms/Login";

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