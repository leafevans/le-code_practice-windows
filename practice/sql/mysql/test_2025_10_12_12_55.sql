USE mysql;

SELECT * FROM user;

CREATE USER 'itcast' @'localhost' IDENTIFIED BY '123456';

CREATE USER 'mambo' @'%' IDENTIFIED BY '123456';

ALTER USER 'mambo' IDENTIFIED WITH caching_sha2_password BY '1234';

DROP USER 'itcast' @'localhost';

DROP USER 'mambo' @'%';

ALTER USER 'hachimi' identified WITH mysql_native_password BY '1234';

SHOW GRANTS FOR 'hachimi' @'%';

GRANT CREATE ON itcast.* TO 'hachimi' @'%';

GRANT SELECT ON itcast.* TO 'hachimi' @'%';

REVOKE SELECT ON itcast.* FROM 'hachimi' @'%';

GRANT
SELECT,
INSERT
,
UPDATE,
DELETE,
ALTER,
DROP,
CREATE ON itcast.* TO 'hachimi' @'%';

REVOKE SELECT ON itcast.* FROM 'hachimi' @'%';

REVOKE UPDATE ON itcast.* FROM 'hachimi' @'%';

REVOKE
INSERT
,
    DELETE,
    CREATE,
DROP,
ALTER ON itcast.*
FROM 'hachimi' @'%';
