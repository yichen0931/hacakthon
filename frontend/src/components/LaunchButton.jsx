
function LaunchButton({setPostRequest, sendPostRequest}) {
    const handleOnClick = () => {
        setPostRequest((prevData)=>({
            ...prevData,
            ['IsDiscount']:true
        }))
        sendPostRequest()
    }

    return(
        <div className="flex justify-end w-[99%]">
            <button className="rounded-2xl bg-pink-500 text-white font-bold w-[150px] h-[50px] self-end hover:bg-pink-700" onClick={handleOnClick}>
                Launch
            </button>
        </div>
    )
}

export default LaunchButton