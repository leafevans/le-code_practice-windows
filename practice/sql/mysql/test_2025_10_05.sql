-- 先删依赖表
DROP TABLE IF EXISTS SC;

-- 后删除被依赖表
DROP TABLE IF EXISTS Course;

DROP TABLE IF EXISTS Student;

-- 创建 Student 表
CREATE TABLE IF NOT EXISTS Student (
    Sno CHAR(9) PRIMARY KEY,
    Sname VARCHAR(20) NOT NULL,
    Ssex CHAR(1) NOT NULL CHECK (Ssex IN ('M', 'F')),
    Sage INT NOT NULL CHECK (Sage > 0),
    Sdept VARCHAR(10)
);

-- 创建 SC 表
CREATE TABLE IF NOT EXISTS Course (
    Cno CHAR(1) PRIMARY KEY,
    Cname VARCHAR(20) NOT NULL,
    Cpno CHAR(1),
    Ccredit INT,
    FOREIGN KEY (Cpno) REFERENCES Course (Cno) ON DELETE
    SET NULL
);

-- 创建 SC 表
CREATE TABLE IF NOT EXISTS SC(
    Sno CHAR(9),
    Cno CHAR(1),
    Grade INT CHECK (
        Grade BETWEEN 0 AND 100
    ),
    PRIMARY KEY (Sno, Cno),
    FOREIGN KEY (Sno) REFERENCES Student (Sno) ON DELETE CASCADE,
    FOREIGN KEY (Cno) REFERENCES Course (Cno) ON DELETE CASCADE
);

-- 插入数据
INSERT INTO Student (Sno, Sname, Ssex, Sage, Sdept)
VALUES ('200215121', 'Michael', 'M', 20, 'CS'),
    ('200215122', 'Sara', 'F', 19, 'CS'),
    ('200215123', 'Grechen', 'F', 18, 'MA'),
    ('200215125', 'Lincoln', 'M', 19, 'IS');

INSERT INTO Course (Cno, Cname, Cpno, Ccredit)
VALUES ('2', 'MA', NULL, 2),
    ('6', 'DP', NULL, 2),
    ('4', 'OS', '6', 3),
    ('7', 'PASCAL', '6', 4),
    ('5', 'DS', '7', 4),
    ('1', 'DB', '5', NULL),
    ('3', 'IS', '1', 4);

INSERT INTO SC (Sno, Cno, Grade)
VALUES ('200215121', '1', 92),
    ('200215121', '2', 85),
    ('200215121', '3', 88),
    ('200215122', '2', 90),
    ('200215122', '3', 80);

-- 验证数据
SELECT 'Student 表数据' AS '验证结果';

SELECT *
FROM Student;

SELECT 'Course 表数据' AS '验证结果';

SELECT *
FROM Course;

SELECT 'SC 表数据' AS '验证结果';

SELECT *
FROM SC;

-- 查询所有学生的学号和姓名
SELECT Sno AS '学号',
    Sname AS '姓名'
FROM Student;

--- 查询所有学生的姓名、学号和所在系
SELECT Sname AS '姓名',
    Sno AS '学号',
    Sdept AS '所在系'
FROM Student;

-- 查询所有学生的完整信息
SELECT *
FROM Student;

-- 查询所有学生的学号、姓名、性别、年龄和所在系
SELECT Sno AS '学号',
    Sname AS '姓名',
    Ssex AS '性别',
    Sage AS '年龄',
    Sdept AS '所在系'
FROM Student;

-- 查询所有学生的姓名和出生年份（假设当前年份为 2007 年）
SELECT Sname AS '姓名',
    2007 - Sage AS '出生年份'
FROM Student;

-- 查询所有学生的姓名、固定标签 'Year of Birth: '、出生年份和小写系别
SELECT Sname AS '姓名',
    'Year of Birth: ' AS '标签',
    2007 - Sage AS '出生年份',
    LOWER(Sdept) AS '小写系别'
FROM Student;

-- 查询所有学生的姓名（别名为 NAME）、固定标签（别名为 BIRTH）、出生年份（别名为 BIRTHDAY）和小写系别（别名为 DEPARTMENT）
SELECT Sname AS 'NAME',
    'Year of Birth: ' AS 'BIRTH',
    2007 - Sage AS 'BIRTHDAY',
    LOWER(Sdept) AS 'DEPARTMENT'
FROM Student;

-- 查询所有选课记录中的学号
SELECT Sno AS '学号'
FROM SC;

-- 查询所有至少选了一门课的不同学生学号
SELECT DISTINCT Sno AS '学号'
FROM SC;

-- 查询所有选课记录中的学号（包括重复值）
SELECT ALL Sno AS '学号'
FROM SC;

-- 查询计算机科学系（CS）的所有学生姓名
SELECT Sname AS '姓名'
FROM Student
WHERE Sdept = 'CS';

-- 查询所有年龄小于 20 岁的学生姓名和年龄
SELECT Sname AS '姓名',
    Sage AS '年龄'
FROM Student
WHERE Sage < 20;

-- 查询所有有不及格成绩（分数低于 60）的不同学生学号
SELECT DISTINCT Sno AS '学号'
FROM SC
WHERE Grade < 60;

-- 查询所有年龄在 20 到 30 岁之间（含）的学生姓名、系别和年龄
SELECT Sname AS '姓名',
    Sdept AS '系别',
    Sage AS '年龄'
FROM Student
WHERE Sage NOT BETWEEN 20 AND 30;
