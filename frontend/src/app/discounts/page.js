'use client'
import { useState, useEffect } from 'react'
import Menu from './components/Menu'
import Sidebar from '@/components/Sidebar'
import OperatingTime from '@/components/OperatingTime'

export default function Discounts() {

    // const [menuItem, setMenuItem] = useState({ 
    //     Username: "",
    //     UserPassword: "",
    //     Firstname: "",
    //     Lastname: "",
    // }) 

    // // get all meal items for a restaurant
    // useEffect(() => {
    //     async function fetchMenuItems() {
    //       const res = await fetch('http://localhost:5001/customer/discount/V001', {
    //           method: 'GET',
    //           credentials: 'include',
    //           headers: {
    //               'Content-Type': 'application/json'
    //           },
    //       })
    //       const data = await res.json()
    //       setMenuItem(data)
    //     }
    //     fetchMenuItems()
    //   }, [])

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
    

    return (
    <div className="flex">
        <Sidebar current="Discounts"/>
        <div className="flex-1 lg:ml-[300px] p-10 overflow-y-auto">
            <OperatingTime startTime="21:00" endTime="21:30"/>
            <Menu menuItems={menuItems}/>
        </div>
    </div>
    )
}