-- 用户笔记表
CREATE TABLE IF NOT EXISTS `user_notes` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '笔记ID',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    `question_id` BIGINT UNSIGNED NOT NULL COMMENT '题目ID',
    `note` TEXT NOT NULL COMMENT '笔记内容',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_question` (`user_id`, `question_id`),
    KEY `idx_user_id` (`user_id`),
    KEY `idx_question_id` (`question_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户笔记表';
