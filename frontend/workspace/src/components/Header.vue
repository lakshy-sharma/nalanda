<template>
    <div class="container md:mx-auto shadow-md bg-gray-300 dark:bg-gray-700 mb-5 pt-5">
        <header class="container text-center">
            <h1 class="text-4xl font-medium">Nalanda</h1>
            <p>
                <i>The Open Library</i><br />
            </p>
        </header>
        <ul class="container md:mx-auto p-6 items-center justify-center md:flex">
            <li class="mr-1 bg-white inline-block py-2 px-4 mt-0.5 mb-0.5 text-slate-600 hover:text-slate-800 font-semibold rounded-md">
                <router-link to="/">
                    Home
                </router-link>
            </li>
            <li class="mr-1 bg-white inline-block py-2 px-4 mt-0.5 mb-0.5 text-slate-600 hover:text-slate-800 font-semibold rounded-md">
                <router-link v-if="store.token === ''" to="/login">
                    Login
                </router-link>
                <a href="javascript:voide(0);" v-else @click="logout" >Logout</a>
            </li>
            <li class="mr-1 bg-white inline-block py-2 px-4 mt-0.5 mb-0.5 text-slate-600 hover:text-slate-800 font-semibold rounded-md">About Us</li>
        </ul>
    </div>
</template>

<script>
import store from "./store.js"
import router from "./../router/index.js"

export default {
    data() {
        return {
            store
        }
    },
    methods: {
        logout() {
            store.token = "";
            // Hit the logout.
            fetch("http://localhost:8000/api/v1/auth/logout")
            .then((response) => response.json())
            .then((response) => {
                if (response.status === "success") {
                    console.log("Logged out")
                }
            })
            router.push("/")
        }
    }
}
</script>