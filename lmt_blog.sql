/*
 Navicat Premium Data Transfer

 Source Server         : 小赖
 Source Server Type    : MySQL
 Source Server Version : 80036
 Source Host           : localhost:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 80036
 File Encoding         : 65001

 Date: 03/01/2025 02:01:04
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for articles
-- ----------------------------
DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_articles_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `fk_articles_user`(`user_id` ASC) USING BTREE,
  CONSTRAINT `fk_articles_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 23 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of articles
-- ----------------------------
INSERT INTO `articles` VALUES (1, '2024-12-23 11:28:33.668', '2024-12-23 11:28:33.668', NULL, '规则三分法', '1、应用：将画面通过两条水平线和两条垂直线分为九等分，重要的视觉元素放在这些线上或交点上，可以增强画面的动态感和吸引力。\r\n2、实践：在拍摄风景、人像时，尝试将地平线放在其中一条水平线上，人物眼睛放在垂直线上的一个交点。', 2);
INSERT INTO `articles` VALUES (2, '2024-12-23 11:40:43.372', '2024-12-23 11:40:43.372', NULL, '光线运用', '1、自然光：黄金时段（日出和日落前后）的光线柔和，适合拍摄人像和风景。\r\n2、人造光：使用闪光灯、反光板等工具来补充或改变光线的方向和强度。', 2);
INSERT INTO `articles` VALUES (3, '2024-12-23 11:42:39.982', '2024-12-23 11:42:39.982', NULL, '拍摄角度', '1、低角度：可以增加主体的威严感，常用于拍摄建筑和人像。\r\n2、高角度：可以展现广阔的视野，适合拍摄风景和城市景观。', 2);
INSERT INTO `articles` VALUES (4, '2024-12-23 11:43:53.588', '2024-12-23 11:43:53.588', NULL, '动态捕捉', '1、快门优先模式：设置合适的快门速度来捕捉动态瞬间。\r\n2、连拍模式：在快速连续拍摄中选择最佳瞬间。', 2);
INSERT INTO `articles` VALUES (5, '2024-12-26 10:01:48.293', '2024-12-30 11:58:34.875', NULL, '《古村柿韵：秋光里的丰收诗画》', '在这幅如诗如画的景象中，一棵柿子树宛如大自然的艺术杰作，骄傲地展示着它丰硕的果实。橙红色的柿子像一盏盏小灯笼，挂满了枝头，在阳光的轻抚下，闪烁着诱人的光泽，仿佛在诉说着秋日的喜悦。柿子树的叶子在果实间若隐若现，那一抹抹绿色为这橙红的世界增添了生机与活力。柿子树背后，是几座古朴的中式建筑，白墙黑瓦，飞檐翘角，散发着岁月沉淀的韵味。建筑的门窗和屋顶在柿子树的映衬下，显得更加典雅。远处的山峦在蓝天的怀抱中绵延起伏，为这幅画面增添了一抹宁静的背景。整个场景沐浴在柔和的光线中，光影交错，营造出一种温暖而祥和的氛围。这不仅仅是一幅描绘柿子树的图画，更是一曲对丰收的赞歌，是自然与人文和谐交融的美丽诗篇。\r\n', 3);
INSERT INTO `articles` VALUES (6, '2024-12-26 10:02:18.850', '2024-12-26 10:02:18.850', NULL, '《柴犬相伴的温馨时光》', '在这张充满温情的画面中，两只柴犬相依相偎，仿佛时间都在这一刻静止。左边的柴犬，吐着粉色的舌头，眼睛明亮有神，似乎在享受着这片刻的宁静与快乐。它的表情充满了纯真与喜悦，就像一个无忧无虑的孩子。右边的柴犬，半眯着眼睛，将头轻轻地靠在同伴的身上，显得十分惬意和放松。它的姿态透露出一种对同伴的信任和依赖，如同找到了最安心的港湾。', 3);
INSERT INTO `articles` VALUES (7, '2024-12-26 10:02:52.073', '2025-01-03 01:08:10.253', NULL, '《月季余晖：黄昏下的浪漫花海》', '在这张如梦如幻的图片中，一片盛开的月季花海在黄昏的余晖中展现出它们最迷人的姿态。夕阳的金辉如轻纱般洒落在每一朵月季花上，给花瓣镀上了一层柔和的光晕，仿佛为它们披上了金色的舞裙。粉色与红色的月季花相互交织，如同恋人在黄昏下的呢喃细语，诉说着甜蜜与浪漫。花海的背景是一片渐变色的天空，从橙黄到淡紫，如同画家在画布上随意挥洒的色彩，自然而又和谐。远处的树木在微光中若隐若现，像是守护这片花海的忠诚卫士。每一朵月季花都在微风中轻轻摇曳，似乎在向世界展示着它们的美丽与芬芳，整个画面充满了诗意与宁静，宛如一幅绝美的自然画卷。\r\n', 3);
INSERT INTO `articles` VALUES (9, '2024-12-26 10:04:41.638', '2024-12-26 10:04:41.638', '2025-01-02 23:47:02.098', '微单相机', '微单相机是一种便于携带的相机，它采用了无反光镜设计，以及紧凑的机身，为摄影师提供方便携带的拍摄体验。\r\n微单相机成像原理为：采用无反光镜取景，所以将光线引导至电子取景器，通过传感器捕获并处理图像，然后再通过液晶屏幕或取景器显示。', 3);
INSERT INTO `articles` VALUES (10, '2024-12-26 10:05:15.278', '2024-12-26 10:05:15.278', '2025-01-03 01:08:53.597', '卡片相机', '卡片相机是一种小巧轻便的相机，其机身和镜头融为一体，适合旅游或日常摄影拍摄。它的成像原理是将光线通过镜头直接传入传感器，然后通过液晶屏幕实时显示画面。卡片相机体积轻巧，方便携带，能够随身携带，方便拍摄。卡片相机的操作简单，不需要特别的技术知识，即可轻松上手。', 3);

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `user_id` bigint UNSIGNED NOT NULL,
  `article_id` bigint UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_comments_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comments
-- ----------------------------
INSERT INTO `comments` VALUES (1, '2024-12-30 16:56:42.760', '2024-12-30 16:56:42.760', NULL, '很可爱的小狗', 3, 4);
INSERT INTO `comments` VALUES (2, '2024-12-30 22:32:16.015', '2024-12-30 22:32:16.015', NULL, '好好看', 3, 7);
INSERT INTO `comments` VALUES (3, '2024-12-30 22:35:41.464', '2024-12-30 22:35:41.464', NULL, '哇哇哇', 3, 3);
INSERT INTO `comments` VALUES (4, '2024-12-30 22:51:27.982', '2024-12-30 22:51:27.982', NULL, '好piu酿', 3, 8);
INSERT INTO `comments` VALUES (5, '2024-12-30 22:53:41.926', '2024-12-30 22:53:41.926', NULL, '月季月季月季', 3, 7);
INSERT INTO `comments` VALUES (6, '2024-12-30 22:57:05.137', '2024-12-30 22:57:05.137', NULL, '哇！', 3, 3);
INSERT INTO `comments` VALUES (8, '2025-01-03 01:18:00.336', '2025-01-03 01:18:00.336', NULL, '好好看的月季', 3, 7);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `email` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `password` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_users_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, '2024-12-10 20:09:24.262', '2024-12-10 20:09:24.262', NULL, 'Random name', 'random@email.com', '$2a$12$fETY9RDWruYyQdZXpwv2aeT3.rWe.48R/YSDRq9IMFtaFRKguy6Lm');
INSERT INTO `users` VALUES (2, '2024-12-22 14:43:46.456', '2024-12-22 14:43:46.456', NULL, 'Ayman', 'info@ayman-elmalah.com', '$2a$12$Ieh2aWepjJRRXNJ6Eh3Naepedfe8HvRsldD4tV7QDP.2ZpT1SBwTm');
INSERT INTO `users` VALUES (3, '2024-12-23 11:55:44.564', '2025-01-02 16:12:22.653', NULL, 'LM', 'info@lmt-elmalah.com', '$2a$12$Tcc22rBN0w/NPGUIMhl5kuhu.U0Jt3pizOh45PIYyyyQB2WtcbH7a');
INSERT INTO `users` VALUES (4, '2025-01-02 21:41:31.888', '2025-01-02 21:41:31.888', NULL, 'LMT', 'info@LLL-elmalah.com', '$2a$12$3o08pEfQOYd91SpjRFs3X.dwaZY0Vus2dNieZKgg0GX7CvqNtLNEm');
INSERT INTO `users` VALUES (5, '2025-01-02 21:48:42.623', '2025-01-02 21:48:42.623', NULL, 'L', 'info@lmm-elmalah.com', '$2a$12$G8KKU3yccGpkW0JfHimZx.1gblr66iWrEP1Br3ikG1QBjaoDk3Mu.');
INSERT INTO `users` VALUES (6, '2025-01-02 21:54:01.705', '2025-01-02 21:54:01.705', NULL, 'LL', 'info@mmm-elmalah.com', '$2a$12$0xwoZe9MPRNxyPXk9u623uxti0r91BIYX6zLAcny8yyqhh71JQsHm');
INSERT INTO `users` VALUES (9, '2025-01-03 00:44:21.300', '2025-01-03 00:54:37.514', NULL, 'LM', 'info@ttl-elmalah.com', '$2a$12$vMWkoX832eAFRUL6CzaTaewG4QwVA987wMcLsOIM7k405tmx85ODm');

SET FOREIGN_KEY_CHECKS = 1;
