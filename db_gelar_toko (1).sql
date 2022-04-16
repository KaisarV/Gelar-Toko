-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Apr 15, 2022 at 03:39 PM
-- Server version: 10.4.22-MariaDB
-- PHP Version: 8.1.2

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_gelar_toko`
--

-- --------------------------------------------------------

--
-- Table structure for table `carts`
--

CREATE TABLE `carts` (
  `Id` int(11) NOT NULL,
  `User_Id` int(11) DEFAULT NULL,
  `Product_Id` int(11) DEFAULT NULL,
  `Quantity` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `carts`
--

INSERT INTO `carts` (`Id`, `User_Id`, `Product_Id`, `Quantity`) VALUES
(18, 9, 7, 3),
(20, 1, 9, 1),
(21, 9, 1, 1),
(22, 9, 1, 50);

-- --------------------------------------------------------

--
-- Table structure for table `chat`
--

CREATE TABLE `chat` (
  `Id` int(11) NOT NULL,
  `Sender_Id` int(11) DEFAULT NULL,
  `Receiver_Id` int(11) DEFAULT NULL,
  `Customer_Id` int(11) DEFAULT NULL,
  `Chat` text DEFAULT NULL,
  `Date` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `chat`
--

INSERT INTO `chat` (`Id`, `Sender_Id`, `Receiver_Id`, `Customer_Id`, `Chat`, `Date`) VALUES
(1, 9, 1, 9, 'tes', '2022-04-15 12:55:47');

-- --------------------------------------------------------

--
-- Table structure for table `feedbacks`
--

CREATE TABLE `feedbacks` (
  `Id` int(11) NOT NULL,
  `User_Id` int(11) DEFAULT NULL,
  `Feedback` text DEFAULT NULL,
  `Date` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `feedbacks`
--

INSERT INTO `feedbacks` (`Id`, `User_Id`, `Feedback`, `Date`) VALUES
(1, 9, 'Aplikasi ini sangat baik\r\n', '2022-03-31 17:43:58'),
(2, 9, 'Halo', '2022-03-31 17:58:36'),
(3, 9, 'Halo', '2022-03-31 18:03:00'),
(4, 9, 'keren', '2022-04-01 18:20:37');

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `Id` int(11) NOT NULL,
  `Name` varchar(255) DEFAULT NULL,
  `Category` varchar(255) DEFAULT NULL,
  `Price` int(11) DEFAULT NULL,
  `Store_Id` int(11) DEFAULT NULL,
  `Is_Blocked` int(11) DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`Id`, `Name`, `Category`, `Price`, `Store_Id`, `Is_Blocked`) VALUES
(1, 'Nike Hyper Venom', 'Olahraga', 1200000, 1, 0),
(2, 'Mouse G-402', 'Komputer dan Elektronik', 350000, 2, 0),
(3, 'Knitted Sweater', 'Pakaian', 70000, 3, 0),
(4, 'Celana Legging', 'Pakaian', 100000, 3, 0),
(5, 'Helm Bogo', 'Hobi', 580000, 4, 0),
(6, 'Keyboard K-580', 'Komputer dan Elektronik', 375000, 2, 0),
(7, 'Masker KN-85', 'Kesehatan', 150000, 5, 0),
(8, 'Hand Sanitizer Carex', 'Kesehatan', 25000, 5, 0),
(9, 'Nike Air Jordan 1 Retro', 'Olahraga', 2600000, 1, 0),
(10, 'Helm KYT', 'Hobi', 700000, 4, 0),
(11, 'Masker N95', 'Kesehatan', 60000, 5, 0),
(14, 'Jam Tangan', 'Aksesoris', 500000, 11, 0),
(15, 'Topi', 'Aksesoris', 500000, 11, 0),
(16, 'Topi', 'Aksesoris', 500000, 11, 0),
(17, 'Headphone', 'Gaming', 100000, 11, 0),
(18, 'Topi', 'Aksesoris', 500000, 11, 0);

-- --------------------------------------------------------

--
-- Table structure for table `product_reviews`
--

CREATE TABLE `product_reviews` (
  `Id` int(11) NOT NULL,
  `User_Id` int(11) DEFAULT NULL,
  `Product_Id` int(11) DEFAULT NULL,
  `Review` text DEFAULT NULL,
  `Rating` int(11) DEFAULT NULL,
  `Date` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `product_reviews`
--

INSERT INTO `product_reviews` (`Id`, `User_Id`, `Product_Id`, `Review`, `Rating`, `Date`) VALUES
(1, 9, 1, 'Barang bagus, tapi kurang', 4, '2022-04-01 18:14:52');

-- --------------------------------------------------------

--
-- Table structure for table `stores`
--

CREATE TABLE `stores` (
  `Id` int(11) NOT NULL,
  `User_Id` int(11) DEFAULT NULL,
  `Name` varchar(255) DEFAULT NULL,
  `Address` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `stores`
--

INSERT INTO `stores` (`Id`, `User_Id`, `Name`, `Address`) VALUES
(1, 1, 'Nike Store', 'Jakarta Barat'),
(2, 2, 'Logitech', 'Kota Bogor'),
(3, 3, 'Rily Store', 'Bekasi Timur'),
(4, 4, 'Istana Helm', 'Kota Batu'),
(5, 5, 'Juragan Masker', 'Blitar'),
(11, 9, 'Continue Gaming', 'J');

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `Id` int(11) NOT NULL,
  `User_Id` int(11) DEFAULT NULL,
  `Product_Id` int(11) DEFAULT NULL,
  `Date` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `Quantity` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`Id`, `User_Id`, `Product_Id`, `Date`, `Quantity`) VALUES
(1, 9, 1, '2022-03-31 14:08:49', 2),
(2, 15, 1, '2022-04-01 15:49:32', 2),
(3, 9, 1, '2022-04-01 15:57:16', 1);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `Id` int(11) NOT NULL,
  `Name` varchar(255) DEFAULT NULL,
  `Phone` varchar(20) DEFAULT NULL,
  `Email` varchar(255) DEFAULT NULL,
  `Password` varchar(255) NOT NULL,
  `Address` varchar(20) DEFAULT NULL,
  `User_Type` int(11) NOT NULL DEFAULT 1,
  `Is_Verified` int(11) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`Id`, `Name`, `Phone`, `Email`, `Password`, `Address`, `User_Type`, `Is_Verified`) VALUES
(1, 'Wanda P', '081232874632', 'wandap@gmail.com', 'wanda123', 'Jakarta Timur', 2, 1),
(2, 'Andi Pranata', '081234278282', 'andipranata@gmail.com', '', 'Kab. Bogor', 2, 1),
(3, 'Dimas Wahyudi', '089512351818', 'dimaswhyd@gmail.com', '', 'Bekasi Timur', 2, 1),
(4, 'Alan Kusuma', '082273482682', 'alankusuma@gmail.com', '', 'Malang', 2, 1),
(5, 'Hadi Wijaya', '087324964329', 'wijayahadi@gmail.com', '', 'Blitar', 2, 1),
(6, 'Agung Tirtayasa', '087329464329', 'atirta@gmail.com', '', 'Bandung', 1, 1),
(7, 'Adi Kusuma', '08124365834', 'adikusuma@gmail.com', '', 'Jakarta Selatan', 1, 1),
(9, 'mikael', '081243268345', 'jajang@gmail.com', 'dadang123', 'Jakarta Utara', 2, 1),
(15, 'mikael jajang', '081243268345', 'dadang@gmail.com', 'jajang123', 'Jakarta Utara', 3, 1),
(17, 'admin', 'admin', 'admin@gmail.com', 'admin123', 'Bogor', 3, 1),
(19, 'kaisar', 'string', 'kaisar.valentino123@gmail.com', 'string', 'string', 1, 1);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `carts`
--
ALTER TABLE `carts`
  ADD PRIMARY KEY (`Id`),
  ADD KEY `id_user` (`User_Id`),
  ADD KEY `id_barang` (`Product_Id`);

--
-- Indexes for table `chat`
--
ALTER TABLE `chat`
  ADD PRIMARY KEY (`Id`),
  ADD KEY `SenderId` (`Sender_Id`),
  ADD KEY `ReceiverId` (`Receiver_Id`),
  ADD KEY `Customer_Id` (`Customer_Id`);

--
-- Indexes for table `feedbacks`
--
ALTER TABLE `feedbacks`
  ADD PRIMARY KEY (`Id`),
  ADD KEY `User_Id` (`User_Id`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`Id`),
  ADD KEY `id_toko` (`Store_Id`);

--
-- Indexes for table `product_reviews`
--
ALTER TABLE `product_reviews`
  ADD PRIMARY KEY (`Id`),
  ADD KEY `User_Id` (`User_Id`),
  ADD KEY `Product_Id` (`Product_Id`);

--
-- Indexes for table `stores`
--
ALTER TABLE `stores`
  ADD PRIMARY KEY (`Id`),
  ADD KEY `id_user` (`User_Id`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`Id`),
  ADD KEY `id_user` (`User_Id`),
  ADD KEY `id_barang` (`Product_Id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`Id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `carts`
--
ALTER TABLE `carts`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=23;

--
-- AUTO_INCREMENT for table `chat`
--
ALTER TABLE `chat`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `feedbacks`
--
ALTER TABLE `feedbacks`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;

--
-- AUTO_INCREMENT for table `product_reviews`
--
ALTER TABLE `product_reviews`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `stores`
--
ALTER TABLE `stores`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=20;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `carts`
--
ALTER TABLE `carts`
  ADD CONSTRAINT `carts_ibfk_1` FOREIGN KEY (`User_Id`) REFERENCES `users` (`Id`),
  ADD CONSTRAINT `carts_ibfk_2` FOREIGN KEY (`Product_Id`) REFERENCES `products` (`Id`);

--
-- Constraints for table `chat`
--
ALTER TABLE `chat`
  ADD CONSTRAINT `chat_ibfk_1` FOREIGN KEY (`Sender_Id`) REFERENCES `users` (`Id`),
  ADD CONSTRAINT `chat_ibfk_2` FOREIGN KEY (`Receiver_Id`) REFERENCES `users` (`Id`),
  ADD CONSTRAINT `chat_ibfk_3` FOREIGN KEY (`Customer_Id`) REFERENCES `users` (`Id`);

--
-- Constraints for table `feedbacks`
--
ALTER TABLE `feedbacks`
  ADD CONSTRAINT `feedbacks_ibfk_1` FOREIGN KEY (`User_Id`) REFERENCES `users` (`Id`);

--
-- Constraints for table `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `products_ibfk_1` FOREIGN KEY (`Store_Id`) REFERENCES `stores` (`Id`);

--
-- Constraints for table `product_reviews`
--
ALTER TABLE `product_reviews`
  ADD CONSTRAINT `product_reviews_ibfk_1` FOREIGN KEY (`User_Id`) REFERENCES `users` (`Id`),
  ADD CONSTRAINT `product_reviews_ibfk_2` FOREIGN KEY (`Product_Id`) REFERENCES `products` (`Id`);

--
-- Constraints for table `stores`
--
ALTER TABLE `stores`
  ADD CONSTRAINT `stores_ibfk_1` FOREIGN KEY (`User_Id`) REFERENCES `users` (`Id`);

--
-- Constraints for table `transactions`
--
ALTER TABLE `transactions`
  ADD CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`User_Id`) REFERENCES `users` (`Id`),
  ADD CONSTRAINT `transactions_ibfk_2` FOREIGN KEY (`Product_Id`) REFERENCES `products` (`Id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
