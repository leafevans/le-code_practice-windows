-- Active: 1768579431897@@127.0.0.1@3306@test
START TRANSACTION;

UPDATE account SET money = money - 1000 WHERE name = '哈基米';

COMMIT;

ROLLBACK;

START TRANSACTION;

UPDATE account SET money = money + 1000 WHERE name = '哈基米';

COMMIT;

START TRANSACTION;

INSERT INTO account (id, name, money) VALUES (3, '奶龙', 2000);

COMMIT;
