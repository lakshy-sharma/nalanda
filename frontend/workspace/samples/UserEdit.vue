<template>
    <div class="container md:mx-auto shadow-md text-center my-2 md:mx-auto">
        <h1 class="mt-3">
            User Editing
        </h1>
        <hr>
        <form-base @userEditEvent="submitHandler" name="userform" event="userEditEvent">
            <text-input v-model="user.first_name" type="text" required="true" label="First Name" :value="user.first_name" name="first-name"></text-input>
            <text-input v-model="user.last_name" type="text" required="true" label="Last Name" :value="user.last_name" name="last-name"></text-input>
            <text-input v-model="user.email" type="email" required="true" label="Email" :value="user.email" name="email"></text-input>
            <text-input v-if="this.user.id === 0" v-model="user.password" type="password" required="true" label="Password" :value="user.email" name="password"></text-input>
            <text-input v-else v-model="user.password" type="password" label="Password" :value="user.email" name="password"></text-input>
            <div class="flex-grid grid-cols-3">
                <input type="submit" class="col-start-2 inline-block bg-slate-200 font-semibold text-slate-600 hover:text-slate-800 rounded-md py-1 px-3 my-2 max-w-sm" value="Submit">
                <router-link to="/admin/users" class="col-start-2 inline-block bg-slate-200 font-semibold text-slate-600 hover:text-slate-800 rounded-md py-1 px-3 my-2 max-w-sm">Cancel</router-link>
            </div>
        </form-base> 
        <hr>
    </div>
</template>

<script>
import Security from './security';
import FormBase from "./forms/FormBase.vue"
import TextInput from './forms/TextInput.vue';

export default {
    beforeMount() {
        Security.requireToken();
        // Proceed only if you receive a valid user id.
        if (parseInt(String(this.$route.params.userId), 10) > 0) {
        }
    },
    data() {
        return {
            user: {
                id: 0,
                first_name: "",
                last_name: "",
                email: "",
                password: ""
            }
        }
    },
    components: {
        "form-base": FormBase,
        "text-input": TextInput
    },
    methods: {
        submitHandler() {
            const payload = {
                id: parseInt(String(this.$route.params.userId), 10),
                first_name: this.user.first_name,
                last_name: this.user.last_name,
                email: this.user.email,
                password: this.user.password
            }
            fetch(`${import.meta.env.VITE_BACKEND_SERVER}/admin/users/save`, Security.postRequestOptions(payload))
            .then((response) => {
                response.json().then((body) => {
                    if(body.status === "success") {
                        console.log("Ok")
                    } else {
                        Toastify({
                            text: "Data Submission Failed",
                            duration: 1000,
                            close: true
                        }).showToast();
                    }
                })
            })
        }
    }
}
</script>