'use client'
import fakePage from '../../assets/Home.svg'
import icon from '../../assets/what-a-steal.svg'
import Image from "next/image";
import { useRouter } from 'next/navigation'

export default function Customer() {
    const router = useRouter()

    return (
        <div className="max-w-[430px] m-auto relative">
            <Image src={fakePage} className="w-[100%] absolute" alt="mock home page"/>
            <a href="./customer/deals">
            <Image src={icon} className="absolute top-[250px] left-[20px] w-[62px] cursor-pointer" alt="what a steal icon"/>
            </a>
        </div>
    )
}