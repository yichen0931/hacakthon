import { useRouter } from 'next/navigation'
import Image from "next/image"
import logo from '../assets/foodpanda-app-icon-square.png';

// interface Props {
//   name

// }

function Item(props) {
  const router = useRouter()
  let cssSelectors = "p-2 mt-2 flex items-center px-4 duration-300 cursor-pointer hover:bg-pink-500 hover:text-white "
  if (props.selected == "true") {
    cssSelectors += "text-pink-500 font-bold border-l-4 border-solid border-pink-500"
  } else {
    cssSelectors += "text-black"    
  }
  
  return (
    <div className={cssSelectors}>
          <i className="bi bi-house-door-fill"></i>
          <button type="button" onClick={() => router.push(props.link ? props.link : "#")}>
            <span className="text-[15px] ml-4 text-black-200  ">{props.name}</span>
          </button>
        </div>
  )
}

function Sidebar(props) {
  return (
      <div>
        <span className="absolute text-black text-4xl top-5 left-4 cursor-pointer">
          <i className="bi bi-filter-left px-2 bg-white-900 rounded-md"></i>
        </span>
        <div className="sidebar fixed top-0 bottom-0 lg:left-0 p-2 w-[300px] overflow-y-auto text-center bg-white-900">
          <div className="text-black-100 text-xl">
              <div className="p-2.5 mt-1 flex items-center">
              <Image src={logo} alt="Logo" width={200} height={100}/>
                {/* <img className="logo" src={pandaLogo}/> */}
                {/* <i className="bi bi-app-indicator px-2 py-1 rounded-md bg-blue-600"></i>
                <h1 className="font-bold text-black-200 text-[15px] ml-3">TailwindCSS</h1>
                <i className="bi bi-x cursor-pointer ml-28 lg:hidden"></i> */}
              </div>
              <div className="my-2 bg-white-600 h-[1px]"></div>
          </div>

          <Item name="Dashboard" selected="true"/>
          <Item name="Marketing"/>
          <Item name="Orders"/>
          <Item name="Rating and Reviews"/>
          <Item name="Reports"/>
          <Item name="Concepts"/>
          <Item name="Notification Centre"/>
          <Item name="Menu Management"/>
          <Item name="Opening Times"/>
          <Item name="Invoices"/>
          <Item name="University"/>
          <Item name="Discounts" link="./discounts"/>

        </div>
      </div>
    );
}

export default Sidebar;