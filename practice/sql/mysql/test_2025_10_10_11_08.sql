CREATE TABLE IF NOT EXISTS student (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
    name VARCHAR(10) COMMENT '姓名',
    no VARCHAR(10) COMMENT '学号'
) COMMENT = '学生表';

INSERT INTO
    student
VALUES (NULL, '黛绮丝', '2000100101'),
    (NULL, '谢逊', '2000100102'),
    (NULL, '殷天正', '2000100103'),
    (NULL, '韦一笑', '2000100104');

CREATE TABLE IF NOT EXISTS course (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
    name VARCHAR(10) COMMENT '课程名称'
) COMMENT = '课程表';

INSERT INTO
    course
VALUES (NULL, 'Java'),
    (NULL, 'PHP'),
    (NULL, 'MySQL'),
    (NULL, 'Hadoop');

CREATE TABLE IF NOT EXISTS student_course (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    studentid INT NOT NULL COMMENT '学生ID',
    courseid INT NOT NULL COMMENT '课程ID',
    CONSTRAINT fk_courseid FOREIGN KEY (courseid) REFERENCES course (id),
    CONSTRAINT fk_studentid FOREIGN KEY (studentid) REFERENCES student (id)
) COMMENT = '学生课程中间表';

INSERT INTO
    student_course
VALUES (NULL, 1, 1),
    (NULL, 1, 2),
    (NULL, 1, 3),
    (NULL, 2, 2),
    (NULL, 2, 3),
    (NULL, 3, 4);
