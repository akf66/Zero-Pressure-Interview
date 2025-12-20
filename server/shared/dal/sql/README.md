# SQL 文件目录

此目录包含所有数据库表的 SQL 定义文件，用于 GORM Gen 代码生成。

## 文件说明

- `users.sql` - 用户表定义
- `interviews.sql` - 面试表定义
- `questions.sql` - 题目表定义
- `storage.sql` - 存储表定义

## 表设计规范

### 通用字段

所有表都应包含以下通用字段：

```sql
`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
`deleted_at` DATETIME DEFAULT NULL COMMENT '删除时间（软删除）',
```

### 索引规范

1. **主键索引**: 使用 `id` 作为主键，类型为 `BIGINT UNSIGNED AUTO_INCREMENT`
2. **唯一索引**: 使用 `uk_` 前缀，如 `uk_username`
3. **普通索引**: 使用 `idx_` 前缀，如 `idx_user_id`
4. **全文索引**: 使用 `ft_` 前缀，如 `ft_title_content`
5. **软删除索引**: 必须为 `deleted_at` 字段创建索引

### 命名规范

1. **表名**: 使用复数形式，小写，下划线分隔，如 `users`, `interviews`
2. **字段名**: 小写，下划线分隔，如 `user_id`, `created_at`
3. **注释**: 所有表和字段都必须有中文注释

### 字符集

- 使用 `utf8mb4` 字符集
- 使用 `utf8mb4_unicode_ci` 排序规则
- 引擎使用 `InnoDB`

## 添加新表

1. 在此目录创建新的 `.sql` 文件
2. 按照上述规范定义表结构
3. 在 `server/shared/dal/config/gen_config.yaml` 中添加表配置
4. 运行代码生成命令

## 示例

```sql
CREATE TABLE IF NOT EXISTS `example` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `name` VARCHAR(100) NOT NULL COMMENT '名称',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` DATETIME DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_status` (`status`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='示例表';