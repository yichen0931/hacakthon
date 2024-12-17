'use client'
import { useRouter } from 'next/navigation'
import backgroundImage from '../../../assets/food-image.jpg'
import ImageWithOverlay from '@/components/ImageWithOverlay'
import VendorCard from '@/components/VendorCard'
import { useEffect, useState } from 'react'
const url = 'http://localhost:5001/customer/discount'

export default function Customer() {
    const router = useRouter()
    const [data, setData] = useState(null)
    const [loading, setLoading] = useState(true)

    useEffect(() => {
        try {
            async function getVendor() {
                const res = await fetch(url, {
                    method: 'GET',
                    headers: {
                            'Content-Type': 'application/json'
                        },
                    })
                const result = await res.json()
                setData(result)
            } 
            getVendor()
            
        } catch (error) {
            console.log(error)
        } finally {
            setLoading(false)
        }
    },[])

    return (
        <div className="max-w-[430px] m-auto relative">
            <ImageWithOverlay imgsrc={backgroundImage} text="Your Last Happy Hour Deals"/>
            <div className="overflow-y-auto">
                {loading && <p>Loading...</p>}
                {!loading && data && (
                    data.map((item) => {
                        const formatTime = (timeString) => {
                            const date = new Date(timeString);
                            let hours = date.getHours();
                            const minutes = date.getMinutes();
                            const ampm = hours >= 12 ? 'pm' : 'am';
                            hours = hours % 12
                            const formattedMin = minutes.toString().padStart(2, '0');
                            return `${hours}:${formattedMin}${ampm}`
                        };

                        return(<VendorCard key={item.VendorID} imgsrc={"/images/"+item.VendorImage} name={item.VendorName} startTime={formatTime(item.DiscountStart)} endTime={formatTime(item.DiscountEnd)} address={item.Address} link={`/customer/deals/${item.VendorID}`}>
                        </VendorCard>)
                    })
                )}
            </div>
        </div>
    )
}
