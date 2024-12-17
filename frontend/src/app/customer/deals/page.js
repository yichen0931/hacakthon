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
            <ImageWithOverlay imgsrc={backgroundImage} text="Your last happy hour deals"/>
            <div className="overflow-y-auto">
                {loading && <p>Loading...</p>}
                {}
                {!loading && data && (
                    data.map((item) => {
                        return(<VendorCard imgsrc={backgroundImage} name="Delifrance" startTime="09:00pm" endTime="9.30pm" address="123 Green Street" link="#"/>)
                    })
                )}
            </div>
        </div>
    )
}
