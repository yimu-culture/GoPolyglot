CREATE TABLE `translation_tasks` (
     `task_id`        INT(11) NOT NULL AUTO_INCREMENT,                                            -- 任务的唯一标识符
     `user_id`        INT(11) NOT NULL,                                                           -- 发起翻译任务的用户ID
     `source_lang`    VARCHAR(10) NOT NULL,                                                       -- 源语言标识符，例如 "en" 或 "zh"
     `target_lang`    VARCHAR(10) NOT NULL,                                                       -- 目标语言标识符，例如 "es" 或 "fr"
     `status`         ENUM('pending', 'in_progress', 'completed') NOT NULL DEFAULT 'pending',     -- 任务状态：待翻译、翻译中、已完成
     `source_doc`     TEXT        NOT NULL,                                                       -- 源文档路径或内容
     `translated_doc` TEXT,                                                                       -- 翻译后的文档路径或内容
     `created_at`     TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,                             -- 任务创建时间
     `updated_at`     TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- 任务最后更新时间
     PRIMARY KEY (`task_id`),                                                                     -- 主键
     INDEX            `idx_user_id` (`user_id`)                                                   -- 为用户ID创建索引，优化查询速度
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
