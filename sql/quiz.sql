CREATE TABLE `quiz` (
    `quiz_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '练习题目id',
    `description` varchar(166) NOT NULL COMMENT '题目描述',
    `type` tinyint NOT NULL COMMENT '唯一题目类型：1 2 待定'
    `options` json NOT NULL COMMENT '所有选项，json格式{"A": "...", "B": "...", "C": "..."}',
    `answer`  char NOT NULL COMMENT '正确答案, 假定所有题目是单选题',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间(自动更新)',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间（自动创建）', 
    PRIMARY KEY (`quiz_id`),
    KEY `idx_tag` (`type`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='测试题目表'
;