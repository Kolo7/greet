CREATE TABLE `mj_historys` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL DEFAULT '0',
  `prompt` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `response` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `hd_image` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `session_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `seed` int NOT NULL DEFAULT '0',
  `width` int NOT NULL DEFAULT '0' COMMENT '图片分辨率宽度',
  `height` int NOT NULL DEFAULT '0' COMMENT '图片分辨率高度',
  `opt_type` tinyint NOT NULL DEFAULT '0' COMMENT '操作类型 1 是generate 2 是u 3 是 v 4 是超分 5 是色彩增强',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status` tinyint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_user_id_created_at` (`user_id`,`created_at`),
  KEY `idx_mjhistory_sessionid` (`session_id`)
) ENGINE=InnoDB AUTO_INCREMENT=14170 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;