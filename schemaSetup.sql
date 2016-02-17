DROP DATABASE IF EXISTS restaurantPaymentAPI;

CREATE DATABASE restaurantPaymentAPI;

use restaurantPaymentAPI;

--
-- Table structure for table `payment`
--

DROP TABLE IF EXISTS `payments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `payments` (
	`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
	`amount` int(11) DEFAULT '0',
	`tipAmount` int(11) DEFAULT '0',
	`tableNo` int(11) DEFAULT '0',
	`paymentReferenceId` int(11) DEFAULT '0',
	`restaurantLocationId` int(11) DEFAULT '0',
	`cardType` ENUM('visa', 'mastercard', 'americanExpress'),
	`created`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;


INSERT into payments (`amount`, `tipAmount`, `tableNo`, `paymentReferenceId`, `restaurantLocationId`, `cardType`) VALUES (50, 10, 100, 400, 12, "visa");

DROP TABLE IF EXISTS `restaurantLocations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `restaurantLocations` (
	`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
	`streetAddress` varchar(200) NOT NULL,
	`county` varchar(200) DEFAULT NULL,
	`created`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;

INSERT into restaurantLocations (streetAddress, county) VALUES ("Marble Arch", "London");
INSERT into restaurantLocations (streetAddress, county) VALUES ("Covent Garden", "London");
