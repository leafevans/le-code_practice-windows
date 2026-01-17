-- Active: 1760096220136@@127.0.0.1@3306@test_db
DROP TABLE IF EXISTS `account`;

CREATE TABLE IF NOT EXISTS `account` (
    id INT AUTO_INCREMENT PRIMARY KEY comment '主键ID',
    name VARCHAR(10) COMMENT '姓名',
    money INT COMMENT '余额'
) COMMENT '账户表';

INSERT INTO
    `account` (id, name, money)
VALUES (NULL, '哈基米', 2000),
    (NULL, '曼波', 2000);

UPDATE `account` SET money = 2000 WHERE name = '哈基米' OR name = '曼波';

SELECT * FROM `account` WHERE name = '哈基米';

UPDATE `account` SET money = money - 1000 WHERE name = '哈基米';

UPDATE `account` SET money = money + 1000 WHERE name = '曼波';

SELECT * FROM `account`;

SET @@autocommit = 1;

SELECT * FROM `account` WHERE name = '哈基米';

UPDATE `account` SET money = money - 1000 WHERE name = '哈基米';

UPDATE `account` SET money = money + 1000 WHERE name = '曼波';

COMMIT;

ROLLBACK;

--

SELECT @@transaction_isolation;

SELECT * FROM `account`;

SET SESSION TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;

START TRANSACTION;

SELECT * FROM `account`;

COMMIT;

SET SESSION TRANSACTION ISOLATION LEVEL READ COMMITTED;

SET SESSION TRANSACTION ISOLATION LEVEL READ COMMITTED;

UPDATE `account` SET money = 2000 WHERE name = '哈基米' OR name = '曼波';

DELETE FROM `account` WHERE id = 3;

START TRANSACTION;

SELECT * FROM `account`;

SELECT * FROM `account`;

COMMIT;

SET SESSION TRANSACTION ISOLATION LEVEL REPEATABLE READ;

SET SESSION TRANSACTION ISOLATION LEVEL REPEATABLE READ;

SET SESSION TRANSACTION ISOLATION LEVEL SERIALIZABLE;

START TRANSACTION;

SELECT * FROM account WHERE id = 3;

INSERT INTO account (id, name, money) VALUES (3, '奶龙', 2000);

SELECT * FROM account WHERE id = 3;

COMMIT;
