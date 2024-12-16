import { useRouter } from 'next/navigation'
import Image from "next/image"
import logo from '../assets/foodpanda-app-icon-square.png'

// To render each item based
// Properties : name, link (optional), indicator (optional)
// Accept two properties name and link
// Name is displayed as text and the link is the link that user will be sent to when pressed
// If no link is provided link is default to "#"
// If not indicator provided it is off or any other value, if indicator value = "on" discount tab will have a green indicator
function Item(props) {
  const router = useRouter()
  let cssSelectors = "p-2 mt-2 flex items-center px-4 duration-300 cursor-pointer hover:bg-pink-500 hover:text-white "
  if (props.selected) {
    cssSelectors += "text-pink-500 font-bold border-l-4 border-solid border-pink-500"
  } else {
    cssSelectors += "text-black"    
  }
  
  return (
    <div className={cssSelectors}>
          <i className="bi bi-house-door-fill"></i>
          <button type="button" onClick={() => router.push(props.link ? props.link : "#")}>
            <span className="text-[15px] ml-4 text-black-200  ">{props.name}</span>
            {props.name == "Discounts" && props.indicator == "on" && <div class="ml-4 w-2 h-2 bg-green-500 rounded-full inline-block"></div>}
          </button>
        </div>
  )
}

// Returns Sidebar
// Properties : current, indicator
// When calling sidebar, use current property to select which item to be selected. i.e. current="Dashboard"
// That item will have a pink text and a left border to indicate it is selected to the user
// To add or edit the name/link for item, change it in the const items variable
// If not indicator provided it is off or any other value, if indicator value = "on" discount tab will have a green indicator
function Sidebar(props) {
  const items = [
    {name:"Dashboard", link:"#"},
    {name:"Marketing", link:"#"},
    {name:"Orders", link:"#"},
    {name:"Rating and Reviews", link:"#"},
    {name:"Reports", link:"#"},
    {name:"Concepts", link:"#"},
    {name:"Notification Centre", link:"#"},
    {name:"Menu Management", link:"#"},
    {name:"Opening Times", link:"#"},
    {name:"Invoices", link:"#"},
    {name:"University", link:"#"},
    {name:"Discounts", link:"./discounts"},

  ]
  return (
      <div>
        <span className="absolute text-black text-4xl top-5 left-4 cursor-pointer">
          <i className="bi bi-filter-left px-2 bg-white-900 rounded-md"></i>
        </span>
        <div className="sidebar fixed top-0 bottom-0 lg:left-0 p-2 w-[300px] overflow-y-auto text-center bg-white-900 border-r-2 border-solid">
          <div className="text-black-100 text-xl">
              <div className="p-3 flex items-start justify-start">
              <Image src={logo} alt="Logo" width="auto" height={50}/>
              </div>
              <div className="my-2 bg-white-600 h-[1px]"></div>
          </div>

          {items.map((item) => (
            <Item name={item.name} link={item.link} selected={props.current == item.name && true} indicator={props.indicator} key={item.name}/>
          ))}

        </div>
      </div>
    );
}

export default Sidebar;