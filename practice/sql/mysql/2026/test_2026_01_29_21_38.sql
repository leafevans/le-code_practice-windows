-- Active: 1773015712865@@127.0.0.1@3307
DROP TABLE IF EXISTS `user`;

CREATE TABLE IF NOT EXISTS `user` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(50) NOT NULL,
    `email` VARCHAR(100) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `phone` VARCHAR(20) NULL,
    `profession` VARCHAR(100) NULL,
    `age` TINYINT UNSIGNED NULL,
    `created_at` DATETIME NOT NULL DEFAULT NOW(),
    `status` TINYINT NOT NULL DEFAULT 1
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

INSERT INTO
    `user` (
        `name`,
        `email`,
        `password`,
        `status`,
        `phone`,
        `profession`,
        `age`
    )
VALUES (
        'alice',
        'alice@example.com',
        'pass123',
        1,
        '13800000001',
        'engineer',
        29
    ),
    (
        'bob',
        'bob@example.com',
        'secret',
        1,
        '13800000002',
        'teacher',
        34
    ),
    (
        'charlie',
        'charlie@example.com',
        'qwerty',
        0,
        '13800000003',
        'student',
        21
    ),
    (
        'dave',
        'dave@example.com',
        '123456',
        1,
        '13800000004',
        'designer',
        27
    ),
    (
        'eve',
        'eve@example.com',
        'letmein',
        1,
        '13800000005',
        'manager',
        31
    );

CREATE INDEX idx_user_name ON `user` (`name`);

CREATE UNIQUE INDEX uniq_user_phone ON `user` (`phone`);

CREATE INDEX idx_user_profession_age_status ON user (profession, age, status);

CREATE INDEX idx_user_email ON user (email);

SHOW INDEX FROM user\G

DROP INDEX idx_user_email ON user;

SHOW GLOBAL STATUS LIKE "Com_______";

SHOW VARIABLES LIKE 'slow_query_log';

SELECT @@slow_query_log;

SELECT @@have_profiling;

SET @@GLOBAL.profiling = 1;

SELECT * FROM user;

SHOW PROFILES;

SHOW PROFILE FOR QUERY 2;

SHOW PROFILE CPU FOR QUERY 2;

EXPLAIN SELECT * FROM user WHERE id = 2;
