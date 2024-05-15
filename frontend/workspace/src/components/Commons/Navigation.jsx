import {Link} from "react-router-dom";
import ThemeSwitcher from "./ThemeSwitcher";
import { useStore } from "./store.js"

function Navigator({location='header'}) {
    let { token, user } = useStore();

    const logout = (event) => {
        fetch(import.meta.env.VITE_BACKEND_SERVER + "/api/v1/auth/logout", Security.postRequestOptions(""))
        .then((response) => response.json())
        .then((response) => {
            if (response) {
                if (response.status === "success") {
                    token = ""
                    user = {}
                    document.cookie = `logged_in=false; max-age=100`
                    document.cookie = `access_token=-1; max-age=100`
                    console.log("Logout Completed");
                } else {
                    console.log(response.message);
                }
            } else {
                console.log(`Bad Response: ${response}`);
            }
        })
        router.push("/")
    }

    return (
        <ul className={location === "header" ? "container md:mx-auto p-6 items-center justify-center md:flex" : "container md:mx-auto p-6 items-center justify-center md:flex"}>
            <li className="mr-1 bg-white inline-block py-2 px-4 mt-0.5 mb-0.5 text-slate-600 hover:text-slate-800 font-semibold rounded-md"><Link to="/">Home</Link></li>
            <li className="mr-1 bg-white inline-block py-2 px-4 mt-0.5 mb-0.5 text-slate-600 hover:text-slate-800 font-semibold rounded-md"><Link to="/aboutme">About Me</Link></li>
            <li className="mr-1 bg-white inline-block py-2 px-4 mt-0.5 mb-0.5 text-slate-600 hover:text-slate-800 font-semibold rounded-md">
                {
                    token === "" ? <Link to="/login">Login</Link> : <a href="javascript:void(0);" onClick={logout}>Logout</a>
                }
            </li>
            {location === "header" ? <li className="mr-1 bg-white inline-block py-2 px-4 mt-0.5 mb-0.5 rounded-md"><ThemeSwitcher /></li> : <></>}

        </ul>
    )
}

export default Navigator