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

-- 创建 Course 表
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
WHERE Sage BETWEEN 20 AND 30;

-- 查询所有年龄不在 20 到 30 岁之间的学生姓名、系别和年龄
SELECT Sname AS '姓名',
    Sdept AS '系别',
    Sage AS '年龄'
FROM Student
WHERE Sage NOT BETWEEN 20 AND 30;

-- 查询计算机科学系（CS）、数学系（MA）或信息系（IS）的所有学生姓名和性别
SELECT Sname AS '姓名',
    Ssex AS '性别'
FROM Student
WHERE Sdept IN ('CS', 'MA', 'IS');

-- 查询不在计算机科学系、数学系或信息系的所有学生姓名和性别
SELECT Sname AS '姓名',
    Ssex AS '性别'
FROM Student
WHERE Sdept NOT IN ('CS', 'MA', 'IS');

--- 查询学号为 '200215121' 的学生的所有信息
SELECT *
FROM Student
WHERE Sno LIKE '200215121';

-- 精确查询学号为 '200215121' 的学生的所有信息
SELECT *
FROM Student
WHERE Sno = '200215121';

-- 查询所有姓名以 'M' 开头的学生的姓名、学号和性别
SELECT Sname AS '姓名',
    Sno AS '学号',
    Ssex AS '性别'
FROM Student
WHERE Sname LIKE 'M%';

-- 查询所有姓名为 'Joh' 开头且长度为 5 个字符的学生姓名
SELECT Sname AS '姓名'
FROM Student
WHERE Sname LIKE 'Joh__';

-- 查询所有姓名中第三个字符为 'r' 的学生姓名和学号
SELECT Sname AS '姓名',
    Sno AS '学号'
FROM Student
WHERE Sname LIKE '__r%';

-- 查询课程名精确为 'DB_Design' 的课程编号和学分
SELECT Cno AS '课程编号',
    Ccredit AS '学分'
FROM Course
WHERE Cname = 'DB_Design';

-- 查询课程名以 'DB_' 开头，且包含 'i' 后跟两个任意字符的课程的所有信息
SELECT *
FROM Course
WHERE Cname LIKE 'DB_%'
    AND Cname LIKE '%i__%';

-- 查询所有成绩为空的选课记录中的学号和课程编号
SELECT Sno AS '学号',
    Cno AS '课程编号'
FROM SC
WHERE Grade IS NULL;

-- 查询所有成绩不为空的选课记录中的学号和课程编号
SELECT Sno AS '学号',
    Cno AS '课程编号'
FROM SC
WHERE Grade IS NOT NULL;

-- 查询计算机科学系（CS）中年龄小于 20 岁的学生姓名
SELECT Sname AS '姓名'
FROM Student
WHERE Sdept = 'CS'
    AND Sage < 20;

-- 查询计算机科学系、数学系或信息系的所有学生姓名和性别
SELECT Sname AS '姓名',
    Ssex AS '性别'
FROM Student
WHERE Sdept IN ('CS', 'MA', 'IS');

-- 查询选修了课程 '3' 的所有学生的学号和成绩，并按成绩降序排列
SELECT Sno AS '学号',
    Grade AS '成绩'
FROM SC
WHERE Cno = '3'
ORDER BY Grade DESC;

-- 查询所有学生的信息，先按系别升序排序，再按年龄降序排序
SELECT *
FROM Student
ORDER BY Sdept ASC,
    Sage DESC;

-- 统计学生总数
SELECT COUNT(*) AS '学生总数'
FROM Student;

-- 统计至少选了一门课的不同学生人数
SELECT COUNT(DISTINCT Sno) AS '选课学生人数'
FROM SC;

-- 计算课程 '1' 的平均成绩
SELECT AVG(Grade) AS '课程 1 平均成绩'
FROM SC
WHERE Cno = '1';

-- 计算学号为 '200215121' 的学生所选课程的总学分
SELECT SUM(C.Ccredit) AS '总学分'
FROM SC
    LEFT JOIN Course AS C ON SC.Cno = C.Cno
WHERE SC.Sno = '200215121';

-- 统计每门课程的选课人数
SELECT Cno AS '课程编号',
    COUNT(Sno) AS '选课人数'
FROM SC
GROUP BY Cno;

-- 查询选修了多于一门课程的学生学号
SELECT Sno AS '学号'
FROM SC
GROUP BY Sno
HAVING COUNT(Cno) > 1;

-- 查询所有选课学生的学号、姓名、性别、系别、课程编号和成绩
SELECT S.Sno AS '学号',
    S.Sname AS '姓名',
    S.Ssex AS '性别',
    S.Sdept AS '系别',
    SC.Cno AS '课程编号',
    SC.Grade AS '成绩'
FROM Student AS S
    RIGHT JOIN SC ON S.Sno = SC.Sno;

-- 查询所有课程的先修课程关系，即每门课程及其直接先修课程的编号
SELECT C.Cno AS '课程编号',
    CP.Cno AS '先修课程编号'
FROM Course C
    LEFT JOIN Course CP ON C.Cpno = CP.Cno;

-- 查询所有学生（包括未选课的学生）的学号、姓名、性别、年龄、系别，以及他们的选课信息
SELECT S.Sno AS '学号',
    S.Sname AS '姓名',
    S.Ssex AS '性别',
    S.Sdept AS '系别',
    SC.Cno AS '课程编号',
    SC.Grade AS '成绩'
