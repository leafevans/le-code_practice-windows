DROP TABLE IF EXISTS dept;

CREATE TABLE IF NOT EXISTS dept (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
    name VARCHAR(50) NOT NULL COMMENT '部门名称'
) COMMENT '部门表';

INSERT INTO
    dept (id, name)
VALUES (1, '研发部'),
    (2, '市场部'),
    (3, '财务部'),
    (4, '销售部'),
    (5, '总经办');

DROP TABLE IF EXISTS emp;

CREATE TABLE IF NOT EXISTS emp (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
    name VARCHAR(50) NOT NULL COMMENT '姓名',
    age INT COMMENT '年龄',
    job VARCHAR(20) COMMENT '职位',
    salary INT COMMENT '薪资',
    entrydate DATE COMMENT '入职时间',
    managerid INT COMMENT '直属部门 ID',
    dept_id INT COMMENT '部门 ID'
) COMMENT '员工表';

INSERT INTO
    emp (
        name,
        age,
        job,
        salary,
        entrydate,
        managerid,
        dept_id
    )
VALUES (
        '路飞',
        66,
        '总裁',
        20000,
        '2000-01-01',
        NULL,
        5
    ),
    (
        '鸣人',
        20,
        '项目经理',
        12500,
        '2005-12-05',
        1,
        1
    ),
    (
        '索隆',
        33,
        '开发',
        8400,
        '2000-11-03',
        2,
        1
    ),
    (
        '佐助',
        48,
        '开发',
        11000,
        '2002-02-05',
        2,
        1
    ),
    (
        '白胡子',
        43,
        '开发',
        10500,
        '2004-09-07',
        3,
        1
    ),
    (
        '小昭',
        19,
        '程序员鼓励师',
        6600,
        '2004-10-12',
        2,
        1
    );

DESC emp;

ALTER TABLE emp
ADD CONSTRAINT fk_emp_dept_id FOREIGN KEY emp (dept_id) REFERENCES dept (id) ON UPDATE CASCADE ON DELETE SET NULL;

UPDATE dept SET id = 6 WHERE id = 1;

DELETE FROM dept WHERE id = 6;

SHOW DATABASES;

SELECT DATABASE();
