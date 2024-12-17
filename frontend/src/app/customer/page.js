'use client'
import fakePage from '../../assets/Home.svg'
import icon from '../../assets/what-a-steal.svg'
import Image from "next/image";
import { useRouter } from 'next/navigation'

export default function Customer() {
    const router = useRouter()

    return (
        <div className="w-[320px] m-auto relative">
            <Image src={fakePage} className="w-[320px] absolute" alt="mock home page"/>
            <a href="./customer/deals">
            <Image src={icon} className="absolute top-[211px] left-[20px] w-[55px] cursor-pointer" alt="what a steal icon"/>
            </a>
        </div>
    )
}