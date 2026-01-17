CREATE TABLE salgrade (
    grade INT,
    losal INT,
    hisal INT
) comment = '薪资等级表';

INSERT INTO
    salgrade (grade, losal, hisal)
VALUES (1, 0, 3660),
    (2, 3001, 5000),
    (3, 5001, 8000),
    (4, 8001, 10000),
    (5, 10001, 15000),
    (6, 15001, 20000),
    (7, 20001, 25000),
    (8, 25001, 30000);

SELECT e.name, e.age, e.job, d.name
FROM emp e, dept d
WHERE
    e.dept_id = d.id;

SELECT e.name, e.age, e.job, d.name
FROM emp e
    LEFT JOIN dept d ON e.dept_id = d.id
WHERE
    e.age < 30;

SELECT DISTINCT d.* FROM dept d JOIN emp e ON d.id = e.dept_id;

SELECT e.*, d.name
FROM emp e
    LEFT JOIN dept d ON e.dept_id = d.id
WHERE
    e.age > 40;

SELECT e.*, s.*
FROM emp e, salgrade s
WHERE
    e.salary >= s.losal
    AND e.salary <= s.hisal;

SELECT e.*, s.grade
FROM emp e, salgrade s
WHERE
    e.salary BETWEEN s.losal AND s.hisal;

SELECT e.*, s.grade
FROM emp e
    INNER JOIN dept d ON e.dept_id = d.id
    INNER JOIN salgrade s ON e.salary BETWEEN s.losal AND s.hisal
WHERE
    d.name = '研发部';

SELECT AVG(e.salary) AS `平均工资`
FROM emp e
    JOIN dept d ON e.dept_id = d.id
WHERE
    d.name = '研发部';

SELECT e1.*
FROM emp e1
    JOIN emp e2 ON e2.name = '夜神月'
    AND e1.salary > e2.salary

SELECT e.*
FROM emp e
WHERE
    e.salary > (
        SELECT AVG(salary)
        FROM emp
    );

-- 查询低于本部门平均工资的员工信息。
SELECT e1.*
FROM emp e1
WHERE
    e1.salary < (
        SELECT AVG(salary)
        FROM emp e2
        WHERE
            e1.dept_id = e2.dept_id
    );

SELECT e.*
FROM emp e
    JOIN (
        SELECT dept_id, AVG(salary) AS avg_sal
        FROM emp
        GROUP BY
            dept_id
    ) d ON e.dept_id = d.dept_id
WHERE
    e.salary < d.avg_sal;

SELECT d.*, e.emp_count
FROM dept d
    LEFT JOIN (
        SELECT dept_id, COUNT(*) AS emp_count
        FROM emp
        GROUP BY
            dept_id
    ) e ON d.id = e.dept_id;

-- 查询所有学生的选课情况，展示出学生名称、学号、课程名称。
SELECT s.name, s.no, c.name AS course_name
FROM
    student s
    LEFT JOIN student_course sc ON s.id = sc.studentid
    LEFT JOIN course c ON sc.courseid = c.id;
