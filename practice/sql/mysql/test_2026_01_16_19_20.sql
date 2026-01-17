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
        '00001',
        '蒙奇·D·路飞',
        '男',
        19,
        '330106200005051234',
        '杭州',
        '2010-05-05'
    ),
    (
        2,
        '00002',
        '罗罗诺亚·索隆',
        '男',
        21,
        '330106199811112345',
        '杭州',
        '2010-11-11'
    ),
    (
        3,
        '00003',
        '漩涡鸣人',
        '男',
        32,
        '110101199010106666',
        '北京',
        '2015-03-23'
    ),
    (
        4,
        '00004',
        '宇智波佐助',
        '男',
        32,
        '110101199007238888',
        '北京',
        '2015-06-01'
    ),
    (
        5,
        '00005',
        '御坂美琴',
        '女',
        14,
        '44030120100502001X',
        '深圳',
        '2023-01-01'
    ),
    (
        6,
        '00006',
        '五条悟',
        '男',
        28,
        '310110199512079999',
        '上海',
        '2018-12-07'
    ),
    (
        7,
        '00007',
        '坂田银时',
        '男',
        27,
        '510101199610107777',
        '成都',
        '2013-04-01'
    ),
    (
        8,
        '00008',
        '蕾姆',
        '女',
        17,
        '500234200602025201',
        '重庆',
        '2021-02-02'
    ),
    (
        9,
        '00009',
        '阿尼亚·福杰',
        '女',
        6,
        '110101201810010000',
        '北京',
        '2024-05-01'
    ),
    (
        10,
        '00010',
        '利姆鲁·特恩佩斯特',
        '男',
        37,
        '42010119850230xxxx',
        '武汉',
        '2020-10-10'
    ),
    (
        11,
        '00011',
        '初音未来',
        '女',
        16,
        '310115200708313939',
        '上海',
        '2007-08-31'
    ),
    (
        12,
        '00012',
        '桐谷和人',
        '男',
        17,
        '410101200810071314',
        '郑州',
        '2022-11-06'
    ),
    (
        13,
        '00013',
        '惣流·明日香·兰格雷',
        '女',
        14,
        '610101201012041234',
        '西安',
        '2015-12-04'
    ),
    (
        14,
        '00014',
        '利威尔·阿克曼',
        '男',
        30,
        '430101199312255678',
        '长沙',
        '2012-09-09'
    ),
    (
        15,
        '00015',
        '绫小路清隆',
        '男',
        16,
        '32010120071020666X',
        '南京',
        '2023-09-01'
    ),
    (
        16,
        '00016',
        '琦玉',
        '男',
        25,
        '110101199801010000',
        '北京',
        '2020-01-01'
    );

SELECT idcard, name, age FROM emp;

SELECT DISTINCT workaddress AS `工作地址` FROM emp;

SELECT DISTINCT age FROM emp ORDER BY age ASC;

SELECT id, name FROM emp WHERE age = 17;

SELECT id, name FROM emp WHERE age > 25;

SELECT id, name FROM emp WHERE age >= 25;

SELECT id, name FROM emp WHERE idcard IS NULL;

SELECT id, name FROM emp WHERE idcard IS NOT NULL;

SELECT id, name FROM emp WHERE age <> 16;

SELECT id, name, age FROM emp WHERE age BETWEEN 15 AND 20;

SELECT id, name, age, gender
FROM emp
WHERE
    gender = '男'
    AND age < 25;

SELECT id, name, age FROM emp WHERE age IN (16, 19, 14);

SELECT id, name FROM emp WHERE name LIKE '___';

SELECT id, name, idcard FROM emp WHERE idcard LIKE '%X';

SELECT COUNT(id) AS `员工数量` FROM emp;

SELECT AVG(age) AS `平均年龄` FROM emp;

SELECT MAX(age) AS `最大年龄` FROM emp;

SELECT MIN(age) AS `最小年龄` FROM emp;

SELECT SUM(age) FROM emp WHERE workaddress = '北京';

SELECT gender, COUNT(id) FROM emp GROUP BY gender;

SELECT gender, AVG(age) FROM emp GROUP BY gender;




