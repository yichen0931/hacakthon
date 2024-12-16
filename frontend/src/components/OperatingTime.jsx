import { useState, useEffect } from 'react'


function OperatingTime(props) {
    const [formData, setFormData] = useState({startTime:'',endTime:''})
    let fetchurl = 'https://localhost:5001/vendor/'+ props.vid
    let posturl = 'https://localhost:5001/vendor/'+ props.vid
    posturl = 'https://pokeapi.co/api/v2/pokemon/ditto'

    // get start and end time when operating time is mounted
    async function fetchVendor() {
        try {
            const res = await fetch(fetchurl, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json'
                },
            })
            if (!res.ok) {
                throw new Error('Network response was not ok')
            }
            const data = await res.json()
            setFormData({startTime:data.StartTime,endTime:data.EndTime})
        } catch (error) {
            console.log(error)
        }
    }
    useEffect(() => {
        fetchVendor()
    },[])

    // update formData state when there is a change in form data
    const handleChange = (e) => {
        const {name, value} = e.target
        setFormData((prevData) => ({
            ...prevData,
            [name]: value
        }))

    }

    async function PostTime() {
        console.log(JSON.stringify(formData))
        try {
            const res = await fetch(posturl, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(formData),
            })
            if (!res.ok) {
                throw new Error('Failed to send data')
            }
        } catch (error) {
            console.error('Error:', error);
        }
    }

    useEffect(() => {
        PostTime()
    },[formData])

    return (
        <>
        <h2 className="text-xl font-semibold p-3">Operating Time</h2>
        <div className="flex border-2 border-solid w-[80%] h-[50px] rounded-2xl items-center">
            <form id="operating-time">
                <span className="p-7">Start Time: </span>
                <input type="time" className="w-[110px] border-b border-solid" value={formData.startTime} onChange={handleChange} name="startTime" id="startTime"/>
                <span className="p-7">End Time: </span>
                <input type="time" className="w-[110px] border-b border-solid" value={formData.endTime} onChange={handleChange} name="endTime" id="endTime"/>
            </form>
        </div>
        </>
    )
}

export default OperatingTime;