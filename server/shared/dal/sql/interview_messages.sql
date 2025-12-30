-- 面试对话记录表
CREATE TABLE IF NOT EXISTS `interview_messages` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '消息ID',
    `interview_id` BIGINT UNSIGNED NOT NULL COMMENT '面试ID',
    `role` VARCHAR(20) NOT NULL COMMENT '角色：interviewer-面试官，candidate-候选人',
    `content` TEXT NOT NULL COMMENT '消息内容',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY `idx_interview_id` (`interview_id`),
    KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='面试对话记录表';
