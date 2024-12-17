
function LaunchButton({setPostRequest, sendPostRequest, indicator}) {
    const handleOnClick = () => {
        if (indicator == "on") {
            setPostRequest((prevData)=>({
                ...prevData,
                Button:"End",
                DiscountStart: "00:00",
                DiscountEnd: "00:00",
            }))
        } else {
            setPostRequest((prevData)=>({
                ...prevData,
                Button:"Launch",
                DiscountStart: "00:00",
                DiscountEnd: "00:00",
            }))
        }
        
    }

    return(
        <div className="flex justify-end w-[99%]">
            <button className="rounded-2xl bg-pink-500 text-white font-bold w-[150px] h-[50px] self-end hover:bg-pink-700" onClick={handleOnClick}>
                {indicator == "on" ? "End" : "Launch Now"}
            </button>
        </div>
    )
}

export default LaunchButton