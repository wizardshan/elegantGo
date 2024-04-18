CREATE TABLE `articles` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `hash_id` varchar(20) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
    `title` varchar(20) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
    `content` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;


INSERT INTO `articles` (`id`, `hash_id`, `title`, `content`, `create_time`, `update_time`) VALUES
(1, 'iryTrPab4l8', '文章标题', '内容', '2024-04-18 11:03:46', '2024-04-18 11:03:46');

CREATE TABLE `users` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `mobile` varchar(11) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
    `nickname` varchar(20) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
    `bio` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

INSERT INTO `users` (`id`, `mobile`, `nickname`, `bio`, `update_time`, `create_time`) VALUES
(1, '13000000000', '昵称', '个人介绍', '2024-04-11 20:02:32', '2021-12-24 21:27:19');