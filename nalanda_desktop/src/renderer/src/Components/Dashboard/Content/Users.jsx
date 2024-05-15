import { useEffect } from 'react';
import Security from './security';
import Toastify from "toastify-js";
import { useState } from 'react';

export function Users() {
    const [users, setUsers] = useState([]);

    useEffect(()=>{
        Security.requireToken();

        fetch(import.meta.env.VITE_BACKEND_SERVER+"/api/v1/admin/users", Security.getRequestOptions())
        .then((response) => {
            response.json().then((body) => {
                if (body.status === "success") {
                    users = body.data.users;
                    console.log(users);
                    setUsers(users);
                } else {
                    Toastify({
                        text: "Couldnt Get Users",
                        duration: 1000,
                        close: true
                    }).showToast();
                }
            })})
    }, [])


    return (
        <div className="container md:mx-auto shadow-md text-center my-2 md:mx-auto">
            <h1 className="mt-3">
                All Users
            </h1>
            <hr />
            <table className="table-auto border-separate">
                <thead>
                    <tr>
                        <th>User</th>
                        <th>Email</th>
                    </tr>
                </thead>
                <tbody>
                    <tr key={index}>
                        <td>
                            <Link to="`/admin/users/${u.ID}`">{{users.FirstName}} {{users.LastName}}</>
                        </td>
                        <td>{{ users.Email }}</td>
                    </tr>
                </tbody>
            </table>
        </div>
    )
}