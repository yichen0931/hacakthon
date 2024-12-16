import MenuCard from './components/MenuCard.js'

export default function Discounts() {
    return (
        <div>
            <h1>Discounts</h1>
            <MenuCard mealId={1} mealName={"Aglio Olio"} mealPrice={10.50}/>
        </div>
    )
}