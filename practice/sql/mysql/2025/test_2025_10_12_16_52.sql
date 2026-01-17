SELECT e.name, e.age, e.job, d.*
FROM emp e
    JOIN dept d ON e.dept_id = d.id;

SELECT e.name, e.age, e.job, d.*
FROM emp e
    LEFT JOIN dept d ON e.dept_id = d.id
WHERE
    e.age < 30;

SELECT DISTINCT d.* FROM dept d JOIN emp e ON d.id = e.dept_id;

SELECT e.*, d.name FROM emp e LEFT JOIN dept d ON e.dept_id = d.id;

SELECT e.*, s.grade
FROM emp e
    JOIN salgrade s ON e.salary BETWEEN s.losal AND s.hisal;

SELECT e.*, s.grade
FROM emp e
    JOIN salgrade s ON e.salary BETWEEN s.losal AND s.hisal
    JOIN dept d ON e.dept_id = d.id
WHERE
    d.name = '研发部';

SELECT AVG(e.salary)
FROM emp e
    LEFT JOIN dept d ON e.dept_id = d.id
WHERE
    d.name = '研发部';

SELECT e1.*
FROM emp e1
    LEFT JOIN emp e2 ON e2.name = '夜神月'
WHERE
    e1.salary > e2.salary;

SELECT e.*
FROM emp e
WHERE
    e.salary > (
        SELECT AVG(salary)
        FROM emp
    );

SELECT e1.*
FROM emp e1
WHERE
    e1.salary < (
        SELECT AVG(e2.salary)
        FROM emp e2
        WHERE
            e1.dept_id = e2.dept_id
    );

SELECT e.*
FROM emp e
    LEFT JOIN (
        SELECT dept_id, AVG(salary) avg_sal
        FROM emp
        GROUP BY
            dept_id
    ) d ON e.dept_id = d.dept_id
WHERE
    e.salary < d.avg_sal;

SELECT d.*, COUNT(e.id)
FROM dept d
    LEFT JOIN emp e ON d.id = e.dept_id
GROUP BY
    d.id;

SELECT s.name, s.no, c.name
FROM
    student s
    LEFT JOIN student_course sc ON s.id = sc.studentid
    LEFT JOIN course c ON sc.courseid = c.id;
