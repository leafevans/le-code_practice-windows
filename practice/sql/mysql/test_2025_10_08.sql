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

SELECT CONCAT("Hello", " World!");

SELECT LOWER("Hello World!");

SELECT UPPER("Hello World!");
