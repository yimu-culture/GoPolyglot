CREATE TABLE `translation_tasks` (
         `id` int(11) NOT NULL AUTO_INCREMENT,
         `user_id` int(11) NOT NULL,
         `source_lang` varchar(10) NOT NULL,
         `target_lang` varchar(10) NOT NULL,
         `status` enum('pending','in_progress','fail','completed') NOT NULL DEFAULT 'pending',
         `source_doc` text NOT NULL,
         `translated_doc` text,
         `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
         `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
         PRIMARY KEY (`id`),
         KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;