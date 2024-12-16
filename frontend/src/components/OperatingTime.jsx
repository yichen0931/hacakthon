import { useState, useEffect } from 'react'


function OperatingTime() {
    async function fetchVendor() {
        try {
            const res = await fetch('https://localhost:5001/', {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json'
                },
            })
            if (!res.ok) {
                throw new Error('Network response was not ok')
            }
            const data = await res.json()
            console.log(data)
        } catch (error) {
            console.log(error)
        }
    }
    useEffect(() => {
        fetchVendor()
    },[])        

    return (
        <>
        <h2 className="text-xl font-semibold p-3">Operating Time</h2>
        <div className="flex border-2 border-solid w-[80%] h-[50px] rounded-2xl items-center">
        <span className="p-7">Start Time: </span>
        <input type="time" className="w-[110px] border-b border-solid"/>
        <span className="p-7">End Time: </span>
        <input type="time" className="w-[110px] border-b border-solid" />
        </div>
        </>
    )
}

export default OperatingTime;