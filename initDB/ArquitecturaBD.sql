CREATE DATABASE  IF NOT EXISTS `project` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `project`;
-- MySQL dump 10.13  Distrib 5.6.13, for osx10.6 (i386)
--
-- Host: 127.0.0.1    Database: project
-- ------------------------------------------------------
-- Server version	5.7.16
--
-- Table structure for table `projects`
--

DROP TABLE IF EXISTS `projects`;

CREATE TABLE `projects` (
  `Project_Id` int(11) NOT NULL AUTO_INCREMENT,
  `Planning_Id` varchar(400) DEFAULT '',
  `Estado` varchar(400) DEFAULT '',
  `Members` varchar(500) DEFAULT '',
  `ProjectLeader` varchar(100) NOT NULL,
  `Title` varchar(50) DEFAULT '',
  `StudyArea` varchar(400) DEFAULT '',
  `Description` varchar(200) DEFAULT '',
  PRIMARY KEY (`Project_Id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


--
-- Dumping data for table `projects`
--
-- Dump completed on 2018-10-22 14:38:05

