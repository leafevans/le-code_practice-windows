-- Active: 1768800322203@@127.0.0.1@3307@practice
-- 创建
CREATE PROCEDURE `sp_student_count`()
BEGIN
SELECT COUNT(*) FROM student;
END;

-- 调用
CALL sp_student_count ();

-- 查看
SELECT *
FROM information_schema.ROUTINES
WHERE
    ROUTINE_SCHEMA = 'practice';

SHOW CREATE PROCEDURE sp_student_count;

DROP PROCEDURE IF EXISTS sp_student_count;

-- 查看系统变量
SHOW VARIABLES;

SHOW SESSION VARIABLES LIKE 'auto%';

SELECT SESSION autocommit;

-- 设置系统变量
SET SESSION autocommit = 0;

SET SESSION autocommit = 1;

-- 赋值
SET @testname = 'Hachimi';

SET @testage := 10;

SET @testgender := '女', @testhobby := 'Golang';

SELECT @testcolor := 'red';

SELECT COUNT(*) INTO @testcount FROM user;
-- 使用
SELECT @testname, @testage, @testgender, @testhobby;

SELECT @testcolor, @testcount;

-- 声明
-- 赋值
CREATE PROCEDURE `sp_student_count`()
BEGIN
DECLARE `student_count` INT DEFAULT 0;
SELECT COUNT(*) INTO student_count FROM student;
SELECT student_count;
END;

CALL sp_student_count ();

--
CREATE PROCEDURE `sp_score_grade`()
BEGIN
DECLARE `score` INT DEFAULT 55;
DECLARE `res` VARCHAR(10) CHARACTER SET utf8mb4;
IF score > 85 THEN 
SET res := '优秀';
ELSEIF score >= 60 THEN
SET res := '及格';
ELSE 
SET res := '不及格';
END IF;
SELECT res;
END;

DROP PROCEDURE IF EXISTS sp_score_grade;

CALL sp_score_grade ();

--
CREATE PROCEDURE `sp_score_grade` (IN score INT, OUT res VARCHAR(10) CHARACTER SET utf8mb4)
BEGIN
IF score >= 85 THEN
SET res := '优秀';
ELSEIF score >= 60 THEN 
SET res := '及格';
ELSE
SET res := '不及格';
END IF;
END;

CALL sp_score_grade (95, @res);

SELECT @res;

CREATE PROCEDURE `sp_halve_score`(INOUT score DOUBLE) 
BEGIN
SET score := score * 0.5;
END;

SELECT @score := 127;

CALL sp_halve_score (@score);

SELECT @score;

--

CREATE PROCEDURE `sp_month_to_quarter`(IN p_month TINYINT)
BEGIN
DECLARE res VARCHAR(10) CHARACTER SET utf8mb4;
CASE 
WHEN p_month BETWEEN 1 AND 3 THEN SET res = '第一季度';
WHEN p_month BETWEEN 4 AND 6 THEN SET res = '第二季度';
WHEN p_month BETWEEN 7 AND 9 THEN SET res = '第三季度';
WHEN p_month BETWEEN 10 AND 12 THEN SET res = '第四季度';
ELSE SET res = '未知';
END CASE;
SELECT CONCAT('您输入的月份为：', p_month, '，它所属的季度为：', res);
END;

CALL sp_month_to_quarter (10);

--
CREATE PROCEDURE `sp_sum_to_n`(IN n INT)
BEGIN
DECLARE total INT DEFAULT 0;
WHILE n > 0 DO
SET total := total + n;
SET n := n - 1;
END WHILE;
SELECT total;
END;

CALL sp_sum_to_n (50);

DROP PROCEDURE sp_sum_to_n;

--
CREATE PROCEDURE `sp_sum_to_n`(IN n INT)
BEGIN
DECLARE total INT DEFAULT 0;
REPEAT
SET total := total + n;
SET n := n - 1;
UNTIL n = 0
END REPEAT;
SELECT total;
END;

CALL sp_sum_to_n (45);

--
CREATE PROCEDURE `sp_sum_to_n`(IN n INT)
BEGIN
DECLARE total INT DEFAULT 0;
sum_loop :LOOP
IF n = 0 THEN 
LEAVE sum_loop;
END IF;
SET total := total + n;
SET n := n - 1;
END LOOP sum_loop;
SELECT total;
END;

CALL sp_sum_to_n (45);

CREATE PROCEDURE `sp_sum_even_to_n`(IN n INT)
BEGIN
DECLARE total INT DEFAULT 0;

sum_loop: LOOP
IF n = 0 THEN 
LEAVE sum_loop;
END IF;

IF MOD(n, 2) = 1 THEN
SET n := n - 1;
ITERATE sum_loop;
END IF;

SET total := total + n;
SET n := n - 1;
END LOOP sum_loop;

SELECT total;
END;

CALL sp_sum_even_to_n (45);

--

DROP PROCEDURE IF EXISTS sp_export_users_by_age;

CREATE PROCEDURE IF NOT EXISTS `sp_export_users_by_age`(IN p_max_age INT)
BEGIN
DECLARE v_name VARCHAR(100) CHARACTER SET utf8mb4;
DECLARE v_profession VARCHAR(100) CHARACTER SET utf8mb4;
DECLARE cur_users CURSOR FOR SELECT name, profession FROM user WHERE age <= p_max_age;
DECLARE EXIT HANDLER FOR SQLSTATE '02000' CLOSE cur_users;

DROP TABLE IF EXISTS tmp_users;

CREATE TABLE IF NOT EXISTS `tmp_users` (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    profession VARCHAR(100)
) CHARACTER SET = utf8mb4;

OPEN cur_users;
WHILE TRUE DO
FETCH cur_users INTO v_name, v_profession;
INSERT INTO tmp_users VALUES (NULL, v_name, v_profession);
END WHILE;
END;

CALL sp_export_users_by_age (30);

SELECT * FROM tmp_users;
