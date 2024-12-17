'use client'
import { useState, useEffect } from 'react'
import Menu from '@/components/Menu'
import Sidebar from '@/components/Sidebar'
import OperatingTime from '@/components/OperatingTime'
import LaunchButton from '@/components/LaunchButton'
import Header from '@/components/Header'

export default function Discounts() {    
    const [sidebarIsOpen, setSidebarIsOpen] = useState(false);
    const [indicator, setIndicator] = useState("")

    // Handle window resize
    useEffect(() => {
        const handleResize = () => {
        if (window.innerWidth >= 768) {
            setSidebarIsOpen(true); // Open sidebar on larger screens
        } else {
            setSidebarIsOpen(false); // Hide sidebar on smaller screens
        }
        };

        // Add event listener on component mount
        window.addEventListener('resize', handleResize);

        // Run once initially to set correct state
        handleResize();

        // Cleanup listener on unmount
        return () => window.removeEventListener('resize', handleResize);
    }, []);

    const [discountStatus, setDiscountStatus] = useState([])

    const url = 'http://localhost:5001/vendor/discount'

    // get all meal items for a restaurant
    useEffect(() => {
        async function fetchDiscountStatus() {
            const res = await fetch(url, {
                method: 'GET',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json'
                },
            })
            if (!res.ok) {
                console.error("Failed to fetch:", res.status, res.statusText);
                return;
            }
            const data = await res.json()
            setDiscountStatus(data || {})
            }
        fetchDiscountStatus()
      }, [])

    const [postRequest, setPostRequest] = useState({
        DiscountStart:"",
        DiscountEnd:"",
        Meals:[],
        Button:"",
    })

    useEffect(() => {
        
        if (discountStatus.length != 0) {
            let startTime = discountStatus[0].discountStart.slice(-8)
            let endTime = discountStatus[0].discountEnd.slice(-8)
            setPostRequest((prevData) => ({
                ...prevData,
                DiscountStart: startTime,
                DiscountEnd: endTime,
            }))
            const nowTime = new Date(Date.now())
            let startSeconds = stringToTimeInMiliseconds(startTime) 
            let endSeconds = stringToTimeInMiliseconds(endTime)
            let nowSeconds = nowTime.getHours() * 3600 + nowTime.getMinutes() * 60 + nowTime.getSeconds()
            if (discountStatus[0].isDiscount || (nowSeconds < endSeconds && nowSeconds > startSeconds)) {
                setIndicator("on")
            } else {
                setIndicator("")
            }
        }
    },[discountStatus])

    async function sendPostRequest() {
        console.log(JSON.stringify(postRequest))
        try {
            const res = await fetch(url, {
                method: 'POST',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(postRequest),
            })
            if (!res.ok) {
                throw new Error('Failed to send data')
            }
        } catch (error) {
            console.error('Error:', error);
        }
    }

    return (
    <>
        <Sidebar current="Discounts" indicator={indicator} isOpen={sidebarIsOpen} setIsOpen={setSidebarIsOpen} className="z-100"/>
        <div className="flex-1 md:ml-[300px] ml-[100px] p-10 overflow-y-auto">
            {/* Main content area */}
                <button
                onClick={() => setSidebarIsOpen(!sidebarIsOpen)}
                className="md:hidden p-2 bg-pink-500 text-white rounded"
                >
                {sidebarIsOpen ? '' : '☰'}
                </button>
            <Header name="Discount" indicator={indicator}/>
            <OperatingTime postRequest={postRequest} setPostRequest={setPostRequest} sendPostRequest={sendPostRequest}/>
            <Menu menuItems={discountStatus}/>
            <LaunchButton postRequest={postRequest} setPostRequest={setPostRequest} sendPostRequest={sendPostRequest}/>
            
        </div>
    </>
    )
}

function stringToTimeInMiliseconds(time) {
    const datevar = new Date("01-01-2000 "+time)
    let seconds = datevar.getHours() * 3600 + datevar.getMinutes() * 60 + datevar.getSeconds()
    return seconds

}