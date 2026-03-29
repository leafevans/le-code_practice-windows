CREATE TABLE `user_log` (
    id INT(11) NOT NULL AUTO_INCREMENT,
    operation VARCHAR(20) NOT NULL COMMENT '操作类型',
    operation_time DATETIME NOT NULL COMMENT '操作时间',
    operation_id INT(11) NOT NULL COMMENT '操作 ID',
    operation_params VARCHAR(500) COMMENT '操作参数',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

---
DROP TRIGGER IF EXISTS `trg_user_after_insert`;

CREATE TRIGGER `trg_user_after_insert`
AFTER INSERT ON user FOR EACH ROW
BEGIN
    INSERT INTO user_log (
        id, operation, operation_time, operation_id, operation_params
    ) VALUES (
        NULL, 'INSERT', NOW(), NEW.id, CONCAT(
            '插入的数据内容为：id = ', NEW.id, 
            ', name = ', NEW.name,
            ', phone = ', NEW.phone,
            ', email = ', NEW.email,
            ', profession = ', NEW.profession
        )
    );
END;

INSERT INTO
    user (
        id,
        name,
        phone,
        email,
        profession,
        age,
        gender,
        status,
        createtime
    )
VALUES (
        25,
        '织田信长',
        '18809091212',
        'odanobunaga@163.com',
        '软件工程',
        23,
        '1',
        '1',
        NOW()
    );

SELECT * FROM user_log;
---
DROP TRIGGER IF EXISTS `trg_user_after_update`;

CREATE TRIGGER IF NOT EXISTS `trg_user_after_update`
AFTER UPDATE ON user FOR EACH ROW
BEGIN
    INSERT INTO user_log (
        id, operation, operation_time, operation_id, operation_params
    ) VALUES (
        NULL, 'UPDATE', NOW(), NEW.id, CONCAT(
            '更新之前的数据内容为：id = ', OLD.id,
            ', name = ', OLD.name,
            ', phone = ', OLD.phone,
            ', email = ', OLD.email,
            ', profession = ', OLD.profession,
            ' | 更新之后的数据内容为：id = ', NEW.id,
            ', name = ', NEW.name,
            ', phone = ', NEW.phone,
            ', email = ', NEW.email,
            ', profession = ', NEW.profession
        )
    );
END;

UPDATE user SET age = 32 WHERE id = 23;

SELECT * FROM user_log;

UPDATE user SET profession = '会计' WHERE id <= 3;

SELECT * FROM user_log;
---
DROP TRIGGER IF EXISTS `trg_user_after_delete`;

CREATE TRIGGER IF NOT EXISTS `trg_user_after_delete`
AFTER DELETE ON `user` FOR EACH ROW
BEGIN
INSERT INTO
    user_log (
        id, operation, operation_time, operation_id, operation_params
    ) 
VALUES (
    NULL, 'DELETE', NOW(), OLD.id, CONCAT(
        '删除的数据内容为：id = ', OLD.id,
        ', name = ', OLD.name,
        ', phone = ', OLD.phone,
        ', email = ', OLD.email,
        ', profession = ', OLD.profession
    )
);
END;

DELETE FROM user WHERE id = 25;

SELECT * FROM user_log;
---
SELECT
    object_type,
    object_schema,
    object_name,
    lock_type,
    lock_duration
FROM performance_schema.metadata_locks;
---
SELECT
    object_schema,
    object_name,
    index_name,
    lock_type,
    lock_mode,
    lock_data
FROM performance_schema.data_locks;
