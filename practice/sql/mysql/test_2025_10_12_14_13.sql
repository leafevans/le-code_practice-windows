DROP TABLE IF EXISTS dept;

DROP TABLE IF EXISTS emp;

CREATE TABLE IF NOT EXISTS dept (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
    name VARCHAR(50) NOT NULL COMMENT '部门名称'
) COMMENT = '部门表';

INSERT INTO
    dept
VALUES (1, '研发部'),
    (2, '市场部'),
    (3, '财务部'),
    (4, '销售部'),
    (5, '总经办');

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
    emp
VALUES (
        1,
        '路飞',
        66,
        '总裁',
        20000,
        '2000-01-01',
        NULL,
        5
    ),
    (
        2,
        '鸣人',
        20,
        '项目经理',
        12500,
        '2005-12-05',
        1,
        1
    ),
    (
        3,
        '索隆',
        33,
        '开发',
        8400,
        '2000-11-03',
        2,
        1
    ),
    (
        4,
        '佐助',
        48,
        '开发',
        11000,
        '2002-02-05',
        2,
        1
    ),
    (
        5,
        '白胡子',
        43,
        '开发',
        10500,
        '2004-09-07',
        3,
        1
    ),
    (
        6,
        '小昭',
        19,
        '程序员鼓励师',
        6600,
        '2004-10-12',
        2,
        1
    );

ALTER TABLE emp
ADD CONSTRAINT fk_emp_dept_id FOREIGN KEY emp (dept_id) REFERENCES dept (id);

ALTER TABLE emp DROP FOREIGN KEY fk_emp_dept_id;

ALTER TABLE emp
ADD CONSTRAINT fk_emp_dept_id FOREIGN KEY (dept_id) REFERENCES dept (id) ON DELETE CASCADE ON UPDATE CASCADE;

UPDATE dept SET id = 1 WHERE id = 6;

SELECT * FROM emp;

DELETE FROM dept WHERE id = 1;

TRUNCATE emp;

ALTER TABLE emp
ADD CONSTRAINT fk_emp_dept_id FOREIGN KEY (dept_id) REFERENCES dept (id) ON DELETE SET NULL ON UPDATE SET NULL;

DELETE FROM dept WHERE id = 1;
