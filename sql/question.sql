CREATE TABLE `question` (
    `question_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '练习题目id',
    `description` varchar(166) NOT NULL COMMENT '题目描述',
    `type` tinyint NOT NULL COMMENT '唯一题目类型, 和病例有关',
    `options` json NOT NULL COMMENT '所有选项，json格式{"A": "...", "B": "...", "C": "..."}',
    `answer`  char NOT NULL COMMENT '正确答案, 假定所有题目是单选题',
    `duration` int NOT NULL DEFAULT 0 COMMENT '答题时长设定',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间(自动更新)',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间（自动创建）', 
    PRIMARY KEY (`question_id`),
    KEY `idx_type` (`type`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='测试题目表'
;