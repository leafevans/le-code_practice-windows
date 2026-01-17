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
    id INT auto_increment PRIMARY KEY COMMENT '主键ID',
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
    CONSTRAINT fk_studentid FOREIGN KEY (studentid) REFERENCES student (id),
    CONSTRAINT fk_courseid FOREIGN KEY (courseid) REFERENCES course (id)
) COMMENT = '学生课程中间表';

INSERT INTO
    student_course
VALUES (NULL, 1, 1),
    (NULL, 1, 2),
    (NULL, 1, 3),
    (NULL, 2, 2),
    (NULL, 2, 3),
    (NULL, 3, 4);

CREATE TABLE IF NOT EXISTS tb_user (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
    name VARCHAR(10) COMMENT '姓名',
    age INT COMMENT '年龄',
    gender CHAR(1) COMMENT '性别',
    phone CHAR(11) COMMENT '手机号'
) COMMENT = '用户基本信息表';

CREATE TABLE IF NOT EXISTS tb_user_edu (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
    degree VARCHAR(20) COMMENT '学历',
    major VARCHAR(50) COMMENT '专业',
    primaryschool VARCHAR(50) COMMENT '小学',
    middleschool VARCHAR(50) COMMENT '中学',
    university VARCHAR(50) COMMENT '大学',
    userid INT UNIQUE COMMENT '用户ID',
    CONSTRAINT fk_userid FOREIGN KEY (userid) REFERENCES tb_user (id)
) COMMENT = '用户教育信息表';

INSERT INTO
    tb_user (id, name, age, gender, phone)
VALUES (
        NULL,
        '黄渤',
        45,
        '1',
        '18800001111'
    ),
    (
        NULL,
        '冰冰',
        35,
        '2',
        '18800002222'
    ),
    (
        NULL,
        '码云',
        55,
        '1',
        '18800008888'
    ),
    (
        NULL,
        '李彦宏',
        50,
        '1',
        '18800009999'
    );

INSERT INTO
    tb_user_edu (
        id,
        degree,
        major,
        primaryschool,
        middleschool,
        university,
        userid
    )
VALUES (
        NULL,
        '本科',
        '舞蹈',
        '静安区第一小学',
        '静安区第一中学',
        '北京舞蹈学院',
        1
    ),
    (
        NULL,
        '硕士',
        '表演',
        '朝阳区第一小学',
        '朝阳区第一中学',
        '北京电影学院',
        2
    ),
    (
        NULL,
        '本科',
        '英语',
        '杭州市第一小学',
        '杭州市第一中学',
        '杭州师范大学',
        3
    ),
    (
        NULL,
        '本科',
        '应用数学',
        '阳泉第一小学',
        '阳泉区第一中学',
        '清华大学',
        4
    );

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

SELECT * FROM emp, dept WHERE emp.dept_id = dept.id;

SELECT * FROM emp INNER JOIN dept ON emp.dept_id = dept.id;

SELECT e.name, d.name FROM emp e JOIN dept d ON e.dept_id = d.id;

SELECT e.*, d.name
FROM emp e
    LEFT JOIN dept d ON e.dept_id = d.id;

SELECT d.*, e.* FROM emp e RIGHT JOIN dept d ON e.dept_id = d.id;

-- 查询员工及其所属领导的名字
SELECT e1.name `职员姓名`, e2.name `上司姓名`
FROM emp e1
    JOIN emp e2 ON e1.managerid = e2.id;

SELECT e1.name `职员姓名`, e2.name `上司姓名`
FROM emp e1
    LEFT JOIN emp e2 ON e1.managerid = e2.id;

SELECT *
FROM emp
WHERE
    salary < 5000
UNION ALL
SELECT *
FROM emp
WHERE
    age > 50;

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
    salary > ALL (
        SELECT salary
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
    salary > any (
        SELECT salary
        FROM emp
        WHERE
            dept_id = (
                SELECT id
                FROM dept
                WHERE
                    name = '总经办'
            )
    );

SELECT *
FROM emp
WHERE (salary, managerid) = (
        SELECT salary, managerid
        FROM emp
        WHERE
            name = '桐人'
    );

SELECT *
FROM emp
WHERE (job, salary) IN (
        SELECT job, salary
        FROM emp
        WHERE
            name = '碇真嗣'
            OR name = '神楽'
    );

SELECT *
FROM (
        SELECT *
        FROM emp
        WHERE
            entrydate > '2006-01-01'
    ) e
    LEFT JOIN dept d ON e.dept_id = d.id;

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

