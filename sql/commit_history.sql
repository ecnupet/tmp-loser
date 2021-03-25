CREATE TABLE `commit_history` (
    `history_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '回答历史记录id',
    `user_id` bigint(20) unsigned NOT NULL COMMENT 'user id',
    `question_id` bigint(20) unsigned NOT NULL COMMENT '测试题目id',
    `choose` char NOT NULL COMMENT '用户的选择', 
    `quiz_id` bigint(10) unsigned NOT NULL COMMENT `所属的测验id`,
    `correct` tinyint(1) NOT NULL NOt NULL DEFAULT '0' COMMENT '知识点正误, 0: 未设置, 1: 正确, 2: 错误',
    `start_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '开始时间，为创建时间',
    `finish_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '完成时间，为最后一次更新时间',
    PRIMARY KEY (`history_id`),
    KEY `idx_user_id` (`user_id`),
    KEY `idx_quiz_id` (`quiz_id`),
    KEY `idx_question_id` (`question_id`)，
    KEY `idx_create_at` (`created_at`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='历史答题记录表'
;