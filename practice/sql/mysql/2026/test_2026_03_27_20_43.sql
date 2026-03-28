CREATE OR REPLACE VIEW `vw_student_top10` AS (
    SELECT id, name
    FROM student
    WHERE
        id <= 10
);

SHOW CREATE VIEW `vw_student_top10`;

SELECT * FROM vw_student_top10;

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

CREATE OR REPLACE VIEW `vw_student_top20` AS (
    SELECT id, name
    FROM student
    WHERE
        id <= 20
);

INSERT INTO `vw_student_top20` VALUES (6, 'Tom');

DELETE FROM student WHERE id = 30;

INSERT INTO `vw_student_top20` VALUES (30, 'Tim');

SELECT * FROM vw_student_top20;

CREATE OR REPLACE VIEW `vw_student_top20` AS
SELECT id, name
FROM student
WHERE
    id <= 20
WITH
    CASCADED CHECK OPTION;

CREATE OR REPLACE VIEW `vw_student_top20` AS
SELECT id, name
FROM student
WHERE
    id <= 20;

CREATE OR REPLACE VIEW `vw_student_10_to_20` AS
SELECT id, name
FROM `vw_student_top20`
WHERE
    id >= 10
WITH
    CASCADED CHECK OPTION;

CREATE OR REPLACE VIEW vw_student_10_to_15 AS
SELECT id, name
FROM vw_student_10_to_20
WHERE
    id <= 15
WITH
    LOCAL CHECK OPTION;

CREATE OR REPLACE VIEW `vw_student_count` AS
SELECT COUNT(*)
FROM student;

SELECT * FROM user;

CREATE OR REPLACE VIEW `vw_user_profile` AS
SELECT
    id,
    name,
    profession,
    age,
    gender,
    status,
    createtime
FROM `user`;

SELECT * FROM `vw_user_profile` LIMIT 100;

CREATE OR REPLACE VIEW `vw_student_course` AS
SELECT
    s.name AS student_name,
    s.no AS student_no,
    c.name AS course_name
FROM
    student s
    LEFT JOIN student_course sc ON s.id = sc.studentid
    LEFT JOIN course c ON c.id = sc.courseid;
