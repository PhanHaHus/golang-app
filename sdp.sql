/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50505
Source Host           : localhost:3306
Source Database       : sdp

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2017-04-12 13:44:48
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for accepting_hosts
-- ----------------------------
DROP TABLE IF EXISTS `accepting_hosts`;
CREATE TABLE `accepting_hosts` (
  `accepting_host_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `password` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `description` varchar(200) COLLATE utf8_unicode_ci DEFAULT NULL,
  `last_login_time` datetime DEFAULT NULL,
  `enabled` tinyint(1) NOT NULL,
  `created_by_id` int(11) NOT NULL,
  `created_time` datetime NOT NULL,
  `updated_time` datetime NOT NULL,
  PRIMARY KEY (`accepting_host_id`),
  KEY `created_by_id` (`created_by_id`),
  CONSTRAINT `accepting_hosts_ibfk_1` FOREIGN KEY (`created_by_id`) REFERENCES `administrators` (`administrator_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of accepting_hosts
-- ----------------------------

-- ----------------------------
-- Table structure for access_rules
-- ----------------------------
DROP TABLE IF EXISTS `access_rules`;
CREATE TABLE `access_rules` (
  `access_rule_id` int(11) NOT NULL AUTO_INCREMENT,
  `application_id` int(11) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `device_id` int(11) DEFAULT NULL,
  `group_id` int(11) DEFAULT NULL,
  `access_rule_type` enum('Block','Accept') COLLATE utf8_unicode_ci NOT NULL,
  `description` varchar(200) COLLATE utf8_unicode_ci DEFAULT NULL,
  `enabled` tinyint(1) NOT NULL,
  `created_by_id` int(11) DEFAULT NULL,
  `created_time` datetime NOT NULL,
  `updated_time` datetime NOT NULL,
  PRIMARY KEY (`access_rule_id`),
  KEY `group_id` (`group_id`),
  KEY `device_id` (`device_id`),
  KEY `user_id` (`user_id`),
  KEY `application_id` (`application_id`),
  KEY `created_by_id` (`created_by_id`),
  CONSTRAINT `access_rules_ibfk_1` FOREIGN KEY (`group_id`) REFERENCES `groups` (`group_id`),
  CONSTRAINT `access_rules_ibfk_2` FOREIGN KEY (`device_id`) REFERENCES `devices` (`device_id`),
  CONSTRAINT `access_rules_ibfk_3` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`),
  CONSTRAINT `access_rules_ibfk_4` FOREIGN KEY (`application_id`) REFERENCES `applications` (`application_id`),
  CONSTRAINT `access_rules_ibfk_5` FOREIGN KEY (`created_by_id`) REFERENCES `administrators` (`administrator_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of access_rules
-- ----------------------------

-- ----------------------------
-- Table structure for access_rule_device_tags
-- ----------------------------
DROP TABLE IF EXISTS `access_rule_device_tags`;
CREATE TABLE `access_rule_device_tags` (
  `access_rule_device_tag_id` int(11) NOT NULL AUTO_INCREMENT,
  `access_rule_id` int(11) NOT NULL,
  `device_tag` varchar(30) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`access_rule_device_tag_id`),
  KEY `access_rule_id` (`access_rule_id`),
  CONSTRAINT `access_rule_device_tags_ibfk_1` FOREIGN KEY (`access_rule_id`) REFERENCES `access_rules` (`access_rule_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of access_rule_device_tags
-- ----------------------------

-- ----------------------------
-- Table structure for access_rule_user_tags
-- ----------------------------
DROP TABLE IF EXISTS `access_rule_user_tags`;
CREATE TABLE `access_rule_user_tags` (
  `access_rule_user_tag_id` int(11) NOT NULL AUTO_INCREMENT,
  `access_rule_id` int(11) NOT NULL,
  `user_tag` varchar(30) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`access_rule_user_tag_id`),
  KEY `access_rule_id` (`access_rule_id`),
  CONSTRAINT `access_rule_user_tags_ibfk_1` FOREIGN KEY (`access_rule_id`) REFERENCES `access_rules` (`access_rule_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of access_rule_user_tags
-- ----------------------------

-- ----------------------------
-- Table structure for administrators
-- ----------------------------
DROP TABLE IF EXISTS `administrators`;
CREATE TABLE `administrators` (
  `administrator_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `email` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `password` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `description` varchar(200) COLLATE utf8_unicode_ci DEFAULT NULL,
  `permission` enum('AH_Admin','Super_Admin','System_Admin') COLLATE utf8_unicode_ci NOT NULL,
  `accepting_host_id` int(11) DEFAULT NULL,
  `enabled` tinyint(1) NOT NULL,
  `created_by_id` int(11) DEFAULT NULL,
  `created_time` datetime NOT NULL,
  `updated_time` datetime NOT NULL,
  PRIMARY KEY (`administrator_id`),
  KEY `accepting_host_id` (`accepting_host_id`),
  KEY `created_by_id` (`created_by_id`),
  CONSTRAINT `administrators_ibfk_1` FOREIGN KEY (`accepting_host_id`) REFERENCES `accepting_hosts` (`accepting_host_id`),
  CONSTRAINT `administrators_ibfk_2` FOREIGN KEY (`created_by_id`) REFERENCES `administrators` (`administrator_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of administrators
-- ----------------------------

-- ----------------------------
-- Table structure for applications
-- ----------------------------
DROP TABLE IF EXISTS `applications`;
CREATE TABLE `applications` (
  `application_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `description` varchar(200) COLLATE utf8_unicode_ci DEFAULT NULL,
  `application_type` enum('SSH','TCP','HTTP','HTTPS') COLLATE utf8_unicode_ci NOT NULL,
  `accepting_host_id` int(11) DEFAULT NULL,
  `ip` varchar(20) COLLATE utf8_unicode_ci NOT NULL,
  `port` int(11) NOT NULL,
  `host_name` varchar(40) COLLATE utf8_unicode_ci DEFAULT NULL,
  `is_valid_user_required` tinyint(1) DEFAULT NULL,
  `is_valid_device_required` tinyint(1) DEFAULT NULL,
  `enabled` tinyint(1) NOT NULL,
  `created_by_id` int(11) DEFAULT NULL,
  `created_time` datetime NOT NULL,
  `updated_time` datetime NOT NULL,
  PRIMARY KEY (`application_id`),
  KEY `created_by_id` (`created_by_id`),
  KEY `accepting_host_id` (`accepting_host_id`),
  CONSTRAINT `applications_ibfk_1` FOREIGN KEY (`created_by_id`) REFERENCES `administrators` (`administrator_id`),
  CONSTRAINT `applications_ibfk_2` FOREIGN KEY (`accepting_host_id`) REFERENCES `accepting_hosts` (`accepting_host_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of applications
-- ----------------------------

-- ----------------------------
-- Table structure for certificate_requests
-- ----------------------------
DROP TABLE IF EXISTS `certificate_requests`;
CREATE TABLE `certificate_requests` (
  `certificate_request_id` int(11) NOT NULL AUTO_INCREMENT,
  `hardware_hash` varchar(200) COLLATE utf8_unicode_ci NOT NULL,
  `device_id` int(11) NOT NULL,
  `initiating_host_id` int(11) DEFAULT NULL,
  `request_token` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `status` enum('CADeclined','Pending','Approved','Declined','CAApproved') COLLATE utf8_unicode_ci NOT NULL,
  `remote_request_id` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL,
  `csr_content` varchar(2048) COLLATE utf8_unicode_ci NOT NULL,
  `public_key_content` varchar(2048) COLLATE utf8_unicode_ci DEFAULT NULL,
  `expired_time` datetime DEFAULT NULL,
  `approve_note` varchar(400) COLLATE utf8_unicode_ci DEFAULT NULL,
  `approved_by_id` int(11) DEFAULT NULL,
  `approved_time` datetime DEFAULT NULL,
  `created_time` datetime NOT NULL,
  `updated_time` datetime NOT NULL,
  PRIMARY KEY (`certificate_request_id`),
  KEY `approved_by_id` (`approved_by_id`),
  KEY `device_id` (`device_id`),
  KEY `initiating_host_id` (`initiating_host_id`),
  CONSTRAINT `certificate_requests_ibfk_1` FOREIGN KEY (`approved_by_id`) REFERENCES `administrators` (`administrator_id`),
  CONSTRAINT `certificate_requests_ibfk_2` FOREIGN KEY (`device_id`) REFERENCES `devices` (`device_id`),
  CONSTRAINT `certificate_requests_ibfk_3` FOREIGN KEY (`initiating_host_id`) REFERENCES `initiating_hosts` (`initiating_host_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of certificate_requests
-- ----------------------------

-- ----------------------------
-- Table structure for devices
-- ----------------------------
DROP TABLE IF EXISTS `devices`;
CREATE TABLE `devices` (
  `device_id` int(11) NOT NULL AUTO_INCREMENT,
  `device_type` enum('Tablet','Desktop','Laptop','Smartphone') COLLATE utf8_unicode_ci NOT NULL,
  `name` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `hardware_hash` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `description` varchar(200) COLLATE utf8_unicode_ci DEFAULT NULL,
  `enabled` tinyint(1) NOT NULL,
  `created_by_id` int(11) DEFAULT NULL,
  `created_time` datetime NOT NULL,
  `updated_time` datetime NOT NULL,
  PRIMARY KEY (`device_id`),
  KEY `created_by_id` (`created_by_id`),
  CONSTRAINT `devices_ibfk_1` FOREIGN KEY (`created_by_id`) REFERENCES `administrators` (`administrator_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of devices
-- ----------------------------

-- ----------------------------
-- Table structure for device_tags
-- ----------------------------
DROP TABLE IF EXISTS `device_tags`;
CREATE TABLE `device_tags` (
  `device_tag_id` int(11) NOT NULL AUTO_INCREMENT,
  `device_id` int(11) NOT NULL,
  `tag` varchar(30) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`device_tag_id`),
  KEY `device_id` (`device_id`),
  CONSTRAINT `device_tags_ibfk_1` FOREIGN KEY (`device_id`) REFERENCES `devices` (`device_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of device_tags
-- ----------------------------

-- ----------------------------
-- Table structure for groups
-- ----------------------------
DROP TABLE IF EXISTS `groups`;
CREATE TABLE `groups` (
  `group_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `parent_group_id` int(11) DEFAULT NULL,
  `description` varchar(200) COLLATE utf8_unicode_ci DEFAULT NULL,
  `enabled` tinyint(1) NOT NULL,
  `created_time` datetime NOT NULL,
  `updated_time` datetime NOT NULL,
  PRIMARY KEY (`group_id`),
  KEY `parent_group_id` (`parent_group_id`),
  CONSTRAINT `groups_ibfk_1` FOREIGN KEY (`parent_group_id`) REFERENCES `groups` (`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of groups
-- ----------------------------

-- ----------------------------
-- Table structure for initiating_hosts
-- ----------------------------
DROP TABLE IF EXISTS `initiating_hosts`;
CREATE TABLE `initiating_hosts` (
  `initiating_host_id` int(11) NOT NULL AUTO_INCREMENT,
  `password` varchar(20) COLLATE utf8_unicode_ci NOT NULL,
  `name` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `description` varchar(200) COLLATE utf8_unicode_ci DEFAULT NULL,
  `last_login_time` datetime NOT NULL,
  `enabled` tinyint(1) NOT NULL,
  `created_by_id` int(11) NOT NULL,
  `created_time` datetime NOT NULL,
  `updated_time` datetime NOT NULL,
  PRIMARY KEY (`initiating_host_id`),
  KEY `created_by_id` (`created_by_id`),
  CONSTRAINT `initiating_hosts_ibfk_1` FOREIGN KEY (`created_by_id`) REFERENCES `administrators` (`administrator_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of initiating_hosts
-- ----------------------------

-- ----------------------------
-- Table structure for reminders
-- ----------------------------
DROP TABLE IF EXISTS `reminders`;
CREATE TABLE `reminders` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `message` varchar(1024) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of reminders
-- ----------------------------
INSERT INTO `reminders` VALUES ('1', 'aS', null, null, null);
INSERT INTO `reminders` VALUES ('2', 'asdfasfd', null, null, null);
INSERT INTO `reminders` VALUES ('12', 'new ', '2017-04-08 07:54:54', '2017-04-08 07:56:06', null);
INSERT INTO `reminders` VALUES ('13', 'remind me ', '2017-04-12 11:12:06', '2017-04-12 11:12:06', null);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `email` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `description` varchar(200) COLLATE utf8_unicode_ci DEFAULT NULL,
  `enabled` tinyint(1) NOT NULL,
  `created_time` datetime NOT NULL,
  `updated_time` datetime NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES ('3', 'Nguyen Thanh Tung', 'tungnt751@viettel.com.vn', 'GPDN', '0', '2017-03-03 11:53:39', '2017-03-03 11:53:39');

-- ----------------------------
-- Table structure for user_groups
-- ----------------------------
DROP TABLE IF EXISTS `user_groups`;
CREATE TABLE `user_groups` (
  `user_id` int(11) NOT NULL,
  `group_id` int(11) NOT NULL,
  PRIMARY KEY (`user_id`,`group_id`),
  KEY `group_id` (`group_id`),
  CONSTRAINT `user_groups_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`),
  CONSTRAINT `user_groups_ibfk_2` FOREIGN KEY (`group_id`) REFERENCES `groups` (`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of user_groups
-- ----------------------------

-- ----------------------------
-- Table structure for user_tags
-- ----------------------------
DROP TABLE IF EXISTS `user_tags`;
CREATE TABLE `user_tags` (
  `user_tag_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `tag` varchar(30) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`user_tag_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `user_tags_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of user_tags
-- ----------------------------
