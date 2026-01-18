-- Active: 1768711665701@@127.0.0.1@3306@test
-- 数据准备
DROP TABLE IF EXISTS account;

CREATE TABLE IF NOT EXISTS account (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
    name VARCHAR(10) COMMENT '姓名',
    money INT COMMENT '余额'
) COMMENT '账号表';

INSERT INTO
    account (id, name, money)
VALUES (NULL, '哈基米', 2000),
    (NULL, '曼波', 2000);

SHOW ENGINES;

UPDATE account SET money = 2000 WHERE name IN ('哈基米', '曼波');

-- 转账操作
-- 查询哈基米账号余额
SELECT * FROM `account` WHERE name = '哈基米';

-- 将哈基米账号余额减 1000
UPDATE account SET money = money - 1000 WHERE name = '哈基米';

-- 将曼波账户余额加 1000
UPDATE account SET money = money + 1000 WHERE name = '曼波';

ROLLBACK;

COMMIT;

START TRANSACTION;

SELECT * FROM account WHERE name = '哈基米';

UPDATE account SET money = money - 1000 WHERE name = '哈基米';

UPDATE account SET money = money + 1000 WHERE name = '曼波';

COMMIT;

ROLLBACK;

-- 查看事务隔离级别
SELECT @@session.transaction_isolation;

SET SESSION TRANSACTION ISOLATION LEVEL READ COMMITTED;

START TRANSACTION;

SELECT * FROM account;

SET SESSION TRANSACTION ISOLATION LEVEL REPEATABLE READ;

START TRANSACTION;

SELECT * FROM account;

SELECT * FROM account;

SET SESSION TRANSACTION ISOLATION LEVEL REPEATABLE READ;

START TRANSACTION;

-- 第一次查询，结果为空
SELECT * FROM account WHERE id = 3;

-- 基于查询结果，尝试插入 id = 3
INSERT INTO account (id, name, money) VALUES (3, '奶龙', 2000);

-- 第二次查询
SELECT * FROM account WHERE id = 3;

SELECT @@session.transaction_isolation;
