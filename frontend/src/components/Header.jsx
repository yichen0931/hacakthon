
function Indicator() {
    return (
        <div className="text-green-500 flex items-center inline-block">
            <div class="w-4 h-4 bg-green-500 rounded-full inline-block"></div>
            &nbsp; Started
        </div>
    )
}

// Returns Header 
// Properties : name, indicator
// name will become the text displayed
// if indicator is set to "on" then it is visible otherwise not visible
// get value will be handled by parent page
function Header(props) {
    return(
        <div className="inline-flex">
            <h1 className="text-3xl font-bold p-3 inline-Block">{props.name}</h1>
            {props.indicator == "on" && Indicator()}
        </div>
    )
}

export default Header