SHOW DATABASES;

SELECT DATABASE();

CREATE DATABASE IF NOT EXISTS test;

DROP DATABASE IF EXISTS test;

USE itcast;

SHOW TABLES;

DESC emp;

SHOW CREATE TABLE emp;

CREATE TABLE IF NOT EXISTS tb_user (
    id INT COMMENT '编号',
    name VARCHAR(50) COMMENT '姓名',
    age INT COMMENT '年龄',
    gender VARCHAR(1) COMMENT '性别'
) COMMENT = '用户表';

CREATE DATABASE IF NOT EXISTS test;

USE test;

CREATE TABLE emp (
    id INT COMMENT '编号',
    workno VARCHAR(10) COMMENT '工号',
    name VARCHAR(10) COMMENT '姓名',
    gender CHAR(1) COMMENT '性别',
    age TINYINT UNSIGNED COMMENT '年龄',
    idcard CHAR(18) COMMENT '身份证号',
    entrydate DATE COMMENT '入职时间'
) COMMENT = '员工表';

ALTER TABLE emp ADD nickname VARCHAR(20) COMMENT '昵称';

ALTER TABLE emp MODIFY nickname VARCHAR(30);

ALTER TABLE emp CHANGE nickname username VARCHAR(30) COMMENT '用户名';

ALTER TABLE emp DROP username;

ALTER TABLE emp RENAME TO emplyee;
