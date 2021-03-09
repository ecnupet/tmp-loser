CREATE TABLE `history` (
    `history_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '回答历史记录id',
    `user_id` bigint(20) unsigned NOT NULL COMMENT 'user id',
    `quiz_id` bigint(20) unsigned NOT NULL COMMENT '测试题目id',
    `choose` char NOT NULL COMMENT '用户的选择', 
    `test_id` bigint(10) unsigned NOT NULL COMMENT `所属的测验id`,
    `correct` tinyint(1) NOT NULL NOt NULL DEFAULT '0' COMMENT '知识点正误, 0: 未设置, 1: 正确, 2: 错误',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间（自动创建）',
    PRIMARY KEY (`history_id`),
    KEY `idx_user_id` (`user_id`),
    KEY `idx_quiz_id` (`quiz_id`),
    KEY `idx_test_id` (`test_id`)，
    KEY `idx_create_at` (`created_at`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='历史答题记录表'
;