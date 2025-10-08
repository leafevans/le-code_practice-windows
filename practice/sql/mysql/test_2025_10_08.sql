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
