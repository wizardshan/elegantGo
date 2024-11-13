-- phpMyAdmin SQL Dump
-- version 4.8.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: 2024-11-13 15:32:50
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
-- Database: `test`
--

DELIMITER $$
--
-- 存储过程
--
CREATE DEFINER=`root`@`localhost` PROCEDURE `proc_batch_insert123` ()  BEGIN
    #定义两个变量
    DECLARE order_no INTEGER(11);
    DECLARE id1 INTEGER(11);
    DECLARE i INTEGER(11);
    SET id1=470;
    SET order_no=127724235;
    SET i=1;
    WHILE i <= 10 DO
        INSERT INTO `t_trade_bill`(`id`, `user_id`, `order_no`, `bill_title`, `bill_type`, `flow_type`, `amount`, `charge`, `trade_time`, `complete_time`, `create_time`, `update_time`, `is_deleted`) VALUES (id1, '200acf0335534e0bb4571eaa9dffe27d', order_no, '余额提现', 1, 2, 10000, 0, NOW(), NOW(), NOW(), NOW(), 0);
        SET id1=id1+1;
        SET order_no=order_no+1;
        SET i=i+1;
    END WHILE;
END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `proc_batch_insert2225` ()  BEGIN
    #定义两个变量
    DECLARE ageVal INTEGER(11);
    DECLARE id1 INTEGER(11);
    DECLARE i INTEGER(11);
    SET id1=470;
    SET ageVal=127724235;
    SET i=1;
    WHILE i <= 10 DO
        INSERT INTO `t_trade_bill`(`id`, `user_id`, `order_no`, `bill_title`, `bill_type`, `flow_type`, `amount`, `charge`, `trade_time`, `complete_time`, `create_time`, `update_time`, `is_deleted`) VALUES (id1, '200acf0335534e0bb4571eaa9dffe27d', aageVal, '余额提现', 1, 2, 10000, 0, NOW(), NOW(), NOW(), NOW(), 0a);
        SET id1=id1+1;
        SET ageVal=ageVal+1;
        SET i=i+1;
    END WHILE;
END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `proc_batch_insert2229` ()  BEGIN
    #定义两个变量
    DECLARE ageVal INTEGER(11);
    DECLARE id1 INTEGER(11);
    DECLARE i INTEGER(11);
    SET id1=470;
    SET ageVal=127724235;
    SET i=1;
    WHILE i <= 10 DO
        INSERT INTO `t_trade_bill`(`id`, `user_id`, `order_no`, `bill_title`, `bill_type`, `flow_type`, `amount`, `charge`, `trade_time`, `complete_time`, `create_time`, `update_time`, `is_deleted`) VALUES (id1, '200acf0335534e0bb4571eaa9dffe27d', aageVal, '余额提现', 1, 2, 10000, 0, NOW(), NOW(), NOW(), NOW(), 0);
        SET id1=id1+1;
        SET ageVal=ageVal+1;
        SET i=i+1;
    END WHILE;
END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `proc_batch_insert235` ()  BEGIN
    #定义两个变量
    DECLARE ageVal INTEGER;
    DECLARE id1 INTEGER;
    DECLARE i INTEGER;
    SET id1=470;
    SET ageVal=127724235;
    SET i=1;
    WHILE i <= 10 DO
        INSERT INTO `t_trade_bill`(`id`, `user_id`, `order_no`, `bill_title`, `bill_type`, `flow_type`, `amount`, `charge`, `trade_time`, `complete_time`, `create_time`, `update_time`, `is_deleted`) VALUES (id1, '200acf0335534e0bb4571eaa9dffe27d', aageVal, '余额提现', 1, 2, 10000, 0, NOW(), NOW(), NOW(), NOW(), 0);
        SET id1=id1+1;
        SET ageVal=ageVal+1;
        SET i=i+1;
    END WHILE;
END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `user_main_pro3` (IN `v_id` INTEGER)  BEGIN
  #定义两个变量
  DECLARE v_userName VARCHAR(50);
  DECLARE v_userName2 VARCHAR(50);
  #set赋值
  SET v_userName = 's';
  #SELECT ... INTO 赋值
  SELECT f_userName INTO v_userName2 FROM t_user_main 
  WHERE f_userId = v_id;
  #DDL语句
  INSERT INTO t_user_main (f_userName) VALUES(CONCAT(v_userName,'*',v_userName2));
END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `user_main_pro31` ()  BEGIN
  #定义两个变量
  DECLARE v_userName INTEGER;
  DECLARE v_userName2 INTEGER;


END$$

DELIMITER ;

-- --------------------------------------------------------

--
-- 表的结构 `articles`
--

CREATE TABLE `articles` (
  `id` int(10) UNSIGNED NOT NULL,
  `hash_id` varchar(20) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `title` varchar(20) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `content` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- 转存表中的数据 `articles`
--

INSERT INTO `articles` (`id`, `hash_id`, `title`, `content`, `create_time`, `update_time`) VALUES
(1, 'iryTrPab4l8', '文章标题', '内容', '2024-04-18 11:03:46', '2024-04-18 11:03:46');

-- --------------------------------------------------------

--
-- 表的结构 `users`
--

CREATE TABLE `users` (
  `id` int(10) UNSIGNED NOT NULL,
  `mobile` varchar(11) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `nickname` varchar(20) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `bio` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- 转存表中的数据 `users`
--

INSERT INTO `users` (`id`, `mobile`, `nickname`, `bio`, `update_time`, `create_time`) VALUES
(1, '13000000000', '昵称', '个人介绍', '2024-04-11 20:02:32', '2021-12-24 21:27:19');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `articles`
--
ALTER TABLE `articles`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `articles`
--
ALTER TABLE `articles`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `users`
--
ALTER TABLE `users`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
