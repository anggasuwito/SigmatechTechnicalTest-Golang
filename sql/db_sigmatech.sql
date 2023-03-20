/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

DROP TABLE IF EXISTS `company`;
CREATE TABLE `company` (
  `id` varchar(36) NOT NULL DEFAULT (uuid()),
  `company_name` varchar(255) DEFAULT NULL,
  `company_fee` int DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `company_asset`;
CREATE TABLE `company_asset` (
  `id` varchar(36) NOT NULL DEFAULT (uuid()),
  `company_id` varchar(36) DEFAULT NULL,
  `asset_name` varchar(255) DEFAULT NULL,
  `otr_price` int DEFAULT NULL,
  `stock_availability` int DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `company_id` (`company_id`),
  CONSTRAINT `company_asset_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `company` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `transaction_history`;
CREATE TABLE `transaction_history` (
  `id` varchar(36) NOT NULL DEFAULT (uuid()),
  `user_id` varchar(36) DEFAULT NULL,
  `transaction_setting_id` varchar(36) DEFAULT NULL,
  `company_asset_id` varchar(36) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `transaction_setting_id` (`transaction_setting_id`),
  KEY `company_asset_id` (`company_asset_id`),
  CONSTRAINT `transaction_history_ibfk_1` FOREIGN KEY (`transaction_setting_id`) REFERENCES `transaction_setting` (`id`),
  CONSTRAINT `transaction_history_ibfk_2` FOREIGN KEY (`company_asset_id`) REFERENCES `company_asset` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `transaction_setting`;
CREATE TABLE `transaction_setting` (
  `id` varchar(36) NOT NULL DEFAULT (uuid()),
  `user_id` varchar(36) DEFAULT NULL,
  `tenor` int DEFAULT NULL COMMENT 'satuan per bulan',
  `max_limit` int DEFAULT NULL COMMENT 'satuan RP',
  `interest` int DEFAULT NULL COMMENT 'satuan persen',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `transaction_setting_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` varchar(36) NOT NULL DEFAULT (uuid()),
  `nik` varchar(255) DEFAULT NULL,
  `full_name` varchar(255) DEFAULT NULL,
  `legal_name` varchar(255) DEFAULT NULL,
  `birth_place` varchar(255) DEFAULT NULL,
  `birth_date` varchar(255) DEFAULT NULL,
  `salary` int DEFAULT NULL,
  `identity_photo` varchar(255) DEFAULT NULL,
  `selfie_photo` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `company` (`id`, `company_name`, `company_fee`, `address`, `created_at`, `updated_at`, `deleted_at`) VALUES
('84b64aba-c6ca-11ed-9138-00ffc2ef2645', 'PT XYZ', 5000, 'Jakarta', '2023-03-20 09:54:23', NULL, NULL);
INSERT INTO `company` (`id`, `company_name`, `company_fee`, `address`, `created_at`, `updated_at`, `deleted_at`) VALUES
('89ce84d8-c6ca-11ed-9138-00ffc2ef2645', 'PT ABC', 25000, 'Bali', '2023-03-20 09:54:31', NULL, NULL);


INSERT INTO `company_asset` (`id`, `company_id`, `asset_name`, `otr_price`, `stock_availability`, `created_at`, `updated_at`, `deleted_at`) VALUES
('a23c00c8-c6ca-11ed-9138-00ffc2ef2645', '84b64aba-c6ca-11ed-9138-00ffc2ef2645', 'Mobil', 1000000, 2, '2023-03-20 09:55:12', NULL, NULL);
INSERT INTO `company_asset` (`id`, `company_id`, `asset_name`, `otr_price`, `stock_availability`, `created_at`, `updated_at`, `deleted_at`) VALUES
('a23c3787-c6ca-11ed-9138-00ffc2ef2645', '84b64aba-c6ca-11ed-9138-00ffc2ef2645', 'Motor', 500000, 10, '2023-03-20 09:55:12', NULL, NULL);
INSERT INTO `company_asset` (`id`, `company_id`, `asset_name`, `otr_price`, `stock_availability`, `created_at`, `updated_at`, `deleted_at`) VALUES
('baf73b5b-c6ca-11ed-9138-00ffc2ef2645', '89ce84d8-c6ca-11ed-9138-00ffc2ef2645', 'Rumah', 900000, 3, '2023-03-20 09:55:54', NULL, NULL);
INSERT INTO `company_asset` (`id`, `company_id`, `asset_name`, `otr_price`, `stock_availability`, `created_at`, `updated_at`, `deleted_at`) VALUES
('baf7661f-c6ca-11ed-9138-00ffc2ef2645', '89ce84d8-c6ca-11ed-9138-00ffc2ef2645', 'Kapal', 800000, 6, '2023-03-20 09:55:54', NULL, NULL);

INSERT INTO `transaction_history` (`id`, `user_id`, `transaction_setting_id`, `company_asset_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
('c0f78366-c6d6-11ed-9138-00ffc2ef2645', '96465f3f-c6bc-11ed-9138-00ffc2ef2645', '46cdf99c-c6ca-11ed-9138-00ffc2ef2645', 'a23c00c8-c6ca-11ed-9138-00ffc2ef2645', '2023-03-20 11:21:58', NULL, NULL);


INSERT INTO `transaction_setting` (`id`, `user_id`, `tenor`, `max_limit`, `interest`, `created_at`, `updated_at`, `deleted_at`) VALUES
('31ea1487-c6ca-11ed-9138-00ffc2ef2645', '96465f3f-c6bc-11ed-9138-00ffc2ef2645', 1, 100000, 5, '2023-03-20 09:52:04', NULL, NULL);
INSERT INTO `transaction_setting` (`id`, `user_id`, `tenor`, `max_limit`, `interest`, `created_at`, `updated_at`, `deleted_at`) VALUES
('46cdf99c-c6ca-11ed-9138-00ffc2ef2645', '96465f3f-c6bc-11ed-9138-00ffc2ef2645', 2, 200000, 10, '2023-03-20 09:52:39', NULL, NULL);
INSERT INTO `transaction_setting` (`id`, `user_id`, `tenor`, `max_limit`, `interest`, `created_at`, `updated_at`, `deleted_at`) VALUES
('46ce39fe-c6ca-11ed-9138-00ffc2ef2645', '96465f3f-c6bc-11ed-9138-00ffc2ef2645', 3, 500000, 15, '2023-03-20 09:52:39', NULL, NULL);
INSERT INTO `transaction_setting` (`id`, `user_id`, `tenor`, `max_limit`, `interest`, `created_at`, `updated_at`, `deleted_at`) VALUES
('46ce4fb5-c6ca-11ed-9138-00ffc2ef2645', '96465f3f-c6bc-11ed-9138-00ffc2ef2645', 6, 700000, 20, '2023-03-20 09:52:39', NULL, NULL),
('6241acfb-c6ca-11ed-9138-00ffc2ef2645', 'f55749e9-c6c9-11ed-9138-00ffc2ef2645', 1, 1000000, 5, '2023-03-20 09:53:25', NULL, NULL),
('6241d66c-c6ca-11ed-9138-00ffc2ef2645', 'f55749e9-c6c9-11ed-9138-00ffc2ef2645', 2, 1200000, 10, '2023-03-20 09:53:25', NULL, NULL),
('6241f3fa-c6ca-11ed-9138-00ffc2ef2645', 'f55749e9-c6c9-11ed-9138-00ffc2ef2645', 3, 1500000, 15, '2023-03-20 09:53:25', NULL, NULL),
('62420ae3-c6ca-11ed-9138-00ffc2ef2645', 'f55749e9-c6c9-11ed-9138-00ffc2ef2645', 6, 2000000, 20, '2023-03-20 09:53:25', NULL, NULL);

INSERT INTO `user` (`id`, `nik`, `full_name`, `legal_name`, `birth_place`, `birth_date`, `salary`, `identity_photo`, `selfie_photo`, `created_at`, `updated_at`, `deleted_at`) VALUES
('96465f3f-c6bc-11ed-9138-00ffc2ef2645', '1111', 'Budi Susanto', 'Budi', 'Jakarta', '1996-03-29', 5000000, 'identity-photo-budi.jpeg', 'selfie-photo-budi.jpeg', '2023-03-20 08:14:39', NULL, NULL);
INSERT INTO `user` (`id`, `nik`, `full_name`, `legal_name`, `birth_place`, `birth_date`, `salary`, `identity_photo`, `selfie_photo`, `created_at`, `updated_at`, `deleted_at`) VALUES
('f55749e9-c6c9-11ed-9138-00ffc2ef2645', '2222', 'Annisa', 'Annisa', 'Bali', '1996-03-29', 10000000, 'identity-photo-annisa.jpeg', 'selfie-photo-annisa.jpeg', '2023-03-20 09:50:22', NULL, NULL);



/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;