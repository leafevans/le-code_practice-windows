CREATE TABLE tb_user (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
    name VARCHAR(10) COMMENT '姓名',
    age INT COMMENT '年龄',
    gender CHAR(1) COMMENT '1: 男, 2: 女',
    phone CHAR(11) COMMENT '手机号'
) COMMENT = '用户基本信息表';

CREATE TABLE tb_user_edu (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
    degree VARCHAR(20) COMMENT '学历',
    major VARCHAR(50) COMMENT '专业',
    primaryschool VARCHAR(50) COMMENT '小学',
    middleschool VARCHAR(50) COMMENT '中学',
    university VARCHAR(50) COMMENT '大学',
    userid INT UNIQUE COMMENT '用户ID',
    CONSTRAINT fk_userid FOREIGN KEY (userid) REFERENCES tb_user (id)
) COMMENT = '用户教育信息表';

INSERT INTO
    tb_user (id, name, age, gender, phone)
VALUES (
        NULL,
        '黄渤',
        45,
        '1',
        '18800001111'
    ),
    (
        NULL,
        '冰冰',
        35,
        '2',
        '18800002222'
    ),
    (
        NULL,
        '码云',
        55,
        '1',
        '18800008888'
    ),
    (
        NULL,
        '李彦宏',
        50,
        '1',
        '18800009999'
    );

INSERT INTO
    tb_user_edu (
        id,
        degree,
        major,
        primaryschool,
        middleschool,
        university,
        userid
    )
VALUES (
        NULL,
        '本科',
        '舞蹈',
        '静安区第一小学',
        '静安区第一中学',
        '北京舞蹈学院',
        1
    ),
    (
        NULL,
        '硕士',
        '表演',
        '朝阳区第一小学',
        '朝阳区第一中学',
        '北京电影学院',
        2
    ),
    (
        NULL,
        '本科',
        '英语',
        '杭州市第一小学',
        '杭州市第一中学',
        '杭州师范大学',
        3
    ),
    (
        NULL,
        '本科',
        '应用数学',
        '阳泉第一小学',
        '阳泉区第一中学',
        '清华大学',
        4
    );
