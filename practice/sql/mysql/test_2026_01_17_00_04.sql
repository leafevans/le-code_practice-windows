USE mysql;

SELECT * FROM user;

CREATE USER 'test_user' @'localhost' IDENTIFIED BY '123456';

CREATE USER 'hachimi' @'%' IDENTIFIED BY '123456';

ALTER USER 'test_user' @'localhost' IDENTIFIED
WITH
    caching_sha2_password BY '654321';

ALTER USER 'hachimi' @'%' IDENTIFIED
WITH
    mysql_native_password BY '654321';

SHOW GRANTS FOR 'hachimi' @'%';

SHOW GRANTS FOR 'test_user' @'localhost';

GRANT ALL ON test.* TO 'hachimi' @'%';

GRANT
ALTER,
DELETE,
CREATE,
SELECT ON test.emp TO 'test_user' @'localhost';

REVOKE ALL ON test.* FROM 'hachimi' @'%';

DROP USER 'hachimi' @'%', 'test_user' @'localhost';

CREATE USER 'mambo' @'localhost' IDENTIFIED BY '123456';

SELECT * FROM user;

DROP USER 'mambo' @'localhost';

CREATE USER 'mambo' @'localhost' IDENTIFIED
WITH
    caching_sha2_password BY '123456';

SHOW PLUGINS;

CREATE USER 'hachimi' @'localhost' IDENTIFIED
WITH
    mysql_native_password BY '654321';

ALTER USER 'mambo' @'localhost' IDENTIFIED
WITH
    mysql_native_password BY '654321';

ALTER USER 'hachimi' @'localhost' IDENTIFIED
WITH
    caching_sha2_password BY '123456';

SHOW GRANTS FOR 'mambo' @'localhost';

SHOW GRANTS FOR 'hachimi' @'localhost';

GRANT ALL ON test.* TO 'mambo' @'localhost';

GRANT
UPDATE,
DELETE,
CREATE,
ALTER,
DROP,
SELECT,
INSERT
    ON test.* TO 'hachimi' @'localhost';

REVOKE ALL ON test.* FROM 'hachimi' @'localhost';

REVOKE
UPDATE,
CREATE,
ALTER,
DELETE ON test.*
FROM 'mambo' @'localhost';

DROP USER 'mambo' @'localhost', 'hachimi' @'localhost';

CREATE USER 'mambo' @'%' IDENTIFIED
WITH
    caching_sha2_password BY '123456';

SELECT * FROM user;

ALTER USER 'mambo' @'%' IDENTIFIED WITH sha256_password BY '654321';

SHOW GRANTS FOR 'mambo' @'%';

GRANT ALL PRIVILEGES ON test.* TO 'mambo' @'%';

REVOKE ALL PRIVILEGES ON test.* FROM 'mambo' @'%';

DROP USER 'mambo' @'%';
