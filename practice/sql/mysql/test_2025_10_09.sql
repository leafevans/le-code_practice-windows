CREATE TABLE `users`(
    id int PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    name VARCHAR(10) NOT NULL UNIQUE COMMENT '姓名',
    age INT CHECK (
        age > 0
        AND age <= 120
    ) COMMENT '年龄',
    `status` CHAR(1) DEFAULT '1' COMMENT '状态',
    gender CHAR(1) COMMENT '性别'
) COMMENT = '用户表';

INSERT INTO `users`(name, age, `status`, gender)
VALUES('TOM', 19, '1', '男'),
    ("Mike", 15, '0', '男');

INSERT INTO users(name, age, `status`, gender)
VALUES(NULL, 19, '1', '女');

INSERT INTO users(name, age, `status`, gender)
VALUES(NULL, 19, '1', '女');

INSERT INTO users(name, age, `status`, gender)
VALUES('Mike', 19, '1', '女');

INSERT INTO users(name, age, `status`, gender)
VALUES('James', 20, '0', '男');

INSERT INTO users(name, age, `status`, gender)
VALUES('Carry', -1, '1', '男');

INSERT INTO users(name, age, gender)
VALUES('Kobe', 18, '男');
