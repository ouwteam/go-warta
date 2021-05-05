-- --------------------------------------------------------
-- Host:                         localhost
-- Server version:               5.7.24 - MySQL Community Server (GPL)
-- Server OS:                    Win64
-- HeidiSQL Version:             10.2.0.5599
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- Dumping structure for table db_warta.channels
DROP TABLE IF EXISTS `channels`;
CREATE TABLE IF NOT EXISTS `channels` (
  `code` varchar(12) NOT NULL,
  `name` varchar(120) DEFAULT NULL,
  PRIMARY KEY (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Dumping data for table db_warta.channels: ~0 rows (approximately)
DELETE FROM `channels`;
/*!40000 ALTER TABLE `channels` DISABLE KEYS */;
INSERT INTO `channels` (`code`, `name`) VALUES
	('DMS_ERR', 'ERROR REPORT'),
	('DMS_LOG', 'REPORT UPDATE/DELETE LOG'),
	('DMS_NOTIF', 'DMS NOTIFICATION');
/*!40000 ALTER TABLE `channels` ENABLE KEYS */;

-- Dumping structure for table db_warta.update_record
DROP TABLE IF EXISTS `update_record`;
CREATE TABLE IF NOT EXISTS `update_record` (
  `last_update_id` bigint(20) unsigned DEFAULT NULL,
  `last_update_at` bigint(20) unsigned DEFAULT NULL
) ENGINE=MEMORY DEFAULT CHARSET=latin1;

-- Dumping data for table db_warta.update_record: 1 rows
DELETE FROM `update_record`;
/*!40000 ALTER TABLE `update_record` DISABLE KEYS */;
INSERT INTO `update_record` (`last_update_id`, `last_update_at`) VALUES
	(354751268, 1620173547);
/*!40000 ALTER TABLE `update_record` ENABLE KEYS */;

-- Dumping structure for table db_warta.users
DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS `users` (
  `chat_id` bigint(20) NOT NULL,
  `username` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`chat_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Dumping data for table db_warta.users: ~0 rows (approximately)
DELETE FROM `users`;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` (`chat_id`, `username`) VALUES
	(909010155, 'kurniawan11');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;

-- Dumping structure for table db_warta.user_channels
DROP TABLE IF EXISTS `user_channels`;
CREATE TABLE IF NOT EXISTS `user_channels` (
  `chat_id` bigint(20) DEFAULT NULL,
  `channel_code` varchar(12) DEFAULT NULL,
  KEY `chat_id_channel_code` (`chat_id`,`channel_code`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Dumping data for table db_warta.user_channels: ~0 rows (approximately)
DELETE FROM `user_channels`;
/*!40000 ALTER TABLE `user_channels` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_channels` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
