import { useState, useEffect } from 'react'

// Making operating time component
// Properties : startTime, endTime
// Parent page will get value from server and populate 
function OperatingTime({discountStatus, setDiscountStatus, PostDiscountStatus}) {
    let posturl = 'https://localhost:5001/vendor/discount/'

    // update formData state when there is a change in form data
    const handleChange = (e) => {
        const {name, value} = e.target
        setDiscountStatus((prevData) => ({
            ...prevData,
            [name]: value
        }))
    }

    useEffect(() => {
        PostDiscountStatus()
    },[discountStatus])

    return (
        <>
        <h2 className="text-xl font-semibold p-3">Operating Time</h2>
        <div className="flex justify-center">
        <div className="flex border-2 border-solid w-[99%] h-[50px] rounded-2xl items-center">
            <form id="operating-time">
                <span className="p-7">Start Time: </span>
                <input type="time" className="w-[110px] border-b border-solid" value={discountStatus.StartTime} onChange={handleChange} name="StartTime" id="StartTime"/>
                <span className="p-7">End Time: </span>
                <input type="time" className="w-[110px] border-b border-solid" value={discountStatus.EndTime} onChange={handleChange} name="EndTime" id="EndTime"/>
            </form>
        </div>
        </div>

        </>
    )
}

export default OperatingTime;