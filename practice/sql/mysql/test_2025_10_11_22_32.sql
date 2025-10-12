-- Active: 1760152257164@@127.0.0.1@3306@test_db
CREATE TABLE emp (
    id INT COMMENT '编号',
    workno VARCHAR(10) COMMENT '工号',
    name VARCHAR(10) COMMENT '姓名',
    gender CHAR(1) COMMENT '性别',
    age TINYINT UNSIGNED COMMENT '年龄',
    idcard CHAR(18) COMMENT '身份证号',
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
        entrydate
    )
VALUES (
        1,
        '1',
        'Hachimi',
        '男',
        18,
        '01234567890123456X',
        '2005-05-14'
    );

INSERT INTO
    emp
VALUES (
        2,
        '2',
        'Mambo',
        '男',
        19,
        '012345678901234567',
        '2005-05-16'
    );

INSERT INTO
    emp
VALUES (
        3,
        '3',
        'Fender',
        '女',
        20,
        '112345678901234567',
        '2004-11-12'
    ),
    (
        4,
        '4',
        'Gibson',
        '女',
        21,
        '212345678901234567',
        '2004-11-01'
    );

DESC emp;

ALTER TABLE emp ADD nickname varchar(20) COMMENT '昵称';

ALTER TABLE emp MODIFY nickname varchar(30);

ALTER TABLE emp CHANGE nickname username varchar(20) COMMENT '用户名';

ALTER TABLE emp DROP username;

ALTER TABLE emp RENAME TO employee;

DROP TABLE IF EXISTS employee;

TRUNCATE emp;

UPDATE emp SET name = 'Kobe', gender = '男' WHERE id = 1;

DELETE FROM emp WHERE id = 4;

SELECT idcard, name, age FROM emp;

