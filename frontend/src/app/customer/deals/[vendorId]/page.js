"use client"
import Image from "next/image";
import backgroundImage from '../../../../assets/food-image.jpg'
import CustomerFoodCard from "@/components/CustomerFoodCard";
import { useEffect, useState } from "react";

export default function VendorPage({ params }) {
    const [menuItems, setMenuItems] = useState([])

    // get all meal items for a restaurant
    useEffect(() => {
        const fetchMenuItems = async() => {
        const res = await fetch(`http://localhost:5001/customer/discount/${params.vendorId}`, {
            method: 'GET',
            credentials: 'include',
            headers: {
                'Content-Type': 'application/json'
            },
        })
        const data = await res.json()
        console.log(data)
        setMenuItems(data)
        }
        fetchMenuItems()
    }, [])

    return (
        <div className="bg-white-100 dark:bg-white-800 py-8">
            <div className="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
                <div className="flex flex-col md:flex-row -mx-4">
                    <div className="md:flex-1 px-4">
                        <div className="w-20% h-10% rounded-lg bg-gray-300 dark:bg-gray-700 mb-4">
                            <Image className="w-20% h-20% object-fit" src={backgroundImage} alt="Product Image"/>
                        </div>
                    </div>
                    <div className="md:flex-1 px-4">
                        <h2 className="text-2xl font-bold text-gray-800 dark:text-white mb-2 text-center">{menuItems[0]?.VendorName}</h2>
                        <p className="deal-info text-center">
                            deals open from <span class="start-time"> 09:00 PM </span> - <span className="end-time">09:30 PM</span>, while stocks last.
                        </p>
                        <div className="mb-4">
                            {menuItems.map((item) => {
                                return (
                                    <CustomerFoodCard key={item.MealID} mealId={item.MealID} mealName={item.MealName} mealDiscountPrice={item.DiscountPrice} mealOriginalPrice={item.MealPrice} mealQuantity={item.Quantity}/>
                                )
                            })}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};
  