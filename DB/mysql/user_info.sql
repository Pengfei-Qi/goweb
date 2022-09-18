/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.235.233
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : 192.168.235.233:3306
 Source Schema         : db1

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 07/09/2022 17:24:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `age` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_info_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_info
-- ----------------------------
INSERT INTO `user_info` VALUES (1, '2022-09-06 11:39:00', '2022-09-07 17:00:16', NULL, '王老五', 20);
INSERT INTO `user_info` VALUES (2, '2022-09-06 11:39:00', '2022-09-07 17:00:16', NULL, '赵六', 21);
INSERT INTO `user_info` VALUES (3, '2022-09-06 16:14:04', '2022-09-07 17:00:16', NULL, '孙权', 22);
INSERT INTO `user_info` VALUES (4, '2022-09-06 16:14:04', '2022-09-07 17:00:16', NULL, '吴越', 23);
INSERT INTO `user_info` VALUES (5, '2022-09-06 17:09:16', '2022-09-07 17:16:59', NULL, 'hello', 27);
INSERT INTO `user_info` VALUES (6, '2022-09-06 17:12:14', '2022-09-07 17:00:16', NULL, '浪潮', 25);
INSERT INTO `user_info` VALUES (7, '2022-09-07 14:22:51', '2022-09-07 17:00:16', NULL, '晓琴', 16);
INSERT INTO `user_info` VALUES (8, '2022-09-07 14:26:48', '2022-09-07 17:00:16', NULL, 'non_existing1', 15);
INSERT INTO `user_info` VALUES (9, '2022-09-07 14:30:16', '2022-09-07 17:00:16', NULL, 'non_existing2', 15);
INSERT INTO `user_info` VALUES (10, '2022-09-07 14:31:22', '2022-09-07 17:00:16', NULL, 'non_existing3', 15);
INSERT INTO `user_info` VALUES (12, '2022-09-07 16:56:14', '2022-09-07 17:00:16', NULL, '吴九', 26);
INSERT INTO `user_info` VALUES (13, '2022-09-07 16:56:34', '2022-09-07 17:16:59', NULL, 'hello', 27);

SET FOREIGN_KEY_CHECKS = 1;
