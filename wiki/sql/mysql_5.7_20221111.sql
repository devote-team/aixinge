
SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for axg_menu
-- ----------------------------
DROP TABLE IF EXISTS `axg_menu`;
CREATE TABLE `axg_menu`  (
  `id` bigint UNSIGNED NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `parent_id` bigint NOT NULL DEFAULT 1,
  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `redirect` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `hidden` tinyint(1) NOT NULL DEFAULT 2,
  `component` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sort` bigint NULL DEFAULT NULL,
  `is_frame` tinyint(1) NOT NULL DEFAULT 2,
  `status` bigint NOT NULL DEFAULT 1,
  `no_cache` tinyint(1) NOT NULL DEFAULT 1,
  `title` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `icon` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;


-- ----------------------------
-- Table structure for axg_role
-- ----------------------------
DROP TABLE IF EXISTS `axg_role`;
CREATE TABLE `axg_role`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `alias` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `status` bigint NOT NULL DEFAULT 1,
  `sort` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of axg_role
-- ----------------------------
INSERT INTO `axg_role` VALUES (1, '2022-05-12 11:27:36', '2022-05-12 11:27:36', NULL, '产品经理', 'ttt', 'string', 1, 0);
INSERT INTO `axg_role` VALUES (2, '2022-05-13 08:27:16', '2022-10-30 18:28:33', NULL, '管理员', 'admin', '系统管理员身份', 1, 0);

-- ----------------------------
-- Table structure for axg_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `axg_role_menu`;
CREATE TABLE `axg_role_menu`  (
  `role_id` bigint NOT NULL,
  `menu_id` bigint NOT NULL
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of axg_role_menu
-- ----------------------------

-- ----------------------------
-- Table structure for axg_user
-- ----------------------------
DROP TABLE IF EXISTS `axg_user`;
CREATE TABLE `axg_user`  (
  `id` bigint UNSIGNED NOT NULL,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `uuid` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `nick_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '系统用户',
  `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `status` bigint NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_users_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of axg_user
-- ----------------------------
INSERT INTO `axg_user` VALUES (1, '2022-02-15 23:25:20', '2022-11-06 20:26:03', NULL, 'f5f10f02-d4fe-472a-b6ed-c84698a3c007', 'admin', '1751ad04bed81acd2d1bc5720b15f0f5', '超级管理员', '1589232642040004608', 1);

-- ----------------------------
-- Table structure for axg_user_role
-- ----------------------------
DROP TABLE IF EXISTS `axg_user_role`;
CREATE TABLE `axg_user_role`  (
  `user_id` bigint UNSIGNED NOT NULL,
  `role_id` bigint UNSIGNED NOT NULL
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of axg_user_role
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
