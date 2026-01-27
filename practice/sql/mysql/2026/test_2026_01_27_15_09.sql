-- 1. 查看当前行数（确认）
SELECT COUNT(*) AS cur_rows FROM user;
-- 1a. 检查表结构并确保 `password` 列足够长（避免 "Data too long" 错误）
SHOW CREATE TABLE `user`;
-- 下面直接扩展 `password` 列到 VARCHAR(255)，保证 SHA2(...,256) 的 64 字符十六进制串可以保存。
ALTER TABLE `user` MODIFY COLUMN `password` VARCHAR(255) NOT NULL;

-- 2. 创建临时源表（只包含需要插入的字段，避免 id 冲突）
DROP TABLE IF EXISTS tmp_user;

CREATE TABLE tmp_user AS
SELECT
    username,
    password,
    name,
    birthday,
    sex
FROM user;

-- 3. 存储过程：按 batch_size 批量插入，直到达到 factor 倍
DELIMITER /
/

DROP PROCEDURE IF EXISTS grow_to_factor /
/

CREATE PROCEDURE grow_to_factor(IN factor INT, IN batch_size INT)
BEGIN
  DECLARE target BIGINT;
  SELECT COUNT(*) INTO target FROM user;
  SET target = target * factor;
  WHILE (SELECT COUNT(*) FROM user) < target DO
    INSERT INTO user (username,password,name,birthday,sex)
    SELECT
      CONCAT(username, '_', FLOOR(RAND()*1000000000)),
      SHA2(CONCAT(password, RAND(), UUID()),256),
      CONCAT(name, '_', FLOOR(RAND()*1000000000)),
      DATE_SUB(birthday, INTERVAL FLOOR(RAND()*36500) DAY),
      FLOOR(RAND()*3)
    FROM tmp_user
    LIMIT batch_size;
  END WHILE;
END
/
/

DELIMITER;

-- 4. 调用（示例：batch_size 可根据内存/事务能力调整，比如 50000）
-- 注意：对 1,000,000 行扩大 50 倍会生成 ~50,000,000 行，执行前请确保磁盘/内存/日志足够。
CALL grow_to_factor (50, 50000);

-- 5. 清理
DROP PROCEDURE IF EXISTS grow_to_factor;

DROP TABLE IF EXISTS tmp_user;

-- 6. 最后确认
SELECT COUNT(*) AS final_rows FROM user;
-- 2. 创建临时源表（只包含需要插入的字段，避免 id 冲突）
DROP TABLE IF EXISTS tmp_user;

CREATE TABLE tmp_user AS
SELECT
    username,
    password,
    name,
    birthday,
    sex
FROM user;

-- 3. 存储过程：按 batch_size 批量插入，直到达到 factor 倍
DELIMITER /
/

DROP PROCEDURE IF EXISTS grow_to_factor /
/

CREATE PROCEDURE grow_to_factor(IN factor INT, IN batch_size INT)
BEGIN
  DECLARE target BIGINT;
  SET target = (SELECT COUNT(*) FROM user) * factor;
  WHILE (SELECT COUNT(*) FROM user) < target DO
    INSERT INTO user (username,password,name,birthday,sex)
    SELECT
      CONCAT(username, '_', FLOOR(RAND()*1000000000)),
      SHA2(CONCAT(password, RAND(), UUID()),256),
      CONCAT(name, '_', FLOOR(RAND()*1000000000)),
      DATE_SUB(birthday, INTERVAL FLOOR(RAND()*36500) DAY),
      FLOOR(RAND()*3)
    FROM tmp_user
    LIMIT batch_size;
  END WHILE;
END
/
/

DELIMITER;

-- 4. 调用（示例：batch_size 可根据内存/事务能力调整，比如 50000）
CALL grow_to_factor (50, 50000);

-- 5. 清理
DROP PROCEDURE IF EXISTS grow_to_factor;

DROP TABLE IF EXISTS tmp_user;

-- 6. 最后确认
SELECT COUNT(*) AS final_rows FROM user;
