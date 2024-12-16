CREATE USER 'user'@'%' IDENTIFIED
WITH mysql_native_password BY 'strongpassword' ;

GRANT ALL PRIVILEGES ON *.* TO 'user'@'%' ;

FLUSH PRIVILEGES;

create database dealsDB;
use dealsDB;

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

CREATE TABLE Vendor (
                        VendorID varchar(50) NOT NULL PRIMARY KEY,
                        VendorName varchar(50) NOT NULL,
                        Address varchar(50) NOT NULL,
                        IsOpen boolean,
                        IsDiscountOpen boolean NOT NULL DEFAULT FALSE,
                        DiscountStart datetime,
                        DiscountEnd datetime
);

CREATE TABLE Rider (
                       RiderID varchar(50) NOT NULL PRIMARY KEY,
                       RiderName varchar(50) NOT NULL,
                       VehiclePlate varchar(50) NOT NULL,
                       Availability boolean NOT NULL DEAFULT FALSE
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
                       OrderStatus varchar(20) ENUM('CART','PENDING','ORDERRECIVEVD','GROUPORDER','PREPARING','PICKED','DELIVERED',)
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

CREATE TABLE Users (
                       UserID varchar(50) NOT NULL PRIMARY KEY,
                       UserName varchar(50) NOT NULL,
                       Password varchar(50),
                       Role varchar(10) NOT NULL ENUM('VENDOR','CUSTOMER','RIDER')
                         FOREIGN KEY (UserID) REFERENCES Sessions(UserID)
);

CREATE TABLE Sessions (
                          SessionID varchar(50) PRIMARY KEY,
                          UserID varchar(50) NOT NULL,
                          SessionExpiry datetime,
                          FOREIGN KEY (UserID) REFERENCES Vendor(VendorID),
                          FOREIGN KEY (UserID) REFERENCES Customer(CustomerID)
);
