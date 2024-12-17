import { useState, useEffect } from 'react'

// Making operating time component
// Properties : startTime, endTime
// Parent page will get value from server and populate 
function OperatingTime({ postRequest, setPostRequest, sendPostRequest}) {
    let startTime, endTime = "00:00"

    try {
        startTime = postRequest.DiscountStart
        endTime = postRequest.DiscountEnd
    } catch (error) {
        console.log(error)
    }

    // update formData state when there is a change in form data
    const handleChange = (e) => {
        const {name, value} = e.target
        setPostRequest((prevData) => ({
            ...prevData,
            [name]: value
        }))
    }

    return (
        <>
        <h2 className="text-xl font-semibold p-3">Operating Time</h2>
        <div className="flex justify-center">
        <div className="flex border-2 border-solid w-[99%] rounded-2xl items-center">
            <form id="operating-time" className='p-2'>
                <span className="p-2">Start Time: </span>
                <input type="time" className="w-[150px] border-b border-solid p-2" value={startTime} onChange={handleChange} name="DiscountStart" id="DiscountStart"/>
                <span className="p-2">End Time: </span>
                <input type="time" className="w-[150px] border-b border-solid p-2" value={endTime} onChange={handleChange} name="DiscountEnd" id="DiscountEnd"/>
            </form>
        </div>
        </div>

        </>
    )
}

export default OperatingTime;