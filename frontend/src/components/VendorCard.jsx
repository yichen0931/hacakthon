import Image from "next/image"
import icon from '../assets/two-wheeler.svg'
import Link from 'next/link'

function VendorCard(props) {
    return(
        // <a href={props.link}>
        //     <div className="w-[95%] m-auto my-2 border-solid border rounded-2xl">
        //         <Image src={props.imgsrc} className="rounded-2xl w-[100%] h-[150px] object-cover"/>
        //         <div className="p-4">
        //             <h2 className=" text-l font-bold">{props.name}</h2>
        //             <p>{props.address}</p>
        //             <div className="flex justify-between">
        //                 <span className=""><Image src={icon} className="inline"/>  S$3.20</span>
        //                 <span className="text-pink-500 font-bold">{props.startTime} - {props.endTime}</span>
        //             </div>
        //         </div>
        //     </div>
        // </a>
        <Link href={props.link}>
            <div className="w-[95%] m-auto my-2 border-solid border rounded-2xl">
                <img src={props.imgsrc} className="rounded-2xl w-[100%] h-[150px] object-cover" alt="vendor-image"/>
                <div className="p-4">
                    <h2 className=" text-l font-bold">{props.name}</h2>
                    <p>{props.address}</p>
                    <div className="flex justify-between">
                        <span className=""><Image src={icon} className="inline" alt="rider-icon"/>  S$3.20</span>
                        <span className="text-pink-500 font-bold">{(props.startTime=="0:00am" && props.endTime=="0:00am") ? "While Stocks Last" : props.startTime+ " - " +  props.endTime}</span>
                    </div>
                </div>
            </div>
        </Link>
    )
}

export default VendorCard