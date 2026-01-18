-- Active: 1768729232514@@127.0.0.1@3306
-- 查看当前数据库支持的存储引擎
SHOW ENGINES;

CREATE TABLE my_myisam (id INT, name VARCHAR(10)) ENGINE = MyISAM;

CREATE TABLE my_memory (id INT, name VARCHAR(10)) ENGINE = Memory;
