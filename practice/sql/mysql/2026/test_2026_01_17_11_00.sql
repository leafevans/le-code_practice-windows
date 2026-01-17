-- Active: 1768579431897@@127.0.0.1@3306@test
SELECT CONCAT("Hello", " World!");

SELECT LOWER("Hello World!")

SELECT UPPER("Hello World!");

SELECT LPAD('Hello World!', 15, '-*-');

SELECT RPAD('Hello World!', 15, '@-@');

SELECT TRIM(" Hello World!");

SELECT SUBSTRING("Hello World!", 1, 5);

SELECT SUBSTRING("Hello World!", 1, 5);

UPDATE emp SET workno = LPAD(workno, 10, '0');

SELECT CEIL(3.14);

SELECT FLOOR(3.14);

SELECT MOD(10, 3);

SELECT RAND();

SELECT ROUND(10 / 3, 3);

SELECT ROUND(3.1415926, 3);

SELECT LPAD( FLOOR(RAND() * 1000000), 6, '0' );

SELECT CURDATE();

SELECT CURRENT_DATE();

SELECT CURTIME();

SELECT CURRENT_TIMESTAMP();

SELECT NOW(), CURRENT_TIMESTAMP();

SELECT YEAR(NOW());

SELECT MONTH(NOW());

SELECT DATE_ADD(NOW(), INTERVAL 70 DAY);

SELECT DAY(NOW());

SELECT DATE_SUB(NOW(), INTERVAL 120 SECOND);

SELECT DATEDIFF("2026-11-12", "2026-5-16");

SELECT id, name, DATEDIFF(NOW(), entrydate) AS `入职天数`
FROM emp
ORDER BY `入职天数`;

SELECT IF(TRUE, 'OK', 'ERROR');

SELECT IFNULL('OK', 'DEFAULT');

SELECT IFNULL(NULL, 'DEFAULT');

SELECT
    name,
    CASE
        WHEN workaddress = '北京' THEN '一线城市'
        WHEN workaddress = '杭州' THEN '一线城市'
        ELSE '二线城市'
    END AS `工作地址`
FROM emp;

SELECT IF(TRUE, 'OK', 'ERROR');

SELECT IFNULL(NULL, 'DEFAULT');

SELECT
    name,
    CASE
        WHEN workaddress = '北京' THEN '一线城市'
        WHEN workaddress = '杭州' THEN '一线城市'
        ELSE '二线城市'
    END AS `工作地址`
FROM emp;

SELECT
    id,
    name,
    CASE workaddress
        WHEN workaddress = '北京' THEN '一线城市'
        WHEN workaddress = '杭州' THEN '一线城市'
        ELSE '二线城市'
    END AS `工作地址`
FROM emp;

CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT COMMENT '编号',
    name VARCHAR(10) NOT NULL UNIQUE COMMENT '姓名',
    age INT CHECK (
        age > 0
        AND age <= 120
    ) COMMENT '年龄',
    status CHAR(1) DEFAULT '1' COMMENT '状态',
    gender CHAR(1) COMMENT '性别'
) COMMENT '用户';

INSERT INTO
    users (name, age, status, gender)
VALUES ('Miku', 19, '1', '女');

INSERT INTO
    users (name, age, status, gender)
VALUES ('Fender', 20, '0', '男');

INSERT INTO users (name, age, gender) VALUES ('Mike', 30, '女');
