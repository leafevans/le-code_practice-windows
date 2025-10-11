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

SELECT @@autocommit;

SET @@autocommit = 1;

SELECT * FROM `account` WHERE name = '哈基米';

UPDATE `account` SET money = money - 1000 WHERE name = '哈基米';

UPDATE `account` SET money = money + 1000 WHERE name = '曼波';

COMMIT;

ROLLBACK;

START TRANSACTION;

SELECT * FROM `account` WHERE name = '哈基米';

UPDATE `account` SET money = money - 1000 WHERE name = '哈基米';

UPDATE `account` SET money = money + 1000 WHERE name = '曼波';

COMMIT;

ROLLBACK;
