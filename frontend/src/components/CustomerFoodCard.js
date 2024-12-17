'use client'
import Image from "next/image";
import { useState, useEffect } from 'react';
import backgroundImage from '../assets/square-food-image.jpg'


const CustomerFoodCard = ({mealId, mealName, mealDiscountPrice, mealOriginalPrice, mealQuantity, onQuantityChange }) => {
    const [quantity, setQuantity] = useState(0);
    const increment = () => {
        if (quantity < mealQuantity) {
            const newQuantity = quantity + 1;
            setQuantity(newQuantity);
            onQuantityChange(mealId, newQuantity, mealDiscountPrice); 
        }
    }
    const decrement = () => {
        if (quantity > 0) {
            const newQuantity = quantity - 1;
            setQuantity(newQuantity);
            onQuantityChange(mealId, newQuantity, mealDiscountPrice); 
        }
    }


    // Handle Button Colour Change for Increment
    const incrementButtonStyle = {
        backgroundColor: quantity === mealQuantity ? '#c0c0c0' : '#EC4899', 
        cursor: quantity === mealQuantity ? 'not-allowed' : 'pointer', 
    };

    // Handle Button Colour Change for Decrement
    const decrementButtonStyle = {
        backgroundColor: quantity === 0 ? '#c0c0c0' : '#EC4899', 
        cursor: quantity === 0 ? 'not-allowed' : 'pointer',
    };

  return (
    <div className="w-full">
        <div className="food-card">
            <div className="food-image flex align-center">
                <Image src={"/images/"+mealId+".jpg"} width={400} height={400}/>
            </div>
            <div className="food-details">
                <h3 className="food-name">{mealName}</h3>
                <div className="price">
                <span className="original-price">${mealOriginalPrice.toFixed(2)}</span>
                <span className="discounted-price">${mealDiscountPrice.toFixed(2)}</span>
                </div>
                <div className="quantity-controls">
                <button onClick={decrement} className="quantity-button w-6 bg-pink-500 hover:bg-pink-700 text-white font-bold py-2 px-4 pb-7 rounded h-8 text-sm" style={decrementButtonStyle}>-</button>
                <span className="quantity">{quantity}</span>
                <button onClick={increment} className="quantity-button w-6 bg-pink-500 hover:bg-pink-700 text-white font-bold py-2 px-4 pb-7 rounded h-8 text-sm" style={incrementButtonStyle}>+</button>
                </div>
            </div>
        </div>
    </div>
  );
};

export default CustomerFoodCard;