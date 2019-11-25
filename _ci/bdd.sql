/*

    DATABASE IS STILL UNDER DEVELOPEMENT

*/

CREATE TABLE `user` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `firstname` varchar(255),
  `lastname` varchar(255),
  `money` float4,
  `level` integer,
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `level` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255)
);

CREATE TABLE `staff` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `email` varchar(255),
  `password` varchar(255),
  `role` integer,
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `role` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255)
);

CREATE TABLE `product` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `price` float4,
  `quantity` integer,
  `image` varchar(255),
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `purchase` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `user_id` integer,
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `purchase_item` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `purchase_id` integer,
  `product_id` integer
);

ALTER TABLE `user` ADD FOREIGN KEY (`level`) REFERENCES `level` (`id`);

ALTER TABLE `staff` ADD FOREIGN KEY (`role`) REFERENCES `role` (`id`);

ALTER TABLE `purchase` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `purchase_item` ADD FOREIGN KEY (`purchase_id`) REFERENCES `purchase` (`id`);

ALTER TABLE `purchase_item` ADD FOREIGN KEY (`product_id`) REFERENCES `product` (`id`);
