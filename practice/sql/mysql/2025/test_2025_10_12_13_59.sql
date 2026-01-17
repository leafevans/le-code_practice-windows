CREATE TABLE IF NOT EXISTS `users` (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    name VARCHAR(10) UNIQUE NOT NULL COMMENT '姓名',
    age INT CHECK (
        age > 0
        AND age <= 120
    ) COMMENT '年龄',
    `status` CHAR(1) DEFAULT '1' COMMENT '状态',
    gender CHAR(1) COMMENT '性别'
) COMMENT = '用户表';

INSERT INTO
    users (name, age, status, gender)
VALUES ('TOM', 19, '1', '男'),
    ('Mike', 15, '0', '男');

INSERT INTO
    users (name, age, status, gender)
VALUES ('Mike', 19, 1, '男');

SELECT * FROM users;

INSERT INTO
    users (name, age, `status`, gender)
VALUES ('James', 20, '0', '男');
