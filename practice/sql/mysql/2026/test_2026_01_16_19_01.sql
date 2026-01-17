-- Active: 1768561452636@@127.0.0.1@3306
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
        '女',
        18,
        '01234567890123456X',
        '2004-11-12'
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
        '男',
        20,
        '112345678901234567',
        '2004-05-14'
    ),
    (
        4,
        '4',
        'Gibson',
        '男',
        21,
        '212345678901234567',
        '2004-11-01'
    );

UPDATE emp SET name = 'Hachimi', age = 20 WHERE id = 1;

DELETE FROM emp WHERE gender IN ('男', '女');

TRUNCATE TABLE emp;
