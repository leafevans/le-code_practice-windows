DROP TABLE IF EXISTS `score`;

CREATE TABLE `score` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(64) NOT NULL,
    `math` TINYINT UNSIGNED NOT NULL DEFAULT 0,
    `english` TINYINT UNSIGNED NOT NULL DEFAULT 0,
    `chinese` TINYINT UNSIGNED NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

INSERT INTO
    `score` (
        `name`,
        `math`,
        `english`,
        `chinese`
    )
VALUES ('张三', 92, 88, 79),
    ('李四', 78, 65, 59),
    ('王五', 58, 55, 62),
    ('赵六', 85, 90, 91),
    ('孙七', 60, 60, 60),
    ('周八', 30, 95, 72),
    ('郑九', 100, 100, 100),
    ('冯十', 64, 59, 48);

SELECT
    name AS `姓名`,
    CASE
        WHEN math >= 85 THEN '优秀'
        WHEN math >= 60 THEN '及格'
        ELSE '不及格'
    END AS `数学`,
    CASE
        WHEN english >= 85 THEN '优秀'
        WHEN english >= 60 THEN '及格'
        ELSE '不及格'
    END AS `英语`,
    CASE
        WHEN chinese >= 85 THEN '优秀'
        WHEN chinese >= 60 THEN '及格'
        ELSE '不及格'
    END AS '语文'
FROM score;
