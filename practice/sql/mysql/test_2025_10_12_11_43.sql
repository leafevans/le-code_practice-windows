DROP TABLE IF EXISTS emp;

CREATE TABLE IF NOT EXISTS emp (
    id int COMMENT '编号',
    workno varchar(10) COMMENT '工号',
    name varchar(10) COMMENT '姓名',
    gender char(1) COMMENT '性别',
    age tinyint UNSIGNED COMMENT '年龄',
    idcard char(18) COMMENT '身份证号',
    workaddress varchar(50) COMMENT '工作地址',
    entrydate DATE COMMENT '入职时间'
) COMMENT = '员工表';

INSERT INTO
    emp (
        id,
        workno,
        name,
        gender,
        age,
        idcard,
        workaddress,
        entrydate
    )
VALUES (
        1,
        '0',
        '劳埃德',
        '男',
        27,
        '310101199801011234',
        '西国情报局',
        '2021-04-01'
    ),
    (
        2,
        '1',
        '约尔',
        '女',
        27,
        '310101199802022345',
        '伯林特市政府',
        '2021-04-02'
    ),
    (
        3,
        '2',
        '阿尼亚',
        '女',
        6,
        '310101201903033456',
        '伊甸学园',
        '2022-09-01'
    ),
    (
        4,
        '3',
        '弗兰基',
        '男',
        30,
        '310101199504044567',
        '弗兰基情报屋',
        '2020-10-10'
    ),
    (
        5,
        '4',
        '西尔维亚',
        '女',
        40,
        '310101198505055678',
        '西国情报局',
        '2015-08-15'
    ),
    (
        6,
        '5',
        '达米安',
        '男',
        6,
        '310101201906066789',
        '伊甸学园',
        '2022-09-01'
    ),
    (
        7,
        '6',
        '五条悟',
        '男',
        28,
        '410101199712077890',
        '咒术高专',
        '2016-04-07'
    ),
    (
        8,
        '7',
        '虎杖悠仁',
        '男',
        15,
        '410101201003208901',
        '咒术高专',
        '2023-06-01'
    ),
    (
        9,
        '8',
        '伏黑惠',
        '男',
        15,
        '410101201012229012',
        '咒术高专',
        '2023-04-10'
    ),
    (
        10,
        '9',
        '钉崎野蔷薇',
        '女',
        16,
        '410101200908070123',
        '咒术高专',
        '2023-04-15'
    ),
    (
        11,
        '10',
        '七海建人',
        '男',
        27,
        '410101199807031234',
        '咒术高专',
        '2021-01-20'
    ),
    (
        12,
        '11',
        '禅院真希',
        '女',
        16,
        '410101200901202345',
        '咒术高专',
        '2022-04-10'
    ),
    (
        13,
        '12',
        '狗卷棘',
        '男',
        17,
        '410101200810233456',
        '咒术高专',
        '2021-04-10'
    ),
    (
        14,
        '13',
        '熊猫',
        '男',
        17,
        '410101200803054567',
        '咒术高专',
        '2021-04-10'
    ),
    (
        15,
        '14',
        '贝姬',
        '女',
        6,
        '310101201908288901',
        '伊甸学园',
        '2022-09-01'
    ),
    (
        16,
        '15',
        '家入硝子',
        '女',
        28,
        '410101199711079012',
        '咒术高专',
        '2016-04-07'
    );

SELECT idcard, name, age FROM emp;

SELECT workaddress AS `工作地址` FROM emp;

SELECT DISTINCT age FROM emp;

SELECT id, name FROM emp WHERE age = 16;

SELECT id, name FROM emp WHERE age > 25;

SELECT id, name, age FROM emp WHERE age >= 25;

SELECT id, name, idcard FROM emp WHERE idcard IS NULL;

SELECT id, name FROM emp WHERE idcard IS NOT NULL;

SELECT id name, age FROM emp WHERE age <> 16;

SELECT id, name, age FROM emp WHERE age BETWEEN 15 AND 20;

SELECT id, name FROM emp WHERE gender = '男' AND age < 25;

SELECT id, name FROM emp WHERE age IN (18, 20, 40)

SELECT id, name FROM emp WHERE name LIKE '___';

SELECT id, name, idcard FROM emp WHERE idcard LIKE '%1';

SELECT COUNT(id) FROM emp;

SELECT AVG(age) FROM emp;

SELECT MAX(age) FROM emp;

SELECT MIN(age) FROM emp;

SELECT SUM(age) FROM emp WHERE workaddress = '西国情报局';

SELECT gender, COUNT(id) FROM emp GROUP BY gender;

SELECT gender, AVG(age) FROM emp GROUP BY gender;

SELECT workaddress, COUNT(id) AS `员工数量`
FROM emp
GROUP BY
    workaddress
HAVING
    COUNT(id) >= 3;

SELECT id, name, age FROM emp ORDER BY age ASC;

DESC emp;

SELECT id, name, entrydate FROM emp ORDER BY entrydate DESC;

SELECT id, name, age, entrydate
FROM emp
ORDER BY age ASC, entrydate DESC;

SELECT id, name FROM emp LIMIT 10;

SELECT id, name FROM emp LIMIT 10, 10;

SELECT id, name, age, gender
FROM emp
WHERE
    age IN (15, 16, 17)
    AND gender = '女';

SELECT id, name, age, gender
FROM emp
WHERE
    age BETWEEN 20 AND 40
    AND name LIKE '___';

SELECT gender, COUNT(id) FROM emp WHERE age < 30 GROUP BY gender;

SELECT name, age, entrydate
FROM emp
WHERE
    age < 30
ORDER BY age ASC, entrydate DESC;

SELECT id, name, age, entrydate
FROM emp
WHERE
    age BETWEEN 20 AND 40
    AND gender = '男'
ORDER BY age ASC, entrydate DESC
LIMIT 5;
