SELECT *
FROM user;

CREATE USER 'itcast' @'localhost' IDENTIFIED BY '123456';

CREATE USER 'hachimi' @'%' IDENTIFIED BY '123456';

ALTER USER 'hachimi' @'%' IDENTIFIED WITH caching_sha2_password by '123456';

DROP USER 'itcast' @'localhost';

SHOW GRANTS FOR 'hachimi' @'%';

GRANT ALL ON itcast.* TO 'hachimi' @'%';

REVOKE ALL ON itcast.*
FROM 'hachimi' @'%';

USE itcast;

SELECT CONCAT('Hello', ' World!');

SELECT LOWER('Hello World!');

SELECT UPPER('Hello World!');

SELECT LPAD('Hello World!', 15, '-*');

SELECT RPAD('Hello World!', 15, '-*');

SELECT TRIM('   Hello World!   ');

SELECT SUBSTRING("Hello World!", 1, 5);

UPDATE emp
SET workno = LPAD(workno, 5, '0');

SELECT CEILING(3.14);

SELECT FLOOR(3.14);

SELECT MOD(10, 3);

SELECT RAND();

SELECT ROUND(3.1415926, 3);

SELECT LPAD(FLOOR(RAND() * 1000000), 6, '0');

SELECT CURRENT_DATE();

SELECT CURDATE();

SELECT CURTIME();

SELECT NOW();

SELECT YEAR(NOW());

SELECT MONTH(NOW());

SELECT DAY(NOW());

SELECT DATE_ADD(NOW(), INTERVAL 70 DAY);

SELECT DATE_SUB(NOW(), INTERVAL 70 MINUTE);

SELECT DATEDIFF("2025-11-12", "2025-11-01");

SELECT id,
    DATEDIFF(CURDATE(), entrydate) '入职日期'
FROM emp
ORDER BY DATEDIFF(CURDATE(), entrydate) DESC;

SELECT IF(TRUE, 'OK', 'ERROR');

SELECT IFNULL("OK", "DEFAULT");

SELECT IFNULL(NULL, "DEFAULT");

SELECT name,
    (
        CASE
            WHEN workaddress = '西国情报局' THEN '一线城市'
            WHEN workaddress = '伯林特市政府' THEN '一线城市'
            ELSE '二线城市'
        END
    ) AS '工作地址'
FROM emp;

SELECT name,
    (
        CASE
            workaddress
            WHEN '西国情报局' THEN '一线城市'
            WHEN '伯林特市政府' THEN '一线城市'
            ELSE '二线城市'
        END
    ) AS '工作地址'
FROM emp;

USE itcast;

CREATE TABLE IF NOT EXISTS score (
    id INT COMMENT ' ID ',
    name VARCHAR(20) COMMENT ' 姓名 ',
    math INT COMMENT ' 数学 ',
    english INT COMMENT ' 英语 ',
    chinese INT COMMENT ' 语文 '
) COMMENT = ' 学员成绩表 ';

INSERT INTO score (id, name, math, english, chinese)
VALUES (5, ' 柯南 ', 99, 60, 88),
    (6, ' 灰原哀 ', 85, 95, 78),
    (7, ' 毛利兰 ', 70, 80, 90),
    (8, ' 江户川柯南 ', 100, 55, 60),
    (9, ' 路飞 ', 65, 40, 75),
    (10, ' 索隆 ', 80, 35, 60),
    (11, ' 娜美 ', 90, 88, 92),
    (12, ' 乔巴 ', 55, 60, 85),
    (13, ' 艾斯 ', 78, 70, 50),
    (14, ' 小哀 ', 95, 99, 100),
    (15, ' 托尼 · 史塔克 ', 100, 100, 100),
    (16, ' 洛基 ', 60, 90, 45),
    (17, ' 弗兰奇 ', 40, 50, 60),
    (18, ' 乔治 ', 88, 77, 66),
    (19, ' 阿尔萨斯 ', 45, 65, 80),
    (20, ' 小新 ', 30, 20, 10);

SELECT name AS '姓名',
    CASE
        WHEN math >= 85 THEN '优秀'
        WHEN math >= 60 THEN '及格'
        ELSE '不及格'
    END AS '数学',
    CASE
        WHEN english >= 85 THEN '优秀'
        WHEN english >= 60 THEN '及格'
        ELSE '不及格'
    END AS '英语',
    CASE
        WHEN chinese >= 85 THEN '优秀'
        WHEN chinese >= 60 THEN '及格'
        ELSE '不及格'
    END AS '语文'
FROM score;
