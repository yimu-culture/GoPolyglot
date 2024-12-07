CREATE TABLE `users`
(
    `id`         INT AUTO_INCREMENT PRIMARY KEY,                                 -- 自增的主键
    `username`   VARCHAR(255) NOT NULL UNIQUE,                                   -- 用户名，唯一
    `password`   VARCHAR(255) NOT NULL,                                          -- 密码，存储哈希后的密码
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,                            -- 记录创建时间
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- 记录更新时间
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
