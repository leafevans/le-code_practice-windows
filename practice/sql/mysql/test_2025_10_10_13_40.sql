CREATE TABLE IF NOT EXISTS dept (
    id INT AUTO_INCREMENT COMMENT 'ID' PRIMARY KEY,
    name VARCHAR(50) NOT NULL COMMENT '部门名称'
) COMMENT = '部门表';

INSERT INTO
    dept
VALUES (1, '研发部'),
    (2, '市场部'),
    (3, '财务部'),
    (4, '销售部'),
    (5, '总经办'),
    (6, '人事部');

CREATE TABLE IF NOT EXISTS emp (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
    name VARCHAR(50) NOT NULL COMMENT '姓名',
    age INT COMMENT '年龄',
    job VARCHAR(20) COMMENT '职位',
    salary INT COMMENT '薪资',
    entrydate DATE COMMENT '入职时间',
    managerid INT COMMENT '直属领导ID',
    dept_id INT COMMENT '部门ID'
) COMMENT = '员工表';

INSERT INTO
    emp (
        id,
        name,
        age,
        job,
        salary,
        entrydate,
        managerid,
        dept_id
    )
VALUES (
        1,
        '漩涡鸣人',
        66,
        '总裁',
        20000,
        '2000-01-01',
        NULL,
        5
    ),
    (
        2,
        '宇智波佐助',
        20,
        '项目经理',
        12500,
        '2005-12-05',
        1,
        1
    ),
    (
        3,
        '艾伦·耶格尔',
        33,
        '开发',
        8400,
        '2000-11-03',
        2,
        1
    ),
    (
        4,
        '蒙奇·D·路飞',
        48,
        '开发',
        11000,
        '2002-02-05',
        2,
        1
    ),
    (
        5,
        '罗罗亚·索隆',
        43,
        '开发',
        10500,
        '2004-09-07',
        3,
        1
    ),
    (
        6,
        '夜神月',
        19,
        '程序员鼓励师',
        6600,
        '2004-10-12',
        2,
        1
    ),
    (
        7,
        '托尼托尼·乔巴',
        60,
        '财务总监',
        8500,
        '2002-09-12',
        1,
        3
    ),
    (
        8,
        '灰原哀',
        19,
        '会计',
        48000,
        '2006-06-02',
        7,
        3
    ),
    (
        9,
        '工藤新一',
        23,
        '出纳',
        5250,
        '2009-05-13',
        7,
        3
    ),
    (
        10,
        '桐人',
        20,
        '市场部总监',
        12500,
        '2004-10-12',
        1,
        2
    ),
    (
        11,
        '琦玉',
        56,
        '职员',
        3750,
        '2006-10-03',
        10,
        2
    ),
    (
        12,
        '碇真嗣',
        19,
        '职员',
        3750,
        '2007-05-09',
        10,
        2
    ),
    (
        13,
        '冈部伦太郎',
        19,
        '职员',
        5500,
        '2009-02-12',
        10,
        2
    ),
    (
        14,
        '皮卡丘',
        88,
        '销售总监',
        14000,
        '2004-10-12',
        1,
        4
    ),
    (
        15,
        '夏目贵志',
        38,
        '销售',
        4600,
        '2004-10-12',
        14,
        4
    ),
    (
        16,
        '神楽',
        40,
        '销售',
        4600,
        '2004-10-12',
        14,
        4
    ),
    (
        17,
        '坂田银时',
        42,
        NULL,
        2000,
        '2011-10-12',
        1,
        NULL
    );

SELECT emp.name, dept.name
FROM emp, dept
WHERE
    emp.dept_id = dept.id;

SELECT emp.name, dept.name
FROM emp
    INNER JOIN dept ON emp.dept_id = dept.id;

SELECT emp.*, dept.name
FROM emp
    LEFT OUTER JOIN dept ON emp.dept_id = dept.id;

SELECT dept.*, emp.*
FROM emp
    RIGHT JOIN dept ON emp.dept_id = dept.id;

SELECT e1.name AS '职员姓名', e2.name AS '上司姓名'
FROM emp AS e1
    JOIN emp AS e2 ON e1.managerid = e2.id;

SELECT e1.name AS '职员姓名', e2.name AS '上司姓名'
FROM emp AS e1
    LEFT JOIN emp AS e2 ON e1.managerid = e2.id;

SELECT *
FROM emp
WHERE
    salary < 5000
UNION
SELECT *
FROM emp
WHERE
    age > 50;

SELECT *
FROM emp
WHERE
    dept_id = (
        SELECT id
        FROM dept
        WHERE
            name = '销售部'
    );

SELECT *
FROM emp
WHERE
    entrydate > (
        SELECT entrydate
        FROM emp
        WHERE
            name = '琦玉'
    );

SELECT *
FROM emp
WHERE
    dept_id IN (
        SELECT id
        FROM dept
        WHERE
            name = '销售部'
            OR name = '市场部'
    );

SELECT *
FROM emp
WHERE
    salary > (
        SELECT MAX(salary)
        FROM emp
        WHERE
            dept_id = (
                SELECT id
                FROM dept
                WHERE
                    name = '研发部'
            )
    );

SELECT *
FROM emp
WHERE
    salary > (
        SELECT MIN(salary)
        FROM emp
        WHERE
            dept_id = (
                SELECT id
                FROM dept
                WHERE
                    name = '总经办'
            )
    );

-- 查询与“碇真嗣”、“神楽”的职位和薪资相同的员工信息。
SELECT *
FROM emp
WHERE (job, salary) IN (
        SELECT job, salary
        FROM emp
        WHERE
            name = '碇真嗣'
            OR name = '神楽'
    );
