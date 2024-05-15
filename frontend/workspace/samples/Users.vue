<template>
        <div class="container md:mx-auto shadow-md text-center my-2 md:mx-auto">
            <h1 class="mt-3">
                All Users
            </h1>
            <hr>
            <table class="table-auto border-separate">
                <thead>
                    <tr>
                        <th>User</th>
                        <th>Email</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="u in this.users" v-bind:key="u.ID">
                        <td>
                            <router-link :to="`/admin/users/${u.ID}`">{{ u.FirstName }} {{ u.LastName }}</router-link>
                        </td>
                        <td>{{ u.Email }}</td>
                    </tr>
                </tbody>
            </table>
        </div>
</template>

<script>
import Security from './security';
import Toastify from "toastify-js";

export default {
    beforeMount() {
        Security.requireToken();

        fetch(import.meta.env.VITE_BACKEND_SERVER+"/api/v1/admin/users", Security.getRequestOptions())
        .then((response) => {
            response.json().then((body) => {
                if (body.status === "success") {
                    this.users = body.data.users
                    console.log(this.users)
                } else {
                    Toastify({
                        text: "Data Fetch Failed",
                        duration: 1000,
                        close: true
                    }).showToast();
                }
            })})
    },
    data() {
        return {
            users: []
        }
    }
}
</script>