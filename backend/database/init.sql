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
                        DiscountStart datetime DEFAULT '2024-12-16 21:00:00',
                        DiscountEnd datetime DEFAULT '2024-12-16 22:00:00',
                        Password varchar(50) NOT NULL,
                        VendorImage varchar(100)
);

CREATE TABLE Meal (
                      MealID varchar(50) PRIMARY KEY,
                      VendorID varchar(50),
                      MealName varchar(50),
                      Description varchar(100),
                      Price float,
                      Availability boolean,
                      SustainabilityCreditScore int,
                      MealImage varchar(100)
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
                          AccumulatedSustainabilityCreditScore INT DEFAULT 0,
                          Password varchar(50) NOT NULL
);

CREATE TABLE Discount (
                          MealID varchar(50) NOT NULL PRIMARY KEY,
                          DiscountedPrice FLOAT NOT NULL ,
                          Quantity INT DEFAULT 0,
                          FOREIGN KEY (MealID) REFERENCES Meal(MealID)
);

CREATE TABLE Orders (
                       OrderID varchar(50) NOT NULL PRIMARY KEY,
                       CustomerID varchar(50) NOT NULL,
                       RiderID varchar(50) NOT NULL,
                       OrderStatus ENUM('CART','PENDING','ORDERRECEIVED','GROUPORDER','PREPARING','PICKED','DELIVERED'),
                       OrderEnd datetime,
                       Total float,
                       DeliveryAddress varchar(100) NOT NULL,
                       FOREIGN KEY (CustomerID) REFERENCES Customer(CustomerID),
                       FOREIGN KEY (RiderID) REFERENCES Rider(RiderID)
);

CREATE TABLE OrderDetail (
                             OrderID varchar(50) NOT NULL,
                             MealID varchar(50) NOT NULL,
                             MealQty INT,
                             MealPrice FLOAT,
                             FOREIGN KEY (OrderID) REFERENCES Orders(OrderID),
                             FOREIGN KEY (OrderID) REFERENCES Orders(OrderID),
                             FOREIGN KEY (MealID) REFERENCES Meal(MealID)
);

CREATE TABLE CustomerSessions (
                          SessionID varchar(50) PRIMARY KEY,
                          CustomerID varchar(50) NOT NULL,
                          SessionExpiry datetime,
                          FOREIGN KEY (CustomerID) REFERENCES Customer(CustomerID)
);

CREATE TABLE VendorSessions (
                                  SessionID varchar(50) PRIMARY KEY,
                                  VendorID varchar(50) NOT NULL,
                                  SessionExpiry datetime,
                                  FOREIGN KEY (VendorID) REFERENCES Vendor(VendorID)
);

-- Insert dummy data into the Vendor table
INSERT INTO Vendor (VendorID, VendorName, Address, IsOpen, IsDiscountOpen, DiscountStart, DiscountEnd, Password, VendorImage)
VALUES
    ('V001', 'Healthy Bites', '123 Green Street', TRUE, FALSE, NULL, NULL, 'password123', 'V001.jpg'),
    ('V002', 'Spice Paradise', '456 Flavor Ave', TRUE, TRUE, '2024-12-20 10:00:00', '2024-12-20 18:00:00','password123', 'V002.jpg'),
    ('V003', 'Dessert Haven', '789 Sweet Lane', FALSE, FALSE, NULL, NULL,'password123','V003.jpg');

-- Insert dummy data into the Meal table for each vendor
INSERT INTO Meal (MealID, VendorID, MealName, Description, Price, Availability, SustainabilityCreditScore, MealImage)
VALUES
-- Meals for Vendor V001
('M001', 'V001', 'Quinoa Salad', 'Healthy salad with quinoa, veggies, and dressing', 8.99, TRUE, 85, 'M001.jpg'),
('M002', 'V001', 'Grilled Chicken Wrap', 'Wrap with grilled chicken and veggies', 7.50, TRUE, 75,'M002.jpg'),
('M003', 'V001', 'Vegetable Soup', 'Warm soup with fresh vegetables', 5.99, TRUE, 80,'M003.jpg'),
-- Meals for Vendor V002
('M004', 'V002', 'Spicy Chicken Curry', 'Rich and spicy chicken curry', 12.50, TRUE, 70,'M004.jpg'),
('M005', 'V002', 'Vegetable Biryani', 'Aromatic rice with mixed vegetables and spices', 10.00, TRUE, 78,'M005.jpg'),
('M006', 'V002', 'Paneer Butter Masala', 'Creamy Indian curry with paneer cubes', 11.00, TRUE, 65,'M006.jpg'),
-- Meals for Vendor V003
('M007', 'V003', 'Chocolate Lava Cake', 'Molten chocolate dessert', 6.50, FALSE, 55,'M007.jpg'),
('M008', 'V003', 'Vanilla Ice Cream', 'Classic vanilla ice cream scoop', 4.00, TRUE, 60,'M008.jpg'),
('M009', 'V003', 'Apple Pie', 'Warm apple pie with cinnamon', 5.00, TRUE, 50,'M009.jpg'),
('M010', 'V003', 'Cheesecake', 'Creamy New York-style cheesecake', 6.00, TRUE, 60,'M010.jpg');


INSERT INTO Customer (CustomerID, CustomerName, Address, Password)
VALUES
    ('C001','Adam','Jurong West Blk 45', 'password123'),
    ('C002','Sandy','Jurong West Blk 20', 'password123'),
    ('C003','Bert','Jurong West Blk 21', 'password123');

INSERT INTO Rider (RiderID, RiderName, VehiclePlate, Availability)
VALUES
    ('R001','Rider1','SBS123',true),
    ('R002','Rider2','SHY3Y3Y', false),
    ('R003','Rider3','JYJ333',true);


