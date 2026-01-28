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

SELECT @@GLOBAL.autocommit;
