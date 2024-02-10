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
                <router-link to="/login">
                    Login
                </router-link>
            </li>
            <li class="mr-1 bg-white inline-block py-2 px-4 mt-0.5 mb-0.5 text-slate-600 hover:text-slate-800 font-semibold rounded-md">About Us</li>
        </ul>
    </div>
</template>