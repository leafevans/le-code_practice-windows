CREATE TABLE IF NOT EXISTS student (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
    name VARCHAR(10) COMMENT '姓名',
    no VARCHAR(10) COMMENT '学号'
) COMMENT '学生表';

INSERT INTO
    student
VALUES (NULL, '黛绮丝', '2000100101'),
    (NULL, '谢逊', '2000100102'),
    (NULL, '殷天正', '2000100103'),
    (NULL, '韦一笑', '2000100104');

CREATE TABLE IF NOT EXISTS course (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
    name VARCHAR(10) COMMENT '课程名称'
) COMMENT '课程表';

INSERT INTO
    course
VALUES (NULL, 'Java'),
    (NULL, 'PHP'),
    (NULL, 'MySQL'),
    (NULL, "Hadoop");

DROP TABLE IF EXISTS student_course;

CREATE TABLE IF NOT EXISTS student_course (
    studentid INT NOT NULL COMMENT '学生 ID',
    courseid INT NOT NULL COMMENT '课程 ID',
    CONSTRAINT fk_studentid FOREIGN KEY (studentid) REFERENCES student (id),
    CONSTRAINT fk_courseid FOREIGN KEY (courseid) REFERENCES course (id),
    PRIMARY KEY (studentid, courseid)
) COMMENT '学生课程中间表';

CREATE TABLE IF NOT EXISTS student_course (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
    studentid INT NOT NULL COMMENT '学生ID',
    courseid INT NOT NULL COMMENT '课程ID',
    CONSTRAINT fk_studentid FOREIGN KEY (studentid) REFERENCES student (id),
    CONSTRAINT fk_courseid FOREIGN KEY (courseid) REFERENCES course (id),
    CONSTRAINT uq_student_course UNIQUE KEY (studentid, courseid)
) COMMENT '学生课程中间表';

INSERT INTO
    student_course (studentid, courseid)
VALUES (1, 1),
    (1, 2),
    (1, 3),
    (2, 2),
    (2, 3),
    (3, 4);

ALTER TABLE student_course
ADD CONSTRAINT uq_student_course UNIQUE KEY (studentid, courseid);

ALTER TABLE student_course
ADD UNIQUE KEY uq_student_course (studentid, courseid);

DESC student_course;
