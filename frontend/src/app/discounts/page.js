'use client'
import { useState, useEffect } from 'react'

export default function Discounts() {
    const [users, setUsers] = useState(null)
    const [newUser, setNewUser] = useState({ 
        Username: "",
        UserPassword: "",
        Firstname: "",
        Lastname: "",
    })
 
    // get all users
    useEffect(() => {
      async function fetchUsers() {
        const res = await fetch('http://localhost:5001/panda/v1/users', {
            method: 'GET',
            credentials: 'include',
            headers: {
                'Content-Type': 'application/json'
            },
        })
        const data = await res.json()
        setUsers(data)
      }
      fetchUsers()
    }, [])

    const createUser = () => {
        fetch("http://localhost:5001/panda/v1/adduser", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(newUser),
        })
          .then((response) => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
                return response.status === 204 ? null : response.json(); // Handle empty response
          })
          .then((createdUser) => {
            setUsers([...users, createdUser]); // update the user list
            setNewUser({ 
                Username: "",
                UserPassword: "",
                Firstname: "",
                Lastname: "",
            }); // Reset input fields
          })
          .catch((error) => console.error("Error creating user:", error));
    };
    
    if (!users) return <div>Loading...</div>
      
    return (
        <div style={{margin: "20px"}}>
            <h1>Users:</h1>
            <ul>
                {users.map((user) => (
                    <li key={user.UserID}>{user.Username}: {user.Firstname} {user.Lastname}</li>
                ))}
            </ul>

            <br/>
            <h1>Add User</h1>
            <input
                type="text"
                placeholder="Username"
                value={newUser.Username}
                onChange={(e) => setNewUser({ ...newUser, Username: e.target.value })}
            />
            <input
                type="text"
                placeholder="Password"
                value={newUser.UserPassword}
                onChange={(e) => setNewUser({ ...newUser, UserPassword: e.target.value })}
            />
            <input
                type="text"
                placeholder="Firstname"
                value={newUser.Firstname}
                onChange={(e) => setNewUser({ ...newUser, Firstname: e.target.value })}
            />
            <input
                type="text"
                placeholder="Lastname"
                value={newUser.Lastname}
                onChange={(e) => setNewUser({ ...newUser, Lastname: e.target.value })}
            />
            <button onClick={createUser}>Add User</button>
        </div>
    )
}