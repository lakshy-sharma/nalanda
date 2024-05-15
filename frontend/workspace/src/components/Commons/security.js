import {useStore} from "./store.js"

let {token, user} = useStore();

let Security = {
    requireToken: function() {
        if (token === "") {
            router.push("/login")
            return false
        }
    },
    postRequestOptions: function(payload) {
        const headers = new Headers();
        headers.append("Content-Type", "application/json")
        headers.append("Authorization", `Bearer ${token}`)

        return {
            method: "POST",
            body: JSON.stringify(payload),
            headers: headers
        }
    },
    getRequestOptions: function() {
        const headers = new Headers();
        headers.append("Content-Type", "application/json")
        headers.append("Authorization", `Bearer ${token}`)

        return {
            method: "GET",
            headers: headers
        }
    }
}

export default Security