START TRANSACTION;

UPDATE `account` SET money = money - 1000 WHERE name = '哈基米';

COMMIT;

--

UPDATE `account` SET money = 2000 WHERE name = '哈基米' OR name = '曼波';

DELETE FROM `account` WHERE id = 3;

--

START TRANSACTION

UPDATE `account` SET money = money + 1000 WHERE name = '哈基米';

COMMIT;

--

START TRANSACTION;

INSERT INTO `account` (id, name, money) value (3, '奶龙', 2000);

COMMIT;
