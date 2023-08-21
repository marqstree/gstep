/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80012
 Source Host           : localhost:3306
 Source Schema         : gstep

 Target Server Type    : MySQL
 Target Server Version : 80012
 File Encoding         : 65001

 Date: 20/08/2023 09:47:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for hr_department
-- ----------------------------
DROP TABLE IF EXISTS `hr_department`;
CREATE TABLE `hr_department`  (
  `id` int(11) NOT NULL,
  `name` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `parent_id` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hr_department
-- ----------------------------
INSERT INTO `hr_department` VALUES (1, 'gstep公司', 0);
INSERT INTO `hr_department` VALUES (2, '信息技术部', 1);
INSERT INTO `hr_department` VALUES (3, '网络室', 2);
INSERT INTO `hr_department` VALUES (4, '开发室', 2);
INSERT INTO `hr_department` VALUES (5, '人力资源部', 1);

-- ----------------------------
-- Table structure for hr_user
-- ----------------------------
DROP TABLE IF EXISTS `hr_user`;
CREATE TABLE `hr_user`  (
  `id` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户id',
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '姓名',
  `position` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '职位',
  `position_code` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '职位编码',
  `is_leader` tinyint(4) NOT NULL COMMENT '是否是部门负责人',
  `department_id` int(11) NOT NULL DEFAULT 0 COMMENT '部门id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hr_user
-- ----------------------------
INSERT INTO `hr_user` VALUES ('001', '刘协', '董事长', '01', 1, 1);
INSERT INTO `hr_user` VALUES ('101', '刘备', '主任', '02', 1, 3);
INSERT INTO `hr_user` VALUES ('102', '关羽', '网络员', '03', 0, 3);
INSERT INTO `hr_user` VALUES ('103', '张飞', '网络员', '03', 0, 3);
INSERT INTO `hr_user` VALUES ('201', '孙权', '主任', '02', 1, 4);
INSERT INTO `hr_user` VALUES ('202', '周瑜', '程序员', '04', 0, 4);
INSERT INTO `hr_user` VALUES ('301', '曹操', '主任', '02', 1, 5);
INSERT INTO `hr_user` VALUES ('302', '郭嘉', '人事员', '05', 0, 5);

-- ----------------------------
-- Table structure for process
-- ----------------------------
DROP TABLE IF EXISTS `process`;
CREATE TABLE `process`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `template_id` int(11) NULL DEFAULT NULL,
  `start_user_id` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `state` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '状态:started,finished',
  `finished_at` datetime NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_id`(`id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 88 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of process
-- ----------------------------

-- ----------------------------
-- Table structure for task
-- ----------------------------
DROP TABLE IF EXISTS `task`;
CREATE TABLE `task`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `process_id` int(11) NULL DEFAULT NULL,
  `step_id` int(11) NULL DEFAULT NULL,
  `title` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `category` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `form` varchar(10000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `audit_method` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `state` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '状态:started,pass,refuse,withdraw',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 120 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of task
-- ----------------------------

-- ----------------------------
-- Table structure for task_assignee
-- ----------------------------
DROP TABLE IF EXISTS `task_assignee`;
CREATE TABLE `task_assignee`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `task_id` int(11) NULL DEFAULT NULL,
  `user_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `state` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '状态:started,pass,refuse',
  `submit_index` int(11) NULL DEFAULT NULL,
  `form` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 204 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of task_assignee
-- ----------------------------

-- ----------------------------
-- Table structure for template
-- ----------------------------
DROP TABLE IF EXISTS `template`;
CREATE TABLE `template`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `template_id` int(11) NULL DEFAULT NULL,
  `title` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `version` int(5) NULL DEFAULT NULL,
  `root_step` varchar(10000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_id`(`id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 63 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of template
-- ----------------------------

-- ----------------------------
-- View structure for department
-- ----------------------------
DROP VIEW IF EXISTS `department`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `department` AS with recursive `ct` as (select `a`.`id` AS `id`,`a`.`name` AS `name`,`a`.`parent_id` AS `parent_id` from `hr_department` `a` where (`a`.`id` = 1) union all select `b`.`id` AS `id`,`b`.`name` AS `name`,`b`.`parent_id` AS `parent_id` from (`hr_department` `b` join `ct` on((`ct`.`id` = `b`.`parent_id`)))) select concat(`ct`.`id`,'') AS `id`,`ct`.`name` AS `name`,concat(`ct`.`parent_id`,'') AS `parent_id` from `ct`;

-- ----------------------------
-- View structure for position
-- ----------------------------
DROP VIEW IF EXISTS `position`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `position` AS select distinct `hr_user`.`position` AS `title`,`hr_user`.`position_code` AS `code` from `hr_user` where ((length(ifnull(`hr_user`.`position_code`,'')) > 0) and (length(ifnull(`hr_user`.`position`,'')) > 0));

-- ----------------------------
-- View structure for user
-- ----------------------------
DROP VIEW IF EXISTS `user`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `user` AS select `hr_user`.`id` AS `id`,`hr_user`.`name` AS `name`,`hr_user`.`position` AS `position_title`,`hr_user`.`position_code` AS `position_code`,`hr_user`.`is_leader` AS `is_leader`,`hr_user`.`department_id` AS `department_id` from `hr_user`;

SET FOREIGN_KEY_CHECKS = 1;