FROM Student AS S
    LEFT JOIN SC ON S.Sno = SC.Sno;

-- 查询所有选课记录中的学号、姓名、性别、年龄、系别、课程编号和成绩，包括那些在学生表中没有对应记录的学生
SELECT S.Sno AS '学号',
    S.Sname AS '姓名',
    S.Ssex AS '性别',
    S.Sdept AS '系别',
    SC.Cno AS '课程编号',
    SC.Grade AS '成绩'
FROM Student AS S
    RIGHT JOIN SC ON S.Sno = SC.Sno;

-- 查询选修了课程 '2' 且成绩高于 90 分的学生的学号和姓名
SELECT S.Sno AS '学号',
    S.Sname AS '姓名'
FROM Student S
    INNER JOIN SC ON S.Sno = SC.Sno
WHERE SC.Cno = '2'
    AND SC.Grade > 90;

-- 查询所有选课学生的学号、姓名、课程名称和成绩
SELECT S.Sno AS '学号',
    S.Sname AS '姓名',
    C.Cname AS '课程名称',
    SC.Grade
FROM SC
    LEFT JOIN Student AS S ON SC.Sno = S.Sno
    LEFT JOIN Course AS C ON SC.Cno = C.Cno;

-- 查询所有选修了课程 '2' 的学生姓名
SELECT Sname AS '姓名'
FROM Student
WHERE Sno IN (
        SELECT Sno
        FROM SC
        WHERE Cno = '2'
    );

-- 查询所有与名为 'Sara' 的学生同系的学生学号、姓名和系别
SELECT Sno AS '学号',
    Sname AS '姓名',
    Sdept AS '系别'
FROM Student
WHERE Sdept IN (
        SELECT Sdept
        FROM Student
        WHERE Sname = 'Sara'
    );

-- 查询所有选修了课程名为 'IS' 的学生的学号和姓名
SELECT Sno AS '学号',
    Sname AS '姓名'
FROM Student
WHERE Sno IN (
        SELECT Sno
        FROM SC
        WHERE Cno IN (
                SELECT Cno
                FROM Course
                WHERE Cname = 'IS'
            )
    );

-- 查询所有非计算机科学系的学生中，年龄小于计算机科学系中任一学生年龄的学生姓名和年龄
SELECT Sname AS '姓名',
    Sage AS '年龄'
FROM Student
WHERE Sdept <> 'CS'
    AND Sage < ANY (
        SELECT Sage
        FROM Student
        WHERE Sdept = 'CS'
    );

-- 查询所有非计算机科学系的学生中，年龄小于计算机科学系中最大年龄的学生姓名和年龄
SELECT Sname AS '姓名',
    Sage AS '年龄'
FROM Student
WHERE Sdept <> 'CS'
    AND Sage < (
        SELECT MAX(Sage)
        FROM Student
        WHERE Sdept = 'CS'
    );

-- 查询所有选修了课程 '1' 的学生姓名
SELECT Sname AS '姓名'
FROM Student AS S
WHERE EXISTS (
        SELECT 1
        FROM SC
        WHERE S.Sno = SC.Sno
            AND SC.Cno = '1'
    );

-- 查询所有未选修课程 '1' 的学生姓名
SELECT Sname AS '姓名'
FROM Student AS S
WHERE NOT EXISTS (
        SELECT 1
        FROM SC
        WHERE SC.Sno = S.Sno
            AND SC.Cno = '1'
    );

-- 查询所有与名为 'Sara' 的学生同系的学生学号、姓名和系别
SELECT Sno AS '学号',
    Sname AS '姓名',
    Sdept AS '系别'
FROM Student AS S1
WHERE EXISTS (
        SELECT 1
        FROM Student AS S2
        WHERE S1.Sdept = S2.Sdept
            AND S2.Sname = 'Sara'
    );

-- 查询所有选修了所有课程的学生姓名
SELECT Sname AS '姓名'
FROM Student AS S
WHERE NOT EXISTS (
        SELECT 1
        FROM Course AS C
        WHERE NOT EXISTS (
                SELECT 1
                FROM SC
                WHERE SC.Sno = S.Sno
                    AND SC.Cno = C.Cno
            )
    );

-- 查询所有选修了学号为 '200215122' 的学生所选修的所有课程的学生学号
SELECT Sno AS '学号'
FROM SC AS SC1
WHERE NOT EXISTS (
        SELECT 1
        FROM SC AS SC2
        WHERE SC2.Sno = '200215122'
            AND NOT EXISTS (
                SELECT 1
                FROM SC AS SC3
                WHERE SC1.Sno = SC3.Sno
                    AND SC2.Cno = SC3.Cno
            )
    );

-- 查询所有计算机科学系的学生或年龄小于等于 19 岁的学生信息（去除重复）
SELECT *
FROM Student
WHERE Sdept = 'CS'
UNION
SELECT *
FROM Student
WHERE Sage <= 19;

-- 查询所有选修了课程 '1' 或课程 '2' 的学生学号（去除重复）
SELECT Sno AS '学号'
FROM SC
WHERE Cno = '1'
UNION
SELECT Sno AS '学号'
FROM SC
WHERE Cno = '2';
