'use client'
import mapImage from '../../assets/map.jpg'
import deliveryImage from '../../assets/delivery.svg'
import bottomImage from '../../assets/bottom.svg'
import Image from "next/image";
import { useRouter } from 'next/navigation'

export default function CheckoutOngoing() {
    const router = useRouter()

    return (
        <div className="max-w-[430px] m-auto bg-gray-100 min-h-screen flex flex-col">
            {/* First section with images */}
            <div className="relative">
                <Image src={mapImage} className="w-[100%]" alt="map"/>
                <Image src={deliveryImage} className="absolute top-[50px] left-[80px] w-[60%] cursor-pointer" alt="delivery icon"/>
            </div>

            {/* Second section with grey background */}
            <div className="bg-gray-100 flex-1 p-4">
                {/* Your content here */}
                <Image src={bottomImage} className="w-[100%]" alt="map"/>

            </div>
        </div>
    );
}