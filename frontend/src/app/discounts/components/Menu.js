import MenuCard from './MenuCard.js'

const Menu = ({menuItems}) => {
    return (
        <div className="w-full p-4">
            <h2 className="text-2xl font-bold mb-4">Menu</h2>
            <div className="w-full space-y-4">
                {menuItems.map((item) => {
                    if (item.Availability == 1) {
                        return (
                            <MenuCard key={item.MealID} mealName={item.MealName} mealPrice={item.Price} />
                        )
                    }
                })}
            </div>
        </div>
    )
};

export default Menu; 