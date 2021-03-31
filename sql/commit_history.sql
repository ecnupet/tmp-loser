CREATE TABLE `commit_history` (
    `history_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '回答历史记录id',
    `user_name` varchar(100) NOT NULL COMMENT '用户名',
    `question_id` int unsigned NOT NULL COMMENT '测试题目id',
    `choose` char NOT NULL COMMENT '用户的选择', 
    `quiz_id` int unsigned NOT NULL COMMENT '所属的测验id',
    `order` int unsigned NOT NULL COMMENT '在测验中的顺序',
    `correct` tinyint(1) NOT NULL NOt NULL DEFAULT '0' COMMENT '知识点正误, 0: 未设置, 1: 正确, 2: 错误',
    `spend` int NOT NULL DEFAULT 0 COMMENT '默认0',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '完成时间，为最后一次更新时间',
    PRIMARY KEY (`history_id`),
    UNIQUE KEY `unique_quiz_order` (`quiz_id`, `order`),
    KEY `idx_user_name` (`user_name`),
    KEY `idx_quiz_id` (`quiz_id`),
    KEY `idx_question_id` (`question_id`),
    KEY `idx_created_at` (`created_at`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='历史答题记录表'
;