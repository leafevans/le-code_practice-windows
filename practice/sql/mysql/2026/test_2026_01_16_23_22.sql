-- Active: 1768561452636@@127.0.0.1@3306@test
SELECT gender, COUNT(id) FROM emp GROUP BY gender;

SELECT gender, AVG(age) FROM emp GROUP BY gender;

SELECT workaddress, COUNT(id) AS `员工数量`
FROM emp
WHERE
    age < 35
GROUP BY
    workaddress
HAVING
    COUNT(id) >= 2;

SELECT workaddress, COUNT(id) AS `员工数量`
FROM emp
WHERE
    age < 35
GROUP BY
    workaddress
HAVING
    COUNT(id) >= 2;

SELECT id, name, age FROM emp ORDER BY age ASC;

SELECT id, name, entrydate FROM emp ORDER BY entrydate DESC;

SELECT id, name, age, entrydate
FROM emp
ORDER BY age ASC, entrydate DESC;

SELECT id, name FROM emp LIMIT 10;

SELECT id, name FROM emp LIMIT 10, 10;

SELECT id, name FROM emp LIMIT 10, 10;

SELECT id, name FROM emp LIMIT 5, 10;

SELECT id, name, age, gender
FROM emp
WHERE
    age IN (15, 16, 17)
    AND gender = '女';

SELECT id, name, age, gender
FROM emp
WHERE
    age BETWEEN 20 AND 40
    AND name LIKE '____';

SELECT gender, COUNT(id) FROM emp WHERE age < 30 GROUP BY gender;

SELECT id, name, age, entrydate
FROM emp
WHERE
    age <= 30
ORDER BY age ASC, entrydate DESC;

SELECT id, name, age, entrydate
FROM emp
WHERE
    age BETWEEN 20 AND 40
    AND gender = '男'
ORDER BY age ASC, entrydate DESC
LIMIT 5;


