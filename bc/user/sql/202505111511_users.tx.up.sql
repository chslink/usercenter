-- ----------------------------
-- 用户核心表
-- ----------------------------
CREATE TABLE `users` (
                         `uid` VARCHAR(36) NOT NULL COMMENT '用户UUID',
                         `username` VARCHAR(64) NOT NULL COMMENT '用户名',
                         `phone` VARCHAR(20) NOT NULL COMMENT 'E.164格式手机号',
                         `email` VARCHAR(128) NULL COMMENT '邮箱地址',
                         `password_hash` VARCHAR(128) NULL COMMENT 'bcrypt加密后的密码',
                         `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '账户状态 1-正常 2-冻结',
                         `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
                         `updated_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
                         PRIMARY KEY (`uid`),
                         UNIQUE INDEX `udx_phone` (`phone`),
                         UNIQUE INDEX `udx_email` (`email`),
                         INDEX `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

# -- ----------------------------
# -- 第三方认证表
# -- ----------------------------
# CREATE TABLE `third_party_auth` (
#                                     `id` VARCHAR(36) NOT NULL ,
#                                     `user_id` VARCHAR(36) NOT NULL,
#                                     `provider` VARCHAR(20) NOT NULL COMMENT '第三方提供商',
#                                     `open_id` VARCHAR(128) NOT NULL COMMENT '第三方唯一ID',
#                                     `union_id` VARCHAR(128) NULL COMMENT '微信等需要unionID的平台',
#                                     `metadata` JSON NULL COMMENT '扩展信息',
#                                     `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
#                                     PRIMARY KEY (`id`),
#                                     UNIQUE INDEX `udx_provider_openid` (`provider`, `open_id`),
#                                     INDEX `idx_user` (`user_id`),
#                                     CONSTRAINT `fk_user_third_auth`
#                                         FOREIGN KEY (`user_id`) REFERENCES `users` (`uid`) ON DELETE CASCADE
# ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
#
# -- ----------------------------
# -- 安全日志表（审计用）
# -- ----------------------------
# CREATE TABLE `security_logs` (
#                                  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
#                                  `user_id` VARCHAR(36) NOT NULL,
#                                  `event_type` VARCHAR(32) NOT NULL COMMENT '事件类型',
#                                  `ip_address` VARCHAR(45) NOT NULL,
#                                  `user_agent` VARCHAR(512) NULL,
#                                  `metadata` JSON NULL,
#                                  `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
#                                  PRIMARY KEY (`id`),
#                                  INDEX `idx_user_event` (`user_id`, `event_type`)
# ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;