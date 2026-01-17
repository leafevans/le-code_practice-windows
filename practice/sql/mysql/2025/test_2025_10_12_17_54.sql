-- Active: 1760262838940@@127.0.0.1@3306
SELECT * FROM `account`;

START TRANSACTION;

INSERT INTO account (id, name, money) VALUES (3, '奶龙', 2000);

COMMIT;
