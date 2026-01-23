DROP TABLE IF EXISTS `student_course`;

DROP TABLE IF EXISTS `student`;

DROP TABLE IF EXISTS `course`;

CREATE TABLE `course` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `student` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) NOT NULL,
    `no` VARCHAR(20) NOT NULL UNIQUE,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `student_course` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `studentid` INT NOT NULL,
    `courseid` INT NOT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_student` (`studentid`),
    KEY `idx_course` (`courseid`),
    CONSTRAINT `fk_sc_student` FOREIGN KEY (`studentid`) REFERENCES `student` (`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_sc_course` FOREIGN KEY (`courseid`) REFERENCES `course` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

INSERT INTO
    `course` (`name`)
VALUES ('Java'),
    ('PHP'),
    ('MySQL'),
    ('Hadoop');

INSERT INTO
    `student` (`name`, `no`)
VALUES ('黛绮丝', '2000100101'),
    ('谢弛', '2000100102'),
    ('殷天正', '2000100103'),
    ('韦一笑', '2000100104');

INSERT INTO
    `student_course` (`studentid`, `courseid`)
VALUES (1, 1),
    (1, 2),
    (1, 3),
    (2, 2),
    (2, 3),
    (3, 4);

EXPLAIN
SELECT s.*, c.*
FROM
    student s,
    course c,
    student_course sc
WHERE
    s.id = sc.studentid
    AND c.id = sc.courseid;

EXPLAIN
SELECT *
FROM student
WHERE
    id IN (
        SELECT studentid
        FROM student_course
        WHERE
            courseid = (
                SELECT id
                FROM course
                WHERE
                    name = 'MySQL'
            )
    );
