<template>
    <div class="container md:mx-auto shadow-md bg-gray-300 dark:bg-gray-700 mb-5 pt-5">
        <header class="container text-center">
            <h1 class="text-4xl font-medium">Nalanda</h1>
            <p>
                <i>The Open Research Center</i><br />
            </p>
        </header>
        <ul class="container md:mx-auto p-6 items-center justify-center md:flex">
            <li class="mr-1 bg-white inline-block py-2 px-4 mt-0.5 mb-0.5 text-slate-600 hover:text-slate-800 font-semibold rounded-md">
                <router-link to="/">
                    Home
                </router-link>
            </li>
            <li class="mr-1 bg-white inline-block py-2 px-4 mt-0.5 mb-0.5 text-slate-600 hover:text-slate-800 font-semibold rounded-md">About Author</li>
            <li class="mr-1 bg-white inline-block py-2 px-4 mt-0.5 mb-0.5 text-slate-600 hover:text-slate-800 font-semibold rounded-md">Intel Search</li>
            <li class="mr-1 bg-white inline-block py-2 px-4 mt-0.5 mb-0.5 text-slate-600 hover:text-slate-800 font-semibold rounded-md">
                <router-link to="/books">Books</router-link>
            </li>
            <li v-if="store.token !== ''" class="mr-1 bg-white inline-block py-2 px-4 mt-0.5 mb-0.5 text-slate-600 hover:text-slate-800 font-semibold rounded-md">
                <Menu as="div" class="relative inline-block text-left">
                    <div>
                        <MenuButton class="inline-flex">
                            Admin Utils
                            <ChevronDownIcon class="h-5 w-5 text-gray-400" aria-hidden="true" />
                        </MenuButton>
                    </div>
                    <transition enter-active-class="transition ease-out duration-100" enter-from-class="transform opacity-0 scale-95" enter-to-class="transform opacity-100 scale-100" leave-active-class="transition ease-in duration-75" leave-from-class="transform opacity-100 scale-100" leave-to-class="transform opacity-0 scale-95">
                        <MenuItems class="absolute right-0 z-10 mt-2 w-56 origin-top-right rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
                            <div class="py-1">
                                <MenuItem v-slot="{ active }">
                                    <router-link :class="[active ? 'bg-gray-100 text-gray-900' : 'text-gray-700', 'block px-4 py-2 text-sm']" to="/admin/users">Manage Users</router-link>
                                </MenuItem>
                                <MenuItem v-slot="{ active }">
                                    <router-link :class="[active ? 'bg-gray-100 text-gray-900' : 'text-gray-700', 'block px-4 py-2 text-sm']" to="/admin/users/0">Add User</router-link>
                                </MenuItem>
                                <MenuItem v-slot="{ active }">
                                    <router-link :class="[active ? 'bg-gray-100 text-gray-900' : 'text-gray-700', 'block px-4 py-2 text-sm']" to="/admin/books">Manage Books</router-link>
                                </MenuItem>
                                <MenuItem v-slot="{ active }">
                                    <router-link :class="[active ? 'bg-gray-100 text-gray-900' : 'text-gray-700', 'block px-4 py-2 text-sm']" to="/admin/books/0">Add Book</router-link>
                                </MenuItem>
                            </div>
                        </MenuItems>
                    </transition>
                </Menu>
            </li>

            <li class="mr-1 bg-white inline-block py-2 px-4 mt-0.5 mb-0.5 text-slate-600 hover:text-slate-800 font-semibold rounded-md">
                <router-link v-if="store.token === ''" to="/login">
                    Login
                </router-link>
                <a href="javascript:void(0);" v-else @click="logout" >Logout</a>
            </li>
        </ul>
    </div>
</template>

<script setup>
import { Menu, MenuButton, MenuItem, MenuItems } from '@headlessui/vue'
import { ChevronDownIcon } from '@heroicons/vue/20/solid'
</script>

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
            const requestOptions = {
                method: "POST",
                headers: new Headers({
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${store.token}`,
                })
            }
            fetch(import.meta.env.VITE_BACKEND_SERVER + "/api/v1/auth/logout", requestOptions)
            .then((response) => response.json())
            .then((response) => {
                if (response) {
                    if (response.status === "success") {
                        store.token = ""
                        store.user = {}
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
    }
}
</script>