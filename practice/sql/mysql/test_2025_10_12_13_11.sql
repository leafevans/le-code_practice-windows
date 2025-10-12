SELECT CONCAT('Hello ', 'World!');

SELECT LOWER('Hello World!');

SELECT UPPER('Hello World!');

SELECT LPAD('Hello World!', 15, ' ');

SELECT RPAD('Hello World!', 15, ' ');

SELECT TRIM(' Hello World! ');

SELECT SUBSTRING('Hello World!', 1, 5);

UPDATE emp SET workno = LPAD(workno, 5, '0');

SELECT * FROM emp;

SELECT CEIL(3.15);

SELECT FLOOR(1.2);

SELECT MOD(10, 3);

SELECT RAND();

SELECT ROUND(3.1415926, 2);

SELECT LPAD( ROUND(RAND() * 1000000), 6, '0' );

SELECT CURRENT_DATE();

SELECT CURRENT_TIME();

SELECT CURRENT_TIMESTAMP();

SELECT NOW();

SELECT YEAR(NOW()), MONTH(NOW()), DAY(NOW()), HOUR(NOW()), MINUTE(NOW()), SECOND(NOW());

SELECT DATE_ADD(NOW(), INTERVAL 70 DAY);

SELECT DATE_SUB(NOW(), INTERVAL 50 HOUR);

SELECT DATEDIFF('2025-05-16', '2005-05-14');

SELECT id, name, DATEDIFF(NOW(), entrydate) AS `入职天数`
FROM emp
ORDER BY DATEDIFF(NOW(), entrydate) DESC;

SELECT IF(TRUE, 'OK', 'ERROR');

SELECT IFNULL('OK', 'DEFAULT');

SELECT IFNULL(NULL, 'DEFAULT');

SELECT
    name,
    CASE
        WHEN workaddress = '西国情报局' THEN '一线城市'
        WHEN workaddress = '伯林特市政府' THEN '一线城市'
        ELSE '二线城市'
    END AS `工作地址`
FROM emp;

SELECT
    name,
    CASE workaddress
        WHEN '西国情报局' THEN '一线城市'
        WHEN '伯林特市政府' THEN '一线城市'
        ELSE '二线城市'
    END AS `工作地址`
FROM emp;
