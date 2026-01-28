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

SELECT * FROM `vw_user_profile`;

CREATE OR REPLACE VIEW `vw_student_course` AS
SELECT
    s.name AS student_name,
    s.no AS student_no,
    c.name AS course_name
FROM
    student s
    LEFT JOIN student_course sc ON s.id = sc.studentid
    LEFT JOIN course c ON sc.courseid = c.id;

SELECT * FROM vw_student_course;
