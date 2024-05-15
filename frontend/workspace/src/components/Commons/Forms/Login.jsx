import { useState } from "react";
import { useStore } from "../store";
  
function LoginForm() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    let { token, user } = useStore();

    const handleChange = (event) => {
        setInputValue(event.target.value);
    };

    const handleSubmit = (event) => {
        event.preventDefault();
        const payload = {
            email: email,
            password: password
        }
        fetch(import.meta.env.VITE_BACKEND_SERVER + "/api/v1/auth/signin", Security.postRequestOptions(payload))
        .then((response) => {
            response.json()
            .then((body) => {
                if (body.status === "success") {
                    token = body.data.access_token;
                    user = {
                        id: body.data.user.id,
                        first_name: body.data.user.first_name
                    }
                    let expiryDays = 1;
                    document.cookie = `logged_in=true; max-age=${(expiryDays * 24 * 60 * 60 * 1000)}`
                    document.cookie = `access_token=${JSON.stringify(body.data.access_token)}; max-age=${(expiryDays * 24 * 60 * 60 * 1000)}`
                    router.push("/")
                    Toastify({
                        text: "Login Success",
                        duration: 500,
                        close: true
                    }).showToast();
                } else {
                    Toastify({
                        text: "Login Failed",
                        duration: 1000,
                        close: true
                    }).showToast();
                }
            })
        })
    }

    return (
        <div className="container md:mx-auto md:my-4 text-center shadow-md">
            <h1 className="text-3xl font-semibold my-2">Login</h1>
            <form method="POST" action="/login" name="loginform" onSubmit={handleSubmit}>
                <div className="grid grid-cols-2 gap-2 justify-items-auto">
                    <label className="mx-2 my-1 justify-self-end">Email: </label>
                    <input type="email" name="email" value={email} onChange={handleChange} required={true} className="mx-1 my-1 border-2 text-gray-700 max-w-xs"/>
                    <label className="mx-2 my-1 justify-self-end">Password: </label>
                    <input type="password" name="password" value={email} onChange={handleChange} required={true} className="mx-1 my-1 border-2 text-gray-700 max-w-xs"/>
                    <button type="submit" className="col-span-2 inline-block bg-slate-200 font-semibold text-slate-600 hover:text-slate-800 rounded-md py-1 px-3 mx-auto my-2 max-w-sm">Login</button>
                </div>
            </form>
        </div>
    )
}

export default LoginForm