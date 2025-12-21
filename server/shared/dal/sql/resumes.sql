-- 简历表
CREATE TABLE IF NOT EXISTS `resumes` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '简历ID',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    `file_url` VARCHAR(500) NOT NULL COMMENT '简历文件URL',
    `parsed_content` TEXT DEFAULT NULL COMMENT '解析后的简历内容（JSON格式）',
    `file_name` VARCHAR(255) DEFAULT NULL COMMENT '原始文件名',
    `file_size` BIGINT DEFAULT NULL COMMENT '文件大小（字节）',
    `file_type` VARCHAR(50) DEFAULT NULL COMMENT '文件类型（pdf/doc/docx等）',
    `parse_status` TINYINT NOT NULL DEFAULT 0 COMMENT '解析状态：0-未解析，1-解析中，2-解析成功，3-解析失败',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` DATETIME DEFAULT NULL COMMENT '删除时间（软删除）',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`),
    KEY `idx_parse_status` (`parse_status`),
    KEY `idx_created_at` (`created_at`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='简历表';