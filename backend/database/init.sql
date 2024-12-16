CREATE USER 'user'@'%' IDENTIFIED
WITH mysql_native_password BY 'strongpassword' ;

GRANT ALL PRIVILEGES ON *.* TO 'user'@'%' ;

FLUSH PRIVILEGES;

create database dealsDB;
use dealsDB;

CREATE TABLE Vendor (
                        VendorID varchar(50) NOT NULL PRIMARY KEY,
                        VendorName varchar(50) NOT NULL,
                        Address varchar(50) NOT NULL,
                        IsOpen boolean,
                        IsDiscountOpen boolean NOT NULL DEFAULT FALSE,
                        DiscountStart datetime,
                        DiscountEnd datetime
);

CREATE TABLE Meal (
                      MealID varchar(50) PRIMARY KEY,
                      VendorID varchar(50),
                      MealName varchar(50),
                      Description varchar(100),
                      Price float,
                      Availability boolean,
                      SustainabilityCreditScore int,
                      FOREIGN KEY (VendorID) REFERENCES Vendor(VendorID)
);

CREATE TABLE Rider (
                       RiderID varchar(50) NOT NULL PRIMARY KEY,
                       RiderName varchar(50) NOT NULL,
                       VehiclePlate varchar(50) NOT NULL,
                       Availability boolean NOT NULL DEFAULT FALSE
);

CREATE TABLE Customer (
                          CustomerID varchar(50) NOT NULL PRIMARY KEY,
                          CustomerName varchar(50) NOT NULL ,
                          Address varchar(50) NOT NULL,
                          AccumulatedSustainabilityCreditScore INT DEFAULT 0
);

CREATE TABLE Discount (
                          MealID varchar(50) NOT NULL PRIMARY KEY,
                          DiscountedPrice FLOAT NOT NULL ,
                          Quantity INT DEFAULT 0,
                          FOREIGN KEY (MealID) REFERENCES Meal(MealID)
);

CREATE TABLE Order (
                       OrderID varchar(50) NOT NULL PRIMARY KEY,
                       CustomerID varchar(50) NOT NULL,
                       RiderID varchar(50) NOT NULL,
                       OrderStatus varchar(20) ENUM('CART','PENDING','ORDERRECIVEVD','GROUPORDER','PREPARING','PICKED','DELIVERED',),
                       OrderEnd datetime,
                       Total float,
                       DeliveryAddress varchar(100) NOT NULL,
                       FOREIGN KEY (CustomerID) REFERENCES Customer(CustomerID),
                       FOREIGN KEY (RiderID) REFERENCES Rider(RiderID)
);

CREATE TABLE OrderDetail (
                             OrderID varchar(50) NOT NULL PRIMARY KEY,
                             MealID varchar(50) NOT NULL,
                             MealQty INT,
                             MealPrice FLOAT,
                             FOREIGN KEY (OrderID) REFERENCES Order(OrderID),
                             FOREIGN KEY (MealID) REFERENCES Meal(MealID)
);

CREATE TABLE Sessions (
                          SessionID varchar(50) PRIMARY KEY,
                          UserID varchar(50) NOT NULL,
                          SessionExpiry datetime,
                          FOREIGN KEY (UserID) REFERENCES Vendor(VendorID),
                          FOREIGN KEY (UserID) REFERENCES Customer(CustomerID)
);

CREATE TABLE Users (
                       UserID varchar(50) NOT NULL PRIMARY KEY,
                       UserName varchar(50) NOT NULL,
                       Password varchar(50),
                       Role varchar(10) NOT NULL ENUM('VENDOR','CUSTOMER','RIDER')
                         FOREIGN KEY (UserID) REFERENCES Sessions(UserID)
);

-- Insert dummy data into the Vendor table
INSERT INTO Vendor (VendorID, VendorName, Address, IsOpen, IsDiscountOpen, DiscountStart, DiscountEnd)
VALUES
    ('V001', 'Healthy Bites', '123 Green Street', TRUE, FALSE, NULL, NULL),
    ('V002', 'Spice Paradise', '456 Flavor Ave', TRUE, TRUE, '2024-12-20 10:00:00', '2024-12-20 18:00:00'),
    ('V003', 'Dessert Haven', '789 Sweet Lane', FALSE, FALSE, NULL, NULL);

-- Insert dummy data into the Meal table for each vendor
INSERT INTO Meal (MealID, VendorID, MealName, Description, Price, Availability, SustainabilityCreditScore)
VALUES
-- Meals for Vendor V001
('M001', 'V001', 'Quinoa Salad', 'Healthy salad with quinoa, veggies, and dressing', 8.99, TRUE, 85),
('M002', 'V001', 'Grilled Chicken Wrap', 'Wrap with grilled chicken and veggies', 7.50, TRUE, 75),
('M003', 'V001', 'Vegetable Soup', 'Warm soup with fresh vegetables', 5.99, TRUE, 80),
-- Meals for Vendor V002
('M004', 'V002', 'Spicy Chicken Curry', 'Rich and spicy chicken curry', 12.50, TRUE, 70),
('M005', 'V002', 'Vegetable Biryani', 'Aromatic rice with mixed vegetables and spices', 10.00, TRUE, 78),
('M006', 'V002', 'Paneer Butter Masala', 'Creamy Indian curry with paneer cubes', 11.00, TRUE, 65),
-- Meals for Vendor V003
('M007', 'V003', 'Chocolate Lava Cake', 'Molten chocolate dessert', 6.50, FALSE, 55),
('M008', 'V003', 'Vanilla Ice Cream', 'Classic vanilla ice cream scoop', 4.00, TRUE, 60),
('M009', 'V003', 'Apple Pie', 'Warm apple pie with cinnamon', 5.00, TRUE, 50),
('M010', 'V003', 'Cheesecake', 'Creamy New York-style cheesecake', 6.00, TRUE, 60);
