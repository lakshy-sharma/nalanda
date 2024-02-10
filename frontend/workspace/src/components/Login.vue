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

export default {
    name: "login",
    data() {
        return {
            email: "",
            password: ""
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
            fetch("http://localhost:8000/api/v1/login", requestOptions)
            .then((response) => response.json())
            .then((data) => {
                if (data.error) {
                    console.log("Error: ", data.message);
                } else {
                    console.log("Everything is fine.");
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