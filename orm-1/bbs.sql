-- phpMyAdmin SQL Dump
-- version 4.8.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: 2024-11-14 14:40:55
-- 服务器版本： 5.7.19
-- PHP Version: 7.4.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `bbs`
--

-- --------------------------------------------------------

--
-- 表的结构 `comments`
--

CREATE TABLE `comments` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL DEFAULT '0',
  `post_id` int(11) NOT NULL DEFAULT '0',
  `content` varchar(255) NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- 转存表中的数据 `comments`
--

INSERT INTO `comments` (`id`, `user_id`, `post_id`, `content`, `create_time`, `update_time`) VALUES
(1, 1, 1, '评论1', '2024-08-01 00:00:00', '2024-08-02 00:00:00'),
(2, 2, 2, '评论2', '2024-08-01 00:00:00', '2024-08-02 00:00:00');

-- --------------------------------------------------------

--
-- 表的结构 `posts`
--

CREATE TABLE `posts` (
  `id` int(11) NOT NULL,
  `hash_id` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `user_id` int(11) DEFAULT '0',
  `title` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `content` text COLLATE utf8mb4_bin,
  `view` int(11) NOT NULL DEFAULT '0',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- 转存表中的数据 `posts`
--

INSERT INTO `posts` (`id`, `hash_id`, `user_id`, `title`, `content`, `view`, `create_time`, `update_time`) VALUES
(1, 'oKqk6tMl7z', 1, '标题1', '内容1', 100, '2024-08-01 00:00:00', '2024-08-02 00:00:00'),
(2, '02qN7SQyOb', 2, '标题2', '内容2', 200, '2024-08-01 00:00:00', '2024-08-02 00:00:00');

-- --------------------------------------------------------

--
-- 表的结构 `users`
--

CREATE TABLE `users` (
  `id` int(10) UNSIGNED NOT NULL,
  `hash_id` varchar(20) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `mobile` varchar(11) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `password` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `age` int(11) NOT NULL DEFAULT '0',
  `level` int(11) NOT NULL DEFAULT '0',
  `nickname` varchar(20) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `avatar` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `bio` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `amount` int(11) NOT NULL DEFAULT '0',
  `status` varchar(10) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `active` tinyint(1) NOT NULL DEFAULT '1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- 转存表中的数据 `users`
--

INSERT INTO `users` (`id`, `hash_id`, `mobile`, `password`, `age`, `level`, `nickname`, `avatar`, `bio`, `amount`, `status`, `create_time`, `update_time`, `active`) VALUES
(1, 'oKqk6tMl7z', '13000000001', 'a906449d5769fa7361d7ecc6aa3f6d28', 28, 10, '昵称1', '头像1.png', '个人介绍1', 200, 'normal', '2024-08-01 00:00:00', '2024-08-02 00:00:00', 1),
(2, '02qN7SQyOb', '13000000002', 'a906449d5769fa7361d7ecc6aa3f6d28', 29, 20, '昵称2', '头像2.png', '个人介绍2', 100, 'cancel', '2024-08-01 00:00:00', '2024-08-02 00:00:00', 1);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `comments`
--
ALTER TABLE `comments`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `posts`
--
ALTER TABLE `posts`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `mobile` (`mobile`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `comments`
--
ALTER TABLE `comments`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `posts`
--
ALTER TABLE `posts`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `users`
--
ALTER TABLE `users`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
