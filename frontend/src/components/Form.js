"use client"
import { useState } from 'react';
import { useRouter } from 'next/navigation';

const Form = () => {
    const [userID, setUserID] = useState('');
    const [password, setPassword] = useState('');
    const [role, setRole] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    const router = useRouter();
    
    const handleLogin = async (e) => {
        e.preventDefault(); // Make sure this is uncommented
        
        // Prepare the data to send in the POST request
        const data = {
            UserID: userID,
            Password: password,
            Role: role
        };
        console.log(data)
        
        try {
            const response = await fetch('http://localhost:5001/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data),
                credentials: 'include',
            });
            
            const result = await response.json();
            console.log(result)
            
            if (result.authenticated) {
                // Redirect or handle successful login
                console.log("Login successful!")
                const { Role } = result; // Destructure Role from the response
                console.log(Role)
                if (Role === "Vendor") {
                    router.push("/discounts")
                } else if (Role == "Customer") {
                    router.push("/customer")
                } else {
                    setErrorMessage('Unknown role');
                }
                // Redirect or do something here (for example using `router.push('/dashboard')`)
            } else {
                // Display error message if any
                setErrorMessage(result.error || 'Login failed');
            }
        } catch (error) {
            console.log(error);
            setErrorMessage('Something went wrong');
        }
    };

    return (
        <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
          <form onSubmit={handleLogin} method="POST" className="space-y-6">
            
            {/* User ID Field */}
            <div>
              <label htmlFor="userid" className="block text-sm/6 font-medium text-gray-900 dark:text-white">
                User ID
              </label>
              <div className="mt-2">
                <input
                  id="userid"
                  name="userid"
                  type="text"
                  onChange={(e) => setUserID(e.target.value)}
                  className="block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline focus:outline-2 focus:-outline-offset-2 focus:outline-pink-600 sm:text-sm/6 dark:text-white"
                />
              </div>
            </div>

            {/* Password Field */}
            <div>
              <div className="flex items-center justify-between">
                <label htmlFor="password" className="block text-sm/6 font-medium text-gray-900 dark:text-white">
                  Password
                </label>
                <div className="text-sm dark:text-white">
                  <a href="#" className="font-semibold text-pink-600 hover:text-pink-500">
                    Forgot password?
                  </a>
                </div>
              </div>
              <div className="mt-2">
                <input
                  id="password"
                  name="password"
                  type="password"
                  required
                  autoComplete="current-password"
                  onChange={(e) => setPassword(e.target.value)}
                  className="block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 dark:text-white focus:outline focus:outline-2 focus:-outline-offset-2 focus:outline-pink-600 sm:text-sm/6"
                />
              </div>
            </div>

            {/* User Role Dropdown */}
            <div>
              <div className="items-center justify-between">
                <label htmlFor="Role" className="block text-sm/6 font-medium text-gray-900 dark:text-white">
                  Role
                </label>
              </div>
              <div className="mt-2 mb-5 max-w-sm mx-auto">
                <select id="roles" className="border border-gray-300 text-gray-900 text-sm rounded-lg block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white focus:outline focus:outline-2 focus:-outline-offset-2 focus:outline-pink-600" onChange={(e) => setRole(e.target.value)}>
                    <option>Choose your role</option>
                    <option value="Customer">Customer</option>
                    <option value="Vendor">Vendor</option>
                </select>
              </div>
            </div>

            {/* Login Button */}
            <div>
              {errorMessage && <p style={{ color: 'red' }}>{errorMessage}</p>}
              <button
                type="submit"
                className="flex w-full justify-center rounded-md bg-pink-600 px-3 py-1.5 text-sm/6 font-semibold text-white shadow-sm hover:bg-pink-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-pink-600"
              >
                Login
              </button>
            </div>
          </form>

          {/* Sign up section */}
          <p className="mt-10 text-center text-sm/6 text-gray-500">
            Not a member?{' '}
            <a href="#" className="font-semibold text-pink-600 hover:text-pink-500">
              Sign up now
            </a>
          </p>
        </div>
    )
};

export default Form; 