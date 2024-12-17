'use client'
import { useState, useEffect } from 'react'
import Menu from '@/components/Menu'
import Sidebar from '@/components/Sidebar'
import OperatingTime from '@/components/OperatingTime'
import LaunchButton from '@/components/LaunchButton'
import Header from '@/components/Header'

export default function Discounts() {
    const [isOpen, setIsOpen] = useState(false);

    // Handle window resize
    useEffect(() => {
        const handleResize = () => {
        if (window.innerWidth >= 768) {
            setIsOpen(true); // Open sidebar on larger screens
        } else {
            setIsOpen(false); // Hide sidebar on smaller screens
        }
        };

        // Add event listener on component mount
        window.addEventListener('resize', handleResize);

        // Run once initially to set correct state
        handleResize();

        // Cleanup listener on unmount
        return () => window.removeEventListener('resize', handleResize);
    }, []);

    const [menuItem, setMenuItem] = useState([])

    // get all meal items for a restaurant
    useEffect(() => {
        const fetchMenuItems = async() => {
        const res = await fetch('http://localhost:5001/vendor/discount', {
            method: 'GET',
            credentials: 'include',
            headers: {
                'Content-Type': 'application/json'
            },
        })
        const data = await res.json()
        setMenuItem(data)
        }
        fetchMenuItems()
    }, [])

    var menuItems = [
        {
            "MealID": "M007",
            "MealName": "Chocolate Lava Cake",
            "Description": "Molten chocolate dessert",
            "Price": 6.5,
            "Availability": 0,
            "SustainabilityCreditScore": 55
        },
        {
            "MealID": "M008",
            "MealName": "Vanilla Ice Cream",
            "Description": "Classic vanilla ice cream scoop",
            "Price": 4,
            "Availability": 1,
            "SustainabilityCreditScore": 60
        },
        {
            "MealID": "M009",
            "MealName": "Apple Pie",
            "Description": "Warm apple pie with cinnamon",
            "Price": 5,
            "Availability": 1,
            "SustainabilityCreditScore": 50
        },
        {
            "MealID": "M010",
            "MealName": "Cheesecake",
            "Description": "Creamy New York-style cheesecake",
            "Price": 6,
            "Availability": 1,
            "SustainabilityCreditScore": 60
        },
    ]

    const [discountStatus, setDiscountStatus] = useState({StartTime:"21:00",EndTime:"21:30",IsDiscount:false})
    const url = 'https://localhost:5001/vendor/discount'

    const login = {"vendorSessionCookie":"137b4120-2284-4cb6-852d-2ea3f34a3ea4"}

    // get all meal items for a restaurant
    useEffect(() => {
        async function fetchMenuItems() {
          const res = await fetch(url, {
              method: 'GET',
              credentials: 'include',
              headers: {
                  'Content-Type': 'application/json'
              },
          })
          const data = await res.json()
          console.log(data)
          setMenuItem(data)
        }
        fetchMenuItems()
      }, [])

    async function PostDiscountStatus() {
        console.log(JSON.stringify(discountStatus))
        try {
            const res = await fetch(url, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(discountStatus),
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
        <Sidebar current="Discounts" indicator="on" isOpen={isOpen} setIsOpen={setIsOpen} className="z-100"/>
        <div className="flex-1 md:ml-[300px] ml-[100px] p-10 overflow-y-auto">
            {/* Main content area */}
                <button
                onClick={() => setIsOpen(!isOpen)}
                className="md:hidden p-2 bg-pink-500 text-white rounded"
                >
                {isOpen ? '' : '☰'}
                </button>
            <Header name="Discount" indicator="on"/>
            <OperatingTime discountStatus={discountStatus} setDiscountStatus={setDiscountStatus} PostDiscountStatus={PostDiscountStatus}/>
            <Menu menuItems={menuItem}/>
            <LaunchButton discountStatus={discountStatus} setDiscountStatus={setDiscountStatus} PostDiscountStatus={PostDiscountStatus}/>
            
        </div>
    </>
    )
}