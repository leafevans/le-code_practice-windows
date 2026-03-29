CREATE PROCEDURE `sp_student_count`()
BEGIN
SELECT COUNT(*) FROM student;
END;

CALL `sp_student_count` ();

SELECT *
FROM information_schema.ROUTINES
WHERE
    ROUTINE_SCHEMA = 'test';

SHOW CREATE PROCEDURE sp_student_count;

DROP PROCEDURE IF EXISTS sp_student_count;

SET SESSION autocommit = 0;

---
SET @testname := 'Hachimi';

SET @testage := 10;

SET @testgendet := '女', @testhobby = 'Golang';

SELECT @testname, @testage, @testgender, @testhobby;

---

SELECT @testcolor := 'red';

SELECT COUNT(*) INTO @testcount FROM user;

---
CREATE PROCEDURE `sp_student_count`()
BEGIN
DECLARE `student_count` INT DEFAULT 0;
SELECT COUNT(*) INTO student_count FROM student;
SELECT student_count;
END;

CALL `sp_student_count` ();
---
CREATE PROCEDURE `sp_score_grade`()
BEGIN
DECLARE `score` INT DEFAULT 22;
DECLARE `res` VARCHAR(5) CHARSET utf8mb4;
IF score >= 85 THEN 
SET res := '优秀';
ELSEIF score > 60 THEN 
SET res := '及格';
ELSE
SET res := '不及格';
END IF;
SELECT res;
END;

CALL sp_score_grade ();

DROP PROCEDURE sp_score_grade;
---
DROP PROCEDURE IF EXISTS sp_score_grade;

CREATE PROCEDURE sp_score_grade (IN score INT, OUT res VARCHAR(10) CHARSET utf8mb4)
BEGIN
IF score >= 85 THEN 
SET res := '优秀';
ELSEIF score >= 60 THEN 
SET res := '及格';
ELSE
SET res := '不及格';
END IF;
END;

CALL sp_score_grade (88, @res);

SELECT @res;
---
DROP PROCEDURE `sp_halve_score`;

CREATE PROCEDURE `sp_halve_score`(INOUT score DOUBLE)
BEGIN
SET score := score * 0.5;
END;

SELECT @score := 127;

CALL sp_halve_score (@score);

SELECT @score;
---
DROP PROCEDURE IF EXISTS sp_month_to_quarter;

CREATE PROCEDURE sp_month_to_quarter(IN p_month TINYINT)
BEGIN
DECLARE res VARCHAR(10) CHARSET utf8mb4;
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
---
DROP PROCEDURE IF EXISTS `sp_sum_to_n`;

CREATE PROCEDURE IF NOT EXISTS `sp_sum_to_n`(IN n INT)
BEGIN
DECLARE total INT DEFAULT 0;
WHILE n > 0 DO
SET total := total + n;
SET n := n - 1;
END WHILE;
SELECT total;
END;

CALL `sp_sum_to_n` (50);
---
DROP PROCEDURE IF EXISTS `sp_sum_to_n`;

CREATE PROCEDURE IF NOT EXISTS `sp_sum_to_n`(IN n INT)
BEGIN
DECLARE total INT DEFAULT 0;
REPEAT
SET total := total + n;
SET n := n - 1;
UNTIL n = 0
END REPEAT;
SELECT total;
END;

CALL `sp_sum_to_n` (490);
---
DROP PROCEDURE IF EXISTS `sp_sum_to_n`;

CREATE PROCEDURE IF NOT EXISTS `sp_sum_to_n`(IN n INT)
BEGIN
DECLARE total INT DEFAULT 0;
sum_loop: LOOP
SET total := total + n;
SET n := n - 1;
IF n = 0 THEN
LEAVE sum_loop;
END IF;
END LOOP sum_loop;
SELECT total;
END;

CALL sp_sum_to_n (50);
---
DROP PROCEDURE IF EXISTS `sp_sum_even_to_n`;

CREATE PROCEDURE IF NOT EXISTS `sp_sum_even_to_n`(IN n INT)
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

CALL `sp_sum_even_to_n` (50);
---
DROP PROCEDURE IF EXISTS `sp_student_count`;

CREATE PROCEDURE IF NOT EXISTS `sp_student_count`()
BEGIN
DECLARE `student_count` INT DEFAULT 0;
SELECT * INTO student_count FROM student;
SELECT student_count;
END;

CALL sp_student_count ();
---
DROP PROCEDURE IF EXISTS `sp_export_users_by_age`;

CREATE PROCEDURE IF NOT EXISTS `sp_export_users_by_age`(IN p_max_age INT)
BEGIN
DECLARE v_name VARCHAR(100) CHARSET utf8mb4;
DECLARE v_profession VARCHAR(100) CHARSET utf8mb4;
DECLARE cur_users CURSOR FOR SELECT name, profession FROM user WHERE age <= p_max_age;

DROP TABLE IF EXISTS tmp_users;

CREATE TABLE IF NOT EXISTS `tmp_users`(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    profession VARCHAR(100)
) CHARSET = utf8mb4;

OPEN cur_users;
WHILE TRUE DO
FETCH cur_users INTO v_name, v_profession;
INSERT INTO tmp_users VALUES (NULL, v_name, v_profession);
END WHILE;

CLOSE cur_users;
END;

CALL sp_export_users_by_age (40);
---
DROP PROCEDURE IF EXISTS `sp_export_users_by_age`;

CREATE PROCEDURE IF NOT EXISTS `sp_export_users_by_age`(IN p_max_age INT)
BEGIN
DECLARE v_name VARCHAR(100) CHARSET utf8mb4;
DECLARE v_profession VARCHAR(100) CHARSET utf8mb4;
DECLARE cur_users CURSOR FOR SELECT name, profession FROM user WHERE age <= p_max_age;
DECLARE EXIT CURSOR FOR NOT FOUND CLOSE cur_users;

DROP TABLE IF EXISTS `tmp_users`;

CREATE TABLE IF NOT EXISTS `tmp_users`(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    profession VARCHAR(100)
) CHARSET = utf8mb4;

OPEN cur_users;
WHILE TRUE DO
FETCH cur_users INTO v_name, v_profession;
INSERT INTO tmp_users VALUE (NULL, v_name, v_profession);
END WHILE;
END;

CALL `sp_export_users_by_age` (45);
---
DROP FUNCTION IF EXISTS `fn_sum_to_n`;

CREATE FUNCTION IF NOT EXISTS `fn_sum_to_n`(n INT)
RETURNS INT DETERMINISTIC
BEGIN
DECLARE total INT DEFAULT 0;
WHILE n > 0 DO 
SET total := total + n;
SET n := n - 1;
END WHILE;
RETURN total;
END;

SELECT `fn_sum_to_n` (50);
