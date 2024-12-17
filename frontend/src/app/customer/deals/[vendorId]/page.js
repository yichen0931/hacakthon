"use client"
import Image from "next/image";
import backgroundImage from '../../../../assets/food-image.jpg'
import CustomerFoodCard from "@/components/CustomerFoodCard";
import { useEffect, useState, useRef } from "react";
import { useRouter } from 'next/navigation';

export default function VendorPage({ params }) {
    const [menuItems, setMenuItems] = useState([]);
    const [cartItems, setCartItems] = useState([]);
    const router = useRouter();

    // Function to handle updates from CustomerFoodCard
    const handleQuantityUpdate = (mealId, quantity, discountedPrice) => {
        setCartItems((prevCartItems) => {
            const existingItemIndex = prevCartItems.findIndex((item) => item.mealId === mealId);
            if (existingItemIndex !== -1) {
                // Update existing item
                const updatedItems = [...prevCartItems];
                updatedItems[existingItemIndex] = { mealId, quantity, discountedPrice};
                return updatedItems;
            } else {
                // Add new item
                return [...prevCartItems, { mealId, quantity, discountedPrice}];
            }
        });
    };

    // Function to handle checkout
    const newCartItems = useRef([])
    const payment = async() => {
        let total = 0;
        for (let cartItem of cartItems) {
            total += cartItem.quantity * cartItem.discountedPrice; 
            newCartItems.current.push({
                "ID": cartItem.mealId,
                "Qty": cartItem.quantity,
                "Price": cartItem.discountedPrice
            })
        }
        const payload = {
            "Total": total,
            "DeliveryAddress": "foodpanda office",
            "Meal": newCartItems
        }; 
        console.log(payload)

        try {
            const response = await fetch('http://localhost:5001/checkout', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(payload),
                credentials: 'include',
            });
            const result = await response.json();
            console.log(result)
            
            if (result.orderInserted) {
                // Redirect or handle successful login
                console.log("Checkout successful!")
                router.push("/checkoutOngoing")
            } else {
                // Display error message if any
                console.log(error)
                alert("Please try again!")
            }
        } catch (error) {
            console.log(error);
            alert("Please try again!")
        }
    }

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

    const checkoutStyle = {
        position: 'fixed',
        width: "100%",
        bottom: '0px',
        left: '50%',
        transform: 'translateX(-50%)',
        fontWeight: 'bold',
        backgroundColor: '#C21760',
        color: 'white',         
        padding: '15px 30px',
        fontSize: '18px',       
        border: 'none',         
        borderRadius: '5px',  
        cursor: 'pointer',    
        boxShadow: '0px 4px 10px rgba(0, 0, 0, 0.2)',
        zIndex: 1000,
    };

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
                        {menuItems && menuItems.length > 0 ? (
                            <>
                                <h2 className="text-2xl font-bold text-gray-800 dark:text-white mb-2 text-center">{menuItems[0]?.VendorName}</h2>
                                <p className="deal-info text-center">
                                    deals open for limited time only, <span class="stocks">while stocks last</span>
                                </p>
                                <div className="mb-4">
                                    {menuItems.map((item) => {
                                        return (
                                            <CustomerFoodCard key={item.MealID} mealId={item.MealID} mealName={item.MealName} mealDiscountPrice={item.DiscountPrice} mealOriginalPrice={item.MealPrice} mealQuantity={item.Quantity} onQuantityChange={handleQuantityUpdate}/>
                                        )
                                    })}
                                </div>
                                <div>
                                    <button style={checkoutStyle} onClick={payment}>
                                        Checkout
                                    </button>
                                </div>
                            </>
                        ) : (
                            <div className="text-center py-10">
                                <h2 className="text-xl font-bold text-gray-800 stocks">No Deals Available</h2>
                                <p className="text-gray-600">Please check back later!</p>
                            </div>
                        )}
                    </div>
                </div>
            </div>
        </div>
    );
};
  