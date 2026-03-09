SHOW INDEX FROM user;

EXPLAIN
SELECT *
FROM user
WHERE
    profession = "软件工程"
    AND age = 31
    AND status = "0";

EXPLAIN SELECT * FROM user WHERE profession = "软件工程" AND age = 31;

EXPLAIN SELECT * FROM user WHERE profession = "软件工程";

EXPLAIN SELECT * FROM user WHERE age = 31 AND status = '0';

EXPLAIN SELECT * FROM user WHERE status = '0';

EXPLAIN
SELECT *
FROM user
WHERE
    profession = '软件工程'
    AND status = '0';

EXPLAIN
SELECT *
FROM user
WHERE
    profession = '软件工程'
    AND age >= 30
    AND status = '0';

DESC SELECT * FROM user WHERE phone = '17799990015';

EXPLAIN SELECT * FROM user WHERE SUBSTRING(phone, 10, 2) = '15';

EXPLAIN SELECT * FROM user WHERE profession LIKE '软件%';

EXPLAIN SELECT * FROM user WHERE profession LIKE '%件工%';

EXPLAIN SELECT * FROM user WHERE id = 10 OR age = 23;

EXPLAIN SELECT * FROM user WHERE phone = '17799990017' OR age = 23;

CREATE INDEX idx_user_age ON user (age);

EXPLAIN
SELECT *
FROM user
USE INDEX (uniq_user_phone)
WHERE
    phone >= '17799990000';

EXPLAIN SELECT * FROM user WHERE profession IS NULL;

DESC
SELECT *
FROM user
USE INDEX (
    idx_user_profession_age_status
)
WHERE
    profession = '软件工程';
