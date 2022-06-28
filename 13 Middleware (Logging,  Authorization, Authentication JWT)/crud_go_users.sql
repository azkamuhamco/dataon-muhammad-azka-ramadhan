-- MySQL dump 10.13  Distrib 8.0.29, for Win64 (x86_64)
--
-- Host: localhost    Database: crud_go
-- ------------------------------------------------------
-- Server version	8.0.29

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `token` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'Azka Ramadhan','azka@ramadhan.com','mantap','2022-06-22 14:06:59','2022-06-22 14:06:59',NULL,NULL),(2,'Yoga Ramadhan','ramadhan@yoga.com','yoyoyogayoda','2022-06-22 14:07:35','2022-06-22 21:18:32',NULL,NULL),(3,'AzkaRamadhan','azka@ramadhan.com','mantap','2022-06-22 14:12:15','2022-06-22 14:12:15',NULL,NULL),(4,'Hotte','hotte@rasa.com','777422','2022-06-22 01:41:07','2022-06-22 01:41:07',NULL,NULL),(5,'Hotte','hotte@rasa.com','777422','2022-06-22 01:41:10','2022-06-22 01:41:10','2022-06-22 15:11:31',NULL),(6,'AzkaRamadhan','azka@ramadhan.com','mantap','2022-06-22 14:13:06','2022-06-22 14:13:06',NULL,NULL),(7,'Karim','azka@karim.com','mantap','2022-06-22 14:28:21','2022-06-28 10:31:00',NULL,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTYzOTA2NjAsInVzZXJJZCI6N30.q_sMNlZZedLMQRh6Qb8E0HE2skiFSCZSZ4wZPPXV3eY'),(8,'Gatte','gatte@rasa.com','1113333111','2022-06-22 02:06:49','2022-06-22 02:06:49',NULL,NULL),(9,'Yoda','yoda@a.com','eaea','2022-06-22 02:08:18','2022-06-22 02:08:18',NULL,NULL),(10,'Hette','hette@rasa.com','4345435','2022-06-22 02:11:26','2022-06-22 02:11:26',NULL,NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-06-28 10:35:29
