'use client'
import Image from "next/image";
import { useState, useEffect } from 'react';

  


const MenuCard = ({mealId, mealName, mealPrice, postRequest, setPostRequest}) => {
  const [quantity, setQuantity] = useState(0);
  const [price,setPrice] = useState(mealPrice || 0);

  useEffect(() => {
    setPrice(mealPrice || 0);
  }, [mealPrice]);

  const increment = () => {
    setQuantity(quantity + 1); 
  }
  const decrement = () => {
    if (quantity > 0) {
      setQuantity(quantity - 1);
    }
  }

  const handleQuantityChange = (e) => {
    const value = parseInt(e.target.value, 10);
    if (!isNaN(value)) {
      setQuantity(value);
    } else {
      setQuantity(0)
    }
  }

  const handlePriceChange = (e) => {
    const value = parseFloat(e.target.value);
    if (!isNaN(value) && value >= 0) {
      setPrice(value);
    } else {
      setPrice(0)
    }
  }

  useEffect(() => {
    setPostRequest((prevData) => {
      // Safely create a new Meals array
      const updatedMeals = prevData.Meals.map((meal) =>
          meal.MealID === mealId
              ? { ...meal, Quantity: quantity, DiscountPrice: price }
              : meal
      );

      // Check if the meal is new and needs to be added
      const mealExists = prevData.Meals.some((meal) => meal.MealID === mealId);

      const finalMeals = mealExists
          ? updatedMeals
          : [
            ...updatedMeals,
            {
              MealID: mealId,
              DiscountPrice: price,
              Quantity: quantity,
            },
          ];

      return {
        ...prevData,
        Meals: finalMeals, // Update immutably
      };
      console.log("Updated PostRequest:", updatedRequest);
      return updatedRequest
    });
  }, [quantity, price]);


  return (
    <div>
      <div className="relative">
        <div className="relative inset-px rounded-lg bg-white max-lg:rounded-t-[2rem]"></div>
        <div className="flex h-full overflow-hidden rounded-[calc(theme(borderRadius.lg)+1px)] max-lg:rounded-t-[calc(2rem+1px)]">
          
          {/* Menu Item Image */}
          <div className="justify-center max-lg:pb-12 max-lg:pt-12 sm:px-5 sm:py-4 ">
            <Image
              className="w-100 max-lg:max-w-100"
              src={mealId != undefined ? "/images/"+mealId+".jpg" : "/images/sample-image.jpg"}
              // "/images/M001.jpg"
              alt="food image"
              width={200}
              height={200}
            />
          </div>

          {/* Menu Item Details */}
          <div className="px-4 sm:px-4 items-center content-center">
            <p className="mt-2 text-lg font-medium tracking-tight text-gray-950 max-lg:text-center">{mealName}</p>
            <p className="mt-2 max-w-lg text-sm/6 text-gray-600 max-lg:text-center">Original: S${mealPrice.toFixed(2)}</p>
            <div className="flex">
              <div>
                <p className="mt-2 max-w-lg text-sm/6 text-gray-600 max-lg:text-center">Set To: </p>
              </div>
              <div>
                <div className="mx-2 mt-1">
                  <div className="flex items-center rounded-md bg-white pl-3 outline outline-1 -outline-offset-1 outline-gray-300 has-[input:focus-within]:outline has-[input:focus-within]:outline-2 has-[input:focus-within]:-outline-offset-2 has-[input:focus-within]:outline-indigo-600">
                    <div className="shrink-0 select-none text-base text-gray-500 sm:text-sm/6">$</div>
                    <input
                      id="discountedPrice"
                      name="discountedPrice"
                      type="text"
                      placeholder="0.00"
                      className="block min-w-0 grow py-1.5 pl-1 pr-3 text-base text-gray-900 placeholder:text-gray-400 focus:outline focus:outline-0 sm:text-sm/6"
                      onChange={handlePriceChange}
                    />
                  </div>
                </div>
              </div>
            </div>
          </div>

          {/* Menu Item Quantity */}
          <div className="px-4 pt-7 sm:px-10 items-center content-center">
            <p className="my-2 max-w-lg text-sm/6 text-gray-600 max-lg:text-center">Quantity:</p>
            <div className="flex">
              <div>
                <button className="w-10 bg-pink-500 hover:bg-pink-700 text-white font-bold py-2 px-4 pb-7 rounded h-8 text-sm" onClick={decrement}>-</button>
              </div>
              <div>
                <div className="mx-2">
                  <div className="flex items-center rounded-md bg-white pl-3 outline outline-1 -outline-offset-1 outline-gray-300 has-[input:focus-within]:outline has-[input:focus-within]:outline-2 has-[input:focus-within]:-outline-offset-2 has-[input:focus-within]:outline-indigo-600">
                    <input
                      id="discountedQuantity"
                      name="discountedQuantity"
                      type="text"
                      placeholder="0"
                      onChange={handleQuantityChange}
                      value={quantity}
                      className="block min-w-0 w-10 grow py-1.5 pr-3 text-center text-base text-gray-900 placeholder:text-gray-400 focus:outline focus:outline-0 sm:text-sm/6"
                    />
                  </div>
                </div>
              </div>
              <div>
                <button className="w-10 bg-pink-500 hover:bg-pink-700 text-white font-bold py-2 px-4 pb-7 rounded h-8 text-sm" onClick={increment}>+</button>
              </div>
            </div>
          </div>
        </div>
        <div className="pointer-events-none absolute inset-px rounded-lg shadow ring-1 ring-black/5 max-lg:rounded-t-[2rem]"></div>
      </div>
    </div>
  );
};

export default MenuCard;