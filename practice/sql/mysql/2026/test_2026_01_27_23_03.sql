-- Active: 1768800322203@@127.0.0.1@3307@practice
-- 创造视图
CREATE OR REPLACE VIEW `vw_student_top10` AS (
    SELECT id, name
    FROM student
    WHERE
        id <= 10
);

-- 查询视图
SHOW CREATE VIEW vw_student_top10;

SELECT * FROM vw_student_top10;

SELECT * FROM vw_student_top10 WHERE id < 3;

-- 修改视图
CREATE OR REPLACE VIEW `vw_student_top10` AS (
    SELECT id, name, no
    FROM student
    WHERE
        id <= 10
);

ALTER VIEW `vw_student_top10` AS (
    SELECT id, name
    FROM student
    WHERE
        id <= 10
);

DROP VIEW IF EXISTS `vw_student_top10`;

---
CREATE OR REPLACE VIEW `vw_student_top20` AS
SELECT id, name
FROM student
WHERE
    id <= 20;

SHOW CREATE TABLE student;

SELECT * FROM vw_student_top20;

SELECT * FROM student;

INSERT INTO vw_student_top20 VALUES (6, 'Tom');

INSERT INTO vw_student_top20 VALUES (30, 'Tom');

CREATE OR REPLACE VIEW `vw_student_top20` AS
SELECT id, name
FROM student
WHERE
    id <= 20
WITH
    CASCADED CHECK OPTION;

-- CASCADED
CREATE OR REPLACE VIEW `vw_student_top20` AS
SELECT id, name
FROM student
WHERE
    id <= 20;

INSERT INTO vw_student_10_to_20 VALUES (7, 'Hachimi');

INSERT INTO vw_student_10_to_20 VALUES (26, 'Mambo');

SELECT * FROM vw_student_10_to_20;

CREATE OR REPLACE VIEW `vw_student_top20` AS
SELECT id, name
FROM student
WHERE
    id <= 20;

CREATE OR REPLACE VIEW `vw_student_10_to_20` AS
SELECT id, name
FROM vw_student_top20
WHERE
    id >= 10
WITH
    CASCADED CHECK OPTION;

-- LOCAL
CREATE OR REPLACE VIEW `vw_student_10_to_15` AS
SELECT id, name
FROM vw_student_10_to_20
WHERE
    id <= 15
WITH
    LOCAL CHECK OPTION;

INSERT INTO vw_student_10_to_15 VALUES (9, 'Hachimi');
