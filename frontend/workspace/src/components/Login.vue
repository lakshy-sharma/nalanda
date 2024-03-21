<template>
    <div class="container md:mx-auto md:my-4 text-center shadow-md">
        <h1 class="text-3xl font-semibold my-2">Login</h1>
        <form-base method="POST" action="/login" name="loginform" event="submitEvent" v-on:submitEvent="submitHandler" >
            <!--Simple Username and Password based Login.-->
            <text-input 
                v-model="email"
                label="Email" 
                type="email" 
                name="email" 
                required="true"
            />
            <text-input 
                v-model="password"
                label="Password" 
                type="password" 
                name="password" 
                required="true"
            />
            <!--Final Submit Button For Login Form-->
            <div class="flex-grid grid-cols-3">
                <input type="submit" class="col-start-2 inline-block bg-slate-200 font-semibold text-slate-600 hover:text-slate-800 rounded-md py-1 px-3 my-2 max-w-sm" value="Login">
            </div>

        </form-base>
    </div>
</template>

<script>
import FormBase from "./forms/FormBase.vue";
import TextInput from "./forms/TextInput.vue";
import store from "./store.js";
import router from "./../router/index.js";
import Toastify from 'toastify-js'

export default {
    name: "login",
    data() {
        return {
            email: "",
            password: "",
            store,
        }
    },
    methods: {
        submitHandler() {
            console.log("Submission made.");
            const payload = {
                email: this.email,
                password: this.password
            }
            const requestOptions = {
                method: "POST",
                body: JSON.stringify(payload)
            }
            fetch("http://localhost:8000/api/v1/auth/signin", requestOptions)
            .then((response) => response.json())
            .then((response) => {
                if (response.status === "fail") {
                    console.log("Failure: ", response.message);
                    Toastify({
                        text: "Login Failed",
                        duration: 1000,
                        close: true
                    }).showToast();
                } else {
                    store.token = response.access_token;
                    store.user = {
                        id: response.user.id,
                        first_name: response.user.first_name
                    }
                    let expiryDays = 1;
                    document.cookie = `logged_in=true; max-age=${(expiryDays * 24 * 60 * 60 * 1000)}`
                    document.cookie = `access_token=${JSON.stringify(response.access_token)}; max-age=${(expiryDays * 24 * 60 * 60 * 1000)}`
                    router.push("/")
                    console.log("Login Complete")
                }
            })
        }
    },
    components: {
        TextInput,
        FormBase
    },
}
</script>