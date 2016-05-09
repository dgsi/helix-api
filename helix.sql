-- MySQL dump 10.13  Distrib 5.7.11, for osx10.10 (x86_64)
--
-- Host: localhost    Database: helix
-- ------------------------------------------------------
-- Server version	5.7.11

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `opportunitylist`
--

DROP TABLE IF EXISTS `opportunitylist`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `opportunitylist` (
  `OpportunityID` int(11) NOT NULL AUTO_INCREMENT,
  `OpportunityDate` varchar(45) DEFAULT NULL,
  `AccountName` varchar(45) DEFAULT NULL,
  `AccountOwner` varchar(45) DEFAULT NULL,
  `ActivityType` varchar(45) DEFAULT NULL,
  `Remarks` varchar(999) DEFAULT NULL,
  `BudgetAmt` decimal(14,2) DEFAULT NULL,
  `SeqID` int(11) DEFAULT NULL,
  `date_created` datetime DEFAULT NULL,
  `date_updated` datetime DEFAULT NULL,
  PRIMARY KEY (`OpportunityID`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `opportunitylist`
--

LOCK TABLES `opportunitylist` WRITE;
/*!40000 ALTER TABLE `opportunitylist` DISABLE KEYS */;
INSERT INTO `opportunitylist` VALUES (2,'2015-12-12','shanename','shaneown','typeas','remarks',200.00,10,NULL,NULL),(3,'2015-12-12','AccountName','AccountOwner','ActivityType','Remarks',200.00,10,NULL,NULL),(5,'16-04-26','SHANE',NULL,'marian','some remarks',200.00,NULL,'2016-04-26 15:57:02',NULL),(6,'$today','$AccountName','$AccountOwner','$ActivityType','$Remarks',0.00,NULL,'2016-04-26 16:08:17',NULL),(7,'1616-0404-2626','SHANE','RASHAN','marian','some remarks',200.00,NULL,'2016-04-26 16:08:29',NULL),(9,'2016-0404-2626','SHANE','RASHAN','marian','some remarks',200.00,NULL,'2016-04-26 16:09:04',NULL),(10,'2016-04-26','SHANE','RASHAN','marian','some remarks',200.00,NULL,'2016-04-26 16:09:18',NULL);
/*!40000 ALTER TABLE `opportunitylist` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payment_request`
--

DROP TABLE IF EXISTS `payment_request`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `payment_request` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `TranID` varchar(50) NOT NULL,
  `TranDate` date DEFAULT NULL,
  `Payee` varchar(45) DEFAULT NULL,
  `Particulars` varchar(45) DEFAULT NULL,
  `ProjectID` varchar(45) DEFAULT NULL,
  `Cur` varchar(45) DEFAULT NULL,
  `Amount` varchar(45) DEFAULT NULL,
  `Status` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payment_request`
--

LOCK TABLES `payment_request` WRITE;
/*!40000 ALTER TABLE `payment_request` DISABLE KEYS */;
/*!40000 ALTER TABLE `payment_request` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `purchase_request`
--

DROP TABLE IF EXISTS `purchase_request`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `purchase_request` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `TranID` varchar(45) DEFAULT NULL,
  `TranDate` date DEFAULT NULL,
  `Particulars` varchar(45) DEFAULT NULL,
  `ProjectID` varchar(100) DEFAULT NULL,
  `Cur` varchar(45) DEFAULT NULL,
  `EstimatedCost` varchar(60) DEFAULT NULL,
  `Status` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `purchase_request`
--

LOCK TABLES `purchase_request` WRITE;
/*!40000 ALTER TABLE `purchase_request` DISABLE KEYS */;
/*!40000 ALTER TABLE `purchase_request` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tbl_user`
--

DROP TABLE IF EXISTS `tbl_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tbl_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `clientid` varchar(45) DEFAULT NULL,
  `username` varchar(45) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `companyid` varchar(45) DEFAULT NULL,
  `date_created` datetime DEFAULT NULL,
  `date_updated` datetime DEFAULT NULL,
  `status` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tbl_user`
--

LOCK TABLES `tbl_user` WRITE;
/*!40000 ALTER TABLE `tbl_user` DISABLE KEYS */;
INSERT INTO `tbl_user` VALUES (20,'20160000001','ned','a0633f217ab7dc060292ec5c78a4474469999876e1d84ef2f306bdd8','flanders','2016-05-08 20:35:25','2016-05-08 21:38:54','active'),(21,'20160000002','ted','2da45dedabbdb5faa68039acdb644dbfe2dbb4773315fba0556a2b13','flanders.com','2016-05-08 20:51:43','2016-05-08 20:51:43','active');
/*!40000 ALTER TABLE `tbl_user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-05-09  9:46:51
